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

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var caregoriesResponse []*pb.Category
	for _, category := range categories {
		caregoryResponse := &pb.Category{
			Id: 			category.ID,
			Name: 			category.Name,
			Description: 	category.Description,
		}
		caregoriesResponse = append(caregoriesResponse, caregoryResponse)
	}

	return &pb.CategoryList{
		Categories: caregoriesResponse,
	}, nil
}