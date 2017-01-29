package rokuclient

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

//GetDeviceApps returns a list of apps installed on a device
func GetDeviceApps(endpoint string, client *http.Client) ([]App, error) {
	wrapper := AppWrapper{}
	err := deviceQuery(endpoint, "query/apps", client, &wrapper)

	return wrapper.Apps, err
}

//GetActiveApp returns the current active app
func GetActiveApp(endpoint string, client *http.Client) (ActiveApp, error) {
	app := ActiveApp{}
	err := deviceQuery(endpoint, "query/active-app", client, &app)

	return app, err
}

//LaunchApp launches an App identified by id
func LaunchApp(endpoint, id string, client *http.Client) error {
	err := deviceAction(endpoint, "launch", id, client)
	return err
}

//PressKey allows clients to act as roku remote
func PressKey(endpoint, key string, client *http.Client) error {
	err := deviceAction(endpoint, "keypress", key, client)
	return err
}

func deviceAction(endpoint, action string, param string, client *http.Client) error {
	if !strings.HasSuffix(endpoint, "/") {
		endpoint += "/"
	}
	requestURL := endpoint + action + "/" + param
	_, err := client.Post(requestURL, "", nil)

	return err
}

func deviceQuery(endpoint, query string, client *http.Client, result interface{}) error {
	if !strings.HasSuffix(endpoint, "/") {
		endpoint += "/"
	}
	requestURL := endpoint + query
	response, err := client.Get(requestURL)

	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(content, &result)

	if err != nil {
		return err
	}

	return nil
}
