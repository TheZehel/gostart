# 07 - File Manipulation

Este exemplo mostra como ler, escrever e processar arquivos e diretórios usando os pacotes padrão do Go.

## Executando

```bash
go run main.go
```

## Explicação

- **`os.Open(nome)` / `os.Create(nome)`**: abre um arquivo para leitura ou cria um novo para escrita.

- **`defer file.Close()`**: garante que o arquivo seja fechado ao final da função, evitando vazamentos de descritores.

- **`bufio.NewReader(file)` / `bufio.NewWriter(file)`**: criam buffers para operações de I/O mais eficientes.

- **`io.Copy(dst, src)`**: copia todo o conteúdo de um io.Reader para um io.Writer.

- **`os.ReadFile(path)` / `os.WriteFile(path, dados, perm)`**: lê ou escreve o conteúdo completo de um arquivo de uma vez.

- **`encoding/csv.NewReader` / `csv.NewWriter`**: lê e escreve arquivos CSV linha a linha.

- **`encoding/json.Unmarshal` / `json.MarshalIndent`**: converte entre JSON e structs Go para ler/escrever configurações.

- **`filepath.WalkDir(root, fn)`**: percorre recursivamente um diretório, permitindo processar cada arquivo/entrada.

- **`os.Stat(path)` / `entry.Info()`**: obtém metadados de arquivos (tamanho, permissões, data de modificação).
