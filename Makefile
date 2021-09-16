# Go - директивы работы с папкой API
go-build:
	cd api && go mod download && go build cmd/main.go && rm main && cd ..

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

create-migration:
	migrate create -ext sql -dir ./schema -seq init

test.coverage:
	go tool cover -func=cover.out

swag:
	swag init -g cmd/api/cmd/main.go

lint:
	golangci-lint run

recompose:
	docker-compose down --remove-orphans
	docker rmi video_api:latest
	docker-compose up -d

# ========================== BUILD DEVELOPMENT DOCKER IMAGES 
dev-build-gateway:
	docker --log-level=debug build --pull --file=gateway/docker/development/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:gateway-${IMAGE_TAG} gateway

dev-build-frontend:
	docker --log-level=debug build --pull --file=frontend/docker/development/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:frontend-${IMAGE_TAG} frontend

dev-build-api:
	docker --log-level=debug build --pull --file=api/docker/development/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:api-${IMAGE_TAG} api

dev-build-db:
	docker --log-level=debug build --pull --file=db/docker/development/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:db-${IMAGE_TAG} db

try-dev-build: dev-build-gateway dev-build-frontend dev-build-api dev-build-db

try-dev-deploy:
	docker stack deploy --compose-file docker-compose-swarm.yml video --with-registry-auth --prune


# ========================== BUILD PROD DOCKER IMAGES 
build-prod-gateway:
	docker --log-level=debug build --pull --file=gateway/docker/production/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:gateway-${IMAGE_TAG} gateway
	docker --log-level=debug build --pull --file=gateway/docker/production/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:gateway-latest gateway

build-prod-frontend:
	docker --log-level=debug build --pull --file=frontend/docker/production/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:frontend-${IMAGE_TAG} frontend
	docker --log-level=debug build --pull --file=frontend/docker/production/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:frontend-latest frontend

build-prod-api:
	docker --log-level=debug build --pull --file=api/docker/production/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:api-${IMAGE_TAG} api
	docker --log-level=debug build --pull --file=api/docker/production/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:api-latest api

build-prod-db:
	docker --log-level=debug build --pull --file=db/docker/production/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:db-${IMAGE_TAG} db
	docker --log-level=debug build --pull --file=db/docker/production/Dockerfile --tag=${REGISTRY}/${LOGIN}/${PROJECT}:db-latest db

build-prod: build-prod-gateway build-prod-frontend build-prod-api  build-prod-db

try-build-prod:
	make build-prod


# ========================== DOCKER IMAGES PUSH
push-prod-gateway:
	docker push ${REGISTRY}/${LOGIN}/${PROJECT}:gateway-${IMAGE_TAG}
	docker push ${REGISTRY}/${LOGIN}/${PROJECT}:gateway-latest

push-prod-frontend:
	docker push ${REGISTRY}/${LOGIN}/${PROJECT}:frontend-${IMAGE_TAG}
	docker push ${REGISTRY}/${LOGIN}/${PROJECT}:frontend-latest

push-prod-api:
	docker push ${REGISTRY}/${LOGIN}/${PROJECT}:api-${IMAGE_TAG}
	docker push ${REGISTRY}/${LOGIN}/${PROJECT}:api-latest

push-prod-db:
	docker push ${REGISTRY}/${LOGIN}/${PROJECT}:db-${IMAGE_TAG}
	docker push ${REGISTRY}/${LOGIN}/${PROJECT}:db-latest

push-prod: push-prod-gateway push-prod-frontend push-prod-api push-prod-db

try-push-prod:
	make push-prod


# ========================== DEPLOY PROD
deploy-prod:
	ssh -o StrictHostKeyChecking=no deploy@${HOST} -p ${PORT} 'rm -rf site_${IMAGE_TAG}'
	ssh -o StrictHostKeyChecking=no deploy@${HOST} -p ${PORT} 'mkdir site_${IMAGE_TAG}'

	envsubst < docker-compose-production.yml > docker-compose-production-env.yml
	scp -o StrictHostKeyChecking=no -P ${PORT} docker-compose-production-env.yml deploy@${HOST}:site_${IMAGE_TAG}/docker-compose.yml
	rm -f docker-compose-production-env.yml

	ssh -o StrictHostKeyChecking=no deploy@${HOST} -p ${PORT} 'mkdir site_${IMAGE_TAG}/secrets'
	scp -o StrictHostKeyChecking=no -P ${PORT} ./secrets/production/postgres_db deploy@${HOST}:site_${IMAGE_TAG}/secrets/postgres_db
	scp -o StrictHostKeyChecking=no -P ${PORT} ./secrets/production/postgres_host deploy@${HOST}:site_${IMAGE_TAG}/secrets/postgres_host
	scp -o StrictHostKeyChecking=no -P ${PORT} ./secrets/production/postgres_password deploy@${HOST}:site_${IMAGE_TAG}/secrets/postgres_password
	scp -o StrictHostKeyChecking=no -P ${PORT} ./secrets/production/postgres_user deploy@${HOST}:site_${IMAGE_TAG}/secrets/postgres_user
	scp -o StrictHostKeyChecking=no -P ${PORT} ./secrets/production/sentry_dsn deploy@${HOST}:site_${IMAGE_TAG}/secrets/sentry_dsn

	ssh -o StrictHostKeyChecking=no deploy@${HOST} -p ${PORT} 'cd site_${IMAGE_TAG} && docker stack deploy --compose-file docker-compose.yml video --with-registry-auth --prune'

deploy-prod-clean:
	rm -f docker-compose-production-env.yml

rollback-prod:
	ssh -o StrictHostKeyChecking=no deploy@${HOST} -p ${PORT} 'cd site_${IMAGE_TAG} && docker stack deploy --compose-file docker-compose.yml video --with-registry-auth --prune'

# ==========================

docker-rmi:
	docker rmi ${REGISTRY}/${LOGIN}/${PROJECT}:api-${IMAGE_TAG} ${REGISTRY}/${LOGIN}/${PROJECT}:gateway-${IMAGE_TAG} ${REGISTRY}/${LOGIN}/${PROJECT}:frontend-${IMAGE_TAG}
