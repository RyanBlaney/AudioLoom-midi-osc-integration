package models

import "sync"

type SessionData struct {
	ID           string
	PlaybackData PlaybackData
	Clients      map[string]ClientData // ClientID -> Data
	HostID       string
	Roles        map[string]string // ClientID -> Role
	mu           sync.Mutex
}
