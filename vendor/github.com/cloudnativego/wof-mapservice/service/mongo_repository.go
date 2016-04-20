package service

import (
	"errors"

	"github.com/cloudnativego/cfmgo"
	"github.com/cloudnativego/cfmgo/params"
	"gopkg.in/mgo.v2/bson"
)

// MongoMapRepository anchor struct for mongo repository
type MongoMapRepository struct {
	Collection cfmgo.Collection
}

// NewMongoRepository creates a new mongo map repository
func NewMongoRepository(col cfmgo.Collection) (repo *MongoMapRepository) {
	repo = &MongoMapRepository{
		Collection: col,
	}
	return
}

// GetMap retrieves a map from the mongo database
func (repo *MongoMapRepository) GetMap(mapID string) (gameMap WofMap, err error) {
	repo.Collection.Wake()
	record, err := repo.getMapByID(mapID)
	if err == nil {
		gameMap = convertRecordToGameMap(record)
	}
	return
}

// GetMapList retrieves all maps
func (repo *MongoMapRepository) GetMapList() (maps []WofMap, err error) {
	repo.Collection.Wake()
	var mapRecords []mapRecord
	_, err = repo.Collection.Find(cfmgo.ParamsUnfiltered, &mapRecords)
	if err == nil {
		maps = make([]WofMap, len(mapRecords))
		for k, v := range mapRecords {
			maps[k] = convertRecordToGameMap(v)
		}
	}
	return
}

// UpdateMap updates an individual map
func (repo *MongoMapRepository) UpdateMap(mapID string, gameMap WofMap) (err error) {
	repo.Collection.Wake()

	var recordID bson.ObjectId
	foundMap, err := repo.getMapByID(mapID)
	if err != nil {
		recordID = bson.NewObjectId()
	} else {
		recordID = foundMap.RecordID
	}

	newMapRecord := convertGameMapToMapRecord(gameMap, recordID)
	_, err = repo.Collection.UpsertID(recordID, newMapRecord)
	return
}

func (repo *MongoMapRepository) getMapByID(mapID string) (record mapRecord, err error) {
	var records []mapRecord
	query := bson.M{"map_id": mapID}
	params := &params.RequestParams{
		Q: query,
	}

	count, err := repo.Collection.Find(params, &records)
	if count == 0 {
		err = errors.New("Map record not found.")
	}
	if err == nil {
		record = records[0]
	}
	return
}

type mapRecord struct {
	RecordID bson.ObjectId     `bson:"_id,omitempty",json:"_id"`
	ID       string            `bson:"map_id",json:"map_id"`
	Metadata mapRecordMetadata `bson:"metadata",json:"metadata"`
	Tiles    [][]mapRecordTile `bson:"tiles",json:"tiles"`
}

type mapRecordMetadata struct {
	Author      string `bson:"author",json:"author"`
	Description string `bson:"description",json:"description"`
}

type mapRecordTile struct {
	ID         string `bson:"id",json:"id"`
	Sprite     string `bson:"sprite",json:"sprite"`
	AllowUp    bool   `bson:"allow_up",json:"allow_up"`
	AllowDown  bool   `bson:"allow_down",json:"allow_down"`
	AllowLeft  bool   `bson:"allow_left",json:"allow_left"`
	AllowRight bool   `bson:"allow_right",json:"allow_right"`
	TileName   string `bson:"tile_name",json:"tile_name"`
}
