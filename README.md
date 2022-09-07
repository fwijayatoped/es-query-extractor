# es-query-extractor
To extract your es query and send it to the service based on defined client.

## Supported Client
* oliverev6, currently only support searchservice
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

	qe := queryExtractor.NewOlivere6Session(
		Client.GetService(),
	)

	qe.WithFullPath("fullpath").WithKeyword("keyword").WithUsecase("myusecase")

	ss := elastic.NewSearchService(esClient).Query(elastic.NewBoolQuery().Must(elastic.NewExistsQuery("feri")))
	qe.SendSearchService(*ss)
```
