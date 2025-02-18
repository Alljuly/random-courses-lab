package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

var products []Product
var nextID = 1

func main() {
	for {
		fmt.Println("\n1. Create Product")
		fmt.Println("2. Read Products")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createProduct()
		case 2:
			readProducts()
		case 3:
			updateProduct()
		case 4:
			deleteProduct()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func createProduct() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter product name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter product price: ")
	var price float64
	fmt.Scanln(&price)

	product := Product{ID: nextID, Name: name, Price: price}
	products = append(products, product)
	nextID++
	fmt.Println("Product added successfully!")
}

func readProducts() {
	if len(products) == 0 {
		fmt.Println("No products available.")
		return
	}
	for _, p := range products {
		fmt.Printf("ID: %d | Name: %s | Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}

func updateProduct() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter product ID to update: ")
	var id int
	fmt.Scanln(&id)

	for i, p := range products {
		if p.ID == id {
			fmt.Print("Enter new name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter new price: ")
			var price float64
			fmt.Scanln(&price)

			products[i].Name = name
			products[i].Price = price
			fmt.Println("Product updated successfully!")
			return
		}
	}
	fmt.Println("Product not found.")
}

func deleteProduct() {
	fmt.Print("Enter product ID to delete: ")
	var id int
	fmt.Scanln(&id)

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			fmt.Println("Product deleted successfully!")
			return
		}
	}
	fmt.Println("Product not found.")
}

/*
package http

import (
    "encoding/json"
    "net/http"
    "project/internal/core/service"
    "project/internal/core/port"
)

type ProductHandler struct {
    service port.ProductService
}

func NewProductHandler(service port.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name  string  `json:"name"`
        Price float64 `json:"price"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    product, err := h.service.CreateProduct(req.Name, req.Price)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

*/