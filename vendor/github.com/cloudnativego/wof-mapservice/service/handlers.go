package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func queryMapsHandler(formatter *render.Render, repo mapRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		mapList, err := repo.GetMapList()
		if err == nil {
			formatter.JSON(w, http.StatusOK, mapList)
		} else {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func queryMapDetailsHandler(formatter *render.Render, repo mapRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		mapID := vars["mapId"]

		gameMap, err := repo.GetMap(mapID)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, err.Error())
		} else {
			formatter.JSON(w, http.StatusOK, gameMap)
		}
	}
}

func putMapHandler(formatter *render.Render, repo mapRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		mapID := vars["mapId"]
		payload, _ := ioutil.ReadAll(req.Body)
		var gameMap WofMap
		err := json.Unmarshal(payload, &gameMap)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = repo.UpdateMap(mapID, gameMap)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		formatter.JSON(w, http.StatusCreated, gameMap)
	}
}
