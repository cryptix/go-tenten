package tenten

import (
	"testing"
)

type testCase struct {
	lat, lon float64
	want     string
}

func TestEncode(t *testing.T) {
	table := []testCase{
		testCase{10.0, 10.0, "HF1 GNP 6228"},
		testCase{53.563823, 10.005327, "MKX 4C2 E4WH"},
		testCase{51.09559, 1.12207, "MEQ N6G 7NY5"},
	}

	for i, tc := range table {
		if got := Encode(tc.lat, tc.lon); got != tc.want {
			t.Fatalf("%d failed: got[%s] wanted[%s]", i, got, tc.want)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(53.563823, 10.005327)
	}
}

func BenchmarkEncodeSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeSlow(53.563823, 10.005327)
	}
}
