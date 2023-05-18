# es-query-extractor
To extract your es query and send it to the service based on defined client.

## Supported Client
* oliverev6, currently support search,count,mget,msearch
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
* go-elasticsearchv7, currently support search
```
	goElastic7Client, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:8090"},
	})

	esClient, _ := queryExtractor.Initialize(
		queryExtractor.SetService("search-microservice"),
		queryExtractor.SetGoElasticV7Client(goElastic7Client),
		queryExtractor.SetRateLimiter(20, 1, 5),
	)

	s := queryExtractor.NewGoElastic7Session(
		esClient.GoElastic7Client.GetService(),
		esClient.GoElastic7Client.GetRateLimiter(),
	)

	s.WithKeyword(keyword).WithUsecase(usecase).WithFullPath(fullpath).WithDebugMode()
	s.WithClient(client)
	s.SendSearchRequest(request, source)
```
* go-nsq, currently support search
```
	goNsqProducer, _ := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())

	nsqProducer, _ := queryExtractor.Initialize(
		queryExtractor.SetService("search-microservice"),
		queryExtractor.SetNSQProducer(goNsqProducer),
	)

	s := qe.NewGoNSQSession(
		nsqProducer.NSQProducer.GetService(),
	)

	s.WithKeyword(keyword).WithUsecase(usecase).WithFullPath(fullpath)
	s.WithProducer(producer)
	s.SendSearchRequest(index, header, source)
```