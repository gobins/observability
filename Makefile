.PHONY: build
build:
	@echo ">> building binary"
	docker build -t myapp:latest .

.PHONY: deploy
deploy:
	@echo ">> deploying app"
	kubectl apply -f deployment/

.PHONY: teardown
teardown:
	@echo ">> cleanup deployment"
	kubectl delete -f deployment/