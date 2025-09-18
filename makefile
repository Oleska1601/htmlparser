goose -dir ./internal/database/repo/migrations -s create new_parsings sql  
swag init -g /internal/app/app.go
docker build -t htmlparser -f deploy/Dockerfile . 
docker-compose -f ./deploy/docker-compose.yml up -d