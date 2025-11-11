![CI](https://github.com/rakellkizz/incident-assistant/actions/workflows/ci.yml/badge.svg)
# Incident Assistant (Go, ITIL-like)

Microserviço em Go para analisar incidentes corporativos:
- Matriz impacto x urgência → **Prioridade (P1..P4)** e **SLA sugerido**
- **Opcional:** resumo automático com IA (OpenAI), se `OPENAI_API_KEY` estiver configurada

## Endpoints
- `POST /analyze-incident`
  - body:
    ```json
    { "title": "Erro de login", "description": "Usuários sem login", "impact": "high", "urgency": "critical" }
    ```
  - response:
    ```json
    { "priority":"P1", "sla_minutes":120, "category":"Aplicação", "summary":"...", "used_ai":true }
    ```

- `GET /health` → status

## Rodando
```bash
go mod tidy
go run main.go
# Em outro terminal:
curl -X POST http://localhost:8080/analyze-incident \
  -H "Content-Type: application/json" \
  -d '{"title":"Erro de login","description":"Usuários sem login","impact":"high","urgency":"critical"}'
