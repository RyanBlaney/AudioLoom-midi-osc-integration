package models

type MIDIData struct {
	MIDIChannels       []int
	MIDIPort           string
	MIDIControlChanges map[int]int // CC number -> value
	NodeData           []MIDINote
}

type MIDINote struct {
	NoteNumber int
	Velocity   int
	Duration   float64
}
