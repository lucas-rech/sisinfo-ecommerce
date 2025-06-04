# 🛒 CASI IFRS Vacaria -  E-commerce 
Esse é um modelo de e-commerce desenvolvido com o propósito de ser uma aplicação real e aplicável para o curso de Sistemas de Informação do IFRS campi Vacaria. A ideia é que sirva como base para estudos e desenvolvimento de habilidades práticas em Go, especialmente com o framework Gin e que possa ser utilizado como modelo para a loja virtual do Centro Acadêmico de Sistemas de Informação (CASI) do IFRS Vacaria.


API RESTful desenvolvida em Go com o framework Gin para gerenciar um sistema de e-commerce. A aplicação até o momento possui autenticação JWT, middleware para validação de papéis de usuário (admin e consumidor), controle de produtos, carrinho de compras e documentação interativa via Swagger.

---

## 🚀 Funcionalidades

* Cadastro e login de usuários com autenticação JWT
* CRUD de produtos (somente para administradores)
* Gerenciamento de carrinho de compras (somente para consumidores)
* Middleware para validação de autenticação e autorização
* Documentação interativa via Swagger UI

---

## 🧪 Tecnologias Utilizadas

* [Go (Golang)](https://golang.org/)
* [Gin Gonic](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/)
* [JWT](https://jwt.io/)
* [Swaggo (Swagger para Go)](https://github.com/swaggo/gin-swagger)

---

## 📁 Estrutura do Projeto

```
backend/
├── config/          # Container de injeção de dependência e conexão com DB
├── internal/
│   ├── handler/     # Controladores HTTP
│   ├── service/     # Regras de negócio
│   ├── repository/  # Acesso ao banco de dados
│   └── model/       # Entidades e DTOs
├── middleware/      # Middlewares personalizados (JWT, Roles)
├── router/          # Definição de rotas e grupos
├── docs/            # Documentação Swagger
|── utils/           # Utilitários e funções auxiliares
└── main.go          # Ponto de entrada da aplicação
```

---

## ⚙️ Executando o Projeto

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/sisinfo-ecommerce.git
cd sisinfo-ecommerce/backend
```

2. Instale as dependências:

```bash
go mod tidy
```

3. Gere a documentação Swagger:

```bash
swag init
```

4. Execute o projeto:

```bash
go run main.go
```

Acesse a documentação em:

```
http://localhost:8080/swagger/index.html
```

---

## 🔐 Autenticação

Utilize o cabeçalho `Authorization` no formato:

```
Bearer <seu_token_jwt disponibilizado no login>
```

---

## 📚 Endpoints Principais

### Autenticação

* `POST /api/v1/login`
* `POST /api/v1/login/register`

### Produtos (Admin)

* `POST /api/v1/admin/product`
* `PATCH /api/v1/admin/product/:id`
* `DELETE /api/v1/admin/product/:id`

### Usuário (Admin)
* `GET /api/v1/admin/user/:id`
* `GET /api/v1/admin/user/email/:email`
* `PATCH /api/v1/admin/user/:id`
* `DELETE /api/v1/admin/user/:id`

### Produtos (Aberto)

* `GET /api/v1/products`
* `GET /api/v1/product/:id`

### Carrinho (Cliente)

* `POST /api/v1/cart/item`
* `PATCH /api/v1/cart/item`
* `DELETE /api/v1/cart/item`
* `GET /api/v1/cart/items`

---

## 🧑 Autor

**Lucas Rech**
Projeto desenvolvido para o curso de Sistemas de Informação do IFRS - Campus Vacaria.

---

## 📄 Licença

Este projeto está licenciado sob os termos da licença MIT.
