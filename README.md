# Neoway

Neoway é uma aplicação web desenvolvida para cadastrar e listar clientes de uma empresa fictícia que fornece serviços B2C e B2B. O sistema permite o cadastro de clientes com CPF/CNPJ, consulta, busca e ordenação dos registros, bem como a validação dos documentos cadastrados. O projeto é implementado com uma arquitetura baseada em microserviços, utilizando Go para o backend e Vue.js para o frontend, com o PostgreSQL como banco de dados.

## Arquitetura

A arquitetura do projeto é organizada da seguinte forma:
### **Estrutura de Pastas:**

```
neoway/
│
├── backend/
│   ├── main.go
│   ├── controllers/
│   │   ├── client_controller.go
│   │   └── status_controller.go
│   ├── models/
│   │   ├── client.go
│   │   └── repository.go 
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
│   │   │   ├── ClientList.vue
|   |   |   └── Status.vue
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
│   ├── client_test.go
|   ── integration_test.go
└── README.md
```

---
## Tecnologias

- **Backend**: Go (Golang), Beego framework, PostgreSQL, pacote `cpfcnpj`
- **Frontend**: Vue.js
- **Banco de Dados**: PostgreSQL
- **Conteinerização**: Docker
- **Testes**: Ginkgo, Jest

## Instalação e Configuração

### Gerenciamento de Dependências

### **Configuração do Banco de Dados**

Para execução local basta ter o PostgreSQL instalado que será feita conexão e os bancos de dados `neoway_db` e `neoway_test` são criados automaticamente.

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

Isso iniciará o servidor de desenvolvimento Vue.js e você deve ser capaz de acessar o backend em `http://localhost:8080` e o frontend em `http://localhost:8081`

### **Executar Testes Unitários e de Integração:**

```bash
go test ./tests
```

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

### Acessando a Aplicação

- **Backend**: A API pode ser acessada em `http://localhost:8080`.
- **Frontend**: A interface do usuário pode ser acessada em `http://localhost:8081`.

### Testes de Endpoints na Própia Interface da Aplicação
![Captura de tela 2024-08-14 011853](https://github.com/user-attachments/assets/bd8e347a-324b-4806-b800-0216c9e8f9a7)

### Cadastro de Clientes
![Captura de tela 2024-08-14 013749](https://github.com/user-attachments/assets/5d6d8121-b91e-47ce-970c-305f26818f03)

### Consulta de Todos os Clientes
![Captura de tela 2024-08-14 011912](https://github.com/user-attachments/assets/b0e32731-72b7-475a-a327-5f2dd2d35311)

### Consulta de Cliente Específico
![Captura de tela 2024-08-14 013817](https://github.com/user-attachments/assets/6edbd8fd-dcf9-48a3-8703-a97b932c6997)

### Validação do Número do CPF/CNPJ
![Captura de tela 2024-08-14 013801](https://github.com/user-attachments/assets/b02bd1bc-d779-40bb-bce0-07eb4092b79f)

### Endpoint de Suporte
![Captura de tela 2024-08-14 011929](https://github.com/user-attachments/assets/50c17190-99f4-4661-8193-6a404e90e762)
