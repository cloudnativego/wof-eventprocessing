package service

import "gopkg.in/mgo.v2/bson"

func convertRecordToGameMap(record mapRecord) (gameMap WofMap) {
	gameMap = WofMap{
		ID:       record.ID,
		Metadata: convertRecordMetadataToGameMetadata(record.Metadata),
		Tiles:    convertRecordTilesToGameTiles(record.Tiles),
	}
	return
}

func convertRecordMetadataToGameMetadata(metadata mapRecordMetadata) (gameMetadata MapMetadata) {
	gameMetadata = MapMetadata{
		Author:      metadata.Author,
		Description: metadata.Description,
	}
	return
}

func convertRecordTilesToGameTiles(recordTiles [][]mapRecordTile) (gameTiles [][]MapTile) {
	gameTiles = make([][]MapTile, len(recordTiles))
	for r := 0; r < len(recordTiles); r++ {
		gameTiles[r] = make([]MapTile, len(recordTiles[r]))
		for c := 0; c < len(gameTiles[r]); c++ {
			gameTiles[r][c] = convertRecordTileToGameTile(recordTiles[r][c])
		}
	}
	return
}

func convertRecordTileToGameTile(recordTile mapRecordTile) (gameTile MapTile) {
	gameTile = MapTile{
		ID:         recordTile.ID,
		Sprite:     recordTile.Sprite,
		AllowUp:    recordTile.AllowUp,
		AllowDown:  recordTile.AllowDown,
		AllowLeft:  recordTile.AllowLeft,
		AllowRight: recordTile.AllowRight,
		TileName:   recordTile.TileName,
	}
	return
}

func convertGameMapToMapRecord(gameMap WofMap, recordID bson.ObjectId) (record *mapRecord) {
	record = &mapRecord{
		RecordID: recordID,
		ID:       gameMap.ID,
		Metadata: convertGameMetadataToRecordMetadata(gameMap.Metadata),
		Tiles:    convertGameTilesToRecordTiles(gameMap.Tiles),
	}
	return
}

func convertGameMetadataToRecordMetadata(gameMetadata MapMetadata) (recordMetadata mapRecordMetadata) {
	recordMetadata = mapRecordMetadata{
		Author:      gameMetadata.Author,
		Description: gameMetadata.Description,
	}
	return
}

func convertGameTilesToRecordTiles(gameTiles [][]MapTile) (recordTiles [][]mapRecordTile) {
	recordTiles = make([][]mapRecordTile, len(gameTiles))
	for r := 0; r < len(gameTiles); r++ {
		recordTiles[r] = make([]mapRecordTile, len(gameTiles[r]))
		for c := 0; c < len(recordTiles[r]); c++ {
			recordTiles[r][c] = convertGameTileToRecordTile(gameTiles[r][c])
		}
	}
	return
}

func convertGameTileToRecordTile(gameTile MapTile) (recordTile mapRecordTile) {
	recordTile = mapRecordTile{
		ID:         gameTile.ID,
		Sprite:     gameTile.Sprite,
		AllowUp:    gameTile.AllowUp,
		AllowDown:  gameTile.AllowDown,
		AllowLeft:  gameTile.AllowLeft,
		AllowRight: gameTile.AllowRight,
		TileName:   gameTile.TileName,
	}
	return
}
