package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// WHERE column = value
func WhereEqual[T comparable](db *gorm.DB, column string, value T) *gorm.DB {

	var zero T
	if value == zero {
		return db
	}

	return db.Where(clause.Eq{Column: column, Value: value})
}

// WHERE column = *value
func WherePtr[T any](db *gorm.DB, column string, value *T) *gorm.DB {

	if value == nil {
		return db
	}

	return db.Where(clause.Eq{Column: column, Value: *value})
}

// WHERE column LIKE '%value%'
func WhereLike(db *gorm.DB, column string, value string) *gorm.DB {

	if value == "" {
		return db
	}

	return db.Where(clause.Like{Column: column, Value: "%" + value + "%"})
}

// WHERE column IN (...)
func WhereIn[T any](db *gorm.DB, column string, values []T) *gorm.DB {

	if len(values) == 0 {
		return db
	}

	return db.Where(column+" IN ?", values)
}

// WHERE column >= value
func WhereGte[T comparable](db *gorm.DB, column string, value T) *gorm.DB {

	var zero T
	if value == zero {
		return db
	}

	return db.Where(clause.Gte{Column: column, Value: value})
}

// WHERE column <= value
func WhereLte[T comparable](db *gorm.DB, column string, value T) *gorm.DB {

	var zero T
	if value == zero {
		return db
	}

	return db.Where(clause.Lte{Column: column, Value: value})
}

// WHERE column > value
func WhereGt[T comparable](db *gorm.DB, column string, value T) *gorm.DB {

	var zero T
	if value == zero {
		return db
	}

	return db.Where(clause.Gt{Column: column, Value: value})
}

// WHERE column < value
func WhereLt[T comparable](db *gorm.DB, column string, value T) *gorm.DB {

	var zero T
	if value == zero {
		return db
	}

	return db.Where(clause.Lt{Column: column, Value: value})
}

// WHERE column BETWEEN start AND end
func WhereBetween[T comparable](db *gorm.DB, column string, start, end T) *gorm.DB {

	var zero T
	if start == zero || end == zero {
		return db
	}

	return db.Where(column+" BETWEEN ? AND ?", start, end)
}

// ORDER BY
func OrderBy(db *gorm.DB, column string, desc bool) *gorm.DB {

	if column == "" {
		return db
	}

	return db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: column},
		Desc:   desc,
	})
}

// Pagination
func Paginate(db *gorm.DB, page, limit int) *gorm.DB {

	if limit <= 0 {
		return db
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	return db.Offset(offset).Limit(limit)
}
