package tenten

import (
	"math"
)

const (
	alphabet = "ABCDEFGHJKMNPQRVWXY0123456789"
	base     = 29
	n        = 12
)

// Encode converts a (lat,lon) tuple to a human readable string
//
// optimizations:
//	- using a fixed size array
//	- replace math.Mod() with integer modulo
//	- replaceing math.Floor with  0|(foo), to be pkg math-free (found here http://www.tapper-ware.net/blog/?p=112)
func Encode(lat, lon float64) string {
	var tt [n]byte

	lat += 90.0
	lon += 180.0

	lat *= 10000.0
	lon *= 10000.0

	lat = float64(0 | int(lat))
	lon = float64(0 | int(lon))

	p := lat*3600000.0 + lon
	ttNum := p * base
	var c int

	for i := 1; i < 10; i++ {
		c += (int(p) % base) * i
		p = float64(0 | int(p/base))
	}

	c %= 29
	ttNum += float64(c)
	ttNum = float64(0 | int(ttNum))

	pos := n - 1
	for i := 0; i < 10; i++ {
		d := int(ttNum) % base

		if (i == 4) || (i == 7) {
			tt[pos] = ' '
			pos--
		}

		ttNum = float64(0 | int(ttNum/base))
		tt[pos] = byte(alphabet[int(d)])
		pos--
	}

	return string(tt[:])
}

// encodeSlow is just there for comparison, it's what i started with
// basicly a copy&paste of the js code and than replacing Math.floor with math.Floor
//
// what i learned:
// 	- math.Mod() added per call ~1300ns/op
// 	- string concat is massivly costly
func encodeSlow(lat, lon float64) (tt string) {
	lat += 90.0
	lon += 180.0

	lat *= 10000.0
	lon *= 10000.0

	lat = math.Floor(lat)
	lon = math.Floor(lon)

	p := lat*3600000.0 + lon
	ttNum := p * base
	var c int

	for i := 1; i < 10; i++ {
		c += int(math.Mod(p, base)) * i
		p = math.Floor(p / base)
	}

	c %= 29
	ttNum += float64(c)
	ttNum = math.Floor(ttNum)

	for i := 0; i < 10; i++ {
		d := math.Mod(ttNum, base)

		if (i == 4) || (i == 7) {
			tt = " " + tt

		}

		ttNum = math.Floor(ttNum / base)
		tt = string(alphabet[int(d)]) + tt
	}

	return tt
}
