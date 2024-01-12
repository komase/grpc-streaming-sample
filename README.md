# grpc-straming-sample

## grpc-server
1. cd grpc-server
2. make generate
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) makes reverse proxy and swagger.json from proto file
3. make run_grpc_server
start grpc server
4. make run_grpc_gateway
start http server

## http-server
1. cd http-server
2. go run main.go
start just simple http server

## http-client
1. cd http-client
2. npm install
3. Set the version 6.6.0 using `openapi-generator-cli version`
4. npm run build