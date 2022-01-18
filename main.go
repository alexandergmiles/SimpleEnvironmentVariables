package main

import (
	"fmt"
	"os"
	"strconv"
)

type EnviromentVariables struct {
	Name string
	Port int
}

func main() {
	envVariables, err := getEnvironmentVariables()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Variables found:\nName - %s\nPor%d\n", envVariables.Name, envVariables.Port)
}

func getEnvironmentVariables() (*EnviromentVariables, error) {
	var envCarrier = new(EnviromentVariables)

	name := os.Getenv("TestName")
	port := os.Getenv("TestPort")

	portAsInt, err := strconv.Atoi(port)

	if err != nil {
		return nil, err
	}

	envCarrier.Name = name
	envCarrier.Port = portAsInt

	return envCarrier, nil
}
