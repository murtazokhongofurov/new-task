# Task

## 1. How to install:
```bash
https://github.com/murtazokhongofurov/new-task.git
```

## 2. Add new .env file using env.example
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=username
DB_PASS=password
DB_NAME=dbname
LOW_LEVEL=debug
HTTP_PORT=8080
```

## 3. Create tables:
```bash
make migrate_up
```

## 4. Run task project 
```bash
make run
```
<br>or
<br>

```bash
go run cmd/main.go
``` 

## 5. Open Swagger docs: 
url: http://localhost:8080/v1/docs/index.html