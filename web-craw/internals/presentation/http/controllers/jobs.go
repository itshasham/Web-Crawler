package controllers

import (
	"net/http"
	"web-craw/internals/domain/repository"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	crawlJobRepo repository.CrawlJobRepository
}

// NewJobController creates a new JobController.
func NewJobController(crawlJobRepo repository.CrawlJobRepository) *JobController {
	return &JobController{crawlJobRepo: crawlJobRepo}
}

// GetJobStatus handles GET /crawl/:job_id.
func (h *JobController) GetJobStatus(c *gin.Context) {
	jobID := c.Param("job_id")

	job, err := h.crawlJobRepo.FindByID(c.Request.Context(), jobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch job status"})
		return
	}

	c.JSON(http.StatusOK, job)
}
