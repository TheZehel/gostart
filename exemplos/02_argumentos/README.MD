# 02 - Argumentos da Linha de Comando

Este exemplo mostra como acessar argumentos passados para o programa via terminal.

## Executando

```bash
go run main.go Go é incrível!
```

> ℹ️ **Explicação**
>
> - `os.Args` retorna uma **slice** com todos os argumentos passados pela linha de comando.
> - `os.Args[0]` é sempre o **nome do programa**.
> - `os.Args[1:]` contém apenas os **argumentos reais** fornecidos.
> - `strings.Join(args, " ")` **concatena** os argumentos, separados por espaço.


> 💡 **Nota**
>
> Se você rodar o comando sem argumentos:
>
> ```bash
> go run main.go
> ```
>
> A saída será:
>
> ```
> Uso: go run main.go <arg1> <arg2> ...
> ```
>
> Isso acontece porque o programa espera ao menos um argumento após o nome do arquivo.
>
> ✅ Exemplo correto:
>
> ```bash
> go run main.go Olá mundo em Go
> ```
>
> Saída:
>
> ```
> Argumentos recebidos: Olá mundo em Go
> ```
