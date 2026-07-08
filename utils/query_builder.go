package utils

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func WhereEqual[T comparable](db *gorm.DB, column string, value T) *gorm.DB {

	var zero T
	if value == zero {
		return db
	}

	return db.Where(clause.Eq{Column: column, Value: value})
}

func WhereLike(db *gorm.DB, column string, value string) *gorm.DB {

	if value == "" {
		return db
	}

	return db.Where(clause.Like{Column: column, Value: "%" + value + "%"})
}
