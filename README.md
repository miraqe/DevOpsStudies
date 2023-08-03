# DevOpsStudies
This repository contains Pipedrive public API learning sessions.

Language used: Golang

Apps and tools used:
- IntelliJ
- Postman
- Docker Desktop
- Git bash
- Windows terminal
- Prometheus
- Github Actions

1. Necessary links to my work
Docker image:
 - https://hub.docker.com/repository/docker/mmiraqe/dealflows-go-pipedrive/general

Postman:
 - https://www.postman.com/miraqe/workspace/v1/collection/28507647-b6b8b116-8c24-4fb2-be97-b645abcda40b

2. Instructions on how to access Docker image using Windows Terminal:
   - docker pull mmiraqe/dealflows-go-pipedrive:latest
   - docker run -d -e PIPEDRIVE_API_TOKEN=02a9c777d721453536ee0c7b328d02a4e10fc449 -p 8080:8081 mmiraqe/dealflows-go-pipedrive:latest

Tests can be run straight from Docker Desktop by clicking on the running container name and using Docker built in Terminal:
  - go test

To run tests using Windows Terminal
Access container's shell (replace <CONTAINER_NAME> with the name you see in DOcker Desktop:
  - docker exec -it <CONTAINER_NAME> /bin/bash
Make sure the directory is /app inside the container
  - cd /app
Run tests using go test command. For example:
  - go test ./... -v -cover

You can use other commands within the Terminal to list the files in the directory:
  - ls
To check the content of a specific file (replace <FILE_NAME> with existing file name after fetching the list:
  - cat <FILE_NAME>

3. API endpoints
  - /getDeals: Get deals from the Pipedrive API
  - /addDeal: Add a new deal to the Pipedrive API
  - /changeDeal: Change an existing deal in the Pipedrive API

Additionally, /metrics endpoint can be accessed via localhost:8081/metrics
Custom metric requests added:
   - httpTotalRequests
   - httpDuration

4. Using POSTMAN to check endpoints:
  - I use Postman Desktop when application is running
  - To see endpoints in web URL: https://www.postman.com/miraqe/workspace/v1/collection/28507647-b6b8b116-8c24-4fb2-be97-b645abcda40b

6. Known issues

The endpoints for /addDeal and /changeDeal are not functioning as expected when running the application locally on localhost (broswer)
While the application is calling correct methods, PUT and POST respectively, the browser tries to handle these methods as GET which results in BAD REQUEST.
Application passes tests in Docker, locally in IntelliJ, Github Actions and in Postman, which suggests the POST and PUT method for endpoints
are called correctly and the application connects to Pipedrive API without any issue.
