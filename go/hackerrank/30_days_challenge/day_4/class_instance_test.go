package main

import "testing"

const errorMsg = "sent value %v, expected value %v, found %v"

var tests = []struct {
	person         person
	expectedAge    int
	expectedString string
}{
	{person{-1}, 0, "You are young."},
	{person{12}, 12, "You are young."},
	{person{17}, 17, "You are a teenager."},
	{person{34}, 34, "You are old."},
}

func TestNewPerson(t *testing.T) {
	t.Parallel()
	for _, test := range tests {
		result := test.person.NewPerson(test.person.age)

		if result.age != test.expectedAge {
			t.Errorf(errorMsg, test.person.age, test.expectedAge, result.age)
		}
	}
}

func TestAmIOld(t *testing.T) {
	t.Parallel()
	for _, test := range tests {
		result := test.person.amIOld()
		if result != test.expectedString {
			t.Errorf(errorMsg, test.person.age, test.expectedString, result)
		}
	}
}

func TestYearPasses(t *testing.T) {
	t.Parallel()
	for _, test := range tests {
		result := test.person.yearPasses()
		if result.age < test.person.age {
			t.Errorf(errorMsg, test.person.age, test.person.age+1, result.age)
		}
	}
}
