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
	WithFullPath(string) Olivere6Contract
	WithKeyword(string) Olivere6Contract
	WithUsecase(string) Olivere6Contract
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
func (b *Olivere6Builder) WithFullPath(fullPath string) Olivere6Contract {
	b.commonAttribute.fullPath = fullPath
	return b
}

// Keyword:
// Define the keyword that will be sent as header request,
// can be empty just in case no keyword usecase
func (b *Olivere6Builder) WithKeyword(keyword string) Olivere6Contract {
	b.keyword = keyword
	return b
}

// Usecase:
// Define the usecase that will be sent as header request
func (b *Olivere6Builder) WithUsecase(usecase string) Olivere6Contract {
	b.commonAttribute.usecase = usecase
	return b
}

// SendSearchService the request via olivere6 client lib
func (b *Olivere6Builder) SendSearchService(searchService elastic.SearchService) {
	go func() {
		searchService.Header("Service", string(b.commonAttribute.service))
		searchService.Header("Usecase", b.commonAttribute.usecase)

		if b.keyword != "" {
			searchService.Header("Keyword", b.keyword)
		}

		if b.commonAttribute.fullPath != "" {
			searchService.Header("Full-Path", b.commonAttribute.fullPath)
		}

		searchService.Do(context.Background())
	}()
}
