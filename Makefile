GOOS ?= $(shell go env GOOS)

.PHONY: frontend
frontend:
	GOOS=$(GOOS) CGO_ENABLED=0 go build -o _build/frontend -v ./cmd/frontend

.PHONY: backend
backend:
	GOOS=$(GOOS) CGO_ENABLED=0 go build -o _build/backend -v ./cmd/backend

.PHONY: build
build: frontend backend

.PHONY: clean
clean:
	rm -fr _build/
