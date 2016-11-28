package main

import (
	"runtime"
	"github.com/andrepinto/sherlock/core"
	"log"
	"fmt"

)

func main(){
	runtime.GOMAXPROCS(1)

	app, configErr := core.Load()
	if configErr != nil {
		log.Fatal(configErr)
	}

	fmt.Println(app)


	//parseEnvironment(os.Environ())

	app.Run()
}



