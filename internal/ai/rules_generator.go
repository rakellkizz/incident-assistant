package ai

import (
	"encoding/json"
	"fmt"
	"os"
)

type ProjectRules struct {
	ProjectName  string   `json:"project_name"`
	AIGuidelines []string `json:"ai_guidelines"`
}

func LoadRules(path string) (*ProjectRules, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var rules ProjectRules
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil, err
	}
	return &rules, nil
}

func PrintRules() {
	rules, err := LoadRules("project_rules.json") // lido da RAIZ do repo
	if err != nil {
		fmt.Println("❌ Erro ao carregar regras:", err)
		return
	}
	fmt.Println("🧠 Regras do Projeto:")
	fmt.Println("Projeto:", rules.ProjectName)
	for _, g := range rules.AIGuidelines {
		fmt.Println(" -", g)
	}
}
