package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var shopItems = make(map[string]float64)

func listHandler(w http.ResponseWriter, r *http.Request) {
	if len(shopItems) == 0 {
		fmt.Fprintf(w, "Shop has no items\n")
	} else {
		fmt.Fprintf(w, "Shop items:\n\n")

		for itemName, itemPrice := range shopItems {
			fmt.Fprintf(w, "%s: $%.2f\n", itemName, itemPrice)
		}
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("item")
	itemPriceStr := r.URL.Query().Get("price")

	if itemName != "" && itemPriceStr != "" {
		if itemPrice, err := strconv.ParseFloat(itemPriceStr, 64); err != nil {
			fmt.Fprintf(w, "Could not convert price to a decimal value\n")
		} else {
			if _, ok := shopItems[itemName]; ok {
				fmt.Fprintf(w, "Item %s already exists in the shop\n", itemName)
			} else {
				shopItems[itemName] = itemPrice

				fmt.Fprintf(w, "Successfully created item %s with price $%.2f\n", itemName, itemPrice)
			}
		}
	}
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	if itemName := r.URL.Query().Get("item"); itemName != "" {
		if itemPrice, ok := shopItems[itemName]; ok {
			fmt.Fprintf(w, "The price of %s is $%.2f\n", itemName, itemPrice)
		} else {
			fmt.Fprintf(w, "Item %s does not exist in the shop\n", itemName)
		}
	} else {
		fmt.Fprintf(w, "No item name was specified\n")
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("item")
	itemPriceStr := r.URL.Query().Get("price")

	if itemName != "" && itemPriceStr != "" {
		if itemPrice, err := strconv.ParseFloat(itemPriceStr, 64); err != nil {
			fmt.Fprintf(w, "Could not convert price to a decimal value\n")
		} else {
			if _, ok := shopItems[itemName]; ok {
				shopItems[itemName] = itemPrice

				fmt.Fprintf(w, "Successfully updated item %s with price $%.2f\n", itemName, itemPrice)
			} else {
				fmt.Fprintf(w, "Item %s does not exist in the shop\n", itemName)
			}
		}
	} else {
		fmt.Fprintf(w, "Item name or price was not specified\n")
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if itemName := r.URL.Query().Get("item"); itemName != "" {
		if _, ok := shopItems[itemName]; ok {
			delete(shopItems, itemName)

			fmt.Fprintf(w, "Successfully deleted item %s from the shop\n", itemName)
		} else {
			fmt.Fprintf(w, "Item %s does not exist in the shop\n", itemName)
		}
	}
}

func main() {
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
