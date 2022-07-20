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
	e.HTTPErrorHandler = _validatorUser.ErroHandlerBooking
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
	e.DELETE("/restaurants", presenter.RestaurantPresenter.DestroyResto, _middleware.JWTMiddleware())
	e.POST("/restaurants/upload", presenter.RestaurantPresenter.UploadImageResto, _middleware.JWTMiddleware())
	e.GET("/restaurants", presenter.RestaurantPresenter.AllResto)
	e.GET("/restaurants/:id", presenter.RestaurantPresenter.DetailResto)
	e.GET("/myresto", presenter.RestaurantPresenter.MyResto, _middleware.JWTMiddleware())

	// comments and ratings
	e.POST("/comments/:id", presenter.CommentPresenter.PostComment, _middleware.JWTMiddleware())
	e.GET("/comments/:id", presenter.CommentPresenter.GetComment)
	e.GET("/comments/rating/:id", presenter.CommentPresenter.GetRating)

	// favourites
	e.POST("/favourites/:id", presenter.FavouritePresenter.PostFav, _middleware.JWTMiddleware())
	e.DELETE("/favourites/:id", presenter.FavouritePresenter.DeleteFav, _middleware.JWTMiddleware())
	e.GET("/favourites", presenter.FavouritePresenter.MyFav, _middleware.JWTMiddleware())

	// admin
	e.GET("/admins/users", presenter.AdminPresenter.AllUser, _middleware.JWTMiddleware())
	e.GET("/admins/restaurants", presenter.AdminPresenter.AllResto, _middleware.JWTMiddleware())
	e.GET("/admins/restaurants/:id", presenter.AdminPresenter.DetailResto, _middleware.JWTMiddleware())
	e.POST("/admins/verif/:id", presenter.AdminPresenter.VerifResto, _middleware.JWTMiddleware())

	// booking
	e.POST("/restaurants/booking/:id", presenter.BookingPresenter.BookingResto, _middleware.JWTMiddleware())
	e.POST("/payment/webhook", presenter.BookingPresenter.AcceptPayment)

	return e

}
