package rokuclient

//Device details
type Device struct {
	Name         string `xml:"user-device-name"`
	ModelName    string `xml:"model-name"`
	SerialNumber string `xml:"serial-number"`
	Endpoint     string
}

//GetDispalName can be used to display device in a list
func (d Device) GetDispalName() string {
	if d.Name != "" {
		return d.Name
	}
	return d.ModelName
}
