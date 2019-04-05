package main

import (
	"testing"
	"time"
)

func TestParseArgs(t *testing.T) {
	var table = []struct {
		in  []string
		out string
	}{
		{[]string{"june", "2002"}, "1 june 2002"},
		{[]string{"a", "b"}, "1 a b"},
	}

	for _, row := range table {
		str, err := ParseArgs(row.in)
		if err != nil {
			t.Fatalf("Expected nil error but instead got %s", err)
		}
		if str != row.out {
			t.Fatalf("Expected [%s] but got [%s] instead", row.out, str)
		}
	}

	_, err := ParseArgs([]string{})
	if err == nil {
		t.Fatal("Expected an error but didn't get one")
	}
}

func TestGetDaysOfMonth(t *testing.T) {
	var localLoc = time.Local
	var table = []struct {
		in  time.Time
		out int
	}{
		{time.Date(2009, 1, 1, 0, 0, 0, 0, localLoc), 31},
		{time.Date(2008, 2, 1, 0, 0, 0, 0, localLoc), 29},
		{time.Date(2019, 2, 1, 0, 0, 0, 0, localLoc), 28},
	}

	for _, row := range table {
		if days := getDaysOfMonth(row.in); days != row.out {
			t.Fatalf("Expected [%d] but got [%d]", row.out, days)
		}
	}
}

func TestOutputDates(t *testing.T) {
	dates, err := outputDates("1 jan 2019")

	if err != nil {
		t.Fatal("Got a non nil error")
	}

	var expected = []string{
		"2019-01-01 (Tue)",
		"2019-01-02 (Wed)",
		"2019-01-03 (Thu)",
		"2019-01-04 (Fri)",
		"2019-01-05 (Sat)",
		"2019-01-06 (Sun)",
		"2019-01-07 (Mon)",
		"2019-01-08 (Tue)",
		"2019-01-09 (Wed)",
		"2019-01-10 (Thu)",
		"2019-01-11 (Fri)",
		"2019-01-12 (Sat)",
		"2019-01-13 (Sun)",
		"2019-01-14 (Mon)",
		"2019-01-15 (Tue)",
		"2019-01-16 (Wed)",
		"2019-01-17 (Thu)",
		"2019-01-18 (Fri)",
		"2019-01-19 (Sat)",
		"2019-01-20 (Sun)",
		"2019-01-21 (Mon)",
		"2019-01-22 (Tue)",
		"2019-01-23 (Wed)",
		"2019-01-24 (Thu)",
		"2019-01-25 (Fri)",
		"2019-01-26 (Sat)",
		"2019-01-27 (Sun)",
		"2019-01-28 (Mon)",
		"2019-01-29 (Tue)",
		"2019-01-30 (Wed)",
		"2019-01-31 (Thu)",
	}

	for index, row := range dates {
		if row != expected[index] {
			t.Fatalf("Expected [%s] for day [%d] but got [%s] instead", expected[index], index, row)
		}
	}
}
