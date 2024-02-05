package repository

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
	Page  int    `json:"page"`
	Field string `json:"field"`
}

func NewPagination(r *http.Request) (*Pagination, error) {
	var (
		sort  = "asc"
		field = "name"
		limit = 10
		page  = 1
	)

	if s := r.URL.Query().Get("sort"); s != "" {
		sort = s
	}

	if f := r.URL.Query().Get("field"); f != "" {
		field = f
	}

	strLimit := r.URL.Query().Get("limit")
	if strLimit != "" {
		intLimit, err := strconv.Atoi(strLimit)

		if err != nil {
			slog.Warn("param limit is invalid", slog.Any("error", err), slog.String("limit", strLimit))
			return nil, errors.New("limit value is invalid")
		}

		if intLimit > 0 {
			limit = intLimit
		}
	}

	strPage := r.URL.Query().Get("page")
	if strPage != "" {
		intPage, err := strconv.Atoi(strPage)

		if err != nil {
			slog.Warn("param page is invalid", slog.Any("error", err), slog.String("page", strPage))
			return nil, errors.New("page value is invalid")
		}

		if intPage > 0 {
			page = intPage
		}
	}

	pagination := &Pagination{Limit: limit, Sort: sort, Page: page, Field: field}
	if err := pagination.validate(); err != nil {
		return nil, err
	}

	return pagination, nil
}

func (p Pagination) Query() string {
	offset := p.Limit * (p.Page - 1)
	query := fmt.Sprintf("SELECT * FROM products ORDER BY %s %s LIMIT %d OFFSET %d", p.Field, p.Sort, p.Limit, offset)
	fmt.Println(query)
	return query
}

func (p Pagination) validate() error {
	if p.Limit <= 0 {
		return errors.New("pagination limit should be > 0")
	}

	if p.Sort != "asc" && p.Sort != "desc" {
		return errors.New("pagination sort should be 'asc' or 'desc'")
	}

	if p.Page <= 0 {
		return errors.New("pagination page should be > 0")
	}

	switch p.Field {
	case "id", "name", "price", "description", "availableQuantity":
		break
	default:
		return errors.New("pagination field should be 'id', 'name', 'price', 'description' or 'availableQuantity'")
	}

	return nil
}
