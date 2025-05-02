# YandexGo\_proj

## 📌 Описание

**YandexGo\_proj** — это сервис для заказа такси, аналогичный Яндекс Go. Он предоставляет пользователям следующие возможности:

* Функционал управления заказом такси
* Регистрация и авторизация пользователей и водителей
* Просмотр истории последних поездок
* Уведомлений о статусе текущего заказа

## 🛠️ Стек технологий

* **Frontend:** Kotlin
* **Backend:** Python, Go
* **Инфраструктура и обмен сообщениями:** Redis, Kafka
* **Контейнеризация и оркестрация:** Docker, Kubernetes
* **Взаимодействие между сервисами:** gRPC, API Gateway, restAPI
* **Документация API:** OpenAPI (Swagger)

## 🧩 Архитектура

````plaintext
[Driver App (Kotlin)]      [User App (Kotlin)]
          \                      /
           \                    /
           [      API Gateway     ]
                   |
         [Backend Services: Go, Python]
               |               |
             [gRPC]        [Redis, Kafka]
```plaintext
[Mobile App (Kotlin)] <--> [API Gateway] <--> [Backend Services: Go, Python]
                                      |          |
                                  [gRPC]      [Redis, Kafka]
````

## 📄 Документация

Swagger UI доступен по адресу:

```plaintext
http://localhost:8080/swagger
```

## 👥 Авторы и контакты

