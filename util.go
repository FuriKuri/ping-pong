package main

import (
	"os"
)

func getArgParameter(name string, defaultValue string) string {
	argsWithoutProg := os.Args[1:]
	for index, element := range argsWithoutProg {
		if element == "--"+name {
			return argsWithoutProg[index+1]
		}
	}
	return defaultValue
}
