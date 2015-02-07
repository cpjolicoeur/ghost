package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gorilla/handlers"
)

const AppVersion = "0.2.0"

func main() {
	app := cli.NewApp()
	app.Name = "ghost"
	app.Usage = "Static file server written in Go"
	app.Version = AppVersion
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Value:  8000,
			Usage:  "port ghost will serve on",
			EnvVar: "GHOST_PORT,PORT",
		},
	}
	app.Action = func(c *cli.Context) {
		fs := http.FileServer(http.Dir("."))
		http.Handle("/", handlers.LoggingHandler(os.Stdout, fs))

		port := c.Int("port")
		fmt.Printf("Starting ghost - http://localhost:%d\n", port)
		log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	}
	app.Run(os.Args)
}
