# simple-gRPC-app
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
Usage of frontend:
-addr string
Server address string (default ":8080")
-env string
Environment variable name
```

## Example
1. Start frontend and backend on two different terminals
```
stiptur@mb02287 simple-gRPC-app % ./_build/frontend -env=SHELL
2023/03/22 11:16:02 Server is listening

stiptur@mb02287 simple-gRPC-app % ./_build/backend
```
2. Check output by using curl or browser
```
stiptur@mb02287 simple-gRPC-app % curl http://localhost:8080/getenv
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





