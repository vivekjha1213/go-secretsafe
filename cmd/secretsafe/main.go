package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vivekjha1213/go-secretsafe/pkg/secretsafe"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: secretsafe [set|get|delete] [namespace] [key] [value]")
		return
	}

	storage, err := secretsafe.NewStorage("")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	cache := secretsafe.NewCache()
	manager := secretsafe.NewSecretManager(storage, cache)

	switch args[0] {
	case "set":
		if len(args) != 4 {
			fmt.Println("Usage: secretsafe set [namespace] [key] [value]")
			return
		}
		err := manager.SetSecret(args[1], args[2], args[3])
		if err != nil {
			log.Fatalf("Failed to set secret: %v", err)
		}
		fmt.Println("Secret set successfully.")
	case "get":
		if len(args) != 3 {
			fmt.Println("Usage: secretsafe get [namespace] [key]")
			return
		}
		secret, err := manager.GetSecret(args[1], args[2])
		if err != nil {
			log.Fatalf("Failed to get secret: %v", err)
		}
		fmt.Println("Secret:", secret)
	case "delete":
		if len(args) != 3 {
			fmt.Println("Usage: secretsafe delete [namespace] [key]")
			return
		}
		err := manager.DeleteSecret(args[1], args[2])
		if err != nil {
			log.Fatalf("Failed to delete secret: %v", err)
		}
		fmt.Println("Secret deleted successfully.")
	default:
		fmt.Println("Unknown command:", args[0])
	}
}
