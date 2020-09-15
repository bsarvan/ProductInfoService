package main

import (
	"container/list"
	"context"
	"fmt"
	"log"

	"github.com/bsarvan/productInfo/service/productInfopb"
	"google.golang.org/grpc"
)

var productIDList = list.New()

func main() {
	fmt.Println("Hello, I'm client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := productInfopb.NewProductInfoClient(cc)

	// Add the product to the product catalogue
	doAddProductRequest(c)

	//Fetching the products for ID
	doGetProductRequest(c)
}

func doAddProductRequest(c productInfopb.ProductInfoClient) {

	fmt.Println("Adding a product to the catalogue")

	product := &productInfopb.Product{
		Name:        "iphone5",
		Description: "Apple Iphone",
	}

	res, err := c.AddProduct(context.Background(), product)
	if err != nil {
		log.Fatalf("error while adding product: %v", err)
	}

	fmt.Printf("Response from ProductInfo Service: %v\n", res.GetValue())
	productIDList.PushBack(res.GetValue())
}

func doGetProductRequest(c productInfopb.ProductInfoClient) {

	fmt.Println("Getting the product from Product Info Service")
	for e := productIDList.Front(); e != nil; e = e.Next() {
		id := fmt.Sprintf("%v", e.Value)
		fmt.Printf("Fetching Product for ProductID - %v\n", id)
		req := &productInfopb.ProductID{
			Value: id,
		}

		res, err := c.GetProduct(context.Background(), req)
		if err != nil {
			log.Fatalf("error getting the product for id: %v", err)
		}

		fmt.Printf("Received the response from server: %v\n", res)
	}
}
