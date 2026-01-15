package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func RegisterQuizRoutes(router *gin.Engine) {
	router.POST("/quiz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"questions": []gin.H{
				{
					"id":   1,
					"text": "What is the primary output of photosynthesis?",
					"options": []string{
						"Carbon Dioxide",
						"Oxygen",
						"Nitrogen",
						"Helium",
					},
					"correctAnswer": 1,
				},
				{
					"id":   2,
					"text": "Which organelle is known as the powerhouse of the cell?",
					"options": []string{
						"Nucleus",
						"Ribosome",
						"Mitochondria",
						"Golgi Apparatus",
					},
					"correctAnswer": 2,
				},
				{
					"id":   3,
					"text": "Who formulated the laws of motion?",
					"options": []string{
						"Isaac Newton",
						"Albert Einstein",
						"Galileo Galilei",
						"Nikola Tesla",
					},
					"correctAnswer": 0,
				},
			},
		})
	})
}
