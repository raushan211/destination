package main

import "fmt"

/*
input  = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
*/
func main() {
	input := [][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}
	dict := make(map[string]int)

	for _, innerObject := range input {

		for i, k := range innerObject {
			if i == 0 {
				dict[k] = 0
			} else {

				if _, ok := dict[k]; ok {
					delete(dict, k)
				} else {
					dict[k] = 1
				}
			}
		}

	}
	//fmt.Println(dict)
	result := ""

	for key, value := range dict {
		if value == 1 {
			result = key
			break
		}
	}

	fmt.Println(result)
}
