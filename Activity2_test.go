package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestMathFunc(t *testing.T) {
	g := Goblin(t)

	g.Describe("Activity 2", func() {
		g.Assert(modValue(10.0, 1.0)).Equal(0.0)
		g.Assert(modValue(10.0, 3.0)).Equal(1.0)
		g.Assert(modValue(10.0, 4.0)).Equal(2.0)
		g.Assert(modValue(10.0, 7.0)).Equal(3.0)
		g.Assert(modValue(10.0, 6.0)).Equal(4.0)
	})
}
