package rokuclient

//App represents an app or screensaver
type App struct {
	AppID   string `xml:"id,attr"`
	Name    string `xml:",innerxml"`
	Type    string `xml:"type,attr"`
	Version string `xml:"version,attr"`
}

//AppWrapper we need this just for unmarshalling
type AppWrapper struct {
	Apps []App `xml:"app"`
}
