FROM node:16-alpine as builder

WORKDIR /app

COPY ./package.json ./yarn.lock ./
# RUN npm install --force && yarn install && yarn cache clean
RUN yarn install && yarn cache clean

COPY ./ ./
RUN yarn build

FROM nginx:1.21-alpine
#FROM nginx:latest

#RUN apk add --no-cache curl
RUN apk update && apk upgrade && apk add bash && apk add --no-cache curl && apk add lsof

COPY docker/production/nginx/conf.d /etc/nginx/conf.d

WORKDIR /app

COPY --from=builder /app/build ./

HEALTHCHECK --interval=5s --timeout=3s --start-period=1s CMD curl --fail http://127.0.0.1/health || exit 1
