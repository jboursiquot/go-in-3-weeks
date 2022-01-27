default:
	@go run main.go&
	@sleep 1
	@for number in 1 2 3 4 5 6; do \
		curl -H "X-Request-Id: $(uuidgen)" localhost:8080 -i; echo; echo "---"; \
  done
