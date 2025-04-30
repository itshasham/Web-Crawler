package controllers

import (
	"net/http"
	"web-craw/internals/domain/repository"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articleRepo repository.ArticleRepository
}

// NewArticleController creates a new ArticleController.
func NewArticleController(articleRepo repository.ArticleRepository) *ArticleController {
	return &ArticleController{articleRepo: articleRepo}
}

// GetArticles handles GET /articles.
func (h *ArticleController) GetArticles(c *gin.Context) {
	source := c.Query("source")
	tag := c.Query("tag")

	filters := make(map[string]interface{})
	if source != "" {
		filters["source"] = source
	}
	if tag != "" {
		filters["tags"] = tag
	}

	articles, err := h.articleRepo.Search(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch articles"})
		return
	}

	c.JSON(http.StatusOK, articles)
}
