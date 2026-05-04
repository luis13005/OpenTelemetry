package controller

import (
	"cepclima/internal/entity"
	"cepclima/internal/usecase"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var cepRegex = regexp.MustCompile(`^[0-9]{5}-?[0-9]{3}$`)

type CepController struct {
	useCase usecase.CepUsecase
}

func NewCepController(useCase usecase.CepUsecase) *CepController {
	return &CepController{useCase: useCase}
}

func (c *CepController) GetClima(ctx *gin.Context) {
	var input entity.InputCep

	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !cepRegex.MatchString(input.Cep) {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	cepTratado := strings.ReplaceAll(input.Cep, "-", "")

	clima, err := c.useCase.GetCepClima(ctx.Request.Context(), cepTratado)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, clima)
}
