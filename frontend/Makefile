
react-clear:
	docker run --rm -v ${PWD}/frontend:/app -w /app alpine sh -c 'rm -rf .ready build'

react-init: frontend-yarn-install

react-yarn-install:
	docker-compose run --rm node-cli yarn install

react-yarn-upgrade:
	docker-compose run --rm node-cli yarn upgrade

react-ready:
	docker run --rm -v ${PWD}/frontend:/app -w /app alpine touch .ready

react-check: react-lint react-test

# error in stylelint
react-lint:
	docker-compose run --rm node-cli yarn eslint
	docker-compose run --rm node-cli yarn stylelint 

# [DEP0148] DeprecationWarning: Use of deprecated folder mapping "./2020/" in the "exports" field module resolution of the packag ...
react-eslint-fix:
	docker-compose run --rm node-cli yarn eslint-fix

react-pretty:
	docker-compose run --rm node-cli yarn prettier --write .

react-test:
	docker-compose run --rm node-cli yarn test --watchAll=false

react-test-watch:
	docker-compose run --rm node-cli yarn test



# ========================== BUILD DEVELOPMENT DOCKER IMAGES 
dev-build-gateway:
	docker --log-level=debug build --pull --file=gateway/docker/development/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-gateway:${IMAGE_TAG} gateway

dev-build-frontend:
	docker --log-level=debug build --pull --file=frontend/docker/development/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-frontend:${IMAGE_TAG} frontend

dev-build-frontend-node:
	docker --log-level=debug build --pull --file=frontend/docker/development/node/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-frontend-node:${IMAGE_TAG} frontend

dev-build-api:
	docker --log-level=debug build --pull --file=api/docker/development/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-api:${IMAGE_TAG} api

dev-build-db:
	docker --log-level=debug build --pull --file=db/docker/development/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-db:${IMAGE_TAG} db

try-dev-build: dev-build-gateway dev-build-frontend dev-build-frontend-node dev-build-api dev-build-db

try-dev-deploy:
	docker stack deploy --compose-file docker-compose-swarm.yml video --with-registry-auth --prune


# ========================== BUILD PROD DOCKER IMAGES 
build-prod-gateway:
	docker --log-level=debug build --pull --file=gateway/docker/production/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-gateway:${IMAGE_TAG} gateway
	docker --log-level=debug build --pull --file=gateway/docker/production/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-gateway:latest gateway

build-prod-frontend:
	docker --log-level=debug build --pull --file=frontend/docker/production/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-frontend:${IMAGE_TAG} frontend
	docker --log-level=debug build --pull --file=frontend/docker/production/nginx/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-frontend:latest frontend

build-prod-api:
	docker --log-level=debug build --pull --file=api/docker/production/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-api:${IMAGE_TAG} api
	docker --log-level=debug build --pull --file=api/docker/production/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-api:latest api

build-prod-db:
	docker --log-level=debug build --pull --file=db/docker/production/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-db:${IMAGE_TAG} db
	docker --log-level=debug build --pull --file=db/docker/production/Dockerfile --tag=${REGISTRY}/${LOGIN}/video-db:latest db

build-prod: build-prod-gateway build-prod-frontend build-prod-api  build-prod-db

try-build-prod:
	make build-prod


# ========================== DOCKER IMAGES PUSH
push-prod-gateway:
	docker push ${REGISTRY}/${LOGIN}/video-gateway:${IMAGE_TAG}
	docker push ${REGISTRY}/${LOGIN}/video-gateway:latest

push-prod-frontend:
	docker push ${REGISTRY}/${LOGIN}/video-frontend:${IMAGE_TAG}
	docker push ${REGISTRY}/${LOGIN}/video-frontend:latest

push-prod-api:
	docker push ${REGISTRY}/${LOGIN}/video-api:${IMAGE_TAG}
	docker push ${REGISTRY}/${LOGIN}/video-api:latest

push-prod-db:
	docker push ${REGISTRY}/${LOGIN}/video-db:${IMAGE_TAG}
	docker push ${REGISTRY}/${LOGIN}/video-db:latest

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
#	scp -o StrictHostKeyChecking=no -P ${PORT} ${JWT_PUBLIC_KEY} deploy@${HOST}:site_${IMAGE_TAG}/secrets/jwt_public.key
#	scp -o StrictHostKeyChecking=no -P ${PORT} ${JWT_PRIVATE_KEY} deploy@${HOST}:site_${IMAGE_TAG}/secrets/jwt_private.key

	ssh -o StrictHostKeyChecking=no deploy@${HOST} -p ${PORT} 'cd site_${IMAGE_TAG} && docker stack deploy --compose-file docker-compose.yml video --with-registry-auth --prune'
