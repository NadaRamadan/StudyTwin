package main

import(
	"Gemini "
	"github.com/gin-gonic/gin"
)

function main(){

	r := gin.Default()

    routes.RegisterQuizRoutes(r)
	routes.RegisterSummaryRoutes(r)

	r.Run(":8080")

}