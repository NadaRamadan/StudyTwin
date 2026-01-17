package routes

import (
	"net/http"

	"github.com/NadaRamadan/StudyTwin/DTOs"
	"github.com/gin-gonic/gin"
)

func RegisterQuizRoutes(router *gin.Engine) {
	router.POST("/quiz", func(c *gin.Context) {
		var req dto.QuizRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		questions := []dto.QuizQuestion{
			{
				ID:            1,
				Text:          "What is the primary output of photosynthesis?",
				Options:       []string{"Carbon Dioxide", "Oxygen", "Nitrogen", "Helium"},
				CorrectAnswer: 1,
			},
			{
				ID:            2,
				Text:          "Which organelle is known as the powerhouse of the cell?",
				Options:       []string{"Nucleus", "Ribosome", "Mitochondria", "Golgi Apparatus"},
				CorrectAnswer: 2,
			},
			{
				ID:            3,
				Text:          "Who formulated the laws of motion?",
				Options:       []string{"Isaac Newton", "Albert Einstein", "Galileo Galilei", "Nikola Tesla"},
				CorrectAnswer: 0,
			},
		}

		c.JSON(http.StatusOK, dto.QuizResponse{
			Questions: questions,
		})
	})
}
