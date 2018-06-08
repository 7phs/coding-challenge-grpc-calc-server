package main

import (
		"github.com/7phs/coding-challenge-grpc-calc-server/config"
	"os"
	"fmt"
	"github.com/7phs/coding-challenge-grpc-calc-server/calc"
)

// Dump a help message.
func help() {
	fmt.Println("Usage: client CMD ARG [ARG ARG ... ARG]")
	fmt.Println("      CMD - a name of arithmetical operation: add, sub, mul, div, fib")
	fmt.Println("      ARG - an integer numeric applied for an arithmetical operation")
	fmt.Println("")
	fmt.Println("      Default service port:", config.DEFAULT_PORT,". Set the new value of an environment variable", config.CONFIG_PORT,", to change it.")
}

func main() {
	if len(os.Args)<=1 {
		help()
		return
	}

	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("failed to get configuration parameters: ", err)
		return
	}

	command, err := calc.NewCommand(os.Args[1:])
	if err!=nil {
		fmt.Println("failed to create a command: ", err)
		return
	}

	result, err := NewClient(conf).Execute(command)
	if err!=nil {
		fmt.Println("failed to execute a command: ", err)
		return
	}

	fmt.Println(result)
}