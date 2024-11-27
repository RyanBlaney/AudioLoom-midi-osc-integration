package models

type PlaybackData struct {
	IsPlaying       bool
	IsRecording     bool
	Position        float64
	Tempo           float64
	TimeSignature   string
	Bar             int
	Beat            int
	SampleRate      int
	SyncMode        string
	LastSyncedAt    int64
	LoopStart       float64
	LoopEnd         float64
	MarkerPositions map[string]float64
}
