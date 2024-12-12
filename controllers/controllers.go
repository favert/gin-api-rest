package controllers

import (
	"net/http"

	"github.com/favert/api-go-gin/database"
	"github.com/favert/api-go-gin/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ExibeTodosAlunos godoc
//
//	@Summary		Exibir todos os alunos
//	@Description	Exibe alunos
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: ": "E ai " + nome + ", tudo beleza?",
	})
}

// CriaNovoAluno godoc
//
//	@Summary		Cria um aluno
//	@Description	Recebe json e cria aluno
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			aluno	body		models.Aluno	true	"Cria aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidaDadosDeAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletaUmAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)
	//c.JSON(http.StatusOK, aluno)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})

}

// BuscaUmAlunoPorID godoc
//
//	@Summary		Exibir um aluno
//	@Description	Exibe aluno pelo id
//	@Tags			alunos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID do Aluno"
//	@Success		200	{object}	models.Aluno
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/alunos/{id} [get]
func BuscaUmAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno nao encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidaDadosDeAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)

}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno nao encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
