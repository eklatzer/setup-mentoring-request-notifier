package main

import (
	"log"

	"setup-mentoring-request-notifier/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
