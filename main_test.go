package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	empty := []string{}
	cases := []struct {
		Line      string
		Condition string
		Commands  []string
	}{
		{"//break", "", empty},
		{"//break if x == 1", "x == 1", empty},
		{"//break : print \"hi\"", "", []string{"print \"hi\""}},
		{"//break if y != x : print x; continue", "y != x", []string{"print x", "continue"}},
		{"xxx //break", "", empty},
		{"xxx //break if x == 1", "x == 1", empty},
		{"xxx //break : print \"hi\"", "", []string{"print \"hi\""}},
		{"xxx //break if y != x : print x; continue", "y != x", []string{"print x", "continue"}},
	}

	for _, c := range cases {
		bp := ParseLine(c.Line)
		if bp.Condition != c.Condition {
			t.Fatalf("condition: %s != %s", bp.Condition, c.Condition)
		}
		if len(bp.Commands) != len(c.Commands) {
			t.Fatalf("incorrect number of commands: %d != %d", len(bp.Commands), len(c.Commands))
		}
		for i, cmd := range bp.Commands {
			if c.Commands[i] != cmd {
				t.Fatalf("command: %s != %s", c.Commands[i], cmd)
			}
		}
	}
}
