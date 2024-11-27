package models

type SessionData struct {
	ID           string
	PlaybackData PlaybackData
	Clients      map[string]ClientData
	HostID       string
}
