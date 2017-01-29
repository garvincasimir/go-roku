package rokuclient

//ActiveApp is the current ActiveApp
type ActiveApp struct {
	App         App `xml:"app"`
	ScreenSaver App `xml:"screensaver"`
}
