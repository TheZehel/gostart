// main.go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

// Post representa o payload de exemplo para requisições GET/POST
type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id,omitempty"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func main() {
    fmt.Println("=== Exemplo GET ===")
    getExample()

    fmt.Println("\n=== Exemplo POST ===")
    postExample()

    fmt.Println("\n=== Exemplo de Requisição Customizada ===")
    customRequestExample()
}

func getExample() {
    // 1. Envia GET para a API pública
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        log.Fatalf("Erro em GET: %v", err)
    }
    defer resp.Body.Close()

    // 2. Lê o corpo da resposta
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Erro ao ler body: %v", err)
    }

    // 3. Verifica o status
    fmt.Printf("Status: %s\n", resp.Status)

    // 4. Converte JSON em struct
    var post Post
    if err := json.Unmarshal(body, &post); err != nil {
        log.Fatalf("Erro no Unmarshal: %v", err)
    }

    // 5. Exibe os dados do post
    fmt.Printf("Post ID %d: %q — %q\n", post.ID, post.Title, post.Body)
}

func postExample() {
    // 1. Cria payload
    newPost := Post{
        UserID: 42,
        Title:  "Meu Novo Post",
        Body:   "Este é o corpo do meu post em Go!",
    }

    // 2. Converte struct para JSON
    jsonData, err := json.Marshal(newPost)
    if err != nil {
        log.Fatalf("Erro no Marshal: %v", err)
    }

    // 3. Envia POST com corpo JSON
    resp, err := http.Post(
        "https://jsonplaceholder.typicode.com/posts",
        "application/json; charset=UTF-8",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        log.Fatalf("Erro em POST: %v", err)
    }
    defer resp.Body.Close()

    // 4. Lê e exibe status + corpo
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Erro ao ler response body: %v", err)
    }
    fmt.Printf("Status: %s\n", resp.Status)
    fmt.Printf("Resposta do servidor: %s\n", string(body))
}

func customRequestExample() {
    // 1. Cria cliente com timeout
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    // 2. Monta requisição GET customizada
    req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/2", nil)
    if err != nil {
        log.Fatalf("Erro ao criar request: %v", err)
    }
    // 3. Define headers
    req.Header.Set("Accept", "application/json")

    // 4. Envia a requisição
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Erro em client.Do: %v", err)
    }
    defer resp.Body.Close()

    // 5. Lê e processa resposta
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Erro ao ler body: %v", err)
    }
    fmt.Printf("Status: %s\nCorpo: %s\n", resp.Status, string(body))
}
