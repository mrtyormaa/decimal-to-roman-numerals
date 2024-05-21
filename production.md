# Steps to Make the decimal-to-roman API Production Ready

`decimal-to-roman` API converts decimal number ranges to Roman numerals. This document discusses the next steps required to make the API production ready. This has been discussed in two sections:
- Functional 
- Non-Functional

The following assumptions to help in reaching some concrete actionable plans:
- decimal-to-roman API will be served to about 10M users worldwide.
- API should have good latency, i.e., the API should respond within 100ms for 95% of requests.
- Service downtime should ideally be zero i.e. service should 99.99999% (7 nines) availability.

## Open Issues
There are several requirements to be completed. You can view the full details on the [GitHub issues page](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues). The summary of the requirements is as follows:

### Bugs
1. **Swagger: `swagger/*any`**
   - **Issue #:** [#10](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/10)
   - **Description:** A bug related to the Swagger library. `/swagger/` does not work and we have to go to `/swagger/index.html` to access swagger.

2. **Incorrect/Misleading Error Message**
   - **Issue #:** [#8](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/8)
   - **Description:** A bug regarding error messages. Although, we catch the exception, the error message is misleading.

### Enhancements
1. **Error Message Granularization**
   - **Issue #:** [#15](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/15)
   - **Description:** Error messages should be more specific to the exact issue instead of having one error message being served to 3 different cases.
2. **Error Message Internationalization**
   - **Issue #:** [#14](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/14)
   - **Description:** Feature request for supporting multiple languages in error messages.

3. **Implement User Authentication**
   - **Issue #:** [#13](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/13)
   - **Description:** Adding user authentication feature.

4. **Implement a Rate Limiter**
   - **Issue #:** [#12](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/12)
   - **Description:** Implementing a rate limiter feature.

5. **Implement a Load Balancer**
   - **Issue #:** [#11](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/11)
   - **Description:** Feature request for a load balancer.

### Documentation
1. **Swagger Example: Max Appears Before Min**
   - **Issue #:** [#9](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/9)
   - **Description:** Documentation improvement needed for Swagger examples.

### General
1. **Incongruent Behaviour between /GET and /POST**
   - **Issue #:** [#7](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues/7)
   - **Description:** Discrepancy between GET and POST behavior, low priority.
2. **Code Refactoring**
    - Move type `AppError` to package `types`. At the moment it resides in package `roman`.
3. **Middleware Improvements**
    - Due to time constraints, the middleware code has not been thoroughly tested. Code quality needs to be improved. 
    - More metrics can be exported for observability.

## Functional Requirements

## 1. API Endpoints

### 1.1 GET Endpoint - COMPLETED
Please find the details [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#api-documentation).

### 1.2 POST Endpoint - COMPLETED
Please find the details [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#api-documentation).

## 2. Validation - COMPLETED

- Ensure the provided integer is within the acceptable range for Roman numerals (1 to 3999).
- Handle errors for invalid inputs, returning appropriate HTTP status codes (e.g., 400 for bad requests).
- Extensive Tests: Unit tests, Integration Tests, Load Tests. The full description can be found [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#testing).
    - Unit Tests: 96% code coverage for the project.
    - Integration Tests: Both endpoints tested for cases like valid, invalid, and edge cases.
    - Load Tests: Basic implementation. Perform a total of 1000 requests, distributed among the goroutines, and check the status code and validate the response body.
    - Compliance Tests: Not required at the moment. But if we implement user authentication and start storing data, these tests should be implemented.
    - Regression Tests: GitHub Actions have been implemented for this project. You can view the tests and the results [here](https://github.com/mrtyormaa/decimal-to-roman-numerals/actions). For further details on the CI/CD practices for the project refer [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#cicd-process).
- TODO: Integration with load-testing tools like `k6` for more thorough tests.

## 3. Rate Limiting - TODO

- Implement rate limiting to prevent abuse and ensure fair usage among users.
- Define rate limits per user/IP, e.g., 100 requests per minute.
    - We can do this by implementing it via libraries provided by Go. `golang.org/x/time/rate`
    - This can also be done via services like Cloudflare, Amazon API Gateway, etc.

## 4. API Documentation - COMPLETED

- Provide comprehensive API documentation using tools like Swagger or Postman.
- Include detailed descriptions of endpoints, parameters, request/response formats, and error codes.
- Provide examples for common use cases.
- Further details can be found [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#api-documentation).

## 5. Internationalization - TODO

- Support multiple languages in error messages and documentation to cater to a global audience.
- The documentation and the error messages should be translated.
    - This can be achieved by various libraries. There is no standard library by Go, but there are many open-source alternatives. `github.com/nicksnyder/go-i18n/v2/i18n`

## 6. Security - TODO
Authentication is not a requirement, but it can be useful if we want to make the service accessible via authorization for monetization.
- Implement HTTPS to encrypt data in transit.
- Use API keys or OAuth for authentication to restrict access to authorized users.
- Protect against common web vulnerabilities (e.g., SQL injection, XSS) using security best practices and frameworks. SQL injection is not an issue for the project at the moment, as we don't have any sql databases. But this might change in the future.
    - The project tries to have a very naive implementation to handle XSS. It is not enough for production quality. We need to test this more and we should use standard frameworks for this, if possible.

Below are some approaches to achieve this.
| **Strategy**             | **Description**                                            | **Libraries/Tools**                                                 |
|--------------------------|------------------------------------------------------------|---------------------------------------------------------------------|
| **Basic Authentication** | Simple method using username and password.                 | Standard library (`net/http`)                                       |
| **Token-Based (JWT)**    | Uses JSON Web Tokens for stateless authentication.         | `jwt-go`                                                            |
| **OAuth 2.0**            | Uses third-party services for authentication.              | `golang.org/x/oauth2`                                               |
| **Session-Based**        | Maintains user sessions on the server.                     | `gorilla/sessions`                                                  |
| **API Key Authentication**| Uses API keys passed in headers or query parameters.       | Custom middleware                                                   |
| **LDAP Authentication**  | Authenticates against an LDAP server.                      | `github.com/go-ldap/ldap`                                           |
| **Third-Party Auth**     | Integrates with services like Firebase, Auth0, etc.        | Firebase SDK, Auth0 SDK                                             |

# Non-functional Requirements (Quality Metrics)

## 1. Scalability - TODO

If the API needs to serve millions of users, multiple instances of the API container is needed. This can be achieved in the following way.

### Horizontal Scalability

- **Stateless Services**: The API service has been designed to be stateless. Hence, the state information can be stored in a distributed cache (e.g., Redis) or database for high scalability.
- **Auto-Scaling**: Auto-scaling groups (e.g., AWS Auto Scaling, Azure VM Scale Sets, Google Cloud Instance Groups) can be used to automatically adjust the number of running instances based on traffic load.
    - A `k8s` deployment artifact is present in the `feature/kubernetesIntegration` branch. This is not complete and has not been properly tested yet. A complete implementation will entail proper configuration of the deployment, service, and the ingress manifests.

### Load Balancing

- **Load Balancer**: Load balancer (e.g., AWS ELB, NGINX, HAProxy) can be deployed to distribute incoming requests across multiple instances. Ensure it supports health checks to route traffic only to healthy instances.
    - At the moment, a partial implementation of this is present in `feature/kubernetesIntegration` branch using k8s.
- **DNS Load Balancing**: DNS load balancing (e.g., AWS Route 53) can be used to distribute traffic geographically, reducing latency by directing users to the nearest data center.

## 2. Performance - TODO
In order to achieve the desired 100ms latency for 95% of requests, one or more of the following strategies can be adopted.

### Low Latency

- **Content Delivery Network (CDN)**: CDN (e.g., Cloudflare, AWS CloudFront) can be used to cache and serve responses closer to the user's location, reducing latency.
- **Optimized Code**: Profile and optimize the code to reduce execution time. Use asynchronous processing where appropriate to handle concurrent requests efficiently.
    - The API currently uses integration with `codeclimate` for code analysis. The details can be seen [here](https://codeclimate.com/github/mrtyormaa/decimal-to-roman-numerals). But this needs to be improved as well with various other static code analysis tools.

### Caching

- **In-Memory Caching**: Implement in-memory caching (e.g., Redis, Memcached) for frequently requested data to reduce database load and improve response times.
- **HTTP Caching**: Use HTTP caching headers (e.g., ETag, Cache-Control) to enable client-side caching and reduce redundant requests.

## 3. Availability - TODO
If the API needs to serve users across the globe and also make sure that the services don't suffer downtimes, we can adopt the following strategies.

### High Availability

- **Multi-Region Deployment**: Deploy the application across multiple regions to ensure availability even if one region goes down. Use a global load balancer to route traffic to the healthiest region.
- **Active-Active Failover**: Set up an active-active failover configuration where multiple regions are active simultaneously, providing immediate failover capability.

### Failover Mechanisms

- **Health Checks**: Implement comprehensive health checks for all services to detect and route traffic away from unhealthy instances.
    - A basic `/health` endpoint has been implemented for this. Integration with Prometheus and Grafana to monitor the services with various metrics has been implemented. The details of this can be found [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#logging-and-monitoring).
- **Automatic Failover**: Use cloud provider features to automatically failover to healthy instances or regions in case of failures.
    - Our implementation of k8s will also cater to this.

### Redundancy - Not required but discussing this in case the requirements evolve in the future.

- **Redundant Components**: At present the API does not have any databases. But when necessary, it can be can ensured that all critical components (servers, databases, network paths) have redundant counterparts. RAID configurations for disk redundancy and multi-zone deployment can be used for network redundancy.

## 4. Reliability - Partially Complete

### Monitoring and Logging

- **Monitoring Tools** (COMPLETE): `Prometheus` and `Grafana` are used to track metrics such as response times, error rates, and system resource usage.
- **Centralized Logging** (TODO): Centralized logging using tools like ELK stack (Elasticsearch, Logstash, Kibana) or Splunk  should be implemented to collect, aggregate, and analyze logs for troubleshooting and auditing.
    - Machine-readable error codes have been implemented. For example, `[ERR1001] Invalid JSON`. 

### Automated Recovery - In Progress

- **Self-Healing Infrastructure**: `Kubernetes` can be used to automatically restart failed containers. Configure cloud provider auto-recovery features to restart failed VMs.
    - In progress in `feature/kubernetesIntegration`.
- **Incident Response**: Automated alerting for immediate notification and response to incidents should be setup

## 5. Maintainability - COMPLETED

### Code Quality

- **Code Reviews**: This project uses Pull Requests to ensure high code quality and adherence to standards.
- **Coding Standards**: This is ensured by integration with `codecov`, `codeclimate`, and `go-report`. The links and badges can be found in the README file as well.  [![codecov](https://codecov.io/github/mrtyormaa/decimal-to-roman-numerals/graph/badge.svg?token=WCPsoNnQEy)](https://codecov.io/github/mrtyormaa/decimal-to-roman-numerals)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrtyormaa/decimal-to-roman-numerals)](https://goreportcard.com/report/github.com/mrtyormaa/decimal-to-roman-numerals)
[![Maintainability](https://api.codeclimate.com/v1/badges/dfbf91b073b8fec1f6bf/maintainability)](https://codeclimate.com/github/mrtyormaa/decimal-to-roman-numerals/maintainability).

### Automated Testing

- **CI/CD Pipelines**: This project uses GitHub Actions to automatically run tests on code changes and deploy to production only when all tests pass. [![Test and coverage](https://github.com/mrtyormaa/decimal-to-roman-numerals/actions/workflows/ci.yml/badge.svg)](https://github.com/mrtyormaa/decimal-to-roman-numerals/actions/workflows/ci.yml).

### Versioning

- **Semantic Versioning**: This project uses semantic versioning and is automated via integration with GitHub Actions. This workflow is designed to automate the creation of GitHub Releases. It adheres to [Semantic Versioning](https://semver.org/), which is a versioning scheme that uses a three-part version number: `MAJOR.MINOR.PATCH`. This workflow is triggered whenever a commit tag that starts with "v" (e.g., "v1.0.0", "v0.1.4") is pushed to the repository.

## Deployment Strategy
### Containerization - COMPLETED

- **Docker**: The project uses `Docker` for containerization of the application to ensure consistency across development, testing, and production environments. `Make` is also used to facilitate these steps.
- **Kubernetes**: We are using Kubernetes for orchestration to manage and scale containerized applications. This is still under progress as of writing this document.

### Infrastructure - TODO

- **Cloud Provider**: API can be deployed via Cloud providers that offers global reach and robust infrastructure like AWS, Google Cloud, Azure.
- **Infrastructure as Code (IaC)**: IaC tools like Terraform, Ansible, chef; Puppet can be to automate the deployment.


