package main

import "fmt"

func main() {
	languages := make(map[string]string)
	slice := make([]int, 0)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	delete(languages, "RB")
	for key, value := range languages {
		fmt.Println(key, value)
	}
}
