package esqueryextractor

import (
	"context"

	elastic "gopkg.in/olivere/elastic.v6"
)

type Olivere6Builder struct {
	commonAttribute CommonAttributeContract
	keyword         string
}

// Interface if client using olivere6 client lib
type Olivere6Contract interface {
	Contract
	SendSearchService(searchService elastic.SearchService)
}

const DefaultUsecase = "undefined"

// Start new session
func NewOlivere6Session(service Service, rateLimiter *RateLimiter) Olivere6Contract {
	return &Olivere6Builder{
		commonAttribute: CommonAttributeContract{
			service:     service,
			usecase:     DefaultUsecase,
			rateLimiter: rateLimiter,
		},
	}
}

// Full-Path:
// Define the fullpath that will be sent as header request
func (b *Olivere6Builder) WithFullPath(fullPath string) Contract {
	b.commonAttribute.fullPath = fullPath
	return b
}

// Keyword:
// Define the keyword that will be sent as header request,
// can be empty just in case no keyword usecase
func (b *Olivere6Builder) WithKeyword(keyword string) Contract {
	b.keyword = keyword
	return b
}

// Usecase:
// Define the usecase that will be sent as header request
func (b *Olivere6Builder) WithUsecase(usecase string) Contract {
	b.commonAttribute.usecase = usecase
	return b
}

// Debug-Mode:
// Define that debug mode will be exclude rate limiter
func (b *Olivere6Builder) WithDebugMode() Contract {
	b.commonAttribute.debugMode = true
	return b
}

// SendSearchService the request via olivere6 client lib
func (b *Olivere6Builder) SendSearchService(service elastic.SearchService) {
	service.Profile(true)
	service.Header("Service", string(b.commonAttribute.service))
	service.Header("Usecase", b.commonAttribute.usecase)
	service.Header("Query-Type", "_search")

	if b.keyword != "" {
		service.Header("Keyword", b.keyword)
	}

	if b.commonAttribute.fullPath != "" {
		service.Header("Full-Path", b.commonAttribute.fullPath)
	}

	callback := func() {
		service.Do(context.Background())
	}

	if b.commonAttribute.debugMode {
		go callback()
	} else {
		b.commonAttribute.rateLimiter.AddFunc(callback)
	}
}

// SendCountService the request via olivere6 client lib
func (b *Olivere6Builder) SendCountService(service elastic.CountService) {
	service.Header("Service", string(b.commonAttribute.service))
	service.Header("Usecase", b.commonAttribute.usecase)
	service.Header("Query-Type", "_count")

	if b.keyword != "" {
		service.Header("Keyword", b.keyword)
	}

	if b.commonAttribute.fullPath != "" {
		service.Header("Full-Path", b.commonAttribute.fullPath)
	}

	callback := func() {
		service.Do(context.Background())
	}

	if b.commonAttribute.debugMode {
		go callback()
	} else {
		b.commonAttribute.rateLimiter.AddFunc(callback)
	}
}

// SendMgetService the request via olivere6 client lib
func (b *Olivere6Builder) SendMgetService(service elastic.MgetService) {
	service.Header("Service", string(b.commonAttribute.service))
	service.Header("Usecase", b.commonAttribute.usecase)
	service.Header("Query-Type", "_mget")

	if b.keyword != "" {
		service.Header("Keyword", b.keyword)
	}

	if b.commonAttribute.fullPath != "" {
		service.Header("Full-Path", b.commonAttribute.fullPath)
	}

	callback := func() {
		service.Do(context.Background())
	}

	if b.commonAttribute.debugMode {
		go callback()
	} else {
		b.commonAttribute.rateLimiter.AddFunc(callback)
	}
}

// SendMultiSearchService the request via olivere6 client lib
func (b *Olivere6Builder) SendMultiSearchService(service elastic.MultiSearchService) {
	service.Header("Service", string(b.commonAttribute.service))
	service.Header("Usecase", b.commonAttribute.usecase)
	service.Header("Query-Type", "_msearch")

	if b.keyword != "" {
		service.Header("Keyword", b.keyword)
	}

	if b.commonAttribute.fullPath != "" {
		service.Header("Full-Path", b.commonAttribute.fullPath)
	}

	callback := func() {
		service.Do(context.Background())
	}

	if b.commonAttribute.debugMode {
		go callback()
	} else {
		b.commonAttribute.rateLimiter.AddFunc(callback)
	}
}
