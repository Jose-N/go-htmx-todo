tailwind-watch:
	./tailwindcss -i ./internal/static/css/input.css -o ./internal/static/css/style.css --watch
.PHONY: tailwind-watch

tailwind-build:
	./tailwindcss -i ./internal/static/css/input.css -o ./internal/static/css/style.css --minify
.PHONY: tailwind-build

templ-watch:
	templ generate --watch
.PHONY: templ-watch

templ-build:
	templ generate
.PHONY: templ-build

build:
	make tailwind-build && make templ-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go
.PHONY: build
