package esqueryextractor

import (
	"context"

	elastic "gopkg.in/olivere/elastic.v6"
)

type Olivere6Builder struct {
	commonAttribute CommonAttributeContract
}

// Interface if client using olivere6 client lib
type Olivere6Contract interface {
	WithQueryString(querystring string) Olivere6Contract
	WithExtraAttributes(attributes map[string]string) Olivere6Contract
	SendSearchService(searchService elastic.SearchService)
}

// Start new session
func NewOlivere6Session() Olivere6Contract {
	return &Olivere6Builder{}
}

// Define the query string from user
func (b *Olivere6Builder) WithQueryString(querystring string) Olivere6Contract {
	b.commonAttribute.querystring = querystring
	return b
}

// Define the extra attributes that will be sent as header request
func (b *Olivere6Builder) WithExtraAttributes(attributes map[string]string) Olivere6Contract {
	b.commonAttribute.extraAttributes = attributes
	return b
}

// SendSearchService the request via olivere6 client lib
func (b *Olivere6Builder) SendSearchService(searchService elastic.SearchService) {
	searchService.Header("User-Agent", "ESQueryExtractor")
	go func() {
		for k, v := range b.commonAttribute.extraAttributes {
			searchService.Header(k, v)
		}
		searchService.Do(context.Background())
	}()
}
