package service

import (
	"context"
	"io"

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

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.Id)
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

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}
	for {
		category, err := stream.Recv()
		if err  == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		categories.Categories = append(categories.Categories, &pb.Category{
			Id: categoryResult.ID,
			Name: categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err  == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Category{
			Id: categoryResult.ID,
			Name: categoryResult.Name,
			Description: categoryResult.Description,
		})

		if err != nil {
			return err
		}
	}
}