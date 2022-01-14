package news

import (
	_newsUsecase "alta/cleanarch/businesses/news"
	_newsControllers "alta/cleanarch/controllers"
	_newsResponse "alta/cleanarch/controllers/news/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type NewsController struct {
	newsUseCase _newsUsecase.Usecase
}

func NewNewsController(newsUC _newsUsecase.Usecase) *NewsController {
	return &NewsController{
		newsUseCase: newsUC,
	}
}

func (newsController *NewsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := newsController.newsUseCase.GetAll(ctx)
	if err != nil {
		return _newsControllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return _newsControllers.NewSuccessResponse(c, result)
}

func (newsController *NewsController) GetByID(c echo.Context) error {
	newsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	ctx := c.Request().Context()

	result, err := newsController.newsUseCase.GetByID(ctx, newsId)
	if err != nil {
		return _newsControllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return _newsControllers.NewSuccessResponse(c, _newsResponse.FromDomain(result))
}
