package playback

import (
	"AudioLoom/midi_osc_integration/models"
	"log"
	"sync"
)

type PlaybackCollector struct {
	PlaybackData *models.PlaybackData
	mu           sync.Mutex
}

func NewPlaybackCollector() *PlaybackCollector {
	return &PlaybackCollector{
		PlaybackData: &models.PlaybackData{},
	}
}

func (p *PlaybackCollector) UpdatePlayback(action string, value interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()

	switch action {
	case "start":
		p.PlaybackData.IsPlaying = true
	case "stop":
		p.PlaybackData.IsPlaying = false
	case "position":
		position, ok := value.(float64)
		if ok {
			p.PlaybackData.Position = position
		} else {
			log.Printf("Invalid position value: %v", value)
		}
	default:
		log.Printf("Unknown action: %s", action)
	}
}
