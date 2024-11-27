package osc

import (
	"log"

	"github.com/hypebeast/go-osc/osc"
)

type OSCClient struct {
	client *osc.Client
}

// NewOSCClient initializes a new OSC client
func NewOSCClient(address string, port int) *OSCClient {
	return &OSCClient{
		client: osc.NewClient(address, port),
	}
}

// SendMessage sends a generic OSC message
func (o *OSCClient) SendMessage(path string, args ...interface{}) {
	msg := osc.NewMessage(path)
	for _, arg := range args {
		msg.Append(arg)
	}

	err := o.client.Send(msg)
	if err != nil {
		log.Printf("Failed to send message to %s:%d: %v\n", o.client.IP, o.client.Port, err)
	}
}

// Play sends a "play" command to Ardour
func (o *OSCClient) Play() {
	o.SendMessage("/play/start")
}

// Stop sends a "stop" command to Ardour
func (o *OSCClient) Stop() {
	o.SendMessage("/play/stop")
}

// SetPosition sends a "set position" command to Ardour
func (o *OSCClient) SetPosition(position float64) {
	o.SendMessage("/playhead/position", position)
}
