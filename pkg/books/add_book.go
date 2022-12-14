package books

import (
	"net/http"

	"simple_crud/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type AddBookRequestBody struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Desription string `json:"description`
}

func (h handler) AddBook(ctx *gin.Context) {
	body := AddBookRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Desription = body.Desription
	book.Author = body.Author

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
