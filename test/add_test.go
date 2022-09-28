package test

import(
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	r := Add(1, 2)
	if r !=3 {

		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}

}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}
