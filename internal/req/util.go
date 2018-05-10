package req

import (
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetUUID TODO
func GetUUID(ctx *gin.Context, paramName string) (uuid.UUID, error) {
	t := ctx.Param(paramName)
	if t == "" {
		return uuid.Nil, fmt.Errorf("ID is required, got %#v", t)
	}

	id, err := uuid.FromString(t)
	if err != nil {
		return uuid.Nil, err
	}

	if id == uuid.Nil {
		return id, fmt.Errorf("got nil ID")
	}
	return id, nil
}
