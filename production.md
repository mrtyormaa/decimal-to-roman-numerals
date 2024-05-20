# Steps to make the API Production ready

We will discuss this in 2 sections. One being the functional ascpect of the API and the other non-functional. 

We will make the following assumptions to help us reach to some concrete actionable plans. The assumptions are as follows:
- This will be served to about 10M users all over the world.
- It should have good latency i.e. API responds within 100ms for 95% of requests.
- Service downtime should be ideally zero. 

Let's begin with the functional requirements

## Functional Requirements
There are several actions to be completed. You can view the full details on the [GitHub issues page](https://github.com/mrtyormaa/decimal-to-roman-numerals/issues). The summary of the requirements are as follows:

### Bugs
1. **Swagger: `swagger/*any`**
   - **Issue #:** #10
   - **Description:** A bug related to the Swagger implementation.

2. **Incorrect/Misleading Error Message**
   - **Issue #:** #8
   - **Description:** A bug regarding error messages.

### Enhancements
1. **Error Message Internationalization**
   - **Issue #:** #14
   - **Description:** Feature request for supporting multiple languages in error messages.

2. **Implement User Authentication**
   - **Issue #:** #13
   - **Description:** Adding user authentication feature.

3. **Implement a Rate Limiter**
   - **Issue #:** #12
   - **Description:** Implementing a rate limiter feature.

4. **Implement a Load Balancer**
   - **Issue #:** #11
   - **Description:** Feature request for a load balancer.

### Documentation
1. **Swagger Example: Max Appears Before Min**
   - **Issue #:** #9
   - **Description:** Documentation improvement needed for Swagger examples.

### General
1. **Incongruent Behaviour between /GET and /POST**
   - **Issue #:** #7
   - **Description:** Discrepancy between GET and POST behavior, low priority.

Now to elaborate on different functional ascpects and discuss further:

### 1. API Endpoints

#### 1.1 GET Endpoint - COMPLETED
Please find the details [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#api-documentation)
#### 1.2 POST Endpoint - COMPLETED
Please find the details [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#api-documentation)
## 2. Validation - COMPLETED

- Ensure the provided integer is within the acceptable range for Roman numerals (1 to 3999).
- Handle errors for invalid inputs, returning appropriate HTTP status codes (e.g., 400 for bad request).
- Extensive Tests: Unit tests. Integration Tests. Load Test. The full description can be found [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#testing).
    - Unit Tests: 96% code coverage for the project.
    - Integrations Tests: Both endpoints tested for cases like valid, invalid and  edge cases etc.
    - Load Tests: Basic implementation. Perform total 1000 requests, distributed among the goroutines and check the status code and validate the response body.
    - Compliance Tests: Not required at the moment. But if we implement user authentication and start storing data, these tests should be implemented.
    - Regression Tests: Github actions has been implemented for this project. You can view the tests and the results [here](https://github.com/mrtyormaa/decimal-to-roman-numerals/actions). For further details on the CI/CD practices for the project refer [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#cicd-process) 
- TODO: Integration with load-testing tools like `k6` to do more thorough tests.

## 3. Rate Limiting - todo

- Implement rate limiting to prevent abuse and ensure fair usage among users.
- Define rate limits per user/IP, e.g., 100 requests per minute.
    - We can do this by implementing it via libraries provided by Go.
    - This can also be done via services like cloudfare, Amazon API gateway etc.

## 4. API Documentation - COMPLETED

- Provide comprehensive API documentation using tools like Swagger or Postman.
- Include detailed descriptions of endpoints, parameters, request/response formats, and error codes.
- Provide examples for common use cases.
- Further details can be found [here](https://github.com/mrtyormaa/decimal-to-roman-numerals?tab=readme-ov-file#api-documentation).

## 5. Internationalization - todo

- Support multiple languages in error messages and documentation to cater to a global audience.
- The documentation and the error messages should be translated
    - This can be acheived by various libraries. There is no standard library by go. But there are many open-source alternatives

## 6. Security - todo
This is not a requirement but it can be useful if we want to make the service accessible via authorisation for probably monetization.
- Implement HTTPS to encrypt data in transit.
- Use API keys or OAuth for authentication to restrict access to authorized users.
- Protect against common web vulnerabilities (e.g., SQL injection, XSS) using security best practices and frameworks. 
    - The project tries to have a very naive implementation for these. It is not enough for a production quality. We should use standard frameworks for this.

Below are some approaches to acheive this.
| **Strategy**             | **Description**                                            | **Libraries/Tools**                                                 |
|--------------------------|------------------------------------------------------------|---------------------------------------------------------------------|
| **Basic Authentication** | Simple method using username and password.                 | Standard library (`net/http`)                                       |
| **Token-Based (JWT)**    | Uses JSON Web Tokens for stateless authentication.         | `jwt-go`                                                            |
| **OAuth 2.0**            | Uses third-party services for authentication.              | `golang.org/x/oauth2`                                               |
| **Session-Based**        | Maintains user sessions on the server.                     | `gorilla/sessions`                                                  |
| **API Key Authentication**| Uses API keys passed in headers or query parameters.       | Custom middleware                                                   |
| **LDAP Authentication**  | Authenticates against an LDAP server.                      | `github.com/go-ldap/ldap`                                           |
| **Third-Party Auth**     | Integrates with services like Firebase, Auth0, etc.        | Firebase SDK, Auth0 SDK                                             |

