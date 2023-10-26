package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("Not Found")

var food = map[int]string{
	1: "Hamburguesa",
	2: "Pizza",
	3: "Perro caliente",
}

func main() { ////////////////////////////////////////////////////////////////////////////////////////////////
	found, err := search("34")
	if err == errNotFound {
		fmt.Println("Pudimos controlar el error correctamente")
		return
	} else if err != nil {
		fmt.Println("Search()", err)
		return
	}
	fmt.Println(found)

} ////////////////////////////////////////////////////////////////////////////////////////////////

func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(): %v", err)
	}

	emoji, err := findFood(num)
	if err != nil {
		return "", fmt.Errorf("findFood(): %v", err)
	}
	return emoji, nil
}

func findFood(id int) (string, error) {
	value, ok := food[id]
	if !ok {
		return "", errNotFound
	}
	return value, nil
}
