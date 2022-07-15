package routes

import (
	"capstone/group3/factory"
	_middleware "capstone/group3/features/middlewares"
	_validatorUser "capstone/group3/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	// presenter := factory.InitFactory()
	e := echo.New()

	e.HTTPErrorHandler = _validatorUser.ErrorHandlerUser
	e.HTTPErrorHandler = _validatorUser.ErroHandlerRestaurant

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// users
	e.POST("/users", presenter.UserPresenter.PostUser)
	e.POST("/login", presenter.UserPresenter.LoginAuth)
	e.PUT("/users", presenter.UserPresenter.PutUser, _middleware.JWTMiddleware())
	e.GET("/myprofile", presenter.UserPresenter.GetByMe, _middleware.JWTMiddleware())
	e.DELETE("/users", presenter.UserPresenter.DeleteByID, _middleware.JWTMiddleware())

	// restaurants
	e.POST("/restaurants", presenter.RestaurantPresenter.CreateResto, _middleware.JWTMiddleware())
	e.PUT("/restaurants", presenter.RestaurantPresenter.UpdateResto, _middleware.JWTMiddleware())

	// comments and ratings
	e.POST("/comments/:id", presenter.CommentPresenter.PostComment, _middleware.JWTMiddleware())
	e.GET("/comments/:id", presenter.CommentPresenter.GetComment)
	e.GET("/comments/rating/:id", presenter.CommentPresenter.GetRating)

	return e

}
