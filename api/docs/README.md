# Настройка swagger

## 1. Документируем методы (пример)

([официальный репозиторий](https://github.com/swaggo/swag))
```go
// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:80
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	...
}
```
```go
// @Summary Sign up                                 - краткое описание метода
// @Security ApiKeyAuth                             - признак того, что для доступа нужен ключ авторизации (токен). Используется везде6 гду нужна авторизация
// @Tags Auth                                       - записываем с большой буквы, по тегу будут группироваться хендлеры
// @Description Creating user account               - полное описание, показывается при открытии блока

// @ID signUp                                       - id, с маленькой буквы, дублируем название метода
// @Accept  json                                    - формат отправки данных в запросе
// @Produce  json                                   - формат ответа

// @Param id path integer true "List id"            - параметры, отправляемые в пути {id}
// @Param input body todo.User true "account info"  - параметры, отправляемые в теле
// @Param input body todo.TodoList true "List create information"

// @Success 200 {string} string "id"                 - параметры возвр. успешного ответа "id": 1
// @Success 200 {object} todo.TodoList

// @Failure 400,404 {object} errorResponse          - объект, возвращаемый при ошибке
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse

// @Router /auth/sign-up [post]                     - путь к хендлеру (нюансы со "/" на конце)
// @Router /api/lists/{id}/items/ [post]
// @Router /api/lists/ [post]
// @Router /api/lists/{id} [get]
// @Router /api/lists/{id} [put]
// @Router /api/lists/{id} [delete]
func (h *Handler) signUp(c *gin.Context) {
	...
}
```

## 2. Генерируем папку с документацией  

Скачиваем пакет:  
```
go get -u github.com/swaggo/swag/cmd/swag
```
запускаем инициализацию:
```
/Library/go/go1.16.5/bin/bin/swag init -g ./cmd/main.go
```
в корне должна быдет появиться директория 'docs'   

## 3. Открываем документацию в браузере

В главном хендлере подключаем модули и добавляем адрес для открытия документации в браузере:
```go
package handler

import (
	...
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/p-12s/todo-list-rest-api/docs"
)
...

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	...
}
```
Документация должна будет открыться по адресу:  
http://localhost:80/swagger/index.html  

Создаем пользователя, получаем TOKEN авторизации, вставляем его в авторизационный попап в формате:  
```
Bearer TOKEN
```
Теперь запросы должны выполняться в браузере.  
