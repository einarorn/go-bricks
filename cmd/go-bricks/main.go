package main

import "log"

func main() {
	app, err := newApp()
	if err != nil {
		log.Fatalf("Error starting app: %v", err)
	}
	app.start()
}
