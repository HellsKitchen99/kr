rebuild:
	docker build --no-cache -t kr-app -f ./app/Dockerfile .

another:
	docker run --rm --memory="256m" --env-file .env -p 2222:1111 kr-app