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
	FindCharacter(ctx *gin.Context)
	CreateCharacter(ctx *gin.Context)
	UpdateCharacter(ctx *gin.Context)
	DeleteCharacter(ctx *gin.Context)
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

func (h *handler) FindCharacter(ctx *gin.Context) {
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
	if err := ctx.ShouldBindJSON(&character); err != nil {
		panic(err)
	}

	characterRes, err := h.characterInteractor.CreateCharacter(character)
	if err != nil {
		panic(err)
	}
	response := response.WebResponse{
		Code: http.StatusOK,
		Status: "ok",
		Data: characterRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

func (h *handler) UpdateCharacter(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		panic(err)
	}

	character := entity.Character{}
	if err := ctx.ShouldBindJSON(&character); err != nil {
		panic(err)
	}

	characterRes, err := h.characterInteractor.UpdateCharacter(character, id)
	if err != nil {
		panic(err)
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   characterRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *handler) DeleteCharacter(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := h.characterInteractor.DeleteById(id)
	if err != nil {
		panic(err)
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusNoContent)
}
