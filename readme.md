### Running server Locally
```
    go run cmd/main.go
```

### Running server via docker container
make sure docker being installed if not can do refer the installation process [here](https://docs.docker.com/engine/install)
once installed follow the following instruction
```
    docker build -t vmware .
    docker run -p 8084:8084 vmware
```

### Run Test Cases and Find Code coverage
```
    go test ./... -coverprofile cover.out
```

### API DOCUMENTATION

1. Fetch Assignment 
```
    endpoint :- http://localhost:8084/v1/assignemt/get

    request body :- 
        {
            "sort" : "relevanceScore",  
            "limit" : 3                
        }

    sort key helps to sort the result based on the value passed, it can have value like `url` , `views` , `relevanceScore` ,mandatory field.
    limit key will help us to restrict the number of assignemt to be fetched, can have value between 1 to 200 , mandatory field

    
    response body :- 
            {
                "count": 15,
                "data": [
                    {
                        "url": "www.wikipedia.com/abc1",
                        "views": 11000,
                        "relevanceScore": 0.1
                    },
                    {
                        "url": "www.example.com/abc1",
                        "views": 1000,
                        "relevanceScore": 0.1
                    },
                    {
                        "url": "www.wikipedia.com/abc2",
                        "views": 12000,
                        "relevanceScore": 0.2
                    }
                ],
                "message": "Assignment fetched Successfully"
            }
```