watch:
	npx tailwindcss -i ./static/src/main.css -o ./static/output/main.css --watch

run:
	npx tailwindcss -i ./static/src/main.css -o ./static/output/main.css
	go run .