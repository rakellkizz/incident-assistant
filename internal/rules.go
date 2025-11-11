package internal

import "strings"

// normaliza valores (low/medium/high/critical)
func norm(s string) string {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "low", "baixa": return "low"
	case "medium", "média", "media": return "medium"
	case "high", "alta": return "high"
	case "critical", "critica", "crítica": return "critical"
	default: return "low"
	}
}

// Matriz simples ITIL-like
func PriorityAndSLA(impact, urgency string) (string, int) {
	im := norm(impact)
	ur := norm(urgency)

	table := map[string]map[string]struct{
		P string
		SLA int
	}{
		"low": {
			"low":      {"P4", 1440},
			"medium":   {"P3", 720},
			"high":     {"P3", 480},
			"critical": {"P2", 240},
		},
		"medium": {
			"low":      {"P3", 720},
			"medium":   {"P3", 480},
			"high":     {"P2", 240},
			"critical": {"P2", 180},
		},
		"high": {
			"low":      {"P3", 480},
			"medium":   {"P2", 240},
			"high":     {"P2", 180},
			"critical": {"P1", 120},
		},
		"critical": {
			"low":      {"P2", 240},
			"medium":   {"P2", 180},
			"high":     {"P1", 120},
			"critical": {"P1", 60},
		},
	}

	if row, ok := table[im]; ok {
		if cell, ok := row[ur]; ok {
			return cell.P, cell.SLA
		}
	}
	return "P4", 1440
}

// placeholder de categorização simples por palavras-chave
func Categorize(title, desc string) string {
	txt := strings.ToLower(title + " " + desc)
	switch {
	case strings.Contains(txt, "banco") || strings.Contains(txt, "sql"):
		return "Banco de Dados"
	case strings.Contains(txt, "rede") || strings.Contains(txt, "vpn"):
		return "Infraestrutura"
	case strings.Contains(txt, "login") || strings.Contains(txt, "autentica"):
		return "Aplicação"
	default:
		return "Geral"
	}
}

// resumo fallback (sem IA)
func FallbackSummary(req IncidentRequest) string {
	return "Incidente: " + req.Title + " | Impacto: " + req.Impact + " | Urgência: " + req.Urgency
}
