package character

import (
	"net/http"
	"strconv"

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
	foundCharacters, err := h.characterInteractor.ListCharacters()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	charactersRes := ToListResponse(&foundCharacters)

	ctx.JSON(http.StatusOK, charactersRes)
}

func (h *handler) FindCharacter(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		panic(err)
	}

	foundCharacter, err := h.characterInteractor.FindCharacterById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	characterRes := ToResponse(&foundCharacter)
	ctx.JSON(http.StatusOK, characterRes)
}

func (h *handler) CreateCharacter(ctx *gin.Context) {
	character := entity.Character{}
	if err := ctx.ShouldBindJSON(&character); err != nil {
		panic(err)
	}

	createdCharacter, err := h.characterInteractor.CreateCharacter(character)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	characterRes := ToResponse(&createdCharacter)
	ctx.JSON(http.StatusOK, characterRes)
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

	updatedCharacter, err := h.characterInteractor.UpdateCharacter(character, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	characterRes := ToResponse(&updatedCharacter)
	ctx.JSON(http.StatusOK, characterRes)
}

func (h *handler) DeleteCharacter(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := h.characterInteractor.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.Status(http.StatusNoContent)
}
