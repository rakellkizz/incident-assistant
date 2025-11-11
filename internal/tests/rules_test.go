package tests

import (
	"testing"
	"incident-assistant/internal"
)

func TestPriorityMatrix(t *testing.T) {
	p, sla := internal.PriorityAndSLA("critical", "high")
	if p != "P1" || sla != 120 {
		t.Fatalf("esperava P1/120, obtive %s/%d", p, sla)
	}
}
