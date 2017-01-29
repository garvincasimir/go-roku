package rokuclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var client = http.DefaultClient

func TestGetDeviceApps(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		response := `<apps>
                    <app id="11">Roku Channel Store</app>
                    <app id="12">Netflix</app>
                    <app id="13">Amazon Video on Demand</app>
                    <app id="14">MLB.TV®</app>
                    <app id="26">Free FrameChannel Service</app>
                    <app id="27">Mediafly</app>
                    <app id="28">Pandora</app>
                </apps>`

		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	apps, err := GetDeviceApps(ts.URL, client)

	if err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("%v", apps)
}

func TestActiveApp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		response := `<active-app>
                        <app>Roku</app>
                        <screensaver id="55545" type="ssvr" version="2.0.1">Default screensaver</screensaver>
                    </active-app>`

		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	app, err := GetActiveApp(ts.URL, client)

	if err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("%v", app)
}

func TestLaunchApp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		response := "ok"

		if !strings.HasSuffix(r.RequestURI, "/launch/24") {
			w.WriteHeader(http.StatusBadRequest)
		}

        t.Logf("%v",r.RequestURI)

		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	err := LaunchApp(ts.URL, "24", client)

	if err != nil {
		t.Errorf("%v", err)
	}

}

func TestPressKey(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		response := "ok"

		if !strings.HasSuffix(r.RequestURI, "/presskey/up") {
			w.WriteHeader(http.StatusBadRequest)
		}

        t.Logf("%v",r.RequestURI)

		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	err := PressKey(ts.URL, "up", client)

	if err != nil {
		t.Errorf("%v", err)
	}

}
