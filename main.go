package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//-----------------------------------EMPLOYEE method------------------------------
	//GET METHOD
	r.GET("/EMPLOYEE", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET EMPLOYEE METHOD",
		})
	})

	//POST METHOD
	r.POST("/EMPLOYEE", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST EMPLOYEE METHOD",
		})
	})

	//PUT METHOD
	r.PUT("/EMPLOYEE", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT EMPLOYEE METHOD",
		})
	})

		//DELETE METHOD
		r.DELETE("/EMPLOYEE", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE EMPLOYEE METHOD",
			})
		})
	
		//--------------------------------------------------------------------

		//-----------------------------------CUSTOMER method------------------------------
	//GET METHOD
	r.GET("/CUSTOMER", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET CUSTOMER METHOD",
		})
	})

	//POST METHOD
	r.POST("/CUSTOMER", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST CUSTOMER METHOD",
		})
	})

	//PUT METHOD
	r.PUT("/CUSTOMER", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT CUSTOMER METHOD",
		})
	})

		//DELETE METHOD
		r.DELETE("/CUSTOMER", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE CUSTOMER METHOD",
			})
		})
	
		//--------------------------------------------------------------------
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
