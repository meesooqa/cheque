package services

import "context"

type CategoriesReaderResultItem interface{}

type CategoriesReader interface {
	Read() ([]CategoriesReaderResultItem, error)
}

type CategoriesAdapter interface {
	Convert(line string) (CategoriesReaderResultItem, error)
}

type CategoriesImporter interface {
	Save(ctx context.Context, items []CategoriesReaderResultItem) error
}
