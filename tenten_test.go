package tenten

import (
	"math"
	"testing"
)

type testCase struct {
	lat, lon float64
	want     string
}

var table = []testCase{
	testCase{10.0, 10.0, "HF1 GNP 6228"},
	testCase{53.5638, 10.0053, "MKX 4C2 E4WH"},
	testCase{51.09559, 1.12207, "MEQ N6G 7NY5"},
}

func TestDecode(t *testing.T) {
	for i, tc := range table {
		lat, lon, err := Decode(tc.want)
		if err != nil {
			t.Fatalf("%d - got an Decode error", i)
		}

		if math.Abs(lat-tc.lat) > 0.001 {
			t.Fatalf("%d latitude test failed: got[%f] wanted[%f]", i, lat, tc.lat)
		}

		if math.Abs(lon-tc.lon) > 0.001 {
			t.Fatalf("%d londitude test failed: got[%f] wanted[%f]", i, lon, tc.lon)
		}
	}

	_, _, err := Decode("")
	if err != ErrMalformedTT {
		t.Fatal("Should return ErrIncorrectTT")
	}

	_, _, err = Decode("000000000000")
	if err != ErrMalformedTT {
		t.Fatal("Should return ErrIncorrectTT")
	}

	_, _, err = Decode("HF1 GNP 6229")
	if err != ErrCorruptTT {
		t.Fatal("Should return ErrIncorrectTT")
	}

}

func TestEncode(t *testing.T) {
	for i, tc := range table {
		if got := Encode(tc.lat, tc.lon); got != tc.want {
			t.Fatalf("%d failed: got[%s] wanted[%s]", i, got, tc.want)
		}
	}
}

func TestEncodeSlow(t *testing.T) {
	for i, tc := range table {
		if got := encodeSlow(tc.lat, tc.lon); got != tc.want {
			t.Fatalf("%d failed: got[%s] wanted[%s]", i, got, tc.want)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("MKX 4C2 E4WH")
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(53.563823, 10.005327)
	}
}

func BenchmarkEncodeSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeSlow(53.563823, 10.005327)
	}
}
