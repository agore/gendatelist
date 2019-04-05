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
