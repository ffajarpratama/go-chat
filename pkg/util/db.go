package util

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func CalculateOffset(page, limit int) (result int) {
	result = (page - 1) * limit
	return
}

func TransformSortClause(column, sort string) (result string) {
	if sort == "latest" {
		return fmt.Sprintf("%s DESC", column)
	}

	return column
}

func NewMongoPaginate(page, limit int) *options.FindOptions {
	l := int64(limit)
	skip := int64(page*limit - limit)
	return &options.FindOptions{Limit: &l, Skip: &skip}
}
