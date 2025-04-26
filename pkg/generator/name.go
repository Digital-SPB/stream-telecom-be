package generator

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/go-faker/faker/v4"
)

type nameGenerator struct {
	rand *rand.Rand
}

func newNameGenerator(seed int64) *nameGenerator {
	return &nameGenerator{
		rand: rand.New(rand.NewSource(seed)),
	}
}

func (g *nameGenerator) getCompanyName() string {
	// Get random company type
	companyTypes := []string{"Corp", "Inc", "LLC", "Group", "Labs", "Technologies", "Solutions"}
	companyType := companyTypes[g.rand.Intn(len(companyTypes))]

	// Create deterministic fake data based on our random source
	oldRand := rand.New(rand.NewSource(1)) // backup global rand
	rand.Seed(g.rand.Int63())              // set our seeded rand

	// Generate name components
	words := []string{
		faker.Word(),
		faker.Word(),
	}

	// Restore global rand
	rand.Seed(oldRand.Int63())

	// Process words
	for i, word := range words {
		// Remove any special characters and numbers
		word = strings.Map(func(r rune) rune {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				return r
			}
			return -1
		}, word)
		words[i] = capitalize(word)
	}

	// Combine components
	return fmt.Sprintf("%s %s %s", words[0], words[1], companyType)
}

// GenerateCampaignName generates a deterministic campaign name based on ID
func GenerateCampaignName(id int64) string {
	generator := newNameGenerator(id)
	return generator.getCompanyName()
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return fmt.Sprintf("%c%s", toUpper(s[0]), s[1:])
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}
	return c
} 