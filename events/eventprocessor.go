package events

import (
	"fmt"

	mapservice "github.com/cloudnativego/wof-mapservice/service"
)

// ProcessMovedEvent processes a received moved event, reflecting the change in state.
func ProcessMovedEvent(reality *GameState, newEvent *PlayerMovedEvent) (newReality *GameState, err error) {

	if reality.GameID != newEvent.GameID {
		return nil, fmt.Errorf("Invalid reality, expecting %s but event was for %s", reality.GameID, newEvent.GameID)
	}

	newReality = copyReality(reality)

	targetTile, err := findTargetTile(newEvent.TargetTileID, reality.GameMap)
	if err != nil {
		return nil, err
	}

	if err == nil && !targetTile.Traversable {
		err = fmt.Errorf("Target Tile %s is non-traversable.", newEvent.TargetTileID)
		return nil, err
	}

	player, ok := newReality.Players[newEvent.PlayerID]
	if !ok {
		err = fmt.Errorf("Missing player %s, rejecting move.", newEvent.PlayerID)
		return nil, err
	}

	player.CurrentTileID = newEvent.TargetTileID
	newReality.Players[newEvent.PlayerID] = player
	return
}

// ProcessPlayerJoinedEvent processes a player join event
func ProcessPlayerJoinedEvent(reality *GameState, newEvent *PlayerJoinedEvent) (newReality *GameState, err error) {

	if reality.GameID != newEvent.GameID {
		return nil, fmt.Errorf("Invalid reality, expecting %s but event was for %s", reality.GameID, newEvent.GameID)
	}

	newReality = copyReality(reality)

	newPlayer := PlayerState{
		ID:        newEvent.PlayerID,
		Sprite:    newEvent.Sprite,
		Name:      newEvent.Name,
		Hitpoints: 100, // TODO: this shouldn't be hardcoded. wtf.
	}

	newReality.Players[newPlayer.ID] = newPlayer
	return
}

func copyReality(reality *GameState) (realityCopy *GameState) {
	realityCopy = &GameState{
		GameID:  reality.GameID,
		GameMap: reality.GameMap,
		Players: reality.Players,
	}
	return
}

func findTargetTile(tileID string, gameMap mapservice.WofMap) (tile *mapservice.MapTile, err error) {
	for r := range gameMap.Tiles {
		for c := range gameMap.Tiles[r] {
			if gameMap.Tiles[r][c].ID == tileID {
				return &gameMap.Tiles[r][c], nil
			}
		}
	}
	return nil, fmt.Errorf("Tile %s not found.", tileID)
}
