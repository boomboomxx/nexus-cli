package main

import (
	"log"
	"testing"
)

var tags = []string{
	"1.202421.0.4-prod-403",
	"1.202421.0.4-prod-133",
	"1.202424.1.2",
	"1.202422.1.1",
	"1.202422.1.4-prod-123",
	"1.202422.1.4-prod-321",
	"1.202422.3.1",
	"1.202423.1.5-prod",
	"latest",
}

func Test_SortMixed(t *testing.T) {

	compareStringNumber := func(str1, str2 string) bool {
		return extractNumberFromString(str1) > extractNumberFromString(str2)
	}
	Compare(compareStringNumber).Sort(tags)
	log.Print("tgs", tags)
	if tags[0] != "1.0.1" && tags[1] != "latest" {
		t.Errorf("ordering incorrect when checking mixed tags")
	}
}

func Test_SortAllDigits(t *testing.T) {
	tags := []string{"1.2.1", "1.0.1"}

	compareStringNumber := func(str1, str2 string) bool {
		return extractNumberFromString(str1) < extractNumberFromString(str2)
	}
	Compare(compareStringNumber).Sort(tags)

	if tags[0] != "1.0.1" && tags[1] != "1.2.1" {
		t.Errorf("ordering incorrect in all digits tags")
	}
}

func Test_ToBeDelete(t *testing.T) {
	tobeDelete := extractToDelete(tags, 1)
	log.Print(tobeDelete)
}
