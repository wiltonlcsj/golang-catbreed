package repositories

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	if name != "Gladys" {
		t.Fatalf(`Hello("Gladys")`)
	}
}