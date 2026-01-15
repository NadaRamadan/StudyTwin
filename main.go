package main
import (
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
		"github.com/NadaRamadan/StudyTwin/routes"
)

func main() {
    r := gin.Default()

    routes.RegisterQuizRoutes(r)

    // Swagger endpoint
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":8080")
}
