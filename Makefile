run:
	@templ generate 
	@npx tailwindcss -i ./style.css -o ./dist/tailwind.css 
	@go build -o ./tmp/main.exe .

gen:
	@templ generate 
	@npx tailwindcss -i ./style.css -o ./dist/tailwind.css 