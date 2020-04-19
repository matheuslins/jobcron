package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	s := gocron.NewScheduler()
	s.Every(10).Second().Do(taskWithParams, 1, "hello")
	<- s.Start()
}