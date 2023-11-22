package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/prakharrai1609/naruto/cli/commands"
)

const (
	Reset = "\033[0m"
	Green = "\033[32m"
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		printAvailableCommands()
		break
	case "generate":
		generateCmd.Parse(os.Args[2:])
		commands.Generate()
		break
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(Green + "	naruto help     - Print available commands" + Reset)
	fmt.Println(Green + "	naruto generate - Generate something" + Reset)
}

func printAvailableCommands() {
	fmt.Println("Available commands:")
	fmt.Println("1. Generate")
}
