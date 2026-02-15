package slopeone

import (
	"math"
	"testing"
)

var dataset = []map[string]float64{
	{"squid": 1, "cuttlefish": 0.5, "octopus": 0.2},
	{"squid": 1, "octopus": 0.5, "nautilus": 0.2},
	{"squid": 0.2, "cuttlefish": 0.4, "octopus": 1, "nautilus": 0.4},
	{"cuttlefish": 0.9, "octopus": 0.4, "nautilus": 0.5},
}

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.01
}

func TestPredict(t *testing.T) {
	algorithm := New()
	algorithm.Update(dataset)

	results := algorithm.Predict(map[string]float64{"squid": 0.4})

	expected := map[string]float64{
		"cuttlefish": 0.25,
		"octopus":    0.23,
		"nautilus":   0.1,
	}

	for item, expectedValue := range expected {
		if actual, exists := results[item]; !exists {
			t.Errorf("Missing prediction for %s", item)
		} else if !floatEquals(expectedValue, actual) {
			t.Errorf("Expected %s=%.2f, got %.2f", item, expectedValue, actual)
		}
	}
}

func TestPredictForUser1(t *testing.T) {
	algorithm := New()
	algorithm.Update(dataset)

	u1 := map[string]float64{
		"squid":      1,
		"cuttlefish": 0.5,
		"octopus":    0.2,
	}

	results := algorithm.Predict(u1)

	if actual, exists := results["nautilus"]; !exists {
		t.Errorf("Missing prediction for nautilus")
	} else if !floatEquals(0.26, actual) {
		t.Errorf("Expected nautilus=0.26, got %.2f", actual)
	}
}

func TestPredictForUser2(t *testing.T) {
	algorithm := New()
	algorithm.Update(dataset)

	u2 := map[string]float64{
		"squid":    1,
		"octopus":  0.5,
		"nautilus": 0.2,
	}

	results := algorithm.Predict(u2)

	if actual, exists := results["cuttlefish"]; !exists {
		t.Errorf("Missing prediction for cuttlefish")
	} else if !floatEquals(0.60, actual) {
		t.Errorf("Expected cuttlefish=0.60, got %.2f", actual)
	}
}

func TestPredictForUser3(t *testing.T) {
	algorithm := New()
	algorithm.Update(dataset)

	u3 := map[string]float64{
		"cuttlefish": 0.9,
		"octopus":    0.4,
		"nautilus":   0.5,
	}

	results := algorithm.Predict(u3)

	if actual, exists := results["squid"]; !exists {
		t.Errorf("Missing prediction for squid")
	} else if !floatEquals(0.77, actual) {
		t.Errorf("Expected squid=0.77, got %.2f", actual)
	}
}
