# Simple gRPC app
A simple application with frontend and backend process using gRPC. 
Backend starts a gRPC server and listens for requests. It returns the value of
the requested env variable. 
Frontend starts a http server and renders the backend's output as html.

# Build instructions
1. Compile frontend and backend `make build`
2. cleanup `make clean`

# Usage

## Frontend
```
Usage of ./frontend:
  -addr string
    	Server address string (default ":8080")
```
NOTE: 
```
GRPC_SERVER env variable can be used to set the address if the gRPC server is running on different node.
For example: export GRPC_SERVER="10.10.10.3:9000"
If not set, default value of ":9000" is used as GRPC server
```

## Example
Start frontend and backend on two different terminals

```
stiptur@mb02287 simple-gRPC-app % ./_build/frontend
2023/03/22 11:16:02 Server is listening

stiptur@mb02287 simple-gRPC-app % ./_build/backend
```
Check output by using curl or browser

```
stiptur@mb02287 simple-gRPC-app % curl http://localhost:8080/getenv\?env=SHELL
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Simple app</title>
</head>

<body>
    <p> SHELL=/bin/zsh</p>
</body>

</html>
'''
