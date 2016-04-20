package service

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/cf-tools"
	"github.com/cloudnativego/cfmgo"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	repo := initRepository()
	initRoutes(mx, formatter, repo)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, repo mapRepository) {
	mx.HandleFunc("/api/maps", queryMapsHandler(formatter, repo)).Methods("GET")
	mx.HandleFunc("/api/maps/{mapId}", queryMapDetailsHandler(formatter, repo)).Methods("GET")
	mx.HandleFunc("/api/maps/{mapId}", putMapHandler(formatter, repo)).Methods("PUT")
}

func initRepository() (repo mapRepository) {
	appEnv, _ := cfenv.Current()
	dbServiceURI, err := cftools.GetVCAPServiceProperty("mongodb", "uri", appEnv)
	if err != nil || dbServiceURI == "" {
		if err != nil {
			fmt.Printf("\nError retrieving database configuration: %v\n", err)
		}
		fmt.Println("MongoDB was not detected; configuring FAKE repository...")
		repo = NewFakeRepository()
	} else {
		mapCollection := cfmgo.Connect(cfmgo.NewCollectionDialer, dbServiceURI, "maps")
		fmt.Println("Connecting to MongoDB service: mongodb...")
		repo = NewMongoRepository(mapCollection)
	}
	return
}
