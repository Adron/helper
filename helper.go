package helper

import "strings"

func KnownResult() string {
	return "known"
}

func GetSomeCoolResult() string {
	return "expectation"
}

func FilterIgnoreList(list []string, ignoreList []string) []string {
	var remainingItems []string

	for _, v := range list {
		valueExists := false
		for _, keyIgnoreValue := range ignoreList {
			if v == keyIgnoreValue {
				valueExists = true
			}
		}
		if !valueExists {
			remainingItems = append(remainingItems, v)
		}
	}

	return remainingItems
}

func BuildArrayFromCsv(values string) []string {
	var cleanUpResults []string
	splitValues := strings.Split(values, ",")
	for _, v := range splitValues {
		cleanUpResults = append(cleanUpResults, strings.TrimSpace(v))
	}
	return cleanUpResults
}

func RemoveListOneFromListTwo(listOne []string, listTwo []string) []string {
	var returnList []string
	for _, listOneItem := range listOne {
		exists := false
		for _, listTwoItem := range listTwo {
			if listOneItem == listTwoItem {
				exists = true
			}
		}
		if !exists {
			returnList = append(returnList, listOneItem)
		}
	}
	return returnList
}

func RemoveItemsWithPrefixFromList(list []string, itemPrefix string) []string {
	var returnList []string
	for _, v := range list {
		if !strings.HasPrefix(v, itemPrefix) {
			returnList = append(returnList, v)
		}
	}
	return returnList
}

func RemoveItemsWithPostfixFromList(list []string, itemSuffix string) []string {
	var returnList []string
	for _, v := range list {
		if !strings.HasSuffix(v, itemSuffix) {
			returnList = append(returnList, v)
		}
	}
	return returnList
}

func RemovePrefixItems(list []string, prefix string) []string {
	return RemovePrefixAndOrSuffixItems(list, prefix, "", false)
}

func RemovePrefixAndOrSuffixItems(list []string, prefix string, postfix string, both bool) []string {
	var listItem []string

	if both {
		for _, value := range list {
			if !strings.HasPrefix(value, prefix) && !strings.HasSuffix(value, postfix) {
				listItem = append(listItem, value)
			}
		}
	} else {
		for _, value := range list {
			if !strings.HasPrefix(value, prefix) {
				listItem = append(listItem, value)
			}
		}
	}

	return listItem
}
