# Neoway

Neoway é uma aplicação web desenvolvida para cadastrar e listar clientes de uma empresa fictícia que fornece serviços B2C e B2B. O sistema permite o cadastro de clientes com CPF/CNPJ, consulta, busca e ordenação dos registros, bem como a validação dos documentos cadastrados. O projeto é implementado com uma arquitetura baseada em microserviços, utilizando Go para o backend e Vue.js para o frontend, com o PostgreSQL como banco de dados.

## Arquitetura

A arquitetura do projeto é organizada da seguinte forma:

```
neoway/
│
├── backend/
│   ├── main.go
│   ├── controllers/
│   │   ├── client_controller.go
│   │   └── status_controller.go
│   ├── models/
│   │   └── client.go
│   ├── routes/
│   │   └── routes.go
│   ├── utils/
│   │   └── cpfcnpj.go
│   └── database/
│       └── db.go
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   ├── ClientForm.vue
│   │   │   └── ClientList.vue
|   │   ├── router/         
|   |   │   └── index.js
|   │   ├── store/            
│   │   |   └── index.js
│   │   ├── views/
│   │   │   └── Home.vue
│   │   ├── App.vue
|   |   └── main.js
│   ├── public/
│   │   └── index.html
│   └── package.json
│
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
│
├── tests/
│   ├── backend/
│   │   ├── client_test.go
│   │   └── status_test.go
│   │   └── integration/
│   │       └── api_integration_test.go
│   └── frontend/
│       ├── ClientForm.spec.js
│       └── ClientList.spec.js
│
└── README.md
```

## Tecnologias

- **Backend**: Go (Golang), Beego framework, PostgreSQL, pacote `cpfcnpj`
- **Frontend**: Vue.js
- **Banco de Dados**: PostgreSQL
- **Conteinerização**: Docker
- **Testes**: Ginkgo, Jest
- **Documentação e Monitoramento**: Swagger, Datadog

## Instalação e Configuração

### Gerenciamento de Dependências

**Backend**

Para gerenciar as dependências do backend em Go, utilizamos o módulo Go. As dependências são especificadas no arquivo `go.mod` e podem ser atualizadas usando os seguintes comandos:

1. **Instalar Dependências**:
Navegue até o diretório `backend` e execute:
    
    ```bash
    cd backend
    go mod download
    ```
    
2. **Atualizar Dependências**:
Se precisar adicionar novas dependências ou atualizar as existentes, execute:
    
    ```bash
    go mod tidy
    ```
    

**Frontend**

Para o frontend, utilizamos o npm (Node Package Manager) para gerenciar as dependências do Vue.js. As dependências são listadas no arquivo `package.json`.

1. **Instalar Dependências**:
Navegue até o diretório `frontend` e execute:
    
    ```bash
    cd frontend
    npm install
    ```
    
2. **Atualizar Dependências**:
Para atualizar as dependências, execute:
    
    ```bash
    npm update
    ```

### Inicialização do Projeto

### **Servidor Backend**

Primeiro, você precisa garantir que o servidor backend esteja rodando. Se você ainda não configurou um script para iniciar o backend, você pode fazer isso manualmente. Navegue até o diretório `backend` e execute:

```bash
go run main.go
```

### **Servidor Frontend**

Para iniciar o servidor de desenvolvimento do frontend, navegue até o diretório do frontend e execute o comando:

```bash
npm run serve
```

Isso iniciará o servidor de desenvolvimento Vue.js e você deve ser capaz de acessar o frontend em `http://localhost:8080` (ou a porta configurada).

### **Executar Testes Unitários e de Integração:**

```bash
go test ./tests
```

### Conteinerização

Para a conteinerização da aplicação, utilizamos Docker. O Docker permite empacotar a aplicação e suas dependências em um container, garantindo que a aplicação funcione de maneira consistente em diferentes ambientes.

**Configuração do Docker**

O Docker é configurado através dos arquivos `Dockerfile` e `docker-compose.yml`.

1. **Dockerfile**:
    - O `Dockerfile` define o ambiente de execução da aplicação backend e frontend. Certifique-se de que os arquivos `Dockerfile` para o backend e o frontend estão corretamente configurados.
2. **docker-compose.yml**:
    - O `docker-compose.yml` define e gerencia múltiplos containers. Ele orquestra a construção e a execução dos containers do backend e frontend, além do banco de dados PostgreSQL.

**Executando a Aplicação com Docker**

1. **Criação e Execução dos Containers**:
No diretório raiz do projeto (`neoway`), execute:
    
    ```bash
    docker-compose up --build
    ```
    
    Este comando cria e inicia os containers definidos no `docker-compose.yml`. O parâmetro `--build` força a reconstrução das imagens dos containers, garantindo que todas as alterações sejam aplicadas.
    
2. **Verificação dos Containers**:
Para verificar se os containers estão em execução, use:
    
    ```bash
    docker-compose ps
    ```
    
3. **Parar e Remover os Containers**:
Para parar e remover os containers, execute:
    
    ```bash
    docker-compose down
    ```
    

### **Configuração do Banco de Dados**

O PostgreSQL é configurado no `docker-compose.yml`, e os bancos de dados `neoway_db` e `test_db` são criados automaticamente quando os containers são iniciados.

### Acessando a Aplicação

- **Backend**: A API pode ser acessada em `http://localhost:8080`.
- **Frontend**: A interface do usuário pode ser acessada em `http://localhost:80`.

### Testes de Endpoints com Postman

### Cadastro de Clientes

- **Método:** POST
- **URL:** `http://localhost:8080/clients`
- **Corpo:** JSON
    
    ```json
    {
      "cpf_cnpj": "12345678901",
      "name": "João da Silva"
    }
    ```

### Consulta de Todos os Clientes

- **Método:** GET
- **URL:** `http://localhost:8080/clients`
- **Descrição:** Consulta todos os clientes cadastrados. Permite busca por nome/razão social e ordenação em ordem alfabética.

### Consulta de Cliente Específico

- **Método:** GET
- **URL:** `http://localhost:8080/clients/{cpf_cnpj}`
- **Descrição:** Consulta se um determinado CPF/CNPJ está cadastrado na base de clientes.
- **Parâmetro de URL:** `{cpf_cnpj}` (substitua pelo CPF/CNPJ que deseja verificar)

### Validação do Número do CPF/CNPJ

- **Método:** POST
- **URL:** `http://localhost:8080/validate`
- **Corpo:** JSON
    
    ```json
    {
      "cpf_cnpj": "12345678901"
    }
    ```
    
- **Descrição:** Valida o número do CPF/CNPJ com dígito verificador na consulta e inclusão.

### Endpoint de Suporte

- **Método:** GET
- **URL:** `http://localhost:8080/status`
- **Descrição:** Retorna informações de up-time do servidor e a quantidade de requisições realizadas desde o início de sua execução.

### Exemplo de Testes com Postman

**Cadastro de Clientes:**

1. **Criação de um Cliente**
    - **Método:** POST
    - **URL:** `http://localhost:8080/clients`
    - **Corpo:**
        
        ```json
        {
          "cpf_cnpj": "12345678901",
          "name": "João da Silva"
        }
        ```
        
    - **Resposta Esperada:**
        - Status: 201 Created
        - Corpo: Detalhes do cliente cadastrado

**Consulta de Todos os Clientes:**

1. **Listagem de Todos os Clientes**
    - **Método:** GET
    - **URL:** `http://localhost:8080/clients`
    - **Resposta Esperada:**
        - Status: 200 OK
        - Corpo: Lista de todos os clientes com informações de nome/razão social e CPF/CNPJ

**Consulta de Cliente Específico:**

1. **Verificação de Cliente por CPF/CNPJ**
    - **Método:** GET
    - **URL:** `http://localhost:8080/clients/12345678901`
    - **Resposta Esperada:**
        - Status: 200 OK se o cliente estiver cadastrado
        - Corpo: Detalhes do cliente

**Validação de CPF/CNPJ:**

1. **Validação de CPF/CNPJ**
    - **Método:** POST
    - **URL:** `http://localhost:8080/validate`
    - **Corpo:**
        
        ```json
        {
          "cpf_cnpj": "12345678901"
        }
        ```
        
    - **Resposta Esperada:**
        - Status: 200 OK
        - Corpo: Resultado da validação (Válido/Inválido)

**Endpoint de Suporte:**

1. **Informações de Up-Time e Requisições**
    - **Método:** GET
    - **URL:** `http://localhost:8080/status`
    - **Resposta Esperada:**
        - Status: 200 OK
        - Corpo: Informações sobre up-time e quantidade de requisições

## Documentação

A API é documentada usando Swagger. Acesse a documentação em:
```
http://localhost:8080/swagger/index.html
```