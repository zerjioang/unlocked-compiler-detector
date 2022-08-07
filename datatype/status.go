package datatype

import "fmt"

type VersionStatus struct {
	File         string `json:"file"`
	Errored      bool   `json:"errored,omitempty"`
	AffectedLine string
	Start        int
	End          int
	LineNo       int
	SuggestedFix string
}

func (s VersionStatus) Print() {
	fmt.Println("Unlocked Compiler Version Detected")
	fmt.Println("----------------------------------")
	fmt.Printf("File               : %s\n", s.File)
	fmt.Printf("Affected line (L%d) : %s\n", s.LineNo, s.AffectedLine)
	fmt.Printf("Suggested fix      : %s\n", s.SuggestedFix)
	fmt.Println("Confidence         : Very High")
	fmt.Println("Impact             : Informational")
	fmt.Printf("\n")
}
