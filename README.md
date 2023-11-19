
# Golang clean architecture
# Запуск

```shell
    make up
```

---


Если не получается запустить...
```shell
    make help
```
Тест

```shell
    make test
```
```shell
    make lint
```
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -c ./build/.golangci.yml
