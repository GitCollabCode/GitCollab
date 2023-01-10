package models

var Languages = map[string]bool{
	"JavaScript": true,
	"Python":     true,
	"PHP":        true,
	"JAVA":       true,
	"C":          true,
	"C/C++":      true,
	"C# ":        true,
	"TypeScript": true,
	"Rust":       true,
	"Go":         true,
	"Lua":        true,
	"Shell":      true,
	"Kotlin":     true,
	"Ruby":       true,
}

func Keys(m map[string]bool) []string {
	keys := make([]string, len(m))
	idx := 0
	for k, _ := range m {
		keys[idx] = k
		idx++
	}
	return keys
}
