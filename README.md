<div align="center">

# 🌻💜 Incident Assistant  
### _Analisador inteligente de incidentes ITIL-like_  
#### Go + Docker + GitHub Actions 🚀

---

![go-ci](https://github.com/rakellkizz/incident-assistant/actions/workflows/ci.yml/badge.svg)
![Go Version](https://img.shields.io/badge/Go-1.22-blue?logo=go)
![Docker](https://img.shields.io/badge/Docker-ready-blue?logo=docker)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-online-success)

---

### 💬 Projeto desenvolvido por **Raquel G. de Souza**  
#### 💼 *Demonstração prática de engenharia de software moderna: CI/CD, containerização e automação ITIL-like.*

</div>


## 🧠 Sobre o Projeto

O **Incident Assistant** é um microserviço desenvolvido em **Go (Golang)** que simula a análise de incidentes de TI de forma inteligente e automatizada.  
Ele aplica regras inspiradas em **ITIL**, categorizando, priorizando e calculando **SLAs** com base em uma matriz de **Impacto × Urgência**.

Além disso, conta com **integração opcional com OpenAI**, permitindo gerar resumos automáticos dos incidentes quando configurado com a variável `OPENAI_API_KEY`.

---

## 🧩 Abordagem ITIL-like

O projeto adota uma abordagem **ITIL-like** para o gerenciamento de incidentes, inspirada nas boas práticas do framework **ITIL (Information Technology Infrastructure Library)**.  
Isso significa que o sistema segue a lógica de **Impacto × Urgência**, determinando automaticamente:

- 🔺 **Prioridade (P1–P4)**  
- ⏱️ **SLA estimado** (em minutos)  
- 🧠 **Categoria** do incidente (Aplicação, Banco de Dados, Infraestrutura, etc.)

Essa abordagem “ITIL-like” replica o raciocínio usado em **centrais de serviço corporativas (Service Desk)**, mas de forma **simplificada e automatizada**, ideal para demonstrações de arquitetura e entrevistas técnicas.

---

## ⚙️ Funcionalidades

- Análise de incidentes via API REST (`/analyze-incident`)
- Determinação automática de prioridade e SLA
- Categorização simples por palavras-chave
- Resumo com IA (usando OpenAI) — opcional
- Container seguro (Distroless non-root)
- Testes automatizados via GitHub Actions (CI/CD)

---

## 🧪 Teste Rápido

### Local (Go)
```bash
go mod tidy
go build
./incident-assistant
# GET http://localhost:8080/health

Docker

docker build -t incident-assistant .
docker run -p 8080:8080 incident-assistant

Teste com IA
docker run -p 8080:8080 -e OPENAI_API_KEY="sua_chave_aqui" incident-assistant

🔌 Endpoints
GET /health

Retorna o status da aplicação:
{ "status": "ok" }

POST /analyze-incident

Exemplo de requisição:
{
  "title": "Erro de login",
  "description": "Usuários não conseguem autenticar",
  "impact": "high",
  "urgency": "critical"
}

Exemplo de resposta:
{
  "priority": "P1",
  "sla_minutes": 120,
  "category": "Aplicação",
  "summary": "Incidente: Erro de login | Impacto: high | Urgência: critical",
  "used_ai": false
}

🧩 Estrutura do Projeto

internal/
 ├── handlers.go      # Rotas e handlers da API (Gin)
 ├── models.go        # Estruturas de request e response
 ├── rules.go         # Matriz de prioridade, SLA e categorização
 └── tests/
      └── rules_test.go   # Testes unitários da matriz P×U
ai.go                # Integração com OpenAI
main.go              # Inicialização do servidor e rotas
Dockerfile           # Build (Alpine) + Runtime (Distroless)
.github/workflows/   # Pipeline CI/CD com GitHub Actions

🧩 Tecnologias Utilizadas

GoLang – Backend leve e de alta performance

Gin – Framework web rápido para APIs REST

Docker – Empacotamento e isolamento de ambiente

GitHub Actions – Integração contínua e testes automatizados

OpenAI API (opcional) – Geração de resumo automatizado via IA

🧭 Como Contribuir

Fork o projeto

Crie uma branch: git checkout -b feature/nome-da-feature

Commit: git commit -m "feat: descrição da mudança"

Push: git push origin feature/nome-da-feature

Abra um Pull Request 🚀

💼 Licença

Este projeto foi desenvolvido por Raquel G. de Souza
 como demonstração técnica de arquitetura e automação ITIL-like.
Sinta-se à vontade para usar como referência em portfólios e entrevistas.

---

## 🎯 Como Testar o Projeto

Este projeto foi desenvolvido para permitir **testes rápidos e práticos**, tanto via **Go local**, **Docker** ou diretamente em **requisições HTTP** (sem precisar interface gráfica).

### ✅ 1. Teste via Docker (recomendado)

O método mais rápido — basta ter o **Docker Desktop** instalado.

```bash
# Clonar o repositório
git clone https://github.com/rakellkizz/incident-assistant.git
cd incident-assistant

# Build da imagem
docker build -t incident-assistant .

# Executar o container
docker run -p 8080:8080 incident-assistant

➡️ Depois, abra o navegador e acesse:
http://localhost:8080/health

{ "status": "ok" }

Para testar o endpoint principal:

curl -X POST http://localhost:8080/analyze-incident \
-H "Content-Type: application/json" \
-d '{"title":"Erro de login","description":"Usuários não conseguem autenticar","impact":"high","urgency":"critical"}'
Resposta esperada:
{
  "priority": "P1",
  "sla_minutes": 120,
  "category": "Aplicação",
  "summary": "Incidente: Erro de login | Impacto: high | Urgência: critical",
  "used_ai": false
}

💡 2. Teste com IA (OpenAI)
Caso deseje ver o comportamento com IA ativada (resumo automático):

docker run -p 8080:8080 -e OPENAI_API_KEY="SUA_CHAVE_AQUI" incident-assistant

O campo "used_ai": true indicará que o resumo foi gerado com suporte de IA.

🧪 3. Teste via Go (sem Docker)
Se preferir testar localmente com o Go instalado:
go mod tidy
go build
./incident-assistant

Endpoints disponíveis:

GET /health

POST /analyze-incident

⚙️ 4. Teste automático via GitHub Actions

1- O projeto já possui um pipeline CI configurado em .github/workflows/ci.yml.
Ele executa automaticamente:

2- go test ./... — validação das regras da matriz de prioridade e SLA.

3- Build do container Docker.

4- Teste de “smoke” automático, verificando que /health e /analyze-incident respondem com sucesso.

5- Pode verificar os resultados desses testes na aba Actions do repositório.

Exemplo de execução bem-sucedida:

🔍 5. Teste manual via PowerShell (Windows)
No PowerShell:
$body = @{
  title       = "Erro de login"
  description = "Usuários não conseguem autenticar"
  impact      = "high"
  urgency     = "critical"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/analyze-incident" -Method Post -Body $body -ContentType "application/json"

🌐 6. Teste direto com arquivo JSON
Crie payload.json:

{
  "title": "Erro de login",
  "description": "Usuários não conseguem autenticar",
  "impact": "high",
  "urgency": "critical"
}

Execute:
curl -X POST http://localhost:8080/analyze-incident -H "Content-Type: application/json" --data-binary "@payload.json"

💼 Dica Importante

ste projeto demonstra domínio prático de engenharia de software, automação de processos ITIL-like, e entrega contínua (CI/CD) com Go, Docker e GitHub Actions.
Todas as rotas podem ser testadas em ambiente local ou via pipeline automatizado, com resultados reproduzíveis.


---

💜 Assim é possivel conseguir:
- **rodar tudo em menos de 2 minutos** via Docker;
- **ver o build e os testes automáticos** no GitHub Actions;
- **testar a API manualmente** se quiser analisar as respostas;
- e entender de forma clara que você domina **DevOps, automação e boas práticas ITIL**.

---

## 🤖 Project Rules (AI-guided testing)

Este módulo adiciona **regras de projeto** que **orientam a IA no desenvolvimento de testes**.  
A ideia é padronizar a qualidade e acelerar a criação de testes com diretrizes claras.

**Arquivos principais**
- `project_rules.json` – regras que a IA deve seguir (na raiz do repositório)
- `internal/ai/rules_generator.go` – lê o JSON e formata as regras
- `run_rules.go` – executável que imprime as regras (para demo/CI)

**Exemplo do `project_rules.json`:**
```json
{
  "project_name": "Incident Assistant",
  "test_standards": {
    "naming_convention": "Test<Feature>_<Scenario>",
    "required_cases": ["success", "validation_error", "edge_case"],
    "coverage_target": 0.8
  },
  "assertions": {
    "success": "res.StatusCode == 200",
    "validation_error": "res.StatusCode == 400"
  },
  "ai_guidelines": [
    "Gerar testes somente para funções públicas exportadas.",
    "Usar mocks para dependências externas.",
    "Adicionar mensagens claras em falhas: 'esperava X, obteve Y'."
  ]
}

