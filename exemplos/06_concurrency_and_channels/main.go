// main.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// worker simula um trabalhador que processa jobs e envia resultados
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Printf("Worker %d processando job %d\n", id, j)
        time.Sleep(time.Second) // trabalho simulado
        results <- j * 2
    }
}

// selectExample demonstra multiplexação com select e default
func selectExample() {
    fmt.Println("=== Exemplo Select ===")
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "mensagem de c1"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "mensagem de c2"
    }()

    for i := 0; i < 4; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("Recebido de c1:", msg1)
        case msg2 := <-c2:
            fmt.Println("Recebido de c2:", msg2)
        default:
            fmt.Println("Nenhuma mensagem pronta, fazendo outra coisa...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    // Exemplo básico de goroutines sincronizadas com WaitGroup
    fmt.Println("=== Exemplo WaitGroup e Goroutines ===")
    var wg sync.WaitGroup
    tasks := []string{"tarefa1", "tarefa2", "tarefa3"}
    wg.Add(len(tasks))
    for _, t := range tasks {
        go func(task string) {
            defer wg.Done()
            fmt.Println("Processando", task)
            time.Sleep(time.Duration(len(task)) * 300 * time.Millisecond)
            fmt.Println("Concluído", task)
        }(t)
    }
    wg.Wait()

    // Exemplo de Worker Pool
    fmt.Println("\n=== Exemplo Worker Pool ===")
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    numWorkers := 3

    // Inicia workers
    wg.Add(numWorkers)
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results, &wg)
    }

    // Envia jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Fecha results quando todos os workers terminarem
    go func() {
        wg.Wait()
        close(results)
    }()

    // Recebe e imprime resultados
    for res := range results {
        fmt.Println("Resultado:", res)
    }

    // Demonstração de select/default
    fmt.Println()
    selectExample()
}
