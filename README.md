# CEP Clima — OpenTelemetry + Zipkin

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

## Traces

Acesse [http://localhost:9411](http://localhost:9411) para visualizar o rastreamento distribuído no Zipkin.
