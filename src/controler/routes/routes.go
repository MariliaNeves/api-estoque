package routes

import (
	controler_product "github.com/MariliaNeves/api-estoque/src/controler/routes/product"
	controler_user "github.com/MariliaNeves/api-estoque/src/controler/routes/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users", controler_user.FindUsers)
	r.GET("/user/:id", controler_user.FindUserByID)
	r.POST("/user", controler_user.CreateUser)
	r.PUT("/user/:id", controler_user.UpdateUser)
	r.DELETE("/user/:id", controler_user.DeleteUser)

	r.GET("/products", controler_product.FindProducts)
	r.GET("/product/:id", controler_product.FindProductByID)
	r.POST("/product", controler_product.CreateProduct)
	r.PUT("/product/:id", controler_product.UpdateProduct)
	r.DELETE("/product/:id", controler_product.DeleteProduct)
}
