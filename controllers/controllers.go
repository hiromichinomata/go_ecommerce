package controllers

import "github.com/gin-gonic/gin"

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string){
	return true, ""
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func ProductViewerAdmin() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func SearchProduct() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}