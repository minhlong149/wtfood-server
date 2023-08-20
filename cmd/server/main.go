package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"wtfood/pkg/handler"
	"wtfood/pkg/middleware"
	"wtfood/pkg/util"
)

type App struct {
	Router   *gin.Engine
	Database *gorm.DB
	Config   util.Config
	Logger   util.Logger
}

func (app *App) SetupRouter() {
	handler := handler.SetupHandler(app.Logger, app.Config, app.Database)

	apiRouter := app.Router.Group("/api")
	apiRouter.Use(middleware.ErrorHandler(app.Logger))

	apiRouter.POST("/login", handler.Login)

	ingredientRouter := apiRouter.Group("/ingredients")
	{
		ingredientRouter.GET("/all", handler.GetAllIngredient)
	}

	dishRouter := apiRouter.Group("/dishes")
	{
		dishRouter.GET("/random", handler.GetRandomDish)
		dishRouter.GET("/:dishId/ingredients/:ingredientId", handler.CheckRecipe)
	}

	authenticatedDishRouter := dishRouter.Use(middleware.Authenticated(app.Config.SecretKey))
	{
		authenticatedDishRouter.GET("/all", handler.GetAllDishes)
		authenticatedDishRouter.GET("/:dishId", handler.GetDishById)
		authenticatedDishRouter.GET("/:dishId/ingredients", handler.GetIngredientsByDishId)
	}

	app.Router.NoRoute(gin.WrapH(http.FileServer(gin.Dir("./web", false))))
}

func main() {
	app := &App{
		Logger: util.SetupLogger(),
		Router: gin.Default(),
		Config: util.LoadConfig(),
	}

	app.Database = util.SetupDB(app.Config)
	app.SetupRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: app.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			app.Logger.Error("Listen: %s", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	if err := srv.Shutdown(context.Background()); err != nil {
		app.Logger.Error("Server shutdown: %s\n", err)
	}
}
