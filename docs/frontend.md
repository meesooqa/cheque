## Initial
```
frontend/
├── package.json
├── package-lock.json
├── public/
├── src/
└── ...
```

```bash
cd frontend
npm init -y
npm install react react-dom
npm install --save-dev @babel/core @babel/preset-env @babel/preset-react babel-loader webpack webpack-cli webpack-dev-server
```

## Dev
### Запуск в dev-режиме:
```bash
docker build -t cheque03-react-dev -f Dockerfile.dev .
docker run -p 3000:3000 -v ${PWD}:/app -v /app/node_modules cheque03-react-dev
```

## Production
### Сборка образа:
```bash
docker build -t cheque03-react-app .
```
### Запуск контейнера:
```bash
docker run -p 80:80 cheque03-react-app
```
### Итог:
Приложение доступно на http://localhost
Статические файлы обслуживаются через nginx

## Пример полного цикла:
### Разработка:
```bash
docker-compose up --build
```
* Frontend: http://localhost:3000
* Backend: http://localhost:8080

### Production сборка:
```bash
docker build -t my-app-prod .
docker run -p 80:80 my-app-prod
```
Теперь ваше React-приложение будет работать в Docker, интегрируясь с Go-бэкендом как в разработке, так и в production.