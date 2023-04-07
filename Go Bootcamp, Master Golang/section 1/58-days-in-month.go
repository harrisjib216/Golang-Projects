package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Refactor the previous exercise from the if statement
//  section.
//
//  "Print the number of days in a given month."
//
//  Use a switch statement instead of if statements.
// ---------------------------------------------------------

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Give me a month name")
	// 	return
	// }

	// year := time.Now().Year()
	// leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	// days, month := 28, os.Args[1]

	// if m := strings.ToLower(month); m == "april" ||
	// 	m == "june" ||
	// 	m == "september" ||
	// 	m == "november" {
	// 	days = 30
	// } else if m == "january" ||
	// 	m == "march" ||
	// 	m == "may" ||
	// 	m == "july" ||
	// 	m == "august" ||
	// 	m == "october" ||
	// 	m == "december" {
	// 	days = 31
	// } else if m == "february" {
	// 	if leap {
	// 		days = 29
	// 	}
	// } else {
	// 	fmt.Printf("%q is not a month.\n", month)
	// 	return
	// }

	// fmt.Printf("%q has %d days.\n", month, days)

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	days, month := 28, os.Args[1]

	switch strings.ToLower(os.Args[1]) {
	case "april", "june", "september", "november":
		days = 30
	case "january", "march", "may", "july", "august", "october", "december":
		days = 31
	case "february":
		if time.Now().Year()%4 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		fmt.Printf("%q is not a month.\n", os.Args[1])
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
