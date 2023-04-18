#!/bin/bash
REPO_NAME=todo-service

echo "⚒️ ⚙️ Build and tag image"
docker build -t ${REPO_NAME}  ../../
docker tag ${REPO_NAME} ${REPO_NAME}:latest

echo "⚒️ Docker compose down"
docker-compose -f docker-compose.yaml -p ${REPO_NAME} down;
echo "⚒️️ Docker compose up"
docker-compose -f docker-compose.yaml -p ${REPO_NAME} up -d;

sleep 15
npx newman run TodoApp.postman_collection.json --folder "e2e"
