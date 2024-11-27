package playback

import (
	"log"

	"github.com/hypebeast/go-osc/osc"
)

func HandlePlayback(
	dispatcher *osc.StandardDispatcher,
	collector *PlaybackCollector,
) {
	dispatcher.AddMsgHandler("/playback/start", func(msg *osc.Message) {
		collector.mu.Lock()
		defer collector.mu.Unlock()

		collector.PlaybackData.IsPlaying = true

		log.Printf("Received start playback message: %v\n", msg.Arguments...)
	})

	dispatcher.AddMsgHandler("/playback/stop", func(msg *osc.Message) {
		collector.mu.Lock()
		defer collector.mu.Unlock()

		collector.PlaybackData.IsPlaying = false
		log.Printf("Received stop playback message: %v\n", msg.Arguments...)
	})

	dispatcher.AddMsgHandler("/time", func(msg *osc.Message) {
		collector.mu.Lock()
		defer collector.mu.Unlock()

		position, ok := msg.Arguments[0].(float32)
		if ok {
			collector.PlaybackData.Position = float64(position)
			log.Printf("Position changed to %v\n", msg.Arguments...)
		} else {
			log.Printf("Failed to change position to non-float value: %v\n", msg.Arguments...)
		}
	})

	dispatcher.AddMsgHandler("*", func(msg *osc.Message) {
		log.Printf("Received message: %v\n", msg.Address)
		log.Printf("   %v\n", msg.Arguments...)
	})
}
