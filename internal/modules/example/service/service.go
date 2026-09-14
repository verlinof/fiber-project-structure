package example_service

import (
	"context"

	example_model "github.com/verlinof/fiber-project-structure/internal/modules/example/model"
)

func (s *ExampleService) GetAll(ctx context.Context) ([]example_model.ExampleResponse, error) {
	var items []example_model.ExampleItem
	err := s.db.WithContext(ctx).Find(&items).Error
	if err != nil {
		return nil, err
	}

	res := make([]example_model.ExampleResponse, len(items))
	for i, item := range items {
		res[i] = example_model.ExampleResponse{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return res, nil
}

func (s *ExampleService) GetByID(ctx context.Context, id int) (example_model.ExampleResponse, error) {
	var item example_model.ExampleItem
	err := s.db.WithContext(ctx).Where("id = ?", id).First(&item).Error
	if err != nil {
		return example_model.ExampleResponse{}, err
	}

	return example_model.ExampleResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}, nil
}

func (s *ExampleService) Create(ctx context.Context, req example_model.CreateExampleRequest) (example_model.ExampleResponse, error) {
	item := example_model.ExampleItem{
		Title:       req.Title,
		Description: req.Description,
	}

	err := s.db.WithContext(ctx).Create(&item).Error
	if err != nil {
		return example_model.ExampleResponse{}, err
	}

	return example_model.ExampleResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}, nil
}

func (s *ExampleService) Update(ctx context.Context, id int, req example_model.UpdateExampleRequest) (example_model.ExampleResponse, error) {
	var item example_model.ExampleItem
	err := s.db.WithContext(ctx).Where("id = ?", id).First(&item).Error
	if err != nil {
		return example_model.ExampleResponse{}, err
	}

	item.Title = req.Title
	item.Description = req.Description

	err = s.db.WithContext(ctx).Save(&item).Error
	if err != nil {
		return example_model.ExampleResponse{}, err
	}

	return example_model.ExampleResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}, nil
}

func (s *ExampleService) Delete(ctx context.Context, id int) error {
	var item example_model.ExampleItem
	err := s.db.WithContext(ctx).Where("id = ?", id).First(&item).Error
	if err != nil {
		return err
	}

	return s.db.WithContext(ctx).Delete(&item).Error
}
