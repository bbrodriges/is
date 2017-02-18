package is

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// InRange returns true if value lies between left and right border
func InRange(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
// https://en.wikipedia.org/wiki/Email_address#Valid_email_addresses
func Email(s string) bool {
	if !strings.Contains(s, "@") || s[0] == '@' || s[len(s)-1] == '@' {
		return false
	}
	return true
}

// URL check if the string is an URL.
func URL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || str[0] == '.' {
		return false
	}

	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	if strings.HasPrefix(u.Host, ".") {
		return false
	}

	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return rxURL.MatchString(str)
}

// RequestURL check if the string rawurl, assuming
// it was received in an HTTP request, is a valid
// URL confirm to RFC 3986
func RequestURL(rawurl string) bool {
	url, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}
	if len(url.Scheme) == 0 {
		return false
	}
	return true
}

// RequestURI check if the string rawurl, assuming
// it was received in an HTTP request, is an
// absolute URI or an absolute path.
func RequestURI(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	return err == nil
}

// Alpha check if the string contains only letters (a-zA-Z).
func Alpha(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') {
			return false
		}
	}
	return true
}

// UTFLetter check if the string contains only unicode letter characters.
// Similar to is.Alpha but for all languages.
func UTFLetter(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if !unicode.IsLetter(v) {
			return false
		}
	}

	return true

}

// Alphanumeric check if the string contains only letters and numbers.
func Alphanumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') {
			return false
		}
	}

	return true
}

// UTFLetterNumeric check if the string contains only unicode letters and numbers.
func UTFLetterNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) {
			return false
		}
	}

	return true
}

// Numeric check if the string contains only numbers.
func Numeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if '9' < v || v < '0' {
			return false
		}
	}

	return true
}

// UTFNumeric check if the string contains only unicode numbers of any kind.
// Numbers can be 0-9 but also Fractions ¾,Roman Ⅸ and Hangzhou 〩.
func UTFNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if !unicode.IsNumber(v) {
			return false
		}
	}

	return true
}

// Whole returns true if value is whole number
func Whole(value float64) bool {
	return math.Remainder(value, 1) == 0
}

// Natural returns true if value is natural number (positive and whole)
func Natural(value float64) bool {
	return value > 0 && Whole(value)
}

// UTFDigit check if the string contains only unicode radix-10 decimal digits.
func UTFDigit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}

	return true
}

// Hexadecimal check if the string is a hexadecimal number.
func Hexadecimal(s string) bool {
	_, err := strconv.ParseInt(s, 16, 0)
	return err == nil
}

// Hexcolor check if the string is a hexadecimal color.
func Hexcolor(s string) bool {
	if s == "" {
		return false
	}

	if s[0] == '#' {
		s = s[1:]
	}

	if len(s) != 3 && len(s) != 6 {
		return false
	}

	for _, c := range s {
		if ('F' < c || c < 'A') && ('f' < c || c < 'a') && ('9' < c || c < '0') {
			return false
		}
	}

	return true
}

// RGBcolor check if the string is a valid RGB color in form rgb(RRR, GGG, BBB).
func RGBcolor(s string) bool {
	if s == "" || len(s) < 10 {
		return false
	}

	if s[0:4] != "rgb(" || s[len(s)-1] != ')' {
		return false
	}

	s = s[4 : len(s)-1]
	s = strings.TrimSpace(s)

	for _, p := range strings.Split(s, ",") {
		if len(p) > 1 && p[0] == '0' {
			return false
		}

		p = strings.TrimSpace(p)
		if i, e := strconv.Atoi(p); (255 < i || i < 0) || e != nil {
			return false
		}
	}

	return true
}

// LowerCase check if the string is lowercase.
func LowerCase(s string) bool {
	if len(s) == 0 {
		return false
	}

	return s == strings.ToLower(s)
}

// UpperCase check if the string is uppercase.
func UpperCase(s string) bool {
	if len(s) == 0 {
		return false
	}

	return s == strings.ToUpper(s)
}

// Int check if the string is an integer.
func Int(s string) bool {
	if len(s) == 0 {
		return false
	}

	_, err := strconv.Atoi(s)
	return err == nil
}

// Float check if the string is a float.
func Float(s string) bool {
	_, err := strconv.ParseFloat(s, 0)
	return err == nil
}

// ByteLength check if the string's length (in bytes) falls in a range.
func ByteLength(s string, min, max int) bool {
	return len(s) >= min && len(s) <= max
}

// UUIDv3 check if the string is a UUID version 3.
func UUIDv3(s string) bool {
	return UUID(s) && s[14] == '3'
}

// UUIDv4 check if the string is a UUID version 4.
func UUIDv4(s string) bool {
	return UUID(s) &&
		s[14] == '4' &&
		(s[19] == '8' || s[19] == '9' || s[19] == 'a' || s[19] == 'b')
}

// UUIDv5 check if the string is a UUID version 5.
func UUIDv5(s string) bool {
	return UUID(s) &&
		s[14] == '5' &&
		(s[19] == '8' || s[19] == '9' || s[19] == 'a' || s[19] == 'b')
}

// UUID check if the string is a UUID (version 3, 4 or 5).
func UUID(s string) bool {
	if len(s) != 36 {
		return false
	}

	for i, c := range s {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if c != '-' {
				return false
			}
			continue
		}

		if ('f' < c || c < 'a') && ('9' < c || c < '0') {
			return false
		}
	}

	return true
}

// JSON check if the string is valid JSON (note: uses json.Unmarshal).
func JSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// Multibyte check if the string contains one or more multibyte chars.
func Multibyte(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if v >= utf8.RuneSelf {
			return true
		}
	}

	return false
}

// ASCII check if the string contains ASCII chars only.
func ASCII(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if v >= utf8.RuneSelf {
			return false
		}
	}

	return true
}

// PrintableASCII check if the string contains printable ASCII chars only.
func PrintableASCII(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if v < ' ' || v > '~' {
			return false
		}
	}

	return true
}

// FullWidth check if the string contains any full-width chars.
func FullWidth(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if ('z' < v || v < 'a') && (0x7e < v || v < 20) && (65439 < v || v < 65377) && (65500 < v || v < 65440) && (1048288 < v || v < 65512) {
			return true
		}
	}

	return false
}

// HalfWidth check if the string contains any half-width chars.
func HalfWidth(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 20 && v <= 0x7e) || (v >= 65377 && v <= 65439) || (v >= 65440 && v <= 65500) || (v >= 65512 && v <= 1048288) {
			return true
		}
	}

	return false
}

// VariableWidth check if the string contains a mixture of full and half-width chars.
func VariableWidth(s string) bool {
	if len(s) == 0 {
		return false
	}

	return HalfWidth(s) && FullWidth(s)
}

// Base64 check if a string is base64 encoded.
func Base64(s string) bool {
	if len(s) == 0 {
		return false
	}

	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

// FilePath check is a string is Win or Unix file path and returns it's type.
func FilePath(str string) (bool, int) {
	if rxWinPath.MatchString(str) {
		// check windows path limit see:
		// http://msdn.microsoft.com/en-us/library/aa365247(VS.85).aspx#maxpath
		if len(str[3:]) > 32767 {
			return false, Win
		}
		return true, Win
	} else if rxUnixPath.MatchString(str) {
		return true, Unix
	}
	return false, Unknown
}

// DataURI checks if a string is base64 encoded data URI such as an image
func DataURI(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[:5] != "data:" {
		return false
	}

	var ci int
	for i := range s {
		if s[i] == ',' {
			ci = i
			break
		}
	}

	if s[ci-7:ci] != ";base64" {
		return false
	}

	return Base64(s[ci+1:])
}

// ISO3166Alpha2 checks if a string is valid two-letter country code
func ISO3166Alpha2(str string) bool {
	for _, entry := range ISO3166List {
		if str == entry.Alpha2Code {
			return true
		}
	}
	return false
}

// ISO3166Alpha3 checks if a string is valid three-letter country code
func ISO3166Alpha3(str string) bool {
	for _, entry := range ISO3166List {
		if str == entry.Alpha3Code {
			return true
		}
	}
	return false
}

// DNSName will validate the given string as a DNS name
func DNSName(str string) bool {
	if str == "" || len(strings.Replace(str, ".", "", -1)) > 255 {
		// constraints already violated
		return false
	}
	return rxDNSName.MatchString(str)
}

// DialString validates the given string for usage with the various Dial() functions
func DialString(str string) bool {
	if h, p, err := net.SplitHostPort(str); err == nil && h != "" && p != "" && (DNSName(h) || IP(h)) && Port(p) {
		return true
	}

	return false
}

// IP checks if a string is either IP version 4 or 6.
func IP(str string) bool {
	return net.ParseIP(str) != nil
}

// Port checks if a string represents a valid port
func Port(str string) bool {
	if i, err := strconv.Atoi(str); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

// IPv4 check if the string is an IP version 4.
func IPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}

// IPv6 check if the string is an IP version 6.
func IPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

// MAC check if a string is valid MAC address.
// Possible MAC formats:
// 01:23:45:67:89:ab
// 01:23:45:67:89:ab:cd:ef
// 01-23-45-67-89-ab
// 01-23-45-67-89-ab-cd-ef
// 0123.4567.89ab
// 0123.4567.89ab.cdef
func MAC(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}

// MongoID check if the string is a valid hex-encoded representation of a MongoDB ObjectId.
func MongoID(str string) bool {
	if str == "" || len(str) != 24 {
		return false
	}

	for _, c := range str {
		if ('F' < c || c < 'A') && ('f' < c || c < 'a') && ('9' < c || c < '0') {
			return false
		}
	}

	return true
}

// Latitude check if a string is valid latitude.
func Latitude(str string) bool {
	if str == "" {
		return false
	}

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}

	if 90.0 < f || f < -90.0 {
		return false
	}

	return true
}

// Longitude check if a string is valid longitude.
func Longitude(str string) bool {
	if str == "" {
		return false
	}

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}

	if 180.0 < f || f < -180.0 {
		return false
	}

	return true
}

// SSN will validate the given string as a U.S. Social Security Number
// See: http://stackoverflow.com/a/1517044
func SSN(s string) bool {
	if len(s) != 9 && len(s) != 11 {
		return false
	}

	s = stripNonNumeric(s)

	if len(s) != 9 {
		return false
	}

	if s[:3] == "000" || s[:3] == "666" || s[3:5] == "00" || s[5:] == "0000" {
		return false
	}

	p, err := strconv.ParseInt(s[:3], 10, 0)
	if err != nil {
		return false
	}

	if 900 <= p && p <= 999 {
		return false
	}

	return true
}

// Semver check if string is valid semantic version
func Semver(str string) bool {
	return rxSemver.MatchString(str)
}

// StringLength check string's length (including multi byte strings)
func StringLength(str string, min int, max int) bool {
	slen := utf8.RuneCountInString(str)
	return slen >= min && slen <= max
}

// Exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
