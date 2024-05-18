package routes

import "github.com/labstack/echo/v4"

func (routes *ArticleRoutesImpl) MobileArticleRoutes(apiGroup echo.Group) {
	ArticleGroup := apiGroup.Group("/api/articles")

	ArticleGroup.POST("/create", routes.Handler.CreateArticle)
}