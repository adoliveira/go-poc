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

#Para fazer o push da imagem para o Docker Hub (docker.io) via GitHub Actions, siga estes passos:

Adicione os segredos DOCKERHUB_USERNAME e DOCKERHUB_TOKEN no repositório do GitHub (Settings > Secrets and variables > Actions > New repository secret).

Adicione os steps abaixo ao final do seu workflow de CI, após o build da imagem:

```yml
- name: Login no Docker Hub
run: echo "${{ secrets.DOCKERHUB_TOKEN }}" | docker login -u "${{ secrets.DOCKERHUB_USERNAME }}" --password-stdin

- name: Tag da imagem
run: docker tag go-api:latest jmbsolution/go-api:latest

- name: Push da imagem para o Docker Hub
run: docker push jmbsolution/go-api:latest
```

Esses steps vão:

Fazer login no Docker Hub usando os segredos.
Taguear a imagem local como jmbsolution/go-api:latest.
Fazer o push da imagem para o repositório do Docker Hub.

## Testes

Para rodar os testes unitários do projeto, execute o comando abaixo dentro da pasta `src`:

```sh
cd src
go test ./...
go test -cover ./...
```

## Minikube deploy
Precisa de minikube instalado no ambiente de dev

Para start do minikube
```sh
minikube start
```

Iniciar o dashboard do minikube
```sh
minikube dashboard
```

Para aplicar os manifests do Kubernetes que estão na pasta k8s, utilize o comando abaixo no terminal, dentro da raiz do projeto:
```sh
kubectl apply -f k8s/
```
Esse comando irá criar (ou atualizar) todos os recursos definidos nos arquivos YAML dentro da pasta k8s, incluindo deployment e service.

Se quiser aplicar apenas um arquivo específico, use:
```sh
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

Após aplicar, para verificar os recursos criados:
```sh
kubectl get pods
kubectl get services
```

Se precisar expor o serviço no Minikube, use:
```sh
minikube service go-api-service
```

Assim, você acessa facilmente a API pelo navegador. Se precisar de mais algum comando ou ajuste, só avisar!

## Requisitos
- Go 1.24+

## Próximos passos
- Implementar rotas e lógica de negócio
- Adicionar testes
