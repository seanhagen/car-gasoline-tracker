package stations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	uuid "github.com/satori/go.uuid"
	"github.com/seanhagen/gas-web/internal"
	"github.com/seanhagen/gas-web/internal/models"
)

// Create creates a station in the Google Datastore
func Create(config *internal.Config) gin.HandlerFunc {
	// db := config.DB
	// dot := config.Dot

	return func(ctx *gin.Context) {
		station := &models.Station{
			ID: uuid.NewV4(),
		}

		json := binding.JSON
		err := ctx.ShouldBindWith(station, json)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, station)
	}
}
