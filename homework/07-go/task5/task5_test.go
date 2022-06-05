package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleFact(t *testing.T) {
	var caseList = []int{0, 1, 2, 31, 50}
	var expectList = []string{
		"1",
		"1",
		"2",
		"8222838654177922817725562880000000",
		"30414093201713378043612608166064768844377641568960512000000000000"}

	for i := 0; i < len(caseList); i++ {
		assert.Equal(t, expectList[i], SimpleFact(caseList[i]).String())
	}
}

func TestTreeFact(t *testing.T) {
	var caseList = []int{0, 1, 2, 31, 50}
	var expectList = []string{
		"1",
		"1",
		"2",
		"8222838654177922817725562880000000",
		"30414093201713378043612608166064768844377641568960512000000000000"}

	for i := 0; i < len(caseList); i++ {
		assert.Equal(t, expectList[i], TreeFact(caseList[i]).String())
	}
}

func TestStieveFact(t *testing.T) {
	var caseList = []int{0, 1, 2, 31, 50}
	var expectList = []string{
		"1",
		"1",
		"2",
		"8222838654177922817725562880000000",
		"30414093201713378043612608166064768844377641568960512000000000000"}

	for i := 0; i < len(caseList); i++ {
		assert.Equal(t, expectList[i], StieveFact(caseList[i]).String())
	}
}
