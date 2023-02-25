# goapp
 - A simple go http web server image for use in personal projects
 - Pass the `APPPORT` env variable to expose the web server on that port
 - For using docker-compose, modify the file according to port preference

## Using go
```
    APPPORT=8081 go run main.go
```

## Using Docker
```
    docker build -t goapp .
    docker run --env APPPORT=8083 -p 80:8083 goapp
```

## Using Docker Compose
```
    docker build -t goapp .
    docker-compose up
```