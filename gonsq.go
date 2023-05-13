package esqueryextractor

import (
	"encoding/json"

	nsq "github.com/nsqio/go-nsq"
)

type GoNSQBuilder struct {
	commonAttribute CommonAttributeContract
	keyword         string
	producer        *nsq.Producer
}

type GoNSQContract interface {
	Contract
	WithProducer(producer *nsq.Producer) Contract
	SendSearchRequest(index string, header map[string]string, source interface{})
}

type NSQMessage struct {
	Index  string            `json:"index"`
	Header map[string]string `json:"header"`
	Query  []byte            `json:"query"`
}

func NewGoNSQSession(service Service) GoNSQContract {
	return &GoNSQBuilder{
		commonAttribute: CommonAttributeContract{
			service: service,
			usecase: DefaultService,
		},
	}
}

// Producer:
// Define the go-nsq producer
func (b *GoNSQBuilder) WithProducer(producer *nsq.Producer) Contract {
	b.producer = producer
	return b
}

// Full-Path:
// Define the fullpath that will be sent as header request
func (b *GoNSQBuilder) WithFullPath(fullPath string) Contract {
	b.commonAttribute.fullPath = fullPath
	return b
}

// Keyword:
// Define the keyword that will be sent as header request,
// can be empty just in case no keyword usecase
func (b *GoNSQBuilder) WithKeyword(keyword string) Contract {
	b.keyword = keyword
	return b
}

// Usecase:
// Define the usecase that will be sent as header request
func (b *GoNSQBuilder) WithUsecase(usecase string) Contract {
	b.commonAttribute.usecase = usecase
	return b
}

// Debug-Mode:
// Define that debug mode will be exclude rate limiter
func (b *GoNSQBuilder) WithDebugMode() Contract {
	return b
}

// SendSearchRequest the request via go-nsq lib
func (b *GoNSQBuilder) SendSearchRequest(index string, header map[string]string, source interface{}) {
	if b.producer == nil {
		return
	}

	query := source.(map[string]interface{})
	query["profile"] = true

	header["Service"] = string(b.commonAttribute.service)
	header["Usecase"] = b.commonAttribute.usecase
	header["Query-Type"] = "_search"

	if b.keyword != "" {
		header["Keyword"] = b.keyword
	}

	if b.commonAttribute.fullPath != "" {
		header["Full-Path"] = b.commonAttribute.fullPath
	}

	queryJSON, _ := json.Marshal(query)

	nsqMessage := NSQMessage{
		Index:  index,
		Header: header,
		Query:  queryJSON,
	}

	payload, _ := json.Marshal(nsqMessage)
	b.producer.Publish("ace-query-extractor", payload)
}
