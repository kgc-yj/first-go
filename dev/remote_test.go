package dev

import (
	"fmt"
	"testing"
	"time"
)

func TestRemote(t *testing.T) {
	fmt.Println("...")
}

type Person struct {
	Name string
	Id   int
	Age  int
}

func PrintNow() {
	fmt.Println(time.Now())
}
