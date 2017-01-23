package is

import "strconv"

func stripNonNumeric(s string) string {
	r := []byte(s)
	for i := len(r) - 1; i >= 0; i-- {
		if '9' < r[i] || r[i] < '0' {
			r = append(r[:i], r[i+1:]...)
		}
	}

	return string(r)
}

// See: https://en.wikipedia.org/wiki/Luhn_algorithm
func luhn(s string) bool {
	if len(s) == 0 {
		return false
	}

	var sum int
	var alter bool

	for i := len(s) - 1; i >= 0; i-- {
		mod, err := strconv.Atoi(s[i:(i + 1)])
		if err != nil {
			return false
		}

		if alter {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alter = !alter
		sum += mod
	}

	return sum%10 == 0
}
