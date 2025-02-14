# Quiz App

A simple and efficient quiz application built with **Golang**, providing a **REST API** and a **CLI** interface for user interaction. The application allows users to answer quiz questions, track their progress, and compare their performance with others.

## 📌 **Project Description**
A simple and efficient quiz application built with **Golang**, providing both a **REST API** and a **CLI** interface for user interaction. The application allows users to answer quiz questions, track their progress, and compare their performance with others.

## 🛠️ **Technologies Used**
- **Golang** – Core programming language for backend development.
- **Cobra** – CLI framework for building a simple and powerful command-line interface.
- **Redis** – Used for storing user progress and rankings.
- **Chi Router** – Lightweight and fast router for the REST API.
- **Zap Logger** – Structured logging for better debugging and production logging.
- **Swagger** – API documentation with a visual UI for easy interaction with the API.
- **In-memory storage** – Simple in-memory data storage.

---

## 📂 **Project Structure**

```
/quiz
├── /cmd                    # Application entry points
│   ├── /api                # Initializes the HTTP API
│   ├── /cli                # Command-line interface (CLI) using Cobra
├── /internal
│   ├── /app                # Application layer (Handlers)
│   │   ├── /http           # API endpoints
│   │   │   ├── /question   # Question-related handlers
│   │   │   ├── /rank       # Ranking-related handlers
│   ├── /domain             # Business logic
│   │   ├── /entities       # Core entities
│   │   ├── /usecases       # Use cases (main logic)
│   ├── /infrastructure     # Storage and caching
│   │   ├── /gateways       # Redis integration
│   │   ├── /memory         # In-memory storage (mock for testing)
│
├── /pkg                    # Utility packages (e.g., logging)
├── /configs                # Application configuration
└── go.mod                  # Project dependencies
```

## 🔥 **Main Features**

### API:
- **Retrieve a random question**
- **Submit an answer**
- **Get the leaderboard**
- **Check player's position**

### CLI:
- **Start a new quiz**
- **Answer questions**
- **View the leaderboard**
- **Check player's position**

## 🚀 **How to Run**

1. Clone the repository:
   ```sh
   git clone https://github.com/CMedrado/quiz.git
   cd quiz
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the API:
   ```sh
   go docker-compose up
   ```

4. Run the CLI:
   ```sh
   go run cmd/cli/main.go quiz
   ```

## 📌 **API Documentation**

### 📍 Get a Question
**Endpoint:** `GET /question/?user={user_id}`

#### Request Example:
```sh
curl -X GET "http://localhost:8080/question"\
     -H "Content-Type: application/json" \
     -d '{"user_id": 12345}'
```

#### Response:
```json
{
  "number_question": 1,
  "question": "What is the capital city of France?",
  "answers": ["Paris", "Berlin", "Madrid"]
}
```

---
### 📍 Submit an Answer
**Endpoint:** `POST /question/?user={user_id}`

#### Request Example:
```sh
curl -X POST "http://localhost:8080/question/?user=12345" \
     -H "Content-Type: application/json" \
     -d '{"number_question": 1, "answer": 0}'
```

#### Response:
```json
{
  "is_correct": true
}
```

---
### 📍 Get Player Position
**Endpoint:** `GET /rank/player/?user={user_id}`

#### Request Example:
```sh
curl -X GET "http://localhost:8080/rank/player/?user=12345"
```

#### Response:
```json
{
  "player": {
    "user_id": "12345",
    "position": 3,
    "score": 8
  }
}
```

---
### 📍 Get Leaderboard
**Endpoint:** `GET /rank/`

#### Request Example:
```sh
curl -X GET "http://localhost:8080/rank/"
```

#### Response:
```json
{
  "players": [
    {"user_id": "67890", "position": 1, "score": 10},
    {"user_id": "12345", "position": 3, "score": 8}
  ]
}
```

---

## **Swagger API Documentation**
Once the API is running, you can view the interactive Swagger documentation at:

**URL:** `http://localhost:8080/swagger`

This interface allows you to test the API directly and view the details of each endpoint, request, and response.

---

## 📊 **Logging and Monitoring**
- **Zap Logger** is used for structured and efficient logging.
- Standardized HTTP response error handling.
- Potential to integrate **distributed tracing** for enhanced visibility in production.

## 📌 **Potential Future Improvements**
- Implement **distributed tracing** for better monitoring.
- Further separation between logging and error handling.
- Add **user authentication** for player identification.
- Develop a **web interface** for the quiz.

---

## **Contributions**
Feel free to fork this project, create issues, or submit pull requests. Any suggestions and contributions are welcome!
