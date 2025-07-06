.PHONY: build run stop

run:
	docker-compose up -d
	@echo "Starting services..."
	@./brabus > /dev/null 2>&1 &
	@./banana &
	@echo "Brabus running!"


stop:
		pkill -f "brabus|banana"
		docker-compose down