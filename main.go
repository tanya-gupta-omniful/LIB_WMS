package main

import (
	appinit "WMS/init"
	"context"
	"fmt"
	"time"

	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/shutdown"
)

const (
	modeWorker     = "worker"
	modeHttp       = "http"
	modeMigration  = "migration"
	upMigration    = "up"
	downMigration  = "down"
	forceMigration = "force"
)

func main(){
	fmt.Println("hello")
	err := config.Init(time.Second * 10)
	if err != nil {
		log.Panicf("Error while initialising config, err: %v", err)
		panic(err)
	}
	ctx, err := config.TODOContext()
	if err != nil {
		log.Panicf("Error while getting context from config, err: %v", err)
		panic(err)
	}
	//initoialise connection
    appinit.Initialize(ctx)

	//run server
	runHttpServer(ctx)

}
func runHttpServer(ctx context.Context) {

	server := http.InitializeServer(config.GetString(ctx, "server.port"), 10*time.Second, 10*time.Second, 70*time.Second)

	// Initialize middlewares and routes
//err := router.Initialize(ctx, server)
//if err != nil {
//	log.Errorf(err.Error())
//	panic(err)
//}
//
//err = router.InternalRoutes(ctx, server)
//if err != nil {
//	log.Errorf(err.Error())
//	panic(err)
//}

	log.Infof("Starting server on port" + config.GetString(ctx, "server.port"))

	err := server.StartServer("WM-service")
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

	<-shutdown.GetWaitChannel()
}