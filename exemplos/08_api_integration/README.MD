# 08 - API Integration

Este exemplo mostra como estruturar e consumir APIs de terceiros, incluindo autenticação, tratamento de erros e lógica de retry.

## Executando

```bash
go run main.go
```

## Explicação

- **net/http**: usar `http.NewRequest` para criar chamadas customizadas e definir headers de autenticação (API Key, Bearer Token).

- **encoding/json** (`json.Marshal` / `json.Unmarshal`): converter payloads e respostas JSON em structs Go.

- **context.WithTimeout** / **context.WithCancel**: controlar timeout e cancelamento de requisições.

- **http.Client{ Timeout: ... }**: configurar cliente HTTP com timeout global.

- **interface e struct**: definir um Client interface para facilitar testes e mocks.

- **time.Sleep e loop de retry**: implementar backoff exponencial em caso de respostas de erro (5xx).

- **errors.Wrap** ou **fmt.Errorf("%w")**: empacotar erros com contexto adicional.

- **log.Printf** ou outro logger: registrar detalhes de requisição, resposta e falhas.

- **ioutil.ReadAll** / **io.Reader**: ler corpo da resposta e tratá-lo de forma segura.
