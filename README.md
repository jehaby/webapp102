# Разворачивание

## 1. Установить нужные штуки 

- Install docker https://www.docker.com/community-edition
- Install docker-compose https://docs.docker.com/compose/
- Install "node": ">= 4.0.0", "npm": ">= 3.0.0" (чем свежее тем лучше)

## 2. Зависимости фронтенда

    npm install # (из директории "app")
    
## 3. Запуск фронтенда    

    npm run dev # (там же)

## 4. Запуск бэкенда
    
    docker-compose -f var/docker/dev/docker-compose.yml up -d    

