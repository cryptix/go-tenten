package tenten

import "errors"

var (
	// ErrMalformedTT is returned if the 10:10 code is to short or missing spaces
	ErrMalformedTT = errors.New("10:10 code is to short or missing spaces")

	// ErrCorruptTT is returned if the 10:10 checksum is incorrect
	ErrCorruptTT = errors.New("10:10 checksum is incorrect")
)

var decodeAlpha = map[rune]int{'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'J': 8, 'K': 9, 'M': 10, 'N': 11, 'P': 12, 'Q': 13, 'R': 14, 'V': 15, 'W': 16, 'X': 17, 'Y': 18, '0': 19, '1': 20, '2': 21, '3': 22, '4': 23, '5': 24, '6': 25, '7': 26, '8': 27, '9': 28}

// Decode takes a 12 charachter string and converts them to latitude and longditude
func Decode(tt string) (lat, lon float64, err error) {
	if len(tt) != 12 || tt[3] != ' ' || tt[7] != ' ' {
		return 0, 0, ErrMalformedTT
	}

	tt = tt[:3] + tt[4:7] + tt[8:]

	var ttNum int
	for _, l := range tt {
		ttNum *= base
		if idx, ok := decodeAlpha[l]; ok {
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
		return 0, 0, ErrCorruptTT
	}

	return
}
