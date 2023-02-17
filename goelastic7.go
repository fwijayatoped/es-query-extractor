package esqueryextractor

import (
	"bytes"
	"encoding/json"

	elastic "github.com/elastic/go-elasticsearch/v7"
	esapi "github.com/elastic/go-elasticsearch/v7/esapi"
)

type GoElastic7Builder struct {
	commonAttribute CommonAttributeContract
	keyword         string
	client          *elastic.Client
}

// Interface if client using goelastic7 client lib
type GoElastic7Contract interface {
	Contract
	WithClient(client *elastic.Client) Contract
	SendSearchRequest(searchRequest []func(*esapi.SearchRequest), source interface{})
}

// Start new session
func NewGoElastic7Session(service Service, rateLimiter *RateLimiter) GoElastic7Contract {
	return &GoElastic7Builder{
		commonAttribute: CommonAttributeContract{
			service:     service,
			usecase:     DefaultUsecase,
			rateLimiter: rateLimiter,
		},
	}
}

// Client:
// Define the go-elasticsearch client
func (b *GoElastic7Builder) WithClient(client *elastic.Client) Contract {
	b.client = client
	return b
}

// Full-Path:
// Define the fullpath that will be sent as header request
func (b *GoElastic7Builder) WithFullPath(fullPath string) Contract {
	b.commonAttribute.fullPath = fullPath
	return b
}

// Keyword:
// Define the keyword that will be sent as header request,
// can be empty just in case no keyword usecase
func (b *GoElastic7Builder) WithKeyword(keyword string) Contract {
	b.keyword = keyword
	return b
}

// Usecase:
// Define the usecase that will be sent as header request
func (b *GoElastic7Builder) WithUsecase(usecase string) Contract {
	b.commonAttribute.usecase = usecase
	return b
}

// Debug-Mode:
// Define that debug mode will be exclude rate limiter
func (b *GoElastic7Builder) WithDebugMode() Contract {
	b.commonAttribute.debugMode = true
	return b
}

// SendSearchService the request via olivere6 client lib
func (b *GoElastic7Builder) SendSearchRequest(searchRequest []func(*esapi.SearchRequest), source interface{}) {
	if b.client == nil {
		return
	}

	query := source.(map[string]interface{})
	query["profile"] = true

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		return
	}

	additionalHeader := map[string]string{
		"Service":    string(b.commonAttribute.service),
		"Usecase":    b.commonAttribute.usecase,
		"Query-Type": "_search",
	}

	if b.keyword != "" {
		additionalHeader["Keyword"] = b.keyword
	}

	if b.commonAttribute.fullPath != "" {
		additionalHeader["Full-Path"] = b.commonAttribute.fullPath
	}

	searchRequest = append(searchRequest, b.client.Search.WithHeader(additionalHeader))
	callback := func() {
		b.client.Search(searchRequest...)
	}

	if b.commonAttribute.debugMode {
		go callback()
	} else {
		b.commonAttribute.rateLimiter.AddFunc(callback)
	}
}
