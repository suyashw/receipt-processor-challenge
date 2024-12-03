# Receipt Processor

This document provides instructions to set up, run, and test the Receipt Processor application.

## Prerequisites

Before running the application, ensure you have the following installed:

1. **Go**: Version 1.23 or higher
2. **Docker**: Latest stable version
3. **cURL** or a tool like **Postman** for testing the API

---

## Running the Application

You can run the application either locally using Go or using Docker.

## Option 1: Run Locally

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Start the application:
   ```bash
   go run main.go
   ```

4. The server will start at `http://localhost:8080`.

## Option 2: Run with Docker

1. Build the Docker image:
   ```bash
   docker build -t receipt-processor .
   ```

2. Run the Docker container:
   ```bash
   docker run -p 8080:8080 receipt-processor
   ```

3. The server will be accessible at `http://localhost:8080`.

---

## Testing the API

The application provides two endpoints for testing:

1. **Process Receipts**
   - **Endpoint**: `POST /receipts/process`
   - **Description**: Processes a receipt and returns a unique ID for it.

2. **Get Points**
   - **Endpoint**: `GET /receipts/{id}/points`
   - **Description**: Fetches the points associated with the given receipt ID.

## Example Test Commands

## Process a Receipt

1. Use the `morning-receipt.json` file to test:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d @morning-receipt.json http://localhost:8080/receipts/process
   ```

2. Expected Response:
   ```json
   { "id": "generated-uuid" }
   ```

## Get Points for a Receipt

1. Use the receipt ID from the previous step to fetch points:
   ```bash
   curl http://localhost:8080/receipts/<id>/points
   ```

2. Expected Response (example):
   ```json
   { "points": 50 }
   ```

---

## Unit Testing

The application includes unit tests to verify functionality:

1. Run the tests:
   ```bash
   go test ./...
   ```

2. Expected Output:
   - All tests should pass successfully.

---

## Example Input Files

Two example receipt files are included for testing:

1. `morning-receipt.json`
2. `simple-receipt.json`

These files contain sample receipt data and can be used to validate the application behavior.

---

## Notes

- The application uses in-memory storage, so data will not persist after a restart.
- Refer to the `README.md` for detailed rules on point calculation.

If you encounter any issues, feel free to open an issue on this repository.