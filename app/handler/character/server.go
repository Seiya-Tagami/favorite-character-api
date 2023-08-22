package character

import (
	"net/http"
	"strconv"

	"github.com/Seiya-Tagami/favorite-character-api/data/response"
	"github.com/Seiya-Tagami/favorite-character-api/domain/entity"
	"github.com/Seiya-Tagami/favorite-character-api/usecase/character"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	ListCharacters(ctx *gin.Context)
	FindCharacterById(ctx *gin.Context)
	CreateCharacter(ctx *gin.Context)
}

type handler struct {
	characterInteractor character.Interactor
}

func New(characterInteractor character.Interactor) Handler {
	return &handler{characterInteractor}
}

func (h *handler) ListCharacters(ctx *gin.Context) {
	characters, err := h.characterInteractor.ListCharacters()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   characters,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *handler) FindCharacterById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		panic(err)
	}

	character, err := h.characterInteractor.FindCharacterById(id)
	if err != nil {
		panic(err)
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   character,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *handler) CreateCharacter(ctx *gin.Context) {
	character := entity.Character{}
	err := ctx.ShouldBindJSON(&character)
	if err != nil {
		panic(err)
	}

	h.characterInteractor.CreateCharacter(character)
	response := response.WebResponse{
		Code: http.StatusOK,
		Status: "ok",
		Data: nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}
