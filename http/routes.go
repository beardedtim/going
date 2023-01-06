package http

import (
	"fmt"
	"net/http"

	database "mkc-p/modi/data"

	"github.com/gin-gonic/gin"
)

func FindAllAuthors(c *gin.Context) {
	rows, error := database.FindAllAuthors()

	if error != nil {
		fmt.Println(error)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something whent wrong. Please try again later",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": rows,
	})
}

func CreateAuthor(c *gin.Context) {
	var requestBody database.AuthorCreate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot parse the incoming body. Please modify your request before trying again.",
		})

		return
	}

	fmt.Println(requestBody)

	result, err := database.CreateNewAuthor(requestBody)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There was an error performing your request. Please try it again later.",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": result,
	})
}

func Healthcheck(c *gin.Context) {
	healthy := database.IsHealthy()

	if healthy {
		c.JSON(http.StatusOK, gin.H{
			"healthy": true,
		})
	} else {
		if healthy {
			c.JSON(http.StatusInternalServerError, gin.H{
				"healthy": false,
			})
		}
	}
}
