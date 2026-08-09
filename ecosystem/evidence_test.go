package ecosystem

import (
	"testing"
)

func TestParseEvidenceLink(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantType EvidenceType
		wantID   string
		wantErr  bool
	}{
		{
			name:     "valid signal link",
			input:    "signal:oauth-idp-compatibility",
			wantType: SignalEvidence,
			wantID:   "oauth-idp-compatibility",
			wantErr:  false,
		},
		{
			name:     "valid customer link",
			input:    "customer:acme-corp",
			wantType: CustomerEvidence,
			wantID:   "acme-corp",
			wantErr:  false,
		},
		{
			name:     "valid market link",
			input:    "market:identity-governance",
			wantType: MarketEvidence,
			wantID:   "identity-governance",
			wantErr:  false,
		},
		{
			name:     "valid idea link",
			input:    "idea:IDEA-12345",
			wantType: IdeaEvidence,
			wantID:   "IDEA-12345",
			wantErr:  false,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "no colon",
			input:   "signal-oauth",
			wantErr: true,
		},
		{
			name:    "invalid type",
			input:   "unknown:some-id",
			wantErr: true,
		},
		{
			name:    "empty id",
			input:   "signal:",
			wantErr: true,
		},
		{
			name:     "whitespace handling",
			input:    "  signal:test-id  ",
			wantType: SignalEvidence,
			wantID:   "test-id",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link, err := ParseEvidenceLink(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEvidenceLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if link.Type != tt.wantType {
					t.Errorf("ParseEvidenceLink() type = %v, want %v", link.Type, tt.wantType)
				}
				if link.ID != tt.wantID {
					t.Errorf("ParseEvidenceLink() id = %v, want %v", link.ID, tt.wantID)
				}
			}
		})
	}
}

func TestEvidenceLinkString(t *testing.T) {
	link := EvidenceLink{
		Type: SignalEvidence,
		ID:   "oauth-idp-compatibility",
	}

	got := link.String()
	want := "signal:oauth-idp-compatibility"
	if got != want {
		t.Errorf("EvidenceLink.String() = %v, want %v", got, want)
	}
}

func TestParseEvidenceLinks(t *testing.T) {
	inputs := []string{
		"signal:test-1",
		"invalid",
		"customer:acme",
		"bad:type",
	}

	links, errs := ParseEvidenceLinks(inputs)

	if len(links) != 2 {
		t.Errorf("ParseEvidenceLinks() got %d links, want 2", len(links))
	}
	if len(errs) != 2 {
		t.Errorf("ParseEvidenceLinks() got %d errors, want 2", len(errs))
	}
}

func TestIsValidEvidenceType(t *testing.T) {
	tests := []struct {
		t    EvidenceType
		want bool
	}{
		{SignalEvidence, true},
		{CustomerEvidence, true},
		{MarketEvidence, true},
		{IdeaEvidence, true},
		{EvidenceType("unknown"), false},
		{EvidenceType(""), false},
	}

	for _, tt := range tests {
		got := IsValidEvidenceType(tt.t)
		if got != tt.want {
			t.Errorf("IsValidEvidenceType(%v) = %v, want %v", tt.t, got, tt.want)
		}
	}
}

func TestNoOpResolver(t *testing.T) {
	resolver := &NoOpResolver{}

	link := EvidenceLink{Type: SignalEvidence, ID: "test"}

	// Resolve should return nil
	evidence, err := resolver.Resolve(link)
	if err != nil {
		t.Errorf("NoOpResolver.Resolve() error = %v", err)
	}
	if evidence != nil {
		t.Errorf("NoOpResolver.Resolve() = %v, want nil", evidence)
	}

	// ValidateLinks should return empty
	errs := resolver.ValidateLinks([]EvidenceLink{link})
	if len(errs) != 0 {
		t.Errorf("NoOpResolver.ValidateLinks() = %v, want empty", errs)
	}

	// ResolveAll should return nil entries
	resolved, err := resolver.ResolveAll([]EvidenceLink{link, link})
	if err != nil {
		t.Errorf("NoOpResolver.ResolveAll() error = %v", err)
	}
	if len(resolved) != 2 {
		t.Errorf("NoOpResolver.ResolveAll() len = %d, want 2", len(resolved))
	}
}

func TestComputeEvidenceSummary(t *testing.T) {
	links := []EvidenceLink{
		{Type: SignalEvidence, ID: "signal-1"},
		{Type: SignalEvidence, ID: "signal-2"},
		{Type: CustomerEvidence, ID: "customer-1"},
		{Type: IdeaEvidence, ID: "idea-1"},
	}

	resolved := []*ResolvedEvidence{
		{
			Link:             links[0],
			FrustrationScore: 80,
			CustomerCount:    5,
			CustomerNames:    []string{"Acme", "BigCorp"},
			TotalARR:         100000, // $1000
			TotalVotes:       10,
		},
		{
			Link:             links[1],
			FrustrationScore: 60,
			CustomerCount:    3,
			CustomerNames:    []string{"BigCorp", "SmallCo"}, // BigCorp is duplicate
			TotalARR:         50000,                          // $500
			TotalVotes:       5,
		},
		nil, // Unresolved
		{
			Link:       links[3],
			TotalVotes: 20,
		},
	}

	summary := ComputeEvidenceSummary(links, resolved)

	// Link counts
	if summary.TotalLinks != 4 {
		t.Errorf("TotalLinks = %d, want 4", summary.TotalLinks)
	}
	if summary.SignalCount != 2 {
		t.Errorf("SignalCount = %d, want 2", summary.SignalCount)
	}
	if summary.CustomerCount != 1 {
		t.Errorf("CustomerCount = %d, want 1", summary.CustomerCount)
	}
	if summary.IdeaCount != 1 {
		t.Errorf("IdeaCount = %d, want 1", summary.IdeaCount)
	}

	// Resolution counts
	if summary.ResolvedCount != 3 {
		t.Errorf("ResolvedCount = %d, want 3", summary.ResolvedCount)
	}
	if summary.UnresolvedCount != 1 {
		t.Errorf("UnresolvedCount = %d, want 1", summary.UnresolvedCount)
	}

	// Aggregated metrics
	if summary.TotalFrustrationScore != 140 { // 80 + 60
		t.Errorf("TotalFrustrationScore = %f, want 140", summary.TotalFrustrationScore)
	}
	if summary.AvgFrustrationScore != 70 { // 140 / 2
		t.Errorf("AvgFrustrationScore = %f, want 70", summary.AvgFrustrationScore)
	}
	if summary.TotalARR != 150000 { // 100000 + 50000
		t.Errorf("TotalARR = %d, want 150000", summary.TotalARR)
	}
	if summary.TotalVotes != 35 { // 10 + 5 + 20
		t.Errorf("TotalVotes = %d, want 35", summary.TotalVotes)
	}

	// Deduplicated customers
	if summary.UniqueCustomers != 3 { // Acme, BigCorp, SmallCo
		t.Errorf("UniqueCustomers = %d, want 3", summary.UniqueCustomers)
	}

	// ARR in dollars
	if summary.ARRInDollars() != 1500 { // 150000 / 100
		t.Errorf("ARRInDollars() = %f, want 1500", summary.ARRInDollars())
	}
}

func TestEvidenceSummaryHelpers(t *testing.T) {
	empty := &EvidenceSummary{}
	if empty.HasEvidence() {
		t.Error("empty summary should not have evidence")
	}
	if empty.HasResolvedEvidence() {
		t.Error("empty summary should not have resolved evidence")
	}

	withLinks := &EvidenceSummary{TotalLinks: 1}
	if !withLinks.HasEvidence() {
		t.Error("summary with links should have evidence")
	}

	withResolved := &EvidenceSummary{TotalLinks: 1, ResolvedCount: 1}
	if !withResolved.HasResolvedEvidence() {
		t.Error("summary with resolved should have resolved evidence")
	}
}
