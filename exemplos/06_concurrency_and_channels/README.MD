# 06 - Concurrency and Channels

Este exemplo demonstra como executar tarefas em paralelo usando **goroutines** e comunicar resultados via **canais**, além de sincronizar com **sync.WaitGroup** e multiplexar operações com **select**.

## Executando

```bash
go run main.go
```

## Explicação

- `go func() { … }()`: inicia uma goroutine, executando a função em paralelo.

- `make(chan T)`: cria um canal não-bufferizado para valores do tipo T.

- `make(chan T, n)`: cria um canal bufferizado com capacidade n.

- `ch <- v`: envia o valor v para o canal ch.

- `v := <-ch`: recebe um valor do canal, bloqueando até haver um envio ou o canal ser fechado.

- `close(ch)`: fecha o canal, sinalizando que não serão enviados mais valores.

- `for v := range ch`: itera sobre valores até o canal ser fechado.

- `var wg sync.WaitGroup`: cria um WaitGroup para esperar várias goroutines.

- `wg.Add(n)`: informa quantas goroutines serão esperadas.

- `wg.Done()`: sinaliza que uma goroutine terminou.

- `wg.Wait()`: bloqueia até que todas as goroutines terminem.

- `select`: multiplexer de operações em múltiplos canais, com possibilidade de default para evitar bloqueios.
