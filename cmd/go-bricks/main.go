package main

import "log"

const header = `
____        _      _        
|  _ \     (_)    | |       
| |_) |_ __ _  ___| | _____ 
|  _ <| '__| |/ __| |/ / __|
| |_) | |  | | (__|   <\__ \
|____/|_|  |_|\___|_|\_\___/
 `

func main() {
	app,err := newApp()
	if err != nil {
		log.Fatalf("Error starting app: %v", err)
	}
	app.start()
}
