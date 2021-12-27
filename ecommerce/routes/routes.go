package routes

import (
	"ecommerce/controllers"
	"ecommerce/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// -- routes UoM
	r.GET("/uom", controllers.GetAllUom)
	// Middlewares UoM
	uomMiddlewareRoute := r.Group("/uom")
	uomMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	uomMiddlewareRoute.POST("/", controllers.CreateUom)
	uomMiddlewareRoute.PUT("/:id", controllers.UpdateUom)
	uomMiddlewareRoute.DELETE("/:id", controllers.DeleteUom)

	// routes User
	r.POST("/user/register", controllers.Register)
	r.POST("/user/login", controllers.Login)
	r.GET("/users/:id/review-product", controllers.GetReviewProductByUserId)
	userMiddlewareRoute := r.Group("/user")
	userMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	userMiddlewareRoute.PUT("/:username", controllers.ChangeUserPassword)
	userMiddlewareRoute.PUT("/profil/:username", controllers.ChangeUserData)
	userMiddlewareRoute.GET("/:username", controllers.GetUserByUserName)
	userMiddlewareRoute.DELETE("/:username", controllers.DeleteUser)
	userMiddlewareRoute.GET("/", controllers.GetAllUser)

	// -- routes ProductType
	r.GET("/product-type", controllers.GetAllProductType)
	// Middlewares ProductType
	productTypeMiddlewareRoute := r.Group("/product-type")
	productTypeMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	productTypeMiddlewareRoute.POST("/", controllers.CreateProductType)
	productTypeMiddlewareRoute.PUT("/:id", controllers.UpdateProductType)
	productTypeMiddlewareRoute.DELETE("/:id", controllers.DeleteProductType)

	// routes Payment
	r.GET("/payment", controllers.GetAllProvider)
	// Middlewares Payment
	paymentMiddlewareRoute := r.Group("/payment")
	paymentMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	paymentMiddlewareRoute.POST("/", controllers.CreateProvider)
	paymentMiddlewareRoute.PUT("/:id", controllers.UpdateProvider)
	paymentMiddlewareRoute.DELETE("/:id", controllers.DeleteProvider)

	// routes Sale Order
	r.GET("/sale-order", controllers.GetAllSaleOrder)
	r.GET("/sale-order/:id/sale-order-line", controllers.GetSaleOrderLineBySaleOrderId)
	// Middlewares Sale Order
	saleOrderMiddlewareRoute := r.Group("/sale-order")
	saleOrderMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	saleOrderMiddlewareRoute.POST("/", controllers.CreateSaleOrder)
	saleOrderMiddlewareRoute.PUT("/:id", controllers.UpdateSaleOrder)
	saleOrderMiddlewareRoute.DELETE("/:id", controllers.DeleteSaleOrder)

	// routes Product
	r.GET("/product", controllers.GetAllProduct)
	r.GET("/product/:id", controllers.GetProductById)
	r.GET("/product/:id/sale-order-line", controllers.GetSaleOrderLineByProductId)
	r.GET("/product/:id/review-product", controllers.GetReviewProductByProductId)
	// Middlewares Product
	productMiddlewareRoute := r.Group("/product")
	productMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	productMiddlewareRoute.POST("/", controllers.CreateProduct)
	productMiddlewareRoute.PUT("/:id", controllers.UpdateProduct)
	productMiddlewareRoute.DELETE("/:id", controllers.DeleteProduct)

	// routes Review Product
	r.GET("/review-product", controllers.GetAllReviewProduct)
	// Middlewares Review Product
	reviewProductMiddlewareRoute := r.Group("/review-product")
	reviewProductMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	reviewProductMiddlewareRoute.POST("/", controllers.CreateReviewProduct)
	reviewProductMiddlewareRoute.PUT("/:id", controllers.UpdateReviewProduct)
	reviewProductMiddlewareRoute.DELETE("/:id", controllers.DeleteReviewProduct)

	// routes sale order line
	r.GET("/sale-order-line", controllers.GetAllSaleOrderLine)
	// Middlewares sale order line
	saleOrderLineMiddlewareRoute := r.Group("/sale-order-line")
	saleOrderLineMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	saleOrderLineMiddlewareRoute.POST("/", controllers.CreateSaleOrderLine)
	saleOrderLineMiddlewareRoute.PUT("/:id", controllers.UpdateSaleOrderLine)
	saleOrderLineMiddlewareRoute.DELETE("/:id", controllers.DeleteSaleOrderLine)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
