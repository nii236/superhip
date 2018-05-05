```
docker run -d -p 5432:5432 --name devdb -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=dev -e POSTGRES_DB=dev postgres:10.3-alpine
```