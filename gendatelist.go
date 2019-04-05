package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/araddon/dateparse"
)

func main() {

	var args, err = ParseArgs(os.Args[1:])
	if err != nil {
		usage()
	}

	days, _ := outputDates(args)
	for _, row := range days {
		fmt.Println(row)
	}
}

func outputDates(args string) ([]string, error) {
	var out = []string{}
	t, _ := dateparse.ParseAny(args)
	days := getDaysOfMonth(t)
	for i := 1; i <= days; i++ {
		out = append(out, t.Format("2006-01-02 (Mon)"))
		t = t.Add(time.Hour * 24)
	}

	return out, nil
}

func getDaysOfMonth(t time.Time) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	days := months[t.Month()-1]
	if t.Year()%4 == 0 || t.Year()%100 == 0 || t.Year()%400 == 0 {
		days = days + 1
	}

	return days
}
func ParseArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("Insufficient arguments")
	}

	return fmt.Sprintf("1 %s %s", args[0], args[1]), nil
}

func usage() {
	fmt.Println("Usage: gendatelist month year")
	fmt.Println(" e.g.: gendatelist september 2018")
	fmt.Println(" e.g.: gendatelist jun 2019")
	fmt.Println(" etc.")
	os.Exit(1)
}
