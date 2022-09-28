package test

import (
	"fmt"
	"github.com/Ming0706/tour"
	"testing"
)

func TestTourMain(t *testing.T) {
	main.Main()
}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}

