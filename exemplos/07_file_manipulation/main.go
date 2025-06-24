// main.go
package main

import (
    "bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
)

// Config representa uma estrutura de configuração em JSON
type Config struct {
    AppName string `json:"appName"`
    Version string `json:"version"`
    Debug   bool   `json:"debug"`
}

func main() {
    // 1. Escrita e leitura de arquivo de texto
    writeAndReadTextFile("example.txt")

    // 2. Cópia de arquivo
    if err := copyFile("example.txt", "example_copy.txt"); err != nil {
        log.Fatalf("Erro ao copiar arquivo: %v", err)
    }
    fmt.Println("Arquivo copiado para example_copy.txt")

    // 3. Leitura e escrita de JSON
    readWriteJSON("config.json")

    // 4. Leitura e escrita de CSV
    readWriteCSV("data.csv")

    // 5. Percorrer diretório e exibir metadados
    fmt.Println("\n=== Sumário de arquivos no diretório atual ===")
    walkDir(".")
}

func writeAndReadTextFile(path string) {
    // Cria e escreve no arquivo
    file, err := os.Create(path)
    if err != nil {
        log.Fatalf("Erro ao criar %s: %v", path, err)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    fmt.Fprintln(writer, "Linha 1: Olá, Go!")
    fmt.Fprintln(writer, "Linha 2: Manipulação de arquivos.")
    writer.Flush()
    fmt.Printf("Arquivo %s escrito com sucesso\n", path)

    // Abre para leitura
    f, err := os.Open(path)
    if err != nil {
        log.Fatalf("Erro ao abrir %s: %v", path, err)
    }
    defer f.Close()

    fmt.Printf("Conteúdo de %s:\n", path)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        fmt.Println(">", scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatalf("Erro ao ler %s: %v", path, err)
    }
}

func copyFile(src, dst string) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    // Copia todo o conteúdo
    if _, err := io.Copy(out, in); err != nil {
        return err
    }
    return nil
}

func readWriteJSON(path string) {
    // Exemplo de configuração
    cfg := Config{
        AppName: "GoFileTool",
        Version: "1.0.0",
        Debug:   true,
    }

    // Escreve JSON
    out, err := os.Create(path)
    if err != nil {
        log.Fatalf("Erro ao criar %s: %v", path, err)
    }
    encoder := json.NewEncoder(out)
    encoder.SetIndent("", "  ")
    if err := encoder.Encode(cfg); err != nil {
        log.Fatalf("Erro ao escrever JSON em %s: %v", path, err)
    }
    out.Close()
    fmt.Printf("Configuração JSON gravada em %s\n", path)

    // Lê JSON novamente
    in, err := os.Open(path)
    if err != nil {
        log.Fatalf("Erro ao abrir %s: %v", path, err)
    }
    defer in.Close()

    var cfg2 Config
    decoder := json.NewDecoder(in)
    if err := decoder.Decode(&cfg2); err != nil {
        log.Fatalf("Erro ao decodificar JSON de %s: %v", path, err)
    }
    fmt.Printf("Config lida: %+v\n", cfg2)
}

func readWriteCSV(path string) {
    // Dados de exemplo
    records := [][]string{
        {"nome", "idade", "cidade"},
        {"Alice", "30", "São Paulo"},
        {"Bob", "25", "Rio de Janeiro"},
        {"Carol", "28", "Belo Horizonte"},
    }

    // Escreve CSV
    file, err := os.Create(path)
    if err != nil {
        log.Fatalf("Erro ao criar %s: %v", path, err)
    }
    writer := csv.NewWriter(file)
    if err := writer.WriteAll(records); err != nil {
        log.Fatalf("Erro ao escrever CSV em %s: %v", path, err)
    }
    writer.Flush()
    file.Close()
    fmt.Printf("Arquivo CSV gravado em %s\n", path)

    // Lê CSV
    f, err := os.Open(path)
    if err != nil {
        log.Fatalf("Erro ao abrir %s: %v", path, err)
    }
    defer f.Close()
    reader := csv.NewReader(f)
    fmt.Printf("Conteúdo de %s:\n", path)
    for {
        rec, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("Erro ao ler CSV de %s: %v", path, err)
        }
        fmt.Println(" -", rec)
    }
}

func walkDir(root string) {
    filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        info, err := d.Info()
        if err != nil {
            return nil
        }
        if !info.IsDir() {
            fmt.Printf("%s\t%db\n", path, info.Size())
        }
        return nil
    })
}
