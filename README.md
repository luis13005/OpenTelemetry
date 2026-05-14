# CEP Clima — OpenTelemetry + Zipkin

## Pré-requisitos

- Docker e Docker Compose
- Chave gratuita do [WeatherAPI](https://www.weatherapi.com/)

## Build e execução

Crie um arquivo `.env` na mesma pasta do `docker-compose.yaml`:

```
WEATHER_API=sua_chave_aqui
```

Suba os serviços:

```bash
docker compose up --build
```

## Requisição

```bash
curl -X POST http://localhost:8081/cep \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}'
```

**Resposta:**

```json
{
  "city": "São Paulo",
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

## Traces

Acesse [http://localhost:9411](http://localhost:9411) para visualizar o rastreamento distribuído no Zipkin.
