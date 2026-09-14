package pkg_utils

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}

type OrderBy struct {
	Field     string
	Direction string
}

func ParseQueryParams(ctx *fiber.Ctx) ([]Filter, []OrderBy) {
	filter := ParseFiltersFromContext(ctx)
	orderBy := ParseOrderByFromContext(ctx)

	return filter, orderBy
}

// ParseFiltersFromContext mengubah query params dari context Fiber menjadi slice of Filter.
// Format yang digunakan: q_{fieldName}_{operator}=value
func ParseFiltersFromContext(ctx *fiber.Ctx) []Filter {
	var filters []Filter
	queryParams := ctx.Queries()

	for key, value := range queryParams {
		if strings.HasPrefix(key, "q_") && value != "" {
			parts := strings.SplitN(strings.TrimPrefix(key, "q_"), "_", 2)
			if len(parts) == 2 {
				field := parts[0]
				operator := strings.ToLower(parts[1])
				var filterValue interface{}

				if operator == "in" {
					filterValue = strings.Split(value, ",")
				} else {
					filterValue = value
				}

				filters = append(filters, Filter{
					Field:    field,
					Operator: operator,
					Value:    filterValue,
				})
			}
		}
	}
	return filters
}

// Format yang digunakan: order_by=field1:desc,field2:asc atau order_by=field1,field2 (default asc)
func ParseOrderByFromContext(ctx *fiber.Ctx) []OrderBy {
	var orderBy []OrderBy
	orderParam := ctx.Query("order_by")

	if orderParam == "" {
		return orderBy
	}

	// Split berdasarkan koma untuk multiple order by
	orderFields := strings.Split(orderParam, ",")

	for _, orderField := range orderFields {
		orderField = strings.TrimSpace(orderField)
		if orderField == "" {
			continue
		}

		// Check if field contains direction (field:direction)
		parts := strings.SplitN(orderField, ":", 2)
		field := strings.TrimSpace(parts[0])
		direction := "asc" // default direction

		if len(parts) == 2 {
			direction = strings.ToLower(strings.TrimSpace(parts[1]))
		}

		// Validate direction
		if direction != "asc" && direction != "desc" {
			direction = "asc" // fallback to default
		}

		orderBy = append(orderBy, OrderBy{
			Field:     field,
			Direction: direction,
		})
	}

	return orderBy
}

// ApplyFiltersAndOrder menerapkan filter dan order by ke query GORM
func ApplyFiltersAndOrder(query *gorm.DB, filters []Filter, orderBy []OrderBy, allowedFilters map[string]string, allowedOrderBy map[string]string) (*gorm.DB, error) {
	// Apply filters first
	filteredQuery, err := ApplyFilters(query, filters, allowedFilters)
	if err != nil {
		return nil, err
	}

	// Apply ordering
	orderedQuery, err := ApplyOrdering(filteredQuery, orderBy, allowedOrderBy)
	if err != nil {
		return nil, err
	}

	return orderedQuery, nil
}

// ApplyFilters diubah untuk menangani operator 'in'
func ApplyFilters(query *gorm.DB, filters []Filter, allowedFilters map[string]string) (*gorm.DB, error) {
	for _, filter := range filters {
		actualColumn, ok := allowedFilters[filter.Field]
		if !ok {
			return nil, fmt.Errorf("invalid filter: %s", filter.Field)
		}

		var operator, condition string
		var value interface{} = filter.Value

		switch filter.Operator { // operator sudah di-lowercase oleh parser
		case "eq":
			operator = "="
		case "neq":
			operator = "!="
		case "gt":
			operator = ">"
		case "gte":
			operator = ">="
		case "lt":
			operator = "<"
		case "lte":
			operator = "<="
		case "like":
			operator = "LIKE"
			// Lakukan type assertion yang aman
			if v, ok := filter.Value.(string); ok {
				value = "%" + v + "%"
			} else {
				return nil, fmt.Errorf("invalid value for like operator on field %s", filter.Field)
			}
		case "ilike":
			operator = "ILIKE"
			if v, ok := filter.Value.(string); ok {
				value = "%" + v + "%"
			} else {
				return nil, fmt.Errorf("invalid value for ilike operator on field %s", filter.Field)
			}
		case "in":
			// GORM secara otomatis menangani slice untuk query IN
			condition = fmt.Sprintf("%s IN (?)", actualColumn)
			query = query.Where(condition, value)
			continue // Lanjutkan ke filter berikutnya karena query sudah ditambahkan

		default:
			return nil, fmt.Errorf("invalid operator: %s", filter.Operator)
		}

		condition = fmt.Sprintf("%s %s ?", actualColumn, operator)
		query = query.Where(condition, value)
	}

	return query, nil
}

// ApplyOrdering menerapkan order by ke query GORM
func ApplyOrdering(query *gorm.DB, orderBy []OrderBy, allowedOrderBy map[string]string) (*gorm.DB, error) {
	for _, order := range orderBy {
		// Validate field
		actualColumn, ok := allowedOrderBy[order.Field]
		if !ok {
			return nil, fmt.Errorf("invalid order by field: %s", order.Field)
		}

		// Validate direction (sudah divalidasi di parser, tapi double check)
		direction := strings.ToLower(order.Direction)
		if direction != "asc" && direction != "desc" {
			direction = "asc" // fallback to default
		}

		// Apply ordering
		orderClause := fmt.Sprintf("%s %s", actualColumn, strings.ToUpper(direction))
		query = query.Order(orderClause)
	}
	return query, nil
}
