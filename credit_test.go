package is

import "testing"

func TestCreditCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"5398228707871528", false},
		{"375556917985515", true},
		{"36050234196908", true},
		{"4716461583322103", true},
		{"4716-2210-5188-5662", true},
		{"4929 7226 5379 7141", true},
		{"5398228707871527", true},
	}
	for _, test := range tests {
		actual := CreditCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected CreditCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestVisaCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"5398228707871528", false},
		{"375556917985515", false},
		{"36050234196908", false},
		{"4716213139245217", true},
		{"4716-2210-5188-5662", true},
		{"4929 7226 5379 7141", true},
		{"5398228707871527", false},
	}
	for _, test := range tests {
		actual := VisaCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected VisaCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestMasterCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"5309309013152196", true},
		{"375556917985515", false},
		{"36050234196908", false},
		{"4716213139245217", false},
		{"4716-2210-5188-5662", false},
		{"4929 7226 5379 7141", false},
		{"5398228707871527", true},
	}
	for _, test := range tests {
		actual := MasterCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected MasterCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestAmericanExpressCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"5309309013152196", false},
		{"375556917985515", true},
		{"3491 0149 1820 987", true},
		{"4716213139245217", false},
		{"4716-2210-5188-5662", false},
		{"4929 7226 5379 7141", false},
		{"359822870787152", false},
	}
	for _, test := range tests {
		actual := AmericanExpressCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected AmericanExpressCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestDinersClubCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"5309309013152196", false},
		{"375556917985515", false},
		{"3491 0149 1820 987", false},
		{"30060129447551", true},
		{"4716-2210-5188-5662", false},
		{"3129 7226 5379 71", false},
		{"3032 5156 3490 24", true},
	}
	for _, test := range tests {
		actual := DinersClubCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected DinersClubCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestDiscoverCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"5309309013152196", false},
		{"375556917985515", false},
		{"6011748439365527", true},
		{"30060129447551", false},
		{"4716-2210-5188-5662", false},
		{"6011229282505485", true},
		{"3032 5156 3490 24", false},
	}
	for _, test := range tests {
		actual := DiscoverCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected DiscoverCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestJCBCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3533868143240232", true},
		{"375556917985515", false},
		{"6011748439365527", false},
		{"30060129447551", false},
		{"4716-2210-5188-5662", false},
		{"6011229282505485", false},
		{"180036877154241", true},
		{"2131424111351356", false},
		{"354515246782342", false},
	}
	for _, test := range tests {
		actual := JCBCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected JCBCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}
