package utils


import (
	"testing"
)

func TestGetCurrentDT(t *testing.T) {
	dt := GetCurrentDT()
	t.Log(dt)
}
