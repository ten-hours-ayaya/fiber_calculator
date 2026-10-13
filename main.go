package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ten-hours-ayaya/fiber-calculator/internal/app"
	"github.com/ten-hours-ayaya/fiber-calculator/internal/util"
)

func main() {
	path := os.Args[1]
	input, err := util.LoadFromFile(path)
	if err != nil {
		log.Fatal("Cannot load input from file", err)
		panic(err)
	}

	app := app.NewApp(input.Dividers, input.Targets)
	result, err := app.Calculate(input.Input, input.Streets)
	if err != nil {
		fmt.Printf("Error during calculation: %v\n", err)
	}

	for _, resultStreet := range result.StreetResults {
		if resultStreet == nil {
			continue
		}
		if resultStreet.Divider != nil {
			fmt.Printf("Street: %v, input: %.2f, divider: %v\n", resultStreet.StreetName, resultStreet.Input, resultStreet.Divider.Name)
		} else {
			fmt.Printf("Street: %v, input: %.2f\n", resultStreet.StreetName, resultStreet.Input)
		}
		for _, node := range resultStreet.Nodes {
			if node == nil {
				continue
			}
			if node.Divider != nil {
				fmt.Printf("\tInput: %.2f, target: %.2f, result: %.2f, divider: %v\n", node.Input, node.Target, node.Input-node.Divider.Left, node.Divider.Name)
			} else {
				fmt.Printf("\tInput: %.2f, target: %.2f, result: %.2f\n", node.Input, node.Target, node.Input)
			}
		}
	}
}
