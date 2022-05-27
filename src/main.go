package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/conf"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
)

// Flavour We hard wire flavour at time of build using linker.
var Flavour string

// BaseDir Base direcory of executable
var BaseDir string

func main() {
	// get current base dir
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println("Exec base dir", exPath)
	BaseDir = exPath

	// get env file
	envFileBytes, err := ioutil.ReadFile(filepath.Join(exPath, "env.json"))
	if err != nil {
		fmt.Print(err)
	}

	// setup
	appCtx := framework.Init(Flavour, string(envFileBytes), BaseDir)
	appCtx.InitMongo(conf.MongoConnections)
	appCtx.SetAsMainContext()
	conf.Route(appCtx)
	// load bootstrap
	conf.Bootstrap(appCtx)
	// serve
	appCtx.Listen()
}
