# 03 - Linhas Duplicadas

Este exemplo lê várias linhas de texto da entrada padrão e identifica quais se repetem, exibindo a contagem de repetições.

## Executando

```bash
go run main.go
```

Digite as linhas e finalize com Ctrl+D (Linux/macOS) ou Ctrl+Z (Windows).

```bash
Go
Python
Go
Rust
Go
^D
```

Saída:
```bash
3	Go
```

> ℹ️ **Explicação**
>
> - `bufio.NewScanner(os.Stdin)`: cria um **scanner** que lê linhas do terminal.
> - `scanner.Scan()`: lê cada **linha digitada**.
> - `scanner.Text()`: retorna o **texto da linha** lida.
> - `make(map[string]int)`: cria um **mapa** para contar as ocorrências de cada linha.
> - `counts[linha]++`: **incrementa** a contagem daquela linha no mapa.
> - No final, o programa **exibe apenas as linhas que aparecem mais de uma vez**.

> ⚠️ **Observação**
>
> Se nenhuma linha for digitada mais de uma vez, **nada será exibido** após `Linhas duplicadas:`.


