package main

import (
	"AudioLoom/midi_osc_integration/internal/osc"
	"AudioLoom/midi_osc_integration/internal/playback"
)

func main() {
	collector := playback.NewPlaybackCollector()

	oscServer := osc.NewOSCServer("127.0.0.1", 9000)
	oscServer.Start(collector)
}
