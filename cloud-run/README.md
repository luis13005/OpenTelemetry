# CEP Clima — Cloud Run

API que recebe um CEP brasileiro e retorna a temperatura atual em Celsius, Fahrenheit e Kelvin.

## Deploy

[https://cloud-run-goexpert-463445218085.us-central1.run.app](https://cloud-run-goexpert-463445218085.us-central1.run.app)

## Rodando localmente com Docker

### Build da imagem

```bash
docker build -t cep-clima .
```

### Subir a aplicação

```bash
docker run -p 8080:8080 cep-clima
```

A API estará disponível em `http://localhost:8080`.

### Exemplo de uso

**Deploy em Cloud Run:**
```bash
curl.exe -X POST https://cloud-run-goexpert-463445218085.us-central1.run.app/cep -H "Content-Type: application/json" -d '{\"cep\":\"01310100\"}'
```
https://cloud-run-goexpert-xxpoxwxwra-uc.a.run.app/cep
**Local:**
```powershell
curl.exe -X POST http://localhost:8080/cep -H "Content-Type: application/json" -d '{\"cep\":\"01310100\"}'
```

## Rodando os testes com Docker

```bash
docker build --target builder -t cep-clima-test .
docker run --rm --entrypoint go cep-clima-test test ./...
```
