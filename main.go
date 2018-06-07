package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/seanhagen/gas-web/internal"
	"github.com/seanhagen/gas-web/internal/routes/stations"
)

var (
	// Version is set by the build process, contains semantic version
	Version string

	// Build is set by the build process, contains sha tag of build
	Build string

	// AppName is set by the build process, contains app name
	AppName string

	// Repo is set by the build process, contains the repo where the code for this binary was built from
	Repo string

	// BuildTime is set by the build process
	BuildTime string
)

func setupRouter(config *internal.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.ErrorLogger())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Authorization", "X-Endpoint-API-UserInfo", "Origin", "Content-Length", "Content-Type",
		},
		ExposeHeaders:    []string{"Count", "Content-Length", "Link"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/", getBuildInfo(config))
	r.GET("/_health", handleEverything)
	r.GET("/_ready", handleEverything)

	r.POST("/v1/stations", stations.Create(config))
	r.POST("/v1/find-station", stations.Find(config))

	r.GET("/v1/locations", handleEverything)
	r.GET("/v1/locations/:id", handleEverything)
	r.GET("/v1/location-by-address", handleEverything)
	r.POST("/v1/locations", handleEverything)

	r.GET("/v1/records", handleEverything)
	r.GET("/v1/records/:id", handleEverything)
	r.POST("/v1/records", handleEverything)

	return r
}

func main() {
	config, err := internal.NewConfig(AppName)
	if err != nil {
		log.Panic(err)
	}

	r := setupRouter(config)

	if config.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	err = r.Run(":" + config.Port)
	if err != nil {
		log.Panicf("Unable to start server: %v", err)
	}
}

// getBuildInfo TODO
func getBuildInfo(config *internal.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		info := struct {
			Version   string `json:"version"`
			Build     string `json:"build"`
			AppName   string `json:"appname"`
			Repo      string `json:"repo"`
			BuildTime string `json:"build_time"`
		}{Version, Build, config.AppName, Repo, BuildTime}

		ctx.JSON(http.StatusOK, info)
	}
}

func handleEverything(ctx *gin.Context) {
	resp := struct {
		Msg string `json:"msg"`
	}{"Hello world!"}

	ctx.JSON(http.StatusOK, resp)
}
