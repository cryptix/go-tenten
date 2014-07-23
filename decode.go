package tenten

import (
	// "fmt"
	"errors"
	"strings"
)

// ErrIncorrectTT is returned if an incorrect 10:10 code is incorrect
var ErrIncorrectTT = errors.New("Incorrect TT")

// Decode takes a 12 charachter string and converts them to latitude and longditude
func Decode(tt string) (lat, lon float64, err error) {
	if len(tt) != 12 {
		return 0, 0, ErrIncorrectTT
	}

	tt = strings.Replace(tt, " ", "", 2)

	var ttNum int
	for _, l := range tt {
		ttNum *= base
		if idx := strings.Index(alphabet, string(l)); idx > 0 {
			ttNum += idx
		}
	}

	var (
		p     = ttNum / base
		check = ttNum % base
	)

	lon = float64(p % 3600000)
	lat = float64(p / 3600000)

	lat /= 10000
	lon /= 10000

	lat -= 90
	lon -= 180

	var c int

	for i := 1; i < 10; i++ {
		c += (p % base) * i
		p = int(p / base)
	}

	c %= base

	if c != check {
		return 0, 0, ErrIncorrectTT
	}

	return
}
