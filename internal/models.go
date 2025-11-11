package internal

type IncidentRequest struct {
	Title       string `json:"title"       binding:"required"`
	Description string `json:"description" binding:"required"`
	Impact      string `json:"impact"      binding:"required"`  // low|medium|high|critical
	Urgency     string `json:"urgency"     binding:"required"`  // low|medium|high|critical
}

type IncidentResponse struct {
	Priority   string `json:"priority"`     // P1..P4
	SLAMinutes int    `json:"sla_minutes"`  // sugestão de SLA
	Category   string `json:"category"`     // e.g., "Aplicação", "Infra", etc. (placeholder)
	Summary    string `json:"summary"`      // resumo (regra/IA)
	UsedAI     bool   `json:"used_ai"`
}
