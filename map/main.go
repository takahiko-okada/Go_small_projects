package main

import "fmt"

func main() {
	// [key_type]value_type
	// colors := map[string]string {
	// "red" : "#ff0000",
	// "green" : "#4bf745",
	// }

	// different ways to declare map
	// var colors3 := map[string]string {
	// }

	// colors := make(map[int]string)
	//   colors[10] = "#fffffff"

	//   delete(colors, 10)

	//   fmt.Println(colors)
	// adding key value pair

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
