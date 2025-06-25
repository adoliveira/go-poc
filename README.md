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

## Via Ingress
Instale o Ingress Controller no Minikube:
```sh
minikube addons enable ingress
```

Aplique o manifest:
```sh
kubectl apply -f k8s/ingress.yaml
```

#minikube tunnel
O comando abaixo cria um túnel de rede e expõe o Ingress Controller na porta 80 do seu host, resolvendo o problema de roteamento:
```sh
minikube tunnel
```
Deixe esse comando rodando em um terminal separado.

Agora é possivel testar usando:
```sh
curl -H "Host: go-api.local" http://localhost/user
```

No Minikube (especialmente com o driver docker no macOS), o Ingress geralmente é exposto em http://localhost, e não diretamente no IP da VM. Por isso, o comando curl com -H "Host: go-api.local" para localhost funciona.

Se quiser acessar pelo navegador, basta garantir que o /etc/hosts tenha:
```sh
127.0.0.1 go-api.local
```
ou
```sh
localhost go-api.local
```
Assim, você pode acessar http://go-api.local/user no navegador.

## Requisitos
- Go 1.24+

## Próximos passos
- Implementar rotas e lógica de negócio
- Adicionar testes
