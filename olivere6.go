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
func NewOlivere6Session(service Service) Olivere6Contract {
	return &Olivere6Builder{
		commonAttribute: CommonAttributeContract{
			service: service,
			usecase: DefaultUsecase,
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

// SendSearchService the request via olivere6 client lib
func (b *Olivere6Builder) SendSearchService(service elastic.SearchService) {
	go func() {
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
		service.Do(context.Background())
	}()
}

// SendCountService the request via olivere6 client lib
func (b *Olivere6Builder) SendCountService(service elastic.CountService) {
	go func() {
		service.Header("Service", string(b.commonAttribute.service))
		service.Header("Usecase", b.commonAttribute.usecase)
		service.Header("Query-Type", "_count")

		if b.keyword != "" {
			service.Header("Keyword", b.keyword)
		}

		if b.commonAttribute.fullPath != "" {
			service.Header("Full-Path", b.commonAttribute.fullPath)
		}

		service.Do(context.Background())
	}()
}

// SendMgetService the request via olivere6 client lib
func (b *Olivere6Builder) SendMgetService(service elastic.MgetService) {
	go func() {
		service.Header("Service", string(b.commonAttribute.service))
		service.Header("Usecase", b.commonAttribute.usecase)
		service.Header("Query-Type", "_mget")

		if b.keyword != "" {
			service.Header("Keyword", b.keyword)
		}

		if b.commonAttribute.fullPath != "" {
			service.Header("Full-Path", b.commonAttribute.fullPath)
		}

		service.Do(context.Background())
	}()
}

// SendMultiSearchService the request via olivere6 client lib
func (b *Olivere6Builder) SendMultiSearchService(service elastic.MultiSearchService) {
	go func() {
		service.Header("Service", string(b.commonAttribute.service))
		service.Header("Usecase", b.commonAttribute.usecase)
		service.Header("Query-Type", "_msearch")
		if b.keyword != "" {
			service.Header("Keyword", b.keyword)
		}

		if b.commonAttribute.fullPath != "" {
			service.Header("Full-Path", b.commonAttribute.fullPath)
		}

		service.Do(context.Background())
	}()
}
