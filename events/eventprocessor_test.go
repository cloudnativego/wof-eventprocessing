package events

import (
	"strings"
	"testing"
	"time"

	"code.google.com/p/go-uuid/uuid"
	mapservice "github.com/cloudnativego/wof-mapservice/service"
)

const (
	author = "Hingle McCringleBerry"
)

var (
	player1 PlayerState
	player2 PlayerState
	id1     = uuid.New()
	id2     = uuid.New()
)

func TestLegitMoveModifiesState(t *testing.T) {
	initialState := generateEmptyState(10, author)
	moveEvent := &PlayerMovedEvent{
		GameID:       initialState.GameID,
		PlayerID:     id1,
		TargetTileID: initialState.GameMap.Tiles[0][1].ID, // move down
		Timestamp:    time.Now().Unix(),
	}

	newState, err := ProcessMovedEvent(initialState, moveEvent)

	if err != nil {
		t.Errorf("Received an error when moving and we shouldn't have: %s", err.Error())
		return
	}

	if newState.Players[id1].CurrentTileID != moveEvent.TargetTileID {
		t.Errorf("New state does not reflect intended target tile for move: player current tile: %s, should be %s.",
			newState.Players[id1].CurrentTileID, moveEvent.TargetTileID)
	}
}

func TestMissingPlayerMoveShouldBeRejected(t *testing.T) {
	initialState := generateEmptyState(10, author)
	moveEvent := &PlayerMovedEvent{
		GameID:       initialState.GameID,
		PlayerID:     uuid.New(),                          // bogus player
		TargetTileID: initialState.GameMap.Tiles[0][1].ID, // move down
		Timestamp:    time.Now().Unix(),
	}

	_, err := ProcessMovedEvent(initialState, moveEvent)

	if err == nil {
		t.Errorf("Should have rejected a move from a missing player, but it was allowed.")
	}
}

func TestNotTraversableIsRejectedWithError(t *testing.T) {
	initialState := generateEmptyState(10, author)
	initialState.GameMap.Tiles[0][1].Traversable = false
	moveEvent := &PlayerMovedEvent{
		GameID:       initialState.GameID,
		PlayerID:     id1,
		TargetTileID: initialState.GameMap.Tiles[0][1].ID, // move down
		Timestamp:    time.Now().Unix(),
	}

	_, err := ProcessMovedEvent(initialState, moveEvent)

	if err == nil {
		t.Errorf("We should have gotten an error rejecting the move as non-traversable, but we didn't.")
	}
}

func TestInvalidMoveRejectsWithError(t *testing.T) {
	initialState := generateEmptyState(10, author)
	moveEvent := &PlayerMovedEvent{
		GameID:       initialState.GameID,
		PlayerID:     id1,
		TargetTileID: uuid.New(), // move into the nether regions
		Timestamp:    time.Now().Unix(),
	}

	_, err := ProcessMovedEvent(initialState, moveEvent)

	if err == nil {
		t.Errorf("Should have rejected the move due to an invalid/non-existent tile, but didn't.")
	}

	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("Error should be rejected due to not found, but got '%s' instead", err.Error())
	}
}

func TestWrongRealityEventsAreRejectedWithError(t *testing.T) {
	initialState := generateEmptyState(10, author)

	movedEvent := &PlayerMovedEvent{
		GameID:       uuid.New(),
		PlayerID:     id1,
		TargetTileID: initialState.GameMap.Tiles[0][1].ID, // valid move
		Timestamp:    time.Now().Unix(),
	}

	joinedEvent := &PlayerJoinedEvent{
		GameID:    uuid.New(), // bad reality/ wrong game
		PlayerID:  uuid.New(),
		Name:      "Devontay",
		Sprite:    "knight-64",
		Timestamp: time.Now().Unix(),
	}
	_, err := ProcessMovedEvent(initialState, movedEvent)
	if err == nil {
		t.Errorf("We should have gotten an invalid reality error for moved event, but didn't.")
	}

	_, err = ProcessPlayerJoinedEvent(initialState, joinedEvent)
	if err == nil {
		t.Errorf("We should have gotten an invalid reality error for joined event, but didn't")
	}
}

func TestPlayerJoinedEventUpdatesReality(t *testing.T) {
	initialState := generateEmptyState(10, author)
	joinEvent := &PlayerJoinedEvent{
		GameID:    initialState.GameID,
		PlayerID:  uuid.New(),
		Name:      "Devontay",
		Sprite:    "knight-64",
		Timestamp: time.Now().Unix(),
	}

	newState, err := ProcessPlayerJoinedEvent(initialState, joinEvent)

	if err != nil {
		t.Errorf("Should not have gotten an error adding player, but did: %s", err.Error())
	}

	if len(newState.Players) != 3 {
		t.Errorf("Should have 3 players after add, got %d", len(newState.Players))
	}

	if newState.Players[joinEvent.PlayerID].Name != "Devontay" {
		t.Errorf("Player state reflected incorrectly, got %+v", newState.Players[joinEvent.PlayerID])
	}

	if newState.Players[joinEvent.PlayerID].Sprite != joinEvent.Sprite {
		t.Errorf("Player has the wrong sprite, got %s", newState.Players[joinEvent.PlayerID].Sprite)
	}

	if newState.Players[joinEvent.PlayerID].Hitpoints != 100 {
		t.Errorf("Player didn't get the right amount of hitpoints, got %d", newState.Players[joinEvent.PlayerID].Hitpoints)
	}
}

func generateEmptyState(mapSize int, author string) (state *GameState) {
	gmap := generateTestMap(mapSize, author)
	state = &GameState{
		GameID:  uuid.New(),
		GameMap: gmap,
		Players: generateDefaultPlayers(gmap.Tiles[0][0].ID),
	}
	return
}

func generateDefaultPlayers(spawnTile string) (players map[string]PlayerState) {

	players = make(map[string]PlayerState)

	player1 = PlayerState{
		Hitpoints:     100,
		ID:            id1,
		CurrentTileID: spawnTile,
	}
	player2 = PlayerState{
		Hitpoints:     100,
		ID:            id2,
		CurrentTileID: spawnTile,
	}
	players[id1] = player1
	players[id2] = player2
	return
}

func generateTestMap(size int, author string) (gameMap mapservice.WofMap) {

	gameMap.Metadata.Author = author
	gameMap.Metadata.Description = "Auto-generated Test Map"
	gameMap.ID = uuid.New()

	tiles := make([][]mapservice.MapTile, size)
	for row := 0; row < size; row++ {
		tiles[row] = make([]mapservice.MapTile, size)
		for col := 0; col < size; col++ {
			tiles[row][col] = makeTile()
		}
	}
	gameMap.Tiles = tiles

	return
}

func makeTile() (tile mapservice.MapTile) {
	tile.AllowDown = true
	tile.AllowLeft = true
	tile.AllowRight = true
	tile.AllowUp = true
	tile.Traversable = true
	tile.Sprite = ""
	tile.TileName = "test-tile"
	tile.ID = uuid.New()
	return
}
