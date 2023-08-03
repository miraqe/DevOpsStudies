# DevOpsStudies

Introduction:

The "DevOpsStudies" repository contains a Golang-based application that interacts with the Pipedrive public API. The primary purpose of this application is to learn and practice DevOps principles, including containerization with Docker, continuous integration with GitHub Actions, and monitoring with Prometheus. The application serves as a study project for developers interested in exploring DevOps concepts and applying them in a real-world scenario.

Main Features:

 - Pipedrive API Interaction: The application interacts with the Pipedrive API to perform various operations, such as fetching deals, adding new deals, and changing existing deals. It utilizes HTTP methods like GET, POST, and PUT to communicate with the Pipedrive API and manage deal-related data.

 - Docker Containerization: The application is containerized using Docker to provide a consistent and isolated environment for running the application. Docker allows developers to package the application along with its dependencies, ensuring seamless deployment and portability across different systems.

 - Continuous Integration with GitHub Actions: The repository is integrated with GitHub Actions to automatically run tests and checks whenever new code is pushed to the repository. This ensures that any changes to the codebase are thoroughly tested, maintaining code quality and reliability.

 - Monitoring with Prometheus: The application is instrumented with custom metrics using Prometheus, which provides monitoring and observability capabilities. Developers can monitor important metrics, such as the total number of HTTP requests and their durations, to gain insights into the application's performance.

 - Through this project, developers can gain hands-on experience with key DevOps tools and practices, enabling them to build scalable and resilient applications while adhering to DevOps principles. The following instructions provide a step-by-step guide on how to set up and run the application locally using Docker, explore its API endpoints, and test its functionality.

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
  
 Users can replace the token line with their own Pipedrive API token.

Tests can be run straight from Docker Desktop by clicking on the running container name and using Docker built in Terminal. It scans the source code to find functions starting with "Test" and runs them as test cases. These tests verify the correctness of our API endpoints, ensuring they behave as expected. Command to use:
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

To see documentation about My Endpoints: URL https://documenter.getpostman.com/view/28507647/2s9XxwwuUQ
  - In the browser, click on "Run in Postman", the Postman Desktop application will open (depending that you have it installed) and asks in which environment to save this.

6. Known issues

The endpoints for /addDeal and /changeDeal are not functioning as expected when running the application locally on localhost (broswer)
While the application is calling correct methods, PUT and POST respectively, the browser tries to handle these methods as GET which results in BAD REQUEST.
Application passes tests in Docker, locally in IntelliJ, Github Actions and in Postman, which suggests the POST and PUT method for endpoints
are called correctly and the application connects to Pipedrive API without any issue.
