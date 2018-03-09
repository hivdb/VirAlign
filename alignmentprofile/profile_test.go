package alignmentprofile

import (
	a "github.com/hivdb/nucamino/types/amino"
	"testing"
)

var exampleProfile = AlignmentProfile{
	StopCodonPenalty:         1,
	GapOpeningPenalty:        2,
	GapExtensionPenalty:      3,
	IndelCodonOpeningBonus:   4,
	IndelCodonExtensionBonus: 5,
	ReferenceSequences: ReferenceSeqs{
		"A": a.ReadString("T"),
		"B": a.ReadString("V"),
	},
	GeneIndelScores: GenePositionalIndelScores{
		Gene("A"): PositionalIndelScores{
			3:  [2]int{4, 5},
			6:  [2]int{7, 8},
			-6: [2]int{7, 8},
			-9: [2]int{10, 11},
		},
	},
}

func TestGenes(t *testing.T) {
	genes := exampleProfile.Genes()
	expected := [2]Gene{"A", "B"}

	if len(genes) != len(expected) {
		t.Errorf("Expected %v to be %v", genes, expected)
	}

	for _, exp := range expected {
		found := false
		for _, g := range genes {
			if g == exp {
				found = true
			}
		}
		if !found {
			t.Errorf("Missing key: %v", exp)
		}
	}
}
