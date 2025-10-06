back

eduquest-backend/
├── services/ # Все микросервисы
│ ├── auth-service/
│ ├── user-service/
│ ├── course-service/
│ ├── progress-service/
│ ├── quest-service/
│ ├── shop-service/
│ ├── community-service/
│ ├── notification-service/
│ ├── analytics-service/
│ └── admin-service/
├── shared/ # Shared код между сервисами
│ ├── pkg/
│ │ ├── database/ # DB helpers
│ │ ├── redis/ # Redis client
│ │ ├── kafka/ # Kafka producer/consumer
│ │ ├── jwt/ # JWT utilities
│ │ ├── logger/ # Structured logging (slog)
│ │ ├── grpc_errors/ # gRPC error handling
│ │ ├── validator/ # Request validation
│ │ └── middleware/ # Shared middleware
│ └── proto/ # Protobuf definitions
│ ├── auth/
│ ├── user/
│ ├── course/
│ └── ...
├── infrastructure/ # Инфраструктура
│ ├── docker/
│ │ ├── postgres/
│ │ │ └── init-scripts/
│ │ ├── redis/
│ │ └── kafka/
│ ├── kubernetes/ # K8s manifests (позже)
│ └── scripts/
│ ├── migrate.sh
│ └── seed-data.sh
├── migrations/ # DB migrations для всех сервисов
│ ├── auth/
│ ├── user/
│ └── ...
├── api-gateway/ # API Gateway (опционально)
├── docker-compose.yml # Для локальной разработки
├── docker-compose.prod.yml
├── Makefile # Удобные команды
├── go.work # Go workspace
├── .env.example
└── README.md
