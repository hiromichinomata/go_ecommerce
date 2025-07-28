package controllers

import "github.com/gin-gonic/gin"

func AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func RemoteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}