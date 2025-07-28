package controllers

import "github.com/gin-gonic/gin"

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dummy"})
	}
}
