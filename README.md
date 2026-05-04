# CEP Clima — Serviço A + Serviço B com OTEL + Zipkin

Sistema com dois serviços Go + rastreamento distribuído.

## Arquitetura

```
Cliente → Serviço A (:8081) → Serviço B (:8080) → ViaCEP / WeatherAPI
                ↘                    ↘
           OTEL Collector (:4318)
                    ↘
               Zipkin (:9411)
```

- **Serviço A**: valida o CEP e repassa ao Serviço B com propagação de trace
- **Serviço B**: consulta ViaCEP + WeatherAPI e retorna temperaturas
- **OTEL Collector**: recebe spans via OTLP HTTP e exporta para o Zipkin
- **Zipkin**: visualização do rastreamento distribuído

## Pré-requisitos

- Docker e Docker Compose instalados
- Chave de API do [WeatherAPI](https://www.weatherapi.com/) (gratuita)

## Subindo o ecossistema

```bash
# Na raiz do projeto (onde está o docker-compose.yaml)
WEATHER_API=sua_chave_aqui docker compose up --build
```

Ou crie um arquivo `.env` na raiz:

```
WEATHER_API=sua_chave_aqui
```

E então:

```bash
docker compose up --build
```

## Fazendo uma requisição ao Serviço A

```bash
curl -X POST http://localhost:8081/cep \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}'
```

### Resposta de sucesso (200)

```json
{
  "city": "São Paulo",
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### Erros

| HTTP | Mensagem              | Causa                         |
|------|-----------------------|-------------------------------|
| 422  | `"invalid zipcode"`   | CEP com formato inválido      |
| 404  | `"can not find zipcode"` | CEP não encontrado no ViaCEP |

## Visualizando os traces no Zipkin

1. Abra o navegador em [http://localhost:9411](http://localhost:9411)
2. Clique em **"Run Query"**
3. Você verá o fluxo completo: `service-a → service-b → viacep.GetCep / weatherapi.GetClima`

## Rodando os testes do Serviço B

```bash
cd cloud-run
WEATHER_API=sua_chave_aqui go test ./...
```
