package article

import (
	"net/http"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/controllers/article/request"
	"github.com/blueharvest-alterra/go-back-end/controllers/article/response"
	"github.com/blueharvest-alterra/go-back-end/controllers/base"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/blueharvest-alterra/go-back-end/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	articleUseCase entities.ArticleUseCaseInterface
}

func NewarticleController(articleUseCase entities.ArticleUseCaseInterface) *ArticleController {
	return &ArticleController{
		articleUseCase: articleUseCase,
	}
}

func (ac *ArticleController) Create(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var articleCreate request.CreateArticleRequest
	articleCreate.AdminID = userData.ID
	if err := c.Bind(&articleCreate); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constant.ErrEmptyInput.Error()))
	}

	picture := form.File["picture_file"]
	if len(picture) == 0 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Gambar Artikel Tidak Boleh Kosong"))
	}

	if len(picture) > 1 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Gambar Artikel Hanya Boleh Satu"))
	}
	for _, file := range picture {
		if !utils.IsImageFile(file.Filename) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Format file gambar tidak didukung"))
		}
	}

	article, errUseCase := ac.articleUseCase.Create(articleCreate.ToEntities(), userData, picture)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	articleResponse := response.ArticleResponseFromUseCase(&article)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Article created!", articleResponse))
}

func (ac *ArticleController) GetById(c echo.Context) error {
	articleId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	article, errUseCase := ac.articleUseCase.GetById(articleId)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	articleResponse := response.ArticleResponseFromUseCase(&article)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get article data!", articleResponse))

}

func (ac *ArticleController) Update(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	var articleEdit request.EditArticleRequest
	articleEdit.AdminID = userData.ID
	if err := c.Bind(&articleEdit); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	articleEdit.ID = articleId

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse(constant.ErrEmptyInput.Error()))
	}

	picture := form.File["picture_file"]

	if len(picture) > 1 {
		return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Gambar Tanaman Hanya Boleh Satu"))
	}
	for _, file := range picture {
		if !utils.IsImageFile(file.Filename) {
			return c.JSON(http.StatusBadRequest, base.NewErrorResponse("Format file gambar tidak didukung"))
		}
	}

	article, errUseCase := ac.articleUseCase.Update(articleEdit.ToEntities(), userData, picture)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	articleResponse := response.ArticleResponseFromUseCase(&article)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("article updated!", articleResponse))
}

func (ac *ArticleController) Delete(c echo.Context) error {
	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	articleId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		if uuid.IsInvalidLengthError(err) {
			return c.JSON(http.StatusNotFound, base.NewErrorResponse(err.Error()))
		}
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	article, errUseCase := ac.articleUseCase.Delete(articleId, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	articleResponse := response.ArticleResponseFromUseCase(&article)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success delete article data!", articleResponse))
}

func (ac *ArticleController) GetAll(c echo.Context) error {
	articles, errUseCase := ac.articleUseCase.GetAll(&[]entities.Article{})
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	articleGetAllResponse := response.SliceFromUseCase(&articles)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success get all article data!", articleGetAllResponse))
}
