package main

import (
	"context"
	"log"
	"time"

	"web-craw/internals/domain/entity"
	"web-craw/internals/infrastructure/database"
	"web-craw/internals/infrastructure/logger"
	adapters "web-craw/internals/presentation/adaptors"
	"web-craw/internals/presentation/adaptors/bbc"
	"web-craw/internals/presentation/adaptors/guardian"
	"web-craw/internals/presentation/http"
	httphandler "web-craw/internals/presentation/http/controllers"
	crawler "web-craw/internals/usecases/crawl"
)

func main() {
	// Initialize Logger
	logger.InitLogger()
	logger.Info("Starting Web Crawler Application")

	// Initialize Repositories
	articleRepo := database.NewInMemoryArticleRepo()
	crawlJobRepo := database.NewInMemoryCrawlJobRepo()
	sourceRepo := database.NewInMemorySourceRepo()

	// ✅ Preload Sources into sourceRepo (Guardian and BBC)
	_ = sourceRepo.Save(context.Background(), &entity.Source{
		SourceID:    "guardian-source-id",
		Name:        "Guardian",
		BaseURL:     "https://www.theguardian.com",
		AdapterName: "GuardianAdapter",
	})

	_ = sourceRepo.Save(context.Background(), &entity.Source{
		SourceID:    "bbc-source-id",
		Name:        "BBC",
		BaseURL:     "https://www.bbc.com",
		AdapterName: "BBCAdapter",
	})

	// Correct SiteAdapters map
	siteAdapters := map[string]adapters.SiteAdapter{
		"GuardianAdapter": guardian.NewGuardianAdapter(),
		"BBCAdapter":      bbc.NewBBCAdapter(),
	}

	// Create RateLimiter and Deduplicator
	rateLimiter := crawler.NewRateLimiter(time.Second) // 1 second delay
	deduplicator := crawler.NewDeduplicator()

	// Initialize Crawler Engine correctly
	engine := crawler.NewEngine(
		articleRepo,
		crawlJobRepo,
		sourceRepo,
		siteAdapters,
		rateLimiter,
		deduplicator,
	)

	// Initialize Controllers
	articleController := httphandler.NewArticleController(articleRepo)
	crawlController := httphandler.NewCrawlController(engine)
	jobController := httphandler.NewJobController(crawlJobRepo)

	// Setup Router
	router := http.NewRouter(articleController, crawlController, jobController)

	// Start HTTP Server
	port := "8080"
	logger.Info("Server starting on port " + port)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
