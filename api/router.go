package api

import (
	v1 "exam/api-gateway/api/handlers/v1"
	"exam/api-gateway/config"
	"exam/api-gateway/pkg/logger"
	"exam/api-gateway/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
// @title           Exam
// @version         1.0
// @description     This is test project made for exam

// @contact.name   Usmonov
// @contact.url    https://t.me/abdullohus
// @contact.email  uabdulloh11@gmail.com

// @host      18.181.186.212

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	// create customer
	api.POST("/create-customer", handlerV1.CreateCustomer)

	// get by id customer
	api.GET("/get-customer/:id", handlerV1.GetCustomerById)

	// get all customers with posts and reviews
	api.GET("/list-customer", handlerV1.GetCustomerList)

	// delete customer by id
	api.DELETE("/delete-customer/:id", handlerV1.DeleteCustomerById)

	// update customer
	api.PUT("/update-customer", handlerV1.UpdateCustomer)

	// create post
	api.POST("/create-post", handlerV1.CreatePost)
	api.PUT("/update-post", handlerV1.UpdatePost)
	api.GET("/get-post/:id", handlerV1.GetPost)
	api.GET("/list-post", handlerV1.GetPostList)
	api.DELETE("/delete-post/:id", handlerV1.DeletePostById)
	api.GET("/getpostlist/:id", handlerV1.GetPostByCustomerId)

	api.POST("/create-ranking", handlerV1.CreateReyting)
	r := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, r))

	return router
}
