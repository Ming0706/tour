package test

import (
	"flag"
	"log"
	"testing"
	//"github.com/Ming0706/tour"
)

func TestTourMain(t *testing.T) {
	var name string
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}

/*func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}
*/
