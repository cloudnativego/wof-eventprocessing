package service

import "errors"

// FakeRepository anchor struct for the fake repo implementation
type FakeRepository struct {
	maps map[string]WofMap
}

// NewFakeRepository creates a new fake repo
func NewFakeRepository() *FakeRepository {
	repo := &FakeRepository{}
	repo.maps = make(map[string]WofMap)
	return repo
}

// GetMap returns a game map
func (repo *FakeRepository) GetMap(mapID string) (gameMap WofMap, err error) {
	gameMap, exists := repo.maps[mapID]
	if !exists {
		err = errors.New("Map doesn't exist")
	}
	return
}

// GetMapList gets a list of all maps
func (repo *FakeRepository) GetMapList() (maps []WofMap, err error) {
	maps = make([]WofMap, len(repo.maps))
	idx := 0
	if len(repo.maps) > 0 {
		for _, value := range repo.maps {
			maps[idx] = value
			idx++
		}
	}
	return
}

// UpdateMap updates or creates a map
func (repo *FakeRepository) UpdateMap(mapID string, gameMap WofMap) (err error) {
	repo.maps[mapID] = gameMap
	return
}
