package osc

import (
	"AudioLoom/midi_osc_integration/internal/playback"
	"log"
	"strconv"

	"github.com/hypebeast/go-osc/osc"
)

type OSCServer struct {
	Address string
	Port    int
	server  *osc.Server
}

func NewOSCServer(address string, port int) *OSCServer {
	return &OSCServer{
		Address: address,
		Port:    port,
		server:  &osc.Server{},
	}
}

func (o *OSCServer) Start(collector *playback.PlaybackCollector) {
	addr := o.Address + ":" + strconv.Itoa(o.Port)
	dispatcher := osc.NewStandardDispatcher()
	o.server = &osc.Server{
		Addr:       addr,
		Dispatcher: dispatcher,
	}

	playback.HandlePlayback(dispatcher, collector)

	log.Printf("Starting OSC server on %s\n", addr)
	err := o.server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start OSC server: %v\n", err)
	}
}
