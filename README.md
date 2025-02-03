# GoGirias

GoGirias é um bot para Telegram que seleciona gírias e fornece dicas para os usuários adivinharem. O bot é uma forma divertida de aprender e interagir com a cultura local.

## Funcionalidades

- Jogue adivinhando gírias a partir de dicas.
- Respostas corretas e incorretas são registradas.
- Possibilidade de pular palavras e encerrar o jogo.

## Pré-requisitos

Antes de começar, você precisará ter o seguinte instalado:

- [Go](https://golang.org/dl/) (versão 1.23.2 ou superior)
- [Docker](https://www.docker.com/) (opcional, para executar a imagem)

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/jeffemart/GoGirias.git
   cd GoGirias
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

## Executando a Aplicação

Para executar a aplicação localmente, use o seguinte comando:

```bash
go run main.go
```

## Usando Docker

Para construir e executar a imagem Docker:

1. Construa a imagem:
   ```bash
   docker build -t jeffemart/gogirias .
   ```

2. Execute a imagem:
   ```bash
   docker run -p 3000:3000 jeffemart/gogirias
   ```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir um problema ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).
