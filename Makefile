rabbitmq:
	docker run -d \
		--name rabbitmq \
		-p 5672:5672 \
		-p 15672:15672 \
		-e RABBITMQ_DEFAULT_USER=rabbitmq \
		-e RABBITMQ_DEFAULT_PASS=rabbitmq \
		rabbitmq:3.8.17-management
