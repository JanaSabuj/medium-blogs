# goapp
 - A simple go http web server image for use in personal projects
 - Pass the `APPPORT` env variable to expose the web server on that port
 - For using docker-compose, modify the file according to port preference

### What's in the app
 - If we hit /XYZ - we would get a dynamic text like below - extracted from the URLPath
<img width="573" alt="image" src="https://user-images.githubusercontent.com/39147514/222353686-35cfac89-7f9e-47d3-a8c4-10525d5980be.png">
 - If we hit /Summer - we would get Summer populated
<img width="573" alt="image" src="https://user-images.githubusercontent.com/39147514/222353823-0efdac71-733e-443e-86f8-52eddd0afc5a.png">
 - The root path - extracts "Everything"
<img width="573" alt="image" src="https://user-images.githubusercontent.com/39147514/222353942-d6428298-411d-45f2-9faa-74db8aad93d8.png">


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

## Docker image
- https://hub.docker.com/r/sabujjana/goapp
