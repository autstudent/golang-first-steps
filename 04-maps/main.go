package main

import "fmt"

func main()  {
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"blue": "#232rs2",
	}
	fmt.Println(colors)

	// var colorsMap map[string]string
	colorsMap := make(map[string]string)
	colorsMap["red"] = "#ff0000"
	colorsMap["green"] = "#4bf745"
	colorsMap["blue"] = "#232rs2"
	delete(colorsMap, "red")
	fmt.Println(colorsMap)

	printMap(colorsMap)

}

func printMap(c map[string]string) {
    for color,hex := range c {
		fmt.Println(color + " - " + hex)
    }
}
