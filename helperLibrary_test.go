package helperLibrary

import (
	"testing"
)

var helper Helper

func TestKnownResult(t *testing.T) {
	result := helper.KnownResult()
	if result != "known" {
		t.Error("This is not the known result, test failed!")
	}
}

func TestNewFunction(t *testing.T) {
	result := helper.GetSomeCoolResult()
	if result != "expectation" {
		t.Error("This test has failed expectations!")
	}
}

func TestIgnoreTheseValues(t *testing.T) {
	var list []string
	var ignoreTheseItems []string

	list = append(list, "firstValue")
	list = append(list, "2ndValue")
	list = append(list, "anotherValue")
	list = append(list, "thisThing")

	ignoreTheseItems = append(ignoreTheseItems, "2ndValue")
	ignoreTheseItems = append(ignoreTheseItems, "thisThing")

	result := helper.FilterIgnoreList(list, ignoreTheseItems)

	if len(result) != 2 {
		t.Error("The list has an incorrect number of items.")
	}
}

func TestBuildArrayFromCsv(t *testing.T) {
	var theCommaSeparatedValues = "value1 ,2, 3,7,6, anotherValue"

	result := helper.BuildArrayFromCsv(theCommaSeparatedValues)

	if result[0] != "value1" {
		t.Error("The first value is not correct.")
	}
	if result[1] != "2" {
		t.Error("The second value is not correct.")
	}
	if result[2] != "3" {
		t.Error("The third value is not correct.")
	}
	if result[3] != "7" {
		t.Error("The fourth value is not correct.")
	}
	if result[4] != "6" {
		t.Error("The fifth value is not correct.")
	}
	if result[5] != "anotherValue" {
		t.Error("The sixth value is not correct.")
	}

	if len(result) != 6 {
		t.Error("The returned array does not have the appropriate number of values.")
	}

}

func TestRemoveListOneFromListTwo(t *testing.T){
	listOne := []string{"table1","table2","table3","table4","table5","table6"}
	listTwo := []string{"table3","table6"}
	expectedList := []string{"table1","table2","table4","table5"}

	result := helper.RemoveListOneFromListTwo(listOne, listTwo)

	if len(expectedList) != len(result) {
		t.Errorf("The expected item count is %v but the returned list item count is %v.", len(expectedList), len(result))
	}
}

func TestRemovePrefixItems(t *testing.T) {
	list := []string{"theRemainingItem","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-PostfixStuff","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-ButNoPostCar"}
	prefixToRemove := "PrefaceText"
	expectedResult := []string {"theRemainingItem"}

	result := helper.RemoveItemsWithPrefixFromList(list, prefixToRemove)

	if len(expectedResult) != len(result) {
		t.Errorf("The expected item count is %v but the returned list item count is %v.", len(expectedResult), len(result))
	}
	if expectedResult[0] != result[0] {
		t.Errorf("The expected value of %v returned as value %v.", expectedResult[0], result[0])
	}
}

func TestRemovePostfixItems(t *testing.T) {
	list := []string{"theRemainingItem","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-PostfixStuff","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-ButNoPostCar"}
	postfixToRemove := "PostfixStuff"
	expectedResult := []string {"theRemainingItem","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-ButNoPostCar"}

	result := helper.RemoveItemsWithPostfixFromList(list, postfixToRemove)

	if len(expectedResult) != len(result) {
		t.Errorf("The expected item count is %v but the returned list item count is %v.", len(expectedResult), len(result))
	}
	for k, _ := range result {
		if expectedResult[k] != result[k] {
			t.Errorf("The expected value of %v returned as value %v.", expectedResult[k], result[k])
		}
	}
}

func TestRemoveItemsWithPrefixAndSuffix(t *testing.T) {
	list := []string{"theRemainingItem","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-PostfixStuff","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-ButNoPostCar"}
	expectedResult := []string {"theRemainingItem","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-ButNoPostCar"}

	prefix := "PrefaceText"
	postfix := "PostfixStuff"

	result := helper.RemovePrefixAndOrSuffixItems(list, prefix, postfix, true)

	if result[0] != expectedResult[0] {
		t.Errorf("The expected result %v is not the result received %v.", expectedResult[0], result[0])
	}
}

func TestRemoveItemsWithPrefix(t *testing.T){
	list := []string{"theRemainingItem","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-PostfixStuff","PrefaceText-8dd410fa-fb47-4f08-bbdf-6829f83935fe-ButNoPostCar"}
	expectedResult := []string {"theRemainingItem"}

	prefix := "PrefaceText"

	result := helper.RemovePrefixItems(list, prefix)

	if result[0] != expectedResult[0] {
		t.Errorf("The expected result %v is not the result received %v.", expectedResult[0], result[0])
	}
}

