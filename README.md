# es-query-extractor
To extract your es query and send it to the service based on defined client.

## Supported Client
* oliverev6
```
	esClient, _ := elastic.NewClient(
		elastic.SetURL("http://localhost:8090"),
		elastic.SetMaxRetries(10),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)

	Client, _ := queryExtractor.Initialize(
		queryExtractor.SetService("your service"),
		queryExtractor.SetOlivereV6Client(esClient),
	)

	qe := queryExtractor.NewOlivere6Session().WithExtraAttributes(
		map[string]string{
			"service": Client.GetService(),
		},
	)

	ss := elastic.NewSearchService(esClient).Query(elastic.NewBoolQuery().Must(elastic.NewExistsQuery("feri")))
	qe.Send(*ss)
```
