package main

import (
	iotContainer "github.com/helmutkemper/iotmaker.docker.server/container"
	iotServer "github.com/helmutkemper/iotmaker.docker.server/server"
	"log"
	"net/http"
	"sync"
)

func main() {

	var container = iotContainer.NewWebContainer()

	var project = iotServer.Project{
		ListenAndServer: iotServer.ListenAndServer{
			InAddress: "0.0.0.0:8080",
		},
		DebugServerEnable: true,
		Handle: map[string]iotServer.Handle{
			"/admin/container/listAll": {
				Func: container.ListAll,
				HeaderToAdd: map[iotServer.HeaderList]iotServer.HeaderApplication{
					iotServer.KHeaderListContentType: iotServer.KHeaderApplicationTypeJSon,
				},
				Method: http.MethodGet,
			},
		},
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func(project iotServer.Project) {
		log.Fatalf("server error: %v", iotServer.NewServer(project))
	}(project)

	wg.Wait()

}
