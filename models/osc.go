package models

type OSCData struct {
	OSCAddress string
	OSCPathMap map[string]string // Path -> Action
	OSCState   string            // connected, disconnected, etc.
}
