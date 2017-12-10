```
go get -u github.com/altairsix/eventsource/...
go get github.com/altairsix/eventsource-protobuf/...
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/protoc-gen-gogofast
go get github.com/gogo/protobuf/gogoproto
eventsource dynamodb create-table --name orders --region eu-central-1

go generate main.go
go build .
```

* Can we make an aggregate of multiple repositories?
  * Or are we supposed to always build a single repo (per org)?
* Try side effects (send email, add or update external api)
  * using observers?
