package soupev1

type ClientConnectedEvent struct {
	Slot    int32
	SteamID uint64
}

type ClientDisconnectEvent struct {
	Slot    int32
	SteamID uint64
}
