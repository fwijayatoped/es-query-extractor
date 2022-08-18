package esqueryextractor

import (
	"context"

	elastic "gopkg.in/olivere/elastic.v6"
)

type Olivere6Builder struct {
	c               *Olivere6Client
	commonAttribute CommonAttributeContract
}

func NewOlivere6Session(c *Olivere6Client) Contract {
	return &Olivere6Builder{
		c: c,
	}
}

func (b *Olivere6Builder) WithPath(path string) Contract {
	b.commonAttribute.path = path
	return b
}

func (b *Olivere6Builder) WithQueryString(querystring string) Contract {
	b.commonAttribute.querystring = querystring
	return b
}

func (b *Olivere6Builder) WithExtraAttributes(attributes map[string]string) Contract {
	b.commonAttribute.extraAttributes = attributes
	return b
}

func (b *Olivere6Builder) Send(searchService elastic.SearchService) {
	go func() {
		for k, v := range b.commonAttribute.extraAttributes {
			searchService.Header(k, v)
		}
		searchService.Do(context.Background())
	}()

}
