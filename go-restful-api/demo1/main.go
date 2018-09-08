package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-study-note/go-restful-api/demo1/router"
	"log"
	"net/http"
	"time"
)

var port = ":8080"

func main() {
	// create the gin engine.
	g := gin.New()
	// gin middlewares
	middlewares := []gin.HandlerFunc{}
	// routes
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", port)
	log.Printf(http.ListenAndServe(port, g).Error())
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
