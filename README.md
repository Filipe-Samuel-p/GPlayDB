# 🎬 GPlayDB

## 🧠 Sobre o projeto
O **GPlayDB** é um projeto desenvolvido para fins educacionais durante o minicurso **"Banco de Dados no dia a dia de um Desenvolvedor"**.  
O objetivo é demonstrar, na prática, como desenvolvedores trabalham com **SQL real**, **bancos relacionais**, **ambientes Dockerizados** e **integração com o back-end** usando Go e PostgreSQL.

A aplicação simula o sistema de **assinaturas dos produtos Globo**, como:
- **Globoplay Padrão**
- **Globoplay Premium**
- **Premiere**
- **Telecine**
- **Cartola PRO**

Durante o minicurso, os participantes irão explorar o banco, criar consultas SQL e ver como essas operações se refletem no back-end e na API.

---

## 🧰 Tecnologias Utilizadas

| Tecnologia | Descrição |
|-------------|------------|
| **Go (Golang)** | Linguagem usada no back-end |
| **database/sql** | Pacote nativo para interação com bancos de dados |
| **PostgreSQL** | Banco de dados relacional |
| **Docker Compose** | Orquestração do ambiente |
| **Beekeeper Studio** | Cliente SQL visual para consultas |
| **Postman** | Cliente HTTP para testar a API |
| **GitHub Issues** | Organização das tarefas e consultas no formato de Kanban |

---

## 💻 Contexto do desenvolvimento

Durante o minicurso (2h de duração), os participantes irão percorrer **todas as etapas do fluxo de um desenvolvedor back-end**, desde a configuração do ambiente até a execução de queries no código Go.

### 🔹 Etapa 1 – Onboarding e ambiente
- Clonar o repositório.
- Subir os containers com `docker-compose up`.
- Validar a conexão com o banco e o seed inicial (tabelas e dados fictícios).
- Acessar o banco via **Beekeeper Studio**.

### 🔹 Etapa 2 – Consultas SQL no Beekeeper
- Cada consulta será uma **tarefa (issue)** no GitHub, simulando um **Kanban de time de desenvolvimento**.  
- Os participantes resolverão issues que envolvem:
  - `SELECT`, `WHERE`, `ORDER BY`
  - `JOIN` entre tabelas (`users`, `products`, `categories`, `subscriptions`)
  - `GROUP BY`, `COUNT`, `SUM`, `AVG`
- O foco é **entender o raciocínio SQL** e **analisar dados reais simulados**.

### 🔹 Etapa 3 – SQL no Back-end
- Veremos como o Go se conecta ao banco via **`database/sql`** e **driver PostgreSQL (`lib/pq`)**.
- Serão criadas funções para executar operações SQL reais:
  - Consultas (`Query` e `QueryRow`)
  - Inserções (`Exec` com parâmetros)
  - Atualizações e exclusões
- Criação de endpoints REST simples, por exemplo:
  - `GET /subscriptions`
  - `POST /subscriptions`
  - `PUT /subscriptions/{id}`
  - `DELETE /subscriptions/{id}`
- Testes serão feitos via **Postman**, e as mudanças poderão ser visualizadas no **Beekeeper Studio**.

---

## Modelagem do banco
<img src="image/GPlayDB.pngs" alt="Modelagem do banco">

## 🚀 Como Utilizar

### 1. Clonar o repositório
```bash
git clone https://github.com/seu-usuario/gplaydb.git
cd gplaydb 
````

### 2. Subir o ambiente
Execute o comando:
``` bash
docker-compose up
```
Isso criará automaticamente:
- Um container com o PostgreSQL e o seed inicial.
- O ambiente de desenvolvimento configurado para o Beekeeper Studio.

### 3. Acessar o banco via Beekeeper Studio
- **Host:** localhost  
- **Porta:** 5432  
- **Usuário:** ...  
- **Senha:** ...  
- **Banco:** gplaydb  


### 4. Explorar as consultas e issues
- Acesse o repositório no GitHub e veja as **issues abertas**.  
- Cada issue representa um desafio SQL (ex.: “Listar produtos”, “Calcular faturamento mensal”, “Top assinaturas ativas”).

### 5. Testar a API
Use o **Postman** para realizar as operações:
- **GET /subscriptions** — lista assinaturas (com JOINs)  
- **POST /subscriptions** — cria nova assinatura  
- **PUT /subscriptions/{id}** — atualiza (ex.: ativa/desativa)  
- **DELETE /subscriptions/{id}** — remove a assinatura



