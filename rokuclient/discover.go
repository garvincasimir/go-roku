package rokuclient

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	ssdpHost       = "239.255.255.250:1900"
	maxRespSize    = 2048
	addressPattern = `location:\s?(http://[^/]+/)`
)

const message = "M-SEARCH * HTTP/1.1\r\n" +
	"HOST: " + ssdpHost + "\r\n" +
	"ST:roku:ecp\r\n" +
	"MAN:\"ssdp:discover\"\r\n"

var locationRegex *regexp.Regexp

//Discover roku devices on the network. waitSeconds sets the time limit for discovery.
func Discover(waitSeconds int) []Device {

	//UDP Multicast address
	addr1, err := net.ResolveUDPAddr("udp", ssdpHost)
	if err != nil {
		log.Fatal(err)
	}

	//Listen for incoming messages
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	//Hello?? Anybody out there???
	conn.WriteTo([]byte(message), addr1)

	buffer := make([]byte, maxRespSize)

	var endpoints []string

	//Set a reasonable timeout for reads
	conn.SetReadDeadline(time.Now().Add(time.Duration(waitSeconds) * time.Second))

	timeout := time.After(time.Duration(waitSeconds) * time.Second)

LOOP:
	for {

		select {
		case <-timeout:
			break LOOP
		default:
			//This blocks so we must set a reasonable timeout
			n, ad, err := conn.ReadFrom(buffer)
			if err != nil {
				if err, ok := err.(net.Error); ok && err.Timeout() {
					break LOOP
				} else {
					log.Fatalf("Error reading from socket: %v", err)
				}
			} else {
				log.Printf("Response From: %v  with %v bytes of data", ad.String(), n)
			}

			//Can we express this without so many nested ifs?
			if n > 0 {
				if data := strings.ToLower(string(buffer[0:n])); strings.Contains(data, "roku") {
					if match := locationRegex.FindStringSubmatch(data); len(match) == 2 {
						endpoints = append(endpoints, match[1])
					} else {
						log.Printf("Could not match on result: %v with match %v", data, match)
					}

				}

			}
		}

	}

	if len(endpoints) == 0 {
		log.Fatal("No Devices Found")
	}

	totalDevices := len(endpoints)
	devices := make([]Device, totalDevices)
	ch := make(chan deviceResult, totalDevices)
	client := &http.Client{}

	for i, v := range endpoints {
		go func(ep string, client *http.Client, index int) {
			result := queryDevice(ep, client)
			result.Position = index
			ch <- result
		}(v, client, i)
	}

	for range endpoints {
		if result := <-ch; result.Error == nil {
			devices[result.Position] = result.Device
		} else {
			log.Printf("%v", result.Error)
		}
	}

	return devices
}

func queryDevice(endpoint string, client *http.Client) deviceResult {

	requestURL := endpoint + "query/device-info"
	response, err := client.Get(requestURL)
	result := deviceResult{}

	if resultData, err := ioutil.ReadAll(response.Body); err == nil {
		xml.Unmarshal(resultData, &result.Device)
	}
	result.Error = err

	return result
}

func init() {
	locationRegex = regexp.MustCompile(addressPattern)
}

type deviceResult struct {
	Device   Device
	Error    error
	Position int //No need for append
}
