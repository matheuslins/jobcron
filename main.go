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
	s.Every(5).Minutes().Do(taskWithParams, 1, "hello")
	<- s.Start()
}