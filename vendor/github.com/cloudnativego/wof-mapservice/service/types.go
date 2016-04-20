package service

type mapRepository interface {
	GetMap(mapID string) (gameMap WofMap, err error)
	GetMapList() (maps []WofMap, err error)
	UpdateMap(mapID string, gameMap WofMap) (err error)
}

// WofMap represents a map usable for play within World of FluxCraft
type WofMap struct {
	Tiles    [][]MapTile `json:"tiles"`
	ID       string      `json:"id"`
	Metadata MapMetadata `json:"metadata"`
}

// MapMetadata represents metadata about a map, not the map itself
type MapMetadata struct {
	Author      string `json:"author"`
	Description string `json:"description"`
}

// MapTile represents an individual tile within a map
type MapTile struct {
	ID          string `json:"id"`
	Sprite      string `json:"sprite"`
	AllowUp     bool   `json:"allow_up"`
	AllowDown   bool   `json:"allow_down"`
	AllowLeft   bool   `json:"allow_left"`
	AllowRight  bool   `json:"allow_right"`
	Traversable bool   `json:"traversable"`
	TileName    string `json:"tile_name"`
}
