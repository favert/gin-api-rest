package routes

import (
	"github.com/favert/api-go-gin/controllers"
	docs "github.com/favert/api-go-gin/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaUmAlunoPorID)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaUmAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
