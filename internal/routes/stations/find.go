package stations

import (
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/seanhagen/gas-web/internal"
	"github.com/seanhagen/gas-web/internal/models"
	"github.com/shopspring/decimal"
)

// LocationSearch ...
type LocationSearch struct {
	X decimal.Decimal `json:"lng" db:"x"`
	Y decimal.Decimal `json:"lat" db:"y"`
}

// Find locates the closest station to a given latlng pair
func Find(config *internal.Config) gin.HandlerFunc {
	db := config.DB
	dot := config.Dot

	return func(ctx *gin.Context) {
		query, err := dot.Raw("find-station-by-latlng")
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		log.Printf("got sql query: '%v'", query)

		in := &LocationSearch{}
		b := binding.JSON

		err = ctx.ShouldBindWith(in, b)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		out := &models.Station{}

		res, err := db.NamedExec(query, in)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		spew.Dump(res)

		ctx.JSON(http.StatusOK, out)
	}
}
