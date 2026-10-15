package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type singletonDataBase struct {
	capitals []string
}

func (sd *singletonDataBase) GetCaptitals() []string {
	return sd.capitals
}

var once sync.Once
var instance *singletonDataBase

func NewSingletonDatabase() *singletonDataBase {
	godotenv.Load("CAPITALS")
	capitals := strings.Split(os.Getenv("CAPITALS"), ", ")
	once.Do(func() {
		fmt.Println("Initialized...")
		db := singletonDataBase{capitals: capitals}
		instance = &db
	})
	fmt.Println(&instance)
	return instance
}

func main() {
	fmt.Println(NewSingletonDatabase(), NewSingletonDatabase())
}
