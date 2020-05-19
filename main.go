package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	gocron.Every(1).Minute().Do(taskWithParams, 1, "hello")
	<- gocron.Start()
}