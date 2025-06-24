# go-poc

Projeto de exemplo de uma API REST em Go.

## Estrutura
- `main.go`: ponto de entrada da aplicação
- `handlers/`: camada de controladores HTTP
- `services/`: camada de regras de negócio
- `models/`: definição dos modelos de dados

## Como rodar

```sh
go run main.go
```

## Docker Build
No terminal, navegue até a raiz do projeto (onde está o Dockerfile):
```sh
cd go-poc
```

Construa a imagem Docker (substitua "go-api" pelo nome que preferir):
```sh
docker build -t go-api .
```

Escanear vulnerabilidades.
```sh
docker scout recommendations
#or
docker scout recommendations local://go-api:latest
```

Rode o container expondo a porta 8080:
```sh
docker run -p 8080:8080 go-api
```

Pronto! Sua API estará acessível em http://localhost:8080/user.

## Testes

Para rodar os testes unitários do projeto, execute o comando abaixo dentro da pasta `src`:

```sh
cd src
go test ./...
go test -cover ./...
```

## Requisitos
- Go 1.24+

## Próximos passos
- Implementar rotas e lógica de negócio
- Adicionar testes
