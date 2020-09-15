package main

import (
	_ "github.com/dapperkop/api-server/migrations" // only before init "github.com/dapperkop/blank" module

	_ "github.com/dapperkop/blank"
	"github.com/dapperkop/blank/apiserver/httpserver"
	"github.com/dapperkop/api-server/app"
)

func main() {
	app.Setup()

	httpserver.Run()
}
