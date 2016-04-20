package events

import (
	mapservice "github.com/cloudnativego/wof-mapservice/service"
)

// PlayerMovedEvent represents a player move.
// TODO: document why no source tile ID
type PlayerMovedEvent struct {
	GameID       string `json:"game_id"`
	PlayerID     string `json:"player_id"`
	TargetTileID string `json:"target_tile_id"`
	Timestamp    int64  `json:"timestamp"`
}

// PlayerJoinedEvent indicates a player joining the game
type PlayerJoinedEvent struct {
	GameID    string `json:"game_id"`
	PlayerID  string `json:"player_id"`
	Sprite    string `json:"sprite"`
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
}

// GameState represents the state of a game
type GameState struct {
	GameID  string                 `json:"game_id"`
	GameMap mapservice.WofMap      `json:"game_map"`
	Players map[string]PlayerState `json:"players"`
}

// PlayerState represents the current state of a player
type PlayerState struct {
	Hitpoints     uint   `json:"hit_points"`
	ID            string `json:"player_id"`
	CurrentTileID string `json:"current_tile_id"`
	Name          string `json:"name"`
	Sprite        string `json:"sprite"`
}
