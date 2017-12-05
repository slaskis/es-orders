```
go get -u github.com/altairsix/eventsource/...
go get github.com/altairsix/eventsource-protobuf/...
eventsource dynamodb create-table --name orders --region eu-central-1

go generate main.go
go build .
```
