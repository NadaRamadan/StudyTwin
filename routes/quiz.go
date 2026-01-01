package routes

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterQuizRoutes(router *gin.Engine){
	router.POST("/quiz", func(c *gin.Context)
	c.JSON(http.StatusOK, gin.H{
		"questions": gin.H{
			
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
		}
	},
	})
)
}