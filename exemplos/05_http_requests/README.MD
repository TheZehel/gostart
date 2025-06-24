# 05 - HTTP Requests

Este exemplo mostra como realizar requisições HTTP (GET e POST) em Go, consumindo e enviando dados em JSON.

## Executando

```bash
go run main.go
```

## Explicação

- **net/http**: pacote padrão para criar clientes e servidores HTTP.

- **http.Get(url)**: envia uma requisição GET e retorna um \*http.Response.

- **http.Post(url, contentType, body)**: envia uma requisição POST com corpo, retornando um \*http.Response.

- **defer resp.Body.Close()**: fecha o corpo da resposta após o uso, evitando vazamentos de recursos.

- **io/ioutil.ReadAll(resp.Body)**: lê todo o conteúdo da resposta como slice de bytes.

- **encoding/json (json.Unmarshal / json.Marshal)**: converte entre JSON e structs Go.

- **http.NewRequest(method, url, body)**: cria requisições HTTP customizadas para definir headers e corpo.

- **client := &http.Client{} / client.Do(req)**: envia requisições customizadas, permitindo configuração de timeout e transporte.

- **resp.StatusCode**: verifica o código de status HTTP retornado pelo servidor.

- **log.Fatalf**: exibe mensagem de erro e interrompe o programa em caso de falha.

- **fmt.Println**: imprime resultados (corpo, status, erros) no terminal.
