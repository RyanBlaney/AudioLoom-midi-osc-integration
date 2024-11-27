package main

import (
	"AudioLoom/midi_osc_integration/internal/osc"
	"AudioLoom/midi_osc_integration/internal/playback"
	"time"
)

func main() {
	collector := playback.NewPlaybackCollector()

	oscServer := osc.NewOSCServer("0.0.0.0", 9000)
	go func() {
		oscServer.Start(collector)
	}()

	time.Sleep(time.Second * 3)

	ardour := osc.NewOSCClient("127.0.0.1", 3819)
	ardour.SendMessage("/transport_play")

	reaper := osc.NewOSCClient("127.0.0.1", 9003)
	reaper.SendMessage("/transport/play")

	select {}
}
