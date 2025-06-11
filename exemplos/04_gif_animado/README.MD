# 04 - GIF Animado (Lissajous)

Este exemplo gera uma animação de curvas de Lissajous em formato GIF e envia a saída para o terminal via `stdout`.

## Executando

```bash
go run main.go > animacao.gif
```

> ℹ️ **Explicação**
>
> - `palette`: define as **cores** usadas no GIF (branco e preto).
> - `math.Sin(...)`: gera as **coordenadas X e Y** com base em senoides para desenhar a curva.
> - `image.NewPaletted(...)`: cria um **canvas** para desenhar a imagem animada.
> - `img.SetColorIndex(...)`: define os **pixels** a serem coloridos com base nas coordenadas calculadas.
> - `gif.EncodeAll(...)`: **codifica e escreve** todos os frames como um GIF animado na saída padrão (`stdout`).
> - A animação é criada variando a **fase** da curva em cada quadro (frame).

> ⚠️ **Observação**
>
> O terminal não exibe GIFs. Para ver a animação, **redirecione a saída para um arquivo** com:
>
> ```bash
> go run main.go > animacao.gif
> ```
> Em seguida, **abra `animacao.gif` com um visualizador de imagens**.
