// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// APIClient define methods para o cliente de API
type APIClient interface {
	GetUser(username string) (*User, error)
}

// GitHubClient implementa APIClient para a REST API do GitHub
type GitHubClient struct {
	httpClient *http.Client
	baseURL    string
	authToken  string
	maxRetries int
}

// User representa parte da resposta JSON da API de usuários do GitHub
type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Blog      string `json:"blog"`
	Location  string `json:"location"`
	Email     string `json:"email"`
	PublicRepos int  `json:"public_repos"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

// NewGitHubClient cria um novo GitHubClient lendo o token de GITHUB_TOKEN
func NewGitHubClient() *GitHubClient {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Println("⚠️  GITHUB_TOKEN não definido. A API do GitHub impõe limites de rate limit sem auth.")
	}
	return &GitHubClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    "https://api.github.com",
		authToken:  token,
		maxRetries: 3,
	}
}

// GetUser busca informações de um usuário no GitHub, com retry em erros 5xx
func (c *GitHubClient) GetUser(username string) (*User, error) {
	url := fmt.Sprintf("%s/users/%s", c.baseURL, username)

	// Context com timeout para cancelamento
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var lastErr error
	for attempt := 0; attempt < c.maxRetries; attempt++ {
		// Cria requisição HTTP
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("criando request: %w", err)
		}
		req.Header.Set("Accept", "application/vnd.github.v3+json")
		if c.authToken != "" {
			req.Header.Set("Authorization", "token "+c.authToken)
		}

		// Executa a requisição
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("erro no client.Do: %w", err)
		} else {
			defer resp.Body.Close()
			// Se for erro de servidor (5xx), prepara retry
			if resp.StatusCode >= 500 && resp.StatusCode < 600 {
				lastErr = fmt.Errorf("servidor respondeu %d", resp.StatusCode)
			} else if resp.StatusCode != 200 {
				// Outros status não-sucesso retornam erro imediato
				body, _ := ioutil.ReadAll(resp.Body)
				return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
			} else {
				// Sucesso: decodifica JSON
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return nil, fmt.Errorf("lendo body: %w", err)
				}
				var user User
				if err := json.Unmarshal(body, &user); err != nil {
					return nil, fmt.Errorf("unmarshal JSON: %w", err)
				}
				return &user, nil
			}
		}

		// Exponential backoff antes do próximo retry
		backoff := time.Duration(1<<attempt) * time.Second
		log.Printf("Tentativa %d falhou: %v. Retrying em %s...", attempt+1, lastErr, backoff)
		time.Sleep(backoff)
	}

	return nil, fmt.Errorf("todas as %d tentativas falharam: %w", c.maxRetries, lastErr)
}

func main() {
	client := NewGitHubClient()

	username := "TheZehel"
	fmt.Printf("Buscando dados do usuário %q no GitHub...\n\n", username)

	user, err := client.GetUser(username)
	if err != nil {
		log.Fatalf("Erro ao obter usuário: %v", err)
	}

	// Exibe dados do usuário
	fmt.Printf("Login:       %s\n", user.Login)
	fmt.Printf("ID:          %d\n", user.ID)
	fmt.Printf("Name:        %s\n", user.Name)
	fmt.Printf("Company:     %s\n", user.Company)
	fmt.Printf("Blog:        %s\n", user.Blog)
	fmt.Printf("Location:    %s\n", user.Location)
	fmt.Printf("Public Repos:%d\n", user.PublicRepos)
	fmt.Printf("Followers:   %d\n", user.Followers)
	fmt.Printf("Following:   %d\n", user.Following)
}
