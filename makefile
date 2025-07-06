.PHONY: build run stop

run:
	docker-compose up -d
	@echo "Starting services..."
	@./bin/brabus > /dev/null 2>&1 &
	@./bin/banana &
	@echo "Brabus running! (app2 logs visible, app1 runs silently)"

# Останавливает всё
stop:
		pkill -f "brabus|banana"  # Убивает процессы
		docker-compose down    # Останавливает контейнеры