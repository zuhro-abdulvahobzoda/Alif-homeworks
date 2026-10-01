package main

import (
	in "fmt"
	c "shop/catalog"
)

func main() {
	products := []c.Product{
		{Name: "Shadow Milk Cookie", Price: 25.5},
		{Name: "Silent Salt Cookie", Price: 25.5},
		{Name: "Eternal Sugar Cookie", Price: 25.5},
		{Name: "Burning Spice Cookie", Price: 25.5},
		{Name: "Mystic Flour Cookie", Price: 25.5},
	}

	for i, p := range products {
		in.Printf("%d) %s: %.1f\n", i+1, p.Name, p.Price)
	}
	in.Print("Total: ", c.TotalPrice(products))
}
