# hexlet-go
## Запуск hello.go
```bash
go run main.go
```
## Компиляция и запуск файлов
```bash
go build main.go
./hello
# устанавливаем пакеты
```
## Компиляция и запуск файлов
```bash
GO111MODULE=off go run .
# Компиляция и запуск файлов
```
## Установка пакета
```bash
go get github.com/fatih/color
# Установка пакета
```
## Установка последней версией пакета
```bash
go get github.com/fatih/color@latest
# Установка последней версией пакета
```
## Обновление списка зависимостей
```bash
go mod tidy
# Обновление списка зависимостей
```
## Удаление зависимости
```bash
# достаточно удалить импорт этой зависимости из кода и запустить
go mod tidy
```
## Линтер GO
```bash
# Линтер GO
golangci-lint run
```
