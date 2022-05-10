package main

import (
	"testing"
)

func TestSimpleFact(t *testing.T) {
	var caseList = [5]int{0, 1, 2, 31, 50}
	var resultList = [5]string{
		"1",
		"1",
		"2",
		"8222838654177922817725562880000000",
		"30414093201713378043612608166064768844377641568960512000000000000"}

	for i := 0; i < 5; i++ {
		if SimpleFact(caseList[i]).String() != resultList[i] {
			t.Error("Expected ", resultList[i])
		}
	}
}

func TestTreeFact(t *testing.T) {
	var caseList = [5]int{0, 1, 2, 31, 50}
	var resultList = [5]string{
		"1",
		"1",
		"2",
		"8222838654177922817725562880000000",
		"30414093201713378043612608166064768844377641568960512000000000000"}

	for i := 0; i < 5; i++ {
		if TreeFact(caseList[i]).String() != resultList[i] {
			t.Error("Expected ", resultList[i])
		}
	}
}

func TestStieveFact(t *testing.T) {
	var caseList = [5]int{0, 1, 2, 31, 50}
	var resultList = [5]string{
		"1",
		"1",
		"2",
		"8222838654177922817725562880000000",
		"30414093201713378043612608166064768844377641568960512000000000000"}

	for i := 0; i < 5; i++ {
		if StieveFact(caseList[i]).String() != resultList[i] {
			t.Error("Expected ", resultList[i])
		}
	}
}
