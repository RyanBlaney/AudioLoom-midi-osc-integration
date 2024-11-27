package models

import "sync"

type ClientData struct {
	HostName string
	DAW      string
	OS       string
	IP       string
	Port     int

	SupportsOSC  bool
	SupportsMIDI bool
	OSC          OSCData
	MIDI         MIDIData

	SupportedFormats []string
	IsHost           bool
	DAWVersion       string // Fault tolerance
	ProtocolVersion  string // MIDI1.0 vs MIDI2.0
	mu               sync.Mutex
}
