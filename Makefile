.PHONY: app
app:
	cd ./app && GOOS=js GOARCH=wasm go build -o ../public/static/app.wasm

.PHONY: run
run:
	go run .
