package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/mateusdeitos/fullcycle-graphql/graph/generated"
	"github.com/mateusdeitos/fullcycle-graphql/graph/model"
)

func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	var courses []*model.Course

	for _, course := range r.Resolver.Courses {
		if course.Category.ID == obj.ID {
			courses = append(courses, course)
		}
	}

	return courses, nil
}

func (r *courseResolver) Chapters(ctx context.Context, obj *model.Course) ([]*model.Chapter, error) {
	var chapters []*model.Chapter

	for _, chapter := range r.Resolver.Chapters {
		if chapter.Course.ID == obj.ID {
			chapters = append(chapters, chapter)
		}
	}

	return chapters, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := &model.Category{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: &input.Description,
	}

	r.Resolver.Categories = append(r.Resolver.Categories, category)

	return category, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	var category *model.Category

	for _, c := range r.Resolver.Categories {
		if c.ID == input.CategoryID {
			category = c
			break
		}
	}

	course := &model.Course{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: &input.Description,
		Category:    category,
	}

	r.Resolver.Courses = append(r.Resolver.Courses, course)

	return course, nil
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	var course *model.Course

	for _, c := range r.Resolver.Courses {
		if c.ID == input.CourseID {
			course = c
			break
		}
	}

	var category *model.Category
	for _, c := range r.Resolver.Categories {
		if c.ID == course.Category.ID {
			category = c
			break
		}
	}

	chapter := &model.Chapter{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		Course:   course,
		Category: category,
	}

	r.Resolver.Chapters = append(r.Resolver.Chapters, chapter)

	return chapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
