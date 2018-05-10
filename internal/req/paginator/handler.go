package paginator

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/shopspring/decimal"
)

// Handler deals with parsing the URL query string to construct
// pagination based on the `created_at` field of tables
type Handler struct {
	perPage int
	after   *decimal.Decimal
	before  *decimal.Decimal
}

// buildPagination creates the pagination parts for a database query
func (h *Handler) BuildPagination(table string) (string, string) {
	f := ""
	if h.after != nil {
		// f = fmt.Sprintf("%v.created_at < timestamp 'epoch' + :after * INTERVAL '1 microsecond'", table)
		f = fmt.Sprintf("%v.created_at < to_timestamp(:after)", table)
	}

	if h.before != nil {
		f = fmt.Sprintf("%v.created_at > timestamp 'epoch' + :before * INTERVAL '1 microsecond'", table)
	}

	return f, " limit :per_page"
}

// PaginateQueryArgs creates a map[string]interface for use in constructing a query string
func (h *Handler) PaginateQueryArgs() map[string]interface{} {
	args := make(map[string]interface{})

	args["per_page"] = h.perPage

	if h.after != nil {
		args["after"] = *h.after
	}

	if h.before != nil {
		args["before"] = *h.before
	}

	return args
}

// PopulatePaginate sets the Handler values for use later
func (h *Handler) PopulatePaginate(vals url.Values) {
	after := vals.Get("after")
	if after != "" {
		i, err := strconv.ParseFloat(after, 64)
		if err == nil {
			x := decimal.NewFromFloat(i)
			h.after = &x
		}
	}

	before := vals.Get("before")
	if before != "" {
		i, err := strconv.ParseFloat(before, 64)
		if err == nil {
			x := decimal.NewFromFloat(i)
			h.before = &x
		}
	}

	perPage := vals.Get("per_page")
	if perPage != "" {
		i, _ := strconv.Atoi(perPage)
		h.perPage = i
	} else {
		h.perPage = defaultPerPage
	}
}
