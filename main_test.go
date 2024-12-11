package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/favert/api-go-gin/controllers"
	"github.com/favert/api-go-gin/database"
	"github.com/favert/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas
}

func CriaAlunoMock() models.Aluno {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
	return aluno
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/fabricio", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz: ":"E ai fabricio, tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
	//fmt.Println(string(respostaBody))
	//fmt.Println(mockDaResposta)
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	_ = CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
}

func TestBuscaUmAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	alunoTeste := CriaAlunoMock()
	//fmt.Println(alunoTeste)
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/"+alunoTeste.CPF, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	alunoTeste := CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.ExibeTodosAlunos)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, alunoTeste.Nome, alunoMock.Nome)
	assert.Equal(t, alunoTeste.CPF, alunoMock.CPF)
	assert.Equal(t, alunoTeste.RG, alunoMock.RG)
}

func TestDeletaUmAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	_ = CriaAlunoMock()
	// defer DeletaAlunoMock() // não precisa pois esperamos que delete no teste

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaUmAluno)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	_ = CriaAlunoMock()

	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)

	alunoNovo := models.Aluno{Nome: "Aluno novo", CPF: "01987654321", RG: "987654321"}
	alunoNovoJson, _ := json.Marshal(alunoNovo)

	req, _ := http.NewRequest("PATCH", pathDaBusca, bytes.NewBuffer(alunoNovoJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, alunoNovo.Nome, alunoMockAtualizado.Nome)
	assert.Equal(t, alunoNovo.CPF, alunoMockAtualizado.CPF)
	assert.Equal(t, alunoNovo.RG, alunoMockAtualizado.RG)
}
