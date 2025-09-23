package main

import (
	"fmt"
)

type TrunkItem struct {
	name     string
	quantity int
}

type Car struct {
	brand string
	name  string
	color string
	power int
	trunk []TrunkItem
}

func (c Car) displayCar() {
	fmt.Println("===== Mon véhicule =====")
	fmt.Printf("\tModèle : %s\n", c.name)
	fmt.Printf("\tMarque : %s\n", c.brand)
	fmt.Printf("\tPuissance : %dCH\n", c.power)
	fmt.Printf("\tCouleur : %s\n", c.color)
	fmt.Println("=== Coffre du véhicule ===")
	if len(c.trunk) < 1 {
		fmt.Printf("\tCoffre du véhicule vide...\n")
		return
	}
	for _, item := range c.trunk {
		fmt.Printf("\t- %s x%d\n", item.name, item.quantity)
	}
}
