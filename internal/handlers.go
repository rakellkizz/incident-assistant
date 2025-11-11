package internal

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AISummarizer interface {
	Summarize(title, description string) (string, error)
}

func AnalyzeHandler(ai AISummarizer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req IncidentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
			return
		}

		priority, sla := PriorityAndSLA(req.Impact, req.Urgency)
		category := Categorize(req.Title, req.Description)

		// tenta IA se houver chave, senão usa fallback
		var summary string
		usedAI := false

		if os.Getenv("OPENAI_API_KEY") != "" && ai != nil {
			if s, err := ai.Summarize(req.Title, req.Description); err == nil && s != "" {
				summary = s
				usedAI = true
			}
		}
		if summary == "" {
			summary = FallbackSummary(req)
		}

		c.JSON(http.StatusOK, IncidentResponse{
			Priority:   priority,
			SLAMinutes: sla,
			Category:   category,
			Summary:    summary,
			UsedAI:     usedAI,
		})
	}
}
