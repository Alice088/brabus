.PHONY: build run stop

# Собирает оба приложения
build:
	cd app1 && go build -o ../bin/app1
	cd app2 && go build -o ../bin/app2

# Запускает Docker Compose и приложения
run:
    docker-compose up -d  # Запуск фоновых сервисов
    ./bin/app1 &          # Запуск в фоне
    ./bin/app2 &          # Запуск в фоне
    echo "Apps running!"

# Останавливает всё
stop:
    pkill -f "app1|app2"  # Убивает процессы
    docker-compose down    # Останавливает контейнеры