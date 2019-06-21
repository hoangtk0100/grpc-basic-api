# grpc-basic-api
Create a basic API using gRPC with Protobuf

# How to build
on Terminal run this command 
```
protoc --proto_path=proto --go_out=plugins=grpc:proto service.proto
```

To start server
```
go run server/main.go
```


Then open another terminal window
```
go run client/main.go
``` 

Test on your web browser
http://localhost:8080/add/4/46
http://localhost:8080/mult/4/46
