package service

import (
	"context"
	"github.com/feliperromao/go-grpc/internal/database"
	"github.com/feliperromao/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	response := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}

	return response, nil
}