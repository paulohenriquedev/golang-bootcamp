package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Aluno struct {
	Nome      string  `json:"nome"`
	Nota      float64 `json:"nota"`
	Matricula string  `json:"matricula"`
}

var alunos = []Aluno{
	{Nome: "Ana", Nota: 8.5, Matricula: "202501"},
	{Nome: "Paulo", Nota: 7.0, Matricula: "202502"},
	{Nome: "Jose", Nota: 9.0, Matricula: "202503"},
	{Nome: "Daniel", Nota: 6.5, Matricula: "202504"},
}

func getAlunos(c echo.Context) error {
	return c.JSON(http.StatusOK, alunos)
}

func getAlunoByMatricula(c echo.Context) error {
	matricula := c.Param("matricula")
	for _, aluno := range alunos {
		if aluno.Matricula == matricula {
			return c.JSON(http.StatusOK, aluno)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "Aluno não encontrado")
}

func createAluno(c echo.Context) error {

	aluno := new(Aluno)

	// c.Bind(aluno)  associa o JSON do corpo da requisição à variável aluno
	if err := c.Bind(aluno); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Dados inválidos: "+err.Error())
	}
	alunos = append(alunos, *aluno)

	return c.JSON(http.StatusCreated, aluno)
}

func updateAluno(c echo.Context) error {
	matricula := c.Param("matricula")

	alunoAtualizado := new(Aluno)
	if err := c.Bind(alunoAtualizado); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Dados inválidos: "+err.Error())
	}

	for i, aluno := range alunos {
		if aluno.Matricula == matricula {

			alunos[i].Nome = alunoAtualizado.Nome
			alunos[i].Nota = alunoAtualizado.Nota

			return c.JSON(http.StatusOK, alunos[i])
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Aluno não encontrado")
}

func deleteAlunoByMatricula(c echo.Context) error {
	matricula := c.Param("matricula")
	alunoId := -1
	for i, aluno := range alunos {
		if aluno.Matricula == matricula {
			alunoId = i
			break
		}
	}
	if alunoId == -1 {
		return echo.NewHTTPError(http.StatusNotFound, "Aluno não encontrado")
	}
	alunos = append(alunos[:alunoId], alunos[alunoId+1:]...)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.GET("/alunos", getAlunos)
	e.GET("/alunos/:matricula", getAlunoByMatricula)
	e.POST("/alunos", createAluno)
	e.PATCH("/alunos/:matricula", updateAluno)
	e.DELETE("/alunos/:matricula", deleteAlunoByMatricula)

	e.Logger.Fatal(e.Start(":8080"))

}
