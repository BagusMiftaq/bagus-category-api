# 🚀 Category REST API (Golang)

Assignment 1 – Implementation of a simple 'Category API' using Golang."

---

## 🧰 Tech Stack

- **Golang**
- **net/http** (Standard Library)
- JSON-based REST API
- Deployed on **Leapcell**

---

## 🌍 Base URLs

| Environment | URL                                         |
| ----------- | ------------------------------------------- |
| Local       | http://localhost:8080                       |
| Production  | https://bagus-category-api.apn.leapcell.app |

> ⚠️ Production must use **HTTPS** to prevent HTTP redirects from changing request methods (POST/PUT/DELETE → GET).

---

## ▶️ Run Project Locally

```bash
git clone https://github.com/BagusMiftaq/bagus-category-api.git
cd bagus-category-api
go run main.go
```

Server will run at:

```
http://localhost:8080
```

---

## 📦 API Endpoints

### 🔹 Get All Categories

**GET** `/categories`

#### Response

```json
{
  "status": 200,
  "message": "Succes Get All Data",
  "result": [
    {
      "id": 1,
      "name": "Ensiklopedia Bahasa Semut",
      "desc": "Mystical Science"
    },
    {
      "id": 2,
      "name": "99 Dongeng Sebelum Kerja",
      "desc": "Myth"
    },
    {
      "id": 3,
      "name": "Kisi-Kisi Pertanyaan Alam Kubur",
      "desc": "Spiritual"
    },
    {
      "id": 4,
      "name": "Merantau ke Framework Seberang",
      "desc": "Adventure"
    }
  ]
}
```

---

### 🔹 Get Category By ID

**GET** `/categories/{id}`

#### Example

```
GET /categories/1
```

---

### 🔹 Create Category

**POST** `/categories`

#### Body

```json
{
  "name": "Menjadi Keluarga Gokil",
  "desc": "Mystical Problem"
}
```

---

### 🔹 Update Category

**PUT** `/categories/{id}`

#### Body

```json
{
  "name": "Ensiklopedia Bahasa Kijang",
  "desc": "Mystical Science"
}
```

---

### 🔹 Delete Category

**DELETE** `/categories/{id}`

---

## 📄 Standard API Response Format

All responses follow a consistent JSON structure:

```json
{
  "status": 200,
  "message": "Success",
  "result": {}
}
```

| Field   | Type           | Description                    |
| ------- | -------------- | ------------------------------ |
| status  | number         | HTTP status code               |
| message | string         | Informational response message |
| result  | object / array | Response data (optional)       |

---

## ❌ Error Response Example

```json
{
  "status": 400,
  "message": "Invalid Category ID"
}
```

---

## 🔐 Validation Rules

- Unknown JSON fields are rejected
- Requests must use `Content-Type: application/json`

---

## 👨‍💻 Author

**Bagus Miftaq**

---
