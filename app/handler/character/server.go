package character

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Seiya-Tagami/favorite-character-api/domain/entity"
	rc "github.com/Seiya-Tagami/favorite-character-api/handler/response/character"
	re "github.com/Seiya-Tagami/favorite-character-api/handler/response/errors"
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
		errorRes := re.ToResponse(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRes})
		return
	}

	charactersRes := rc.ToListResponse(&foundCharacters)

	ctx.JSON(http.StatusOK, charactersRes)
}

func (h *handler) FindCharacter(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	foundCharacter, err := h.characterInteractor.FindCharacterById(id)
	if err != nil {
		errorRes := re.ToResponse(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRes})
		return
	}

	characterRes := rc.ToResponse(&foundCharacter)
	ctx.JSON(http.StatusOK, characterRes)
}

func (h *handler) CreateCharacter(ctx *gin.Context) {
	character := entity.Character{}
	if err := ctx.ShouldBindJSON(&character); err != nil {
		errorRes := re.ToResponse(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorRes})
		return
	}

	if err := character.Validate(); err != nil {
		errorRes := re.ToResponse(http.StatusBadRequest, err.Error())
		fmt.Println(errorRes)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorRes})
		return
	}

	createdCharacter, err := h.characterInteractor.CreateCharacter(character)
	if err != nil {
		errorRes := re.ToResponse(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRes})
		return
	}

	characterRes := rc.ToResponse(&createdCharacter)
	ctx.JSON(http.StatusOK, characterRes)
}

func (h *handler) UpdateCharacter(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	character := entity.Character{}
	character.ID = id
	if err := ctx.ShouldBindJSON(&character); err != nil {
		errorRes := re.ToResponse(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorRes})
		return
	}

	if err := character.Validate(); err != nil {
		errorRes := re.ToResponse(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRes})
		return
	}

	updatedCharacter, err := h.characterInteractor.UpdateCharacter(character, id)
	if err != nil {
		errorRes := re.ToResponse(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRes})
		return
	}

	characterRes := rc.ToResponse(&updatedCharacter)
	ctx.JSON(http.StatusOK, characterRes)
}

func (h *handler) DeleteCharacter(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := h.characterInteractor.DeleteById(id)
	if err != nil {
		errorRes := re.ToResponse(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRes})
		return
	}

	ctx.Status(http.StatusNoContent)
}
