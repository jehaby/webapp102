package main

import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/jehaby/webapp102/resolver"
	"github.com/jehaby/webapp102/schema"
)

func main() {
	/* 	cfg := config.C{
	   		config.HTTP{
	   			Addr:   ":8899",
	   			Secret: "secret",
	   		}, // TODO: config
	   		config.DB{Conn: "user=postgres dbname=webapp port=65432 host=localhost sslmode=disable"},
	   	}

	   	app := service.NewApp(cfg)
	   	defer app.Logger.Sync()
	*/
	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/gdebug", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(graphqlDebugPage)
	}))

	http.Handle("/query", &relay.Handler{Schema: graphqlSchema})

	log.Fatal(http.ListenAndServe("localhost:8095", nil))

	/* 	httpApp := http.NewApp(
	   		cfg,
	   		app,
	   	)

	   	httpApp.Start(context.TODO())
	*/
}

var graphqlDebugPage = []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
		</head>
		<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function graphQLFetcher(graphQLParams) {
					return fetch("/query", {
						method: "post",
						body: JSON.stringify(graphQLParams),
						credentials: "include",
					}).then(function (response) {
						return response.text();
					}).then(function (responseBody) {
						try {
							return JSON.parse(responseBody);
						} catch (error) {
							return responseBody;
						}
					});
				}
				ReactDOM.render(
					React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
					document.getElementById("graphiql")
				);
			</script>
		</body>
	</html>
	`)
