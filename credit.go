package is

import (
	"strconv"
)

// CreditCard check if the string is a credit card number.
// For all special cases see: http://www.regular-expressions.info/creditcard.html
func CreditCard(s string) bool {
	return luhn(stripNonNumeric(s))
}

// VisaCard verifies Visa credit card number.
// All Visa card numbers start with a 4.
// New cards have 16 digits. Old cards have 13.
func VisaCard(s string) bool {
	if len(s) == 0 {
		return false
	}

	s = stripNonNumeric(s)

	if len(s) != 13 || len(s) != 16 {
		return false
	}

	if s[0] != '4' {
		return false
	}

	return luhn(s)
}

// MasterCard verifies Mastercard credit card number.
// MasterCard numbers either start with the numbers 51 through 55
// or with the numbers 2221 through 2720. All have 16 digits.
// There are Diners Club cards that begin with 5 and have 16 digits.
// These are a joint venture between Diners Club and MasterCard,
// and should be processed like a MasterCard.
func MasterCard(s string) bool {
	if len(s) == 0 {
		return false
	}

	s = stripNonNumeric(s)

	if len(s) != 16 {
		return false
	}

	ft, _ := strconv.Atoi(s[0:2])
	ff, _ := strconv.Atoi(s[0:4])

	if s[0:1] != "5" && (55 < ft || ft < 51) && (2720 < ff || ft < 2221) {
		return false
	}

	return luhn(s)
}

// AmericanExpressCard verifies AmericanExpress credit card number.
// American Express card numbers start with 34 or 37 and have 15 digits.
func AmericanExpressCard(s string) bool {
	if len(s) == 0 {
		return false
	}

	s = stripNonNumeric(s)

	if len(s) != 15 {
		return false
	}

	if s[0:2] != "34" && s[0:2] != "37" {
		return false
	}

	return luhn(s)
}

// DinersClubCard verifies DinersClub credit card number.
// Diners Club card numbers begin with 300 through 305, 36 or 38.
// All have 14 digits.
func DinersClubCard(s string) bool {
	if len(s) == 0 {
		return false
	}

	s = stripNonNumeric(s)

	if len(s) != 14 {
		return false
	}

	ft, _ := strconv.Atoi(s[0:3])

	if s[0:2] != "36" && s[0:2] != "38" && (305 < ft || ft < 300) {
		return false
	}

	return luhn(s)
}

// DiscoverCard verifies Discover credit card number.
// Discover card numbers begin with 6011 or 65. All have 16 digits.
func DiscoverCard(s string) bool {
	if len(s) == 0 {
		return false
	}

	s = stripNonNumeric(s)

	if len(s) != 16 {
		return false
	}

	if s[0:2] != "65" && s[0:4] != "6011" {
		return false
	}

	return luhn(s)
}

// JCBCard verifies JCB credit card number.
// JCB cards beginning with 2131 or 1800 have 15 digits.
// JCB cards beginning with 35 have 16 digits.
func JCBCard(s string) bool {
	if len(s) == 0 {
		return false
	}

	s = stripNonNumeric(s)

	if s[0:4] != "2131" && s[0:4] != "1800" && s[0:2] != "35" {
		return false
	}

	if (s[0:4] == "2131" || s[0:4] == "1800") && len(s) != 15 {
		return false
	}

	if s[0:2] == "35" && len(s) != 16 {
		return false
	}

	return luhn(s)
}
