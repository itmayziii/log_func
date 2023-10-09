/*
Package main is meant for local testing purposes only, it sets up a minimalistic standalone server to send CloudEvents
to.
*/
package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"log"
	"os"
)

import (
	_ "github.com/itmayziii/log_func"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
