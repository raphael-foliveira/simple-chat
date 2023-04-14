FROM node:18.15-alpine

WORKDIR /app

COPY package.json .

RUN yarn

COPY . .

EXPOSE 5173
