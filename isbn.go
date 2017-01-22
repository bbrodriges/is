package is

import "bytes"

// ISBN check if the string is an ISBN (version 10 or 13).
// If version value is not equal to 10 or 13, it will be check both variants.
func ISBN(s string, version int) bool {
	if version == 10 {
		return ISBN10(s)
	}

	if version == 13 {
		return ISBN13(s)
	}

	return ISBN10(s) || ISBN13(s)
}

// ISBN10 check if the string is an ISBN version 10.
func ISBN10(s string) bool {
	if len(s) < 10 {
		return false
	}

	s = stripISBN(s, 10)

	if len(s) != 10 {
		return false
	}

	var sum int
	for i := 0; i < 9; i++ {
		sum += (i + 1) * int(s[i]-'0')
	}
	if s[9] == 'X' {
		sum += 100 // aka 10 * 10
	} else {
		sum += 10 * int(s[9]-'0')
	}

	return sum%11 == 0
}

// ISBN13 check if the string is an ISBN version 13.
func ISBN13(s string) bool {
	if len(s) < 13 {
		return false
	}

	s = stripISBN(s, 13)

	if len(s) != 13 {
		return false
	}

	var sum int
	for i := 0; i < 12; i++ {
		if i%2 == 0 {
			sum += int(s[i] - '0')
		} else {
			sum += int(s[i]-'0') * 3
		}
	}

	return int(s[12]-'0')-((10-(sum%10))%10) == 0
}

func stripISBN(s string, version int) string {
	b := bytes.NewBuffer(nil)
	for i, c := range s {
		if '0' <= c && c <= '9' {
			b.WriteByte(s[i])
		}
	}

	// special case for ISBN 10
	if version == 10 && (s[len(s)-1] == 'x' || s[len(s)-1] == 'X') {
		b.WriteString("X")
	}

	return b.String()
}
