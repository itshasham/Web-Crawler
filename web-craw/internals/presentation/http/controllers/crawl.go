package controllers

import (
	"net/http"
	crawler "web-craw/internals/usecases/crawl"

	"github.com/gin-gonic/gin"
)

type CrawlController struct {
	engine *crawler.Engine
}

// NewCrawlController creates a new CrawlController.
func NewCrawlController(engine *crawler.Engine) *CrawlController {
	return &CrawlController{engine: engine}
}

// StartCrawl handles POST /crawl.
func (h *CrawlController) StartCrawl(c *gin.Context) {
	var req struct {
		Source string `json:"source" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	jobID, err := h.engine.StartCrawl(c.Request.Context(), req.Source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // ✅ send real error message
		return
	}

	c.JSON(http.StatusOK, gin.H{"job_id": jobID})
}
