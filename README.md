# YandexGo\_proj

## 📌 Описание

**YandexGo_proj** — это сервис для заказа такси, аналогичный Яндекс Go. Он предоставляет пользователям следующие возможности:

* Функционал управления заказом такси
* Регистрация и авторизация пользователей и водителей
* Просмотр истории последних поездок
* Просмотр статуса текущего заказа

## 🛠️ Стек технологий

* **Frontend:** Kotlin
* **Backend:** Python, Go
* **База данных:** PostgreSQL
* **Инфраструктура и межсервисное взаимодействие:** Redis, Kafka, gRPC, REST API, API Gateway
* **Контейнеризация:** Docker, Docker Compose
* **Документация API:** OpenAPI (Swagger)
* **Метрики:** Grafana, Prometheus

## 🧩 Архитектура

Проект построен на микросервисной архитектуре. Пользователь и водитель взаимодействуют с разными сервисами через мобильные приложения, при этом данные хранятся в отдельных PostgreSQL-базах. Координация заказов происходит через Kafka и Redis.
```plaintext
                     +------------------------------+      +------------------------------+
                     |  📱 Mobile-App (User)        |      |  📱 Mobile-App (Driver)      |
                     |  — локально: JWT, история    |      |  — локально: JWT             |
                     +-----------+------------------+      +-----------+------------------+
                                 |                                     |
                     Авторизация, операции с заказом      Авторизация, GPS, операции с заказами
                                 |                                     |
                     +-----------v------------------+      +-----------v------------------+
                     |     🟦 User-Service          |      |     🟩 Driver-Service       |
                     |  REST / gRPC / Kafka ⬆       |      |  REST / gRPC / Kafka ⬇       |
                     |  Создание заказов            |      |  Принятие заказов            |
                     |  → Kafka: topic "orders"     |      |  → Redis: активные водители  |
                     +-----------+------------------+      |  → Redis: активные заказы    |
                                 |                          +-----------+------------------+
                                 |                                      |
         +-----------------------+---------------------------------------+-------+    
         |                                                                       |
+--------v------------------+                                  +-----------------v------------------+
|  PostgreSQL Users         |                                  |  PostgreSQL Drivers                |
|  ID, ФИО, email, телефон, |                                  |  ID, ФИО, email, паспорт, авто,    |
+---------------------------+                                  |  права                             |
                                                               +------------------------------------+


                            +---------------------------+
                            | 🔐 Auth-Service           |
                            | JWT проверка / OAuth      |
                            +---------------------------+

                            +---------------------------+
                            |         Kafka             |
                            |     Topic: "orders"       |
                            +------------+--------------+
                                         |
     Пользователь создаёт заказ → Kafka  | Водитель принимает заказ 
                                         v
                          +--------------+---------------+
                          |            Redis             |
                          |  DB 0: active drivers        |
                          |     — ID, координаты         |
                          |  DB 1: active orders         |
                          |     — ID, статус             |
                          +--------------+---------------+
                                ▲              ▲
                                |              |
                       +--------+              +--------+
                       |                                |
               Driver-Service                  Driver-Service
               пишет координаты              пишет принятый заказ

```

---

## 🌐 API ручки (Admin Service)

**GET /admin/admin**
**Описание:** Доступ к администраторскому ресурсу. Требуется уровень доступа 69.
**Заголовок запроса:**

```json
Authorization: Bearer <JWT-токен>
```

**Ответ:**

```json
{
  "message": "Доступ к администраторскому ресурсу",
  "user": {
    "id": "user_1",
    "email": "admin@example.com",
    "level_access": 69
  }
}
```

## 🌐 API ручки (Auth Service)

### 1. **POST /auth/users/add**
**Описание:** Добавление нового пользователя.
**Тело запроса:**

```json
{
  "id": "string",
  "first_name": "Ivan",
  "last_name": "Petrov",
  "email": "ivan@example.com",
  "phone_number": "+79991234567",
  "level_access": 1,
  "password": "securepassword"
}

```

**Ответ:**

```json
{
  "message": "Пользователь успешно добавлен"
}
```

### 2. **DELETE /auth/users/delete**
**Описание:** Удаление пользователя по ID.
Параметры запроса: id — ID пользователя (в query string)
**Пример запроса:**

```json
DELETE /auth/users/delete?id=user_1
```

**Ответ:**

```json
{
  "message": "Пользователь успешно удален"
}
```

### 3. **POST /auth/drivers/add**
**Описание:** Добавление нового водителя.
**Тело запроса:**

```json
{
  "id": "string", 
  "first_name": "Ivan",
  "last_name": "Ivanov",
  "email": "driver@example.com",
  "phone_number": "+79998887766",
  "level_access": 2,
  "password": "securepassword",
  "driver_license": "1234567890",
  "driver_license_date": "2020-01-01",
  "car_number": "A123BC77",
  "car_model": "Model S",
  "car_marks": "Tesla",
  "car_color": "black"
}
```

**Ответ:**

```json
{
  "message": "Водитель успешно добавлен"
}
```

### 4. **DELETE /auth/drivers/delete**
**Описание:** Удаление водителя по ID.
Параметры запроса: id — ID водителя (в query string)
**Пример запроса:**

```json
DELETE /auth/drivers/delete?id=driver_1
```

**Ответ:**

```json
{
  "message": "Водитель успешно удален"
}
```

## 🌐 API ручки (JWT Service)

### 1. **POST /JWT/token**
**Описание:** Получение токена авторизации.
**Форма запроса (тип application/x-www-form-urlencoded):**

```json
username=ivan@example.com&password=securepassword
```

**Ответ:**

```json
{
  "access_token": "<JWT-токен>",
  "token_type": "bearer"
}
```

### 2. **GET /JWT/users/me**
**Описание:** Получение данных текущего пользователя.
**Заголовок запроса:**

```json
Authorization: Bearer <JWT-токен>
```

**Ответ:**

```json
{
  "id": "user_1",
  "email": "ivan@example.com",
  "level_access": 1
}
```
## 🌐 API ручки (Order Service)

### 1. **POST /do_order/make_order**
**Описание:** Создание заказа такси и поиск ближайших водителей.
Параметры запроса (в query string): user_id, latitude, longitude
**Пример запроса:**

```json
POST /do_order/make_order?user_id=user_1&latitude=55.7558&longitude=37.6173
```

**Ответ:**

```json
{
  "message": "Заказ создан, ближайшие водители найдены"
}
```

## 🌐 API ручки (Driver Service)

### 1. **GET /driver/profile**

**Описание**: Получение профиля водителя.
**Тело запроса**:

```json
{
  "token": "JWT-токен"
}
```

**Ответ**:

```json
{
  "id": "driver_1",
  "username": "Ivan",
  "phone": "+79991234567",
  "car_model": "Model S",
  "email": "ivan@example.com",
  "car_color": "black",
  "car_mark": "Tesla",
  "car_number": "A123BC77"
}
```

---

### 2. **POST /ride/{id}/accept**

**Описание**: Водитель принимает заказ.
**Параметры пути**: `id` — ID поездки.
**Пример запроса**:

```json
{
  "id": "ride_123"
}
```

**Ответ**:

```json
{
  "status": true,
  "message": "Ride accepted"
}
```

---

### 3. **POST /ride/{id}/complete**

**Описание**: Завершение поездки.
**Параметры пути**: `id` — ID поездки.
**Пример запроса**:

```json
{
  "id": "ride_123"
}
```

**Ответ**:

```json
{
  "status": true,
  "message": "Ride accepted"
}
```

---

### 4. **POST /ride/{id}/cancel**

**Описание**: Отмена поездки водителем.
**Параметры пути**: `id` — ID поездки.
**Пример запроса**:

```json
{
  "id": "ride_123"
}
```

**Ответ**:

```json
{
  "status": true,
  "message": "Ride accepted"
}
```

---

### 5. **GET /driver/current\_ride/{id}**

**Описание**: Получение информации о текущей поездке водителя.
**Параметры пути**: `id` — ID водителя.

**Ответ**:

```json
{
  "id": "ride_123",
  "user_id": "user_1",
  "driver_id": "driver_1",
  "start_location": {
    "latitude": 55.75,
    "longitude": 37.61
  },
  "end_location": {
    "latitude": 55.80,
    "longitude": 37.65
  },
  "status": "in_progress",
  "timestamp": 1716740000
}
```

---

### 6. **POST /driver/location**

**Описание**: Отправка текущего местоположения водителя.
**Тело запроса**:

```json
{
  "driver_id": "driver_1",
  "location": {
    "latitude": 55.75,
    "longitude": 37.61
  }
}
```

**Ответ**:

```json
{
  "status": true,
  "message": "Location updated"
}
```

---

### 7. **GET /driver/nearby\_req**

**Описание**: Получение списка ближайших заказов.
**Ответ**:

```json
{
  "ride_requests": [
    {
      "user_id": "user_1",
      "start_location": {
        "latitude": 55.76,
        "longitude": 37.62
      },
      "end_location": {
        "latitude": 55.78,
        "longitude": 37.64
      }
    }
  ]
}
```

---

### 8. **GET /user/{id}**

**Описание**: Получение информации о пользователе.
**Параметры пути**: `id` — ID пользователя.
**Ответ**:

```json
{
  "id": "user_1",
  "username": "Petr",
  "phone": "+79999887766"
}
```

---

### 9. **GET /driver/{id}/rides**

**Описание**: Получение истории поездок водителя.
**Параметры пути**: `id` — ID водителя.
**Ответ**:

```json
{
  "rides": [
    {
      "id": "ride_123",
      "user_id": "user_1",
      "driver_id": "driver_1",
      "start_location": {
        "latitude": 55.75,
        "longitude": 37.61
      },
      "end_location": {
        "latitude": 55.80,
        "longitude": 37.65
      },
      "status": "completed",
      "timestamp": 1716730000
    }
  ]
}
```
---

---

## 🌐 API ручки (Client Service)

### 1. **GET /user/profile**

**Описание**: Получение профиля пользователя.
**Тело запроса**:

```json
{
  "token": "JWT-токен"
}
```

**Ответ**:

```json
{
  "id": "user_1",
  "username": "Ivan",
  "email": "ivan@example.com",
  "phone": "+79991234567"
}
```

---

### 2. **PUT /user/profile**

**Описание**: Обновление профиля пользователя.
**Тело запроса**:

```json
{
  "id": "user_1",
  "username": "Ivan",
  "email": "ivan@example.com",
  "phone": "+79991234567"
}
```

**Ответ**:

```json
{
  "id": "user_1",
  "username": "Ivan",
  "email": "ivan@example.com",
  "phone": "+79991234567"
}
```

---

### 3. **POST /ride/request**

**Описание**: Создание нового заказа.
**Тело запроса**:

```json
{
  "user_id": "user_1",
  "start_location": {
    "latitude": 55.751244,
    "longitude": 37.618423
  },
  "end_location": {
    "latitude": 55.760186,
    "longitude": 37.618711
  }
}
```

**Ответ**:

```json
{
  "id": "ride_123",
  "user_id": "user_1",
  "driver_id": "driver_9",
  "start_location": {...},
  "end_location": {...},
  "status": "pending",
  "timestamp": 1716653720
}
```

---

### 4. **POST /ride/{id}/cancel**

**Описание**: Отмена поездки пользователем.
**Параметры пути**: `id` — ID поездки.
**Пример запроса**:

```json
{
  "id": "ride_123"
}
```

**Ответ**:

```json
{
  "status": true,
  "message": "Ride canceled successfully"
}
```

---

### 5. **GET /ride/{id}**

**Описание**: Получение информации о текущей поездке.
**Параметры пути**: `id` — ID пользователя.
**Пример запроса**:

```json
{
  "id": "user_1"
}
```

**Ответ**:

```json
{
  "id": "ride_123",
  "user_id": "user_1",
  "driver_id": "driver_9",
  "start_location": {...},
  "end_location": {...},
  "status": "in_progress",
  "timestamp": 1716653720
}
```

---

### 6. **GET /ride/history**

**Описание**: Получение истории поездок пользователя.
**Тело запроса**:

```json
{
  "id": "user_1"
}
```

**Ответ**:

```json
{
  "rides": [
    {
      "id": "ride_123",
      "status": "completed",
      ...
    },
    {
      "id": "ride_122",
      "status": "canceled",
      ...
    }
  ]
}
```

---

### 7. **GET /driver/{id}/location**

**Описание**: Получение текущего местоположения водителя.
**Параметры пути**: `id` — ID водителя.
**Пример запроса**:

```json
{
  "id": "driver_9"
}
```

**Ответ**:

```json
{
  "latitude": 55.751244,
  "longitude": 37.618423
}
```

---

### 8. **GET /driver/{id}**

**Описание**: Получение информации о водителе.
**Параметры пути**: `id` — ID водителя.
**Пример запроса**:

```json
{
  "id": "driver_9"
}
```

**Ответ**:

```json
{
  "id": "driver_9",
  "username": "Ivan",
  "phone": "+79991234567",
  "car_model": "Model S",
  "car_make": "Tesla",
  "car_number": "A123BC77",
  "car_color": "black",
  "location": {
    "latitude": 55.751244,
    "longitude": 37.618423
  }
}
```

---

## 📄 Документация
Документация с UI доступна по адресу:

```plaintext
http://79.174.85.58:8080/docs
```

## 👥 Авторы и контакты
* Дылдин Сергей: backend: grpc +metrics (drivers, client services)
* Дробышев Егор: backend: авторизация, sql +metrics (auth service)
* Богданов Михаил: backend: admin, auth, jwt, order services
* Татульян Артём: frontend
* Стрельников Никита: frontend
