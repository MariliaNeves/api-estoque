package routes

import (
	"github.com/MariliaNeves/api-estoque/src/controller"
	controller_product "github.com/MariliaNeves/api-estoque/src/controller/routes/product"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/users", userController.FindUsers)
	r.GET("/user/:id", userController.FindUserByID)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	r.GET("/products", controller_product.FindProducts)
	r.GET("/product/:id", controller_product.FindProductByID)
	r.POST("/product", controller_product.CreateProduct)
	r.PUT("/product/:id", controller_product.UpdateProduct)
	r.DELETE("/product/:id", controller_product.DeleteProduct)
}
