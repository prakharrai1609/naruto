## Naruto Go Backend Framework

Naruto is a lightweight Go backend framework inspired by Express. It provides a simple and flexible structure for building web applications and APIs in the Go programming language.

### How to Test

1. Navigate to the tests directory for the Naruto framework:

    ```bash
    cd naruto/tests/naruto
    ```

2. Run the tests using the following command:

    ```bash
    go test
    ```

### Test Server

1. Navigate to the test server directory:

    ```bash
    cd naruto/tests/test-server/server
    ```

2. Run the test server:

    ```bash
    go run server.go
    ```

### Sample Test cURLs

1. Retrieve user class information:

    ```bash
    curl http://localhost:8080/user/class/
    ```

2. Create an admin:

    ```bash
    curl http://localhost:8080/admin/create/
    ```

Feel free to explore and modify the Naruto framework based on your project requirements. For more information on building routes, middleware, and handling requests and responses, refer to the Naruto documentation.

Happy coding with Naruto! 🍥🍃


Owner: [Prakhar Rai](https://www.twitter.com/prakharrai1609) 