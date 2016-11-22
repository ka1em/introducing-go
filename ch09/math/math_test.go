package math

import (
	"testing"
)

func TestAverage(t *testing.T) {
	v := Average([]float64{1, 1, 2, 2})
	if v != 1.5 {
		t.Error("expected 1.5, got", v)
	}
}
