package game

// An dummy implementation of the MobileConnector that doesn't do anything
type dummyMobileConnector struct{}

func (*dummyMobileConnector) Vibrate()            {}
func (*dummyMobileConnector) ShowEndDialog()      {}
func (*dummyMobileConnector) GetDeviceID() string { return "dummy-device-id" }

// MobileConnector is the interface used interact with the mobile device. Set by the Android or iOS device, if available
type MobileConnector interface {
	Vibrate()
	ShowEndDialog()
	GetDeviceID() string
}
