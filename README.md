# Go API

## Introduction
---
This is a small project to finalize my Golang course. This API handles events and users.
Users can create events and events are asociated with a user. Only event owners 
can delete the event and to enable security the application uses JWT for 
authentication. Last but not least, as DB driver the application works with etcd.

## Usage

First start the etcd container
```bash
docker compose up -d
```

Then get the packages for Go
```bash
go mod tidy
```

then you can start working with the endpoints.

## API Documentation

This documentation provides a detailed overview of the API endpoints defined in the Postman collection. 
The API consists of two main sections: **Events** and **Users**. 
Each section contains various endpoints for performing CRUD (Create, Read, Update, Delete) operations.

---

### Base URL
The base URL for all API requests is:

```
http://{{server}}:{{port}}
```

- `{{server}}`: Typically set to `localhost`.
- `{{port}}`: Typically set to `8080`.

---

## Events

### 1. Add New Event
**Endpoint**: `POST /events`

#### Description:
Adds a new event to the system.

#### Headers:
- `Content-Type`: application/json
- `Authorization`: Bearer `<JWT_TOKEN>`

#### Request Body:
```json
{
    "name": "event name",
    "location": "event location"
}
```

#### Example Request:
```bash
curl -X POST http://localhost:8080/events \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
-d '{
    "name": "event name",
    "location": "event location"
}'
```

---

### 2. Get All Events
**Endpoint**: `GET /events`

#### Description:
Retrieves a list of all events.

#### Example Request:
```bash
curl -X GET http://localhost:8080/events
```

---

### 3. Get Event by ID
**Endpoint**: `GET /events/{id}`

#### Description:
Retrieves a specific event by its ID.

#### Path Parameters:
- `id`: The unique uuid identifier of the event (e.g., `62aca175-8945-4221-a2b3-52e6889eee23`).

#### Example Request:
```bash
curl -X GET http://localhost:8080/events/230591
```

---

### 4. Update Event
**Endpoint**: `PUT /events/{id}`

#### Description:
Updates an existing event.

#### Path Parameters:
- `id`: The unique identifier of the event (e.g., `62aca175-8945-4221-a2b3-52e6889eee23`).

#### Headers:
- `Content-Type`: application/json

#### Request Body:
```json
{
    "name": "event name",
    "location": "event location"
}
```

#### Example Request:
```bash
curl -X PUT http://localhost:8080/events/230590 \
-H "Content-Type: application/json" \
-d '{
    "name": "event name",
    "location": "event location"
}'
```

---

### 5. Delete Event
**Endpoint**: `DELETE /events/{id}`

#### Description:
Deletes a specific event by its ID.

#### Path Parameters:
- `id`: The unique identifier of the event (e.g., `f0f2ca84-b788-413c-85f5-714443c42379`).

#### Headers:
- `Content-Type`: application/json
- `Authorization`: Bearer `<JWT_TOKEN>`

#### Example Request:
```bash
curl -X DELETE http://localhost:8080/events/f0f2ca84-b788-413c-85f5-714443c42379 \
-H "Content-Type: application/json" \
-H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

---

## Users

### 1. Add New User
**Endpoint**: `POST /signup`

#### Description:
Registers a new user in the system.

#### Headers:
- `Content-Type`: application/json

#### Request Body:
```json
{
    "email": "user@email.com",
    "password": "password123"
}
```

#### Example Request:
```bash
curl -X POST http://localhost:8080/signup \
-H "Content-Type: application/json" \
-d '{
    "email": "user@email.com",
    "password": "password123"
}'
```

---

### 2. Login
**Endpoint**: `POST /login`

#### Description:
Authenticates a user and returns a JWT token.

#### Headers:
- `Content-Type`: application/json

#### Request Body:
```json
{
    "email": "user@email.com",
    "password": "password123"
}
```

#### Example Request:
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@email.com",
    "password": "password"
}'
```

---

## Notes
- Replace placeholders like `{{server}}`, `{{port}}`, and `<JWT_TOKEN>` with actual values when making requests.
- Ensure that the JWT token used in authorized requests is valid and not expired.
- The API uses JSON Web Tokens (JWT) for authentication, which are added to the `Authorization` header as a Bearer token.

---

This documentation provides a clear and concise overview of the API's structure and usage. Developers can refer to this guide to understand how to interact with the API effectively.
