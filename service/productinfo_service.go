package main

import (
	"context"
	"fmt"

	"github.com/bsarvan/productInfo/service/productInfopb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*productInfopb.Product
}

func (s *server) AddProduct(ctx context.Context, product *productInfopb.Product) (*productInfopb.ProductID, error) {
	fmt.Printf("Invoking function AddProduct: %v\n", product)

	out := uuid.New()

	product.Id = out.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*productInfopb.Product)
	}

	s.productMap[product.Id] = product
	return &productInfopb.ProductID{Value: product.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, pid *productInfopb.ProductID) (*productInfopb.Product, error) {
	fmt.Printf("Invoking function GetProduct: %v\n", pid)

	value, exists := s.productMap[pid.GetValue()]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Product does not exist.", pid.GetValue())
}
