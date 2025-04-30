package http

import (
	"web-craw/internals/presentation/http/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	articleController *controllers.ArticleController,
	crawlController *controllers.CrawlController,
	jobController *controllers.JobController,
) *gin.Engine {
	router := gin.Default()

	router.GET("/articles", articleController.GetArticles)
	router.POST("/crawl", crawlController.StartCrawl)
	router.GET("/crawl/:job_id", jobController.GetJobStatus)

	return router
}
