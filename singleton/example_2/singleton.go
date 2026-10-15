package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type config struct {
	// COnfig variable
}

var counter int = 1
var singleConfigInstance *config

func getConfigInstance() *config {
	if singleConfigInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if singleConfigInstance == nil {
			fmt.Println("Creating single instance now, and counter is ", counter)
			singleConfigInstance = &config{}
			counter = counter + 1
		} else {
			fmt.Println("Single instance already created-1, returning that one")
		}
	} else {
		fmt.Println("Single instance already created-2, returning that one")
	}
	return singleConfigInstance
}

func main() {
	for range 100 {
		go getConfigInstance()
	}
	fmt.Scanln()
}
