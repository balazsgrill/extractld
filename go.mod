module github.com/balazsgrill/extractld

go 1.18

require (
	github.com/balazsgrill/oauthenticator v0.0.0-20230131174844-7c6759b33a9a
	github.com/balazsgrill/sparqlupdate v0.0.0-20230108043446-cb1acac570f1
)

require github.com/knakk/sparql v0.0.0-20220326141742-15797a7da0ca // indirect

require (
	github.com/balazsgrill/basecamp3 v0.0.0-20230108173528-44f824ab5128
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/knakk/digest v0.0.0-20160404164910-fd45becddc49 // indirect
	github.com/knakk/rdf v0.0.0-20190304171630-8521bf4c5042
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/balazsgrill/basecamp3 => ../basecamp3

replace github.com/balazsgrill/oauthenticator => ../oauthenticator

replace github.com/balazsgrill/sparqlupdate => ../sparqlupdate
