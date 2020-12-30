//nolint:dupl
package externalplugins

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestValidateConfig(t *testing.T) {
	testcases := []struct {
		name           string
		lgtm           *TiCommunityLgtm
		merge          *TiCommunityMerge
		owners         *TiCommunityOwners
		autoresponders *TiCommunityAutoresponder
		blunderbuss    *TiCommunityBlunderbuss
		expected       error
	}{
		{
			name: "https pull owners URL",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			expected: nil,
		},
		{
			name: "http pull owners URL",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "http://bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "http://bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "http://bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "http://bots.tidb.io/ti-community-bot",
			},
			expected: nil,
		},
		{
			name: "invalid lgtm pull owners URL",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "http/bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			expected: fmt.Errorf("parse \"http/bots.tidb.io/ti-community-bot\": invalid URI for request"),
		},
		{
			name: "invalid merge pull owners URL",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "http/bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			expected: fmt.Errorf("parse \"http/bots.tidb.io/ti-community-bot\": invalid URI for request"),
		},
		{
			name: "invalid owners sig endpoint",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "https/bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			expected: fmt.Errorf("parse \"https/bots.tidb.io/ti-community-bot\": invalid URI for request"),
		},
		{
			name: "invalid blunderbuss regex",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   "?[)",
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			expected: fmt.Errorf("error parsing regexp: missing argument to repetition operator: `?`"),
		},
		{
			name: "invalid blunderbuss pull owners",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https/bots.tidb.io/ti-community-bot",
			},
			expected: fmt.Errorf("parse \"https/bots.tidb.io/ti-community-bot\": invalid URI for request"),
		},
		{
			name: "invalid blunderbuss max reviewer count",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   -1,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			expected: fmt.Errorf("max reviewer count must more than 0"),
		},
		{
			name: "invalid blunderbuss grace period duration",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"tidb-community-bots/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			merge: &TiCommunityMerge{
				Repos:              []string{"tidb-community-bots/test-dev"},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			owners: &TiCommunityOwners{
				Repos:       []string{"tidb-community-bots/test-dev"},
				SigEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			autoresponders: &TiCommunityAutoresponder{
				Repos: []string{"tidb-community-bots/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:               []string{"tidb-community-bots/test-dev"},
				MaxReviewerCount:    2,
				ExcludeReviewers:    []string{},
				PullOwnersEndpoint:  "https://bots.tidb.io/ti-community-bot",
				GracePeriodDuration: -1,
			},
			expected: fmt.Errorf("grace period duration must not less than 0"),
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			config := Configuration{
				TiCommunityLgtm: []TiCommunityLgtm{
					*tc.lgtm,
				},
				TiCommunityMerge: []TiCommunityMerge{
					*tc.merge,
				},
				TiCommunityOwners: []TiCommunityOwners{
					*tc.owners,
				},
				TiCommunityAutoresponder: []TiCommunityAutoresponder{
					*tc.autoresponders,
				},
				TiCommunityBlunderbuss: []TiCommunityBlunderbuss{
					*tc.blunderbuss,
				},
			}
			actual := config.Validate()

			if tc.expected == nil && actual != nil {
				t.Errorf("unexpected error: '%v'", actual)
			}
			if tc.expected != nil && actual == nil {
				t.Errorf("expected error '%v', but it is nil", tc.expected)
			}
			if tc.expected != nil && actual != nil && tc.expected.Error() != actual.Error() {
				t.Errorf("expected error '%v', but it is '%v'", tc.expected, actual)
			}
		})
	}
}

func TestLgtmFor(t *testing.T) {
	testcases := []struct {
		name        string
		lgtm        *TiCommunityLgtm
		org         string
		repo        string
		expectEmpty *TiCommunityLgtm
	}{
		{
			name: "Full name",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra/test-dev"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name: "Only org",
			lgtm: &TiCommunityLgtm{
				Repos:              []string{"ti-community-infra"},
				ReviewActsAsLgtm:   true,
				PullOwnersEndpoint: "http://bots.tidb.io/ti-community-bot",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name:        "Can not find",
			lgtm:        &TiCommunityLgtm{},
			org:         "ti-community-infra1",
			repo:        "test-dev1",
			expectEmpty: &TiCommunityLgtm{},
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			config := Configuration{TiCommunityLgtm: []TiCommunityLgtm{
				*tc.lgtm,
			}}
			lgtm := config.LgtmFor(tc.org, tc.repo)

			if tc.expectEmpty != nil {
				assert.DeepEqual(t, lgtm, &TiCommunityLgtm{})
			} else {
				assert.DeepEqual(t, lgtm.Repos, tc.lgtm.Repos)
			}
		})
	}
}

func TestMergeFor(t *testing.T) {
	testcases := []struct {
		name        string
		merge       *TiCommunityMerge
		org         string
		repo        string
		expectEmpty *TiCommunityMerge
	}{
		{
			name: "Full name",
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra/test-dev"},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name: "Only org",
			merge: &TiCommunityMerge{
				Repos:              []string{"ti-community-infra"},
				PullOwnersEndpoint: "http://bots.tidb.io/ti-community-bot",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name:        "Can not find",
			merge:       &TiCommunityMerge{},
			org:         "ti-community-infra1",
			repo:        "test-dev1",
			expectEmpty: &TiCommunityMerge{},
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			config := Configuration{TiCommunityMerge: []TiCommunityMerge{
				*tc.merge,
			}}

			merge := config.MergeFor(tc.org, tc.repo)

			if tc.expectEmpty != nil {
				assert.DeepEqual(t, merge, &TiCommunityMerge{})
			} else {
				assert.DeepEqual(t, merge.Repos, tc.merge.Repos)
			}
		})
	}
}

func TestOwnersFor(t *testing.T) {
	testcases := []struct {
		name        string
		owners      *TiCommunityOwners
		org         string
		repo        string
		expectEmpty *TiCommunityOwners
	}{
		{
			name: "Full name",
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra/test-dev"},
				SigEndpoint: "http://bots.tidb.io/ti-community-bot",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name: "Only org",
			owners: &TiCommunityOwners{
				Repos:       []string{"ti-community-infra"},
				SigEndpoint: "http://bots.tidb.io/ti-community-bot",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name:        "Can not find",
			owners:      &TiCommunityOwners{},
			org:         "ti-community-infra1",
			repo:        "test-dev1",
			expectEmpty: &TiCommunityOwners{},
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			config := Configuration{TiCommunityOwners: []TiCommunityOwners{
				*tc.owners,
			}}

			owners := config.OwnersFor(tc.org, tc.repo)

			if tc.expectEmpty != nil {
				assert.DeepEqual(t, owners, &TiCommunityOwners{})
			} else {
				assert.DeepEqual(t, owners.Repos, tc.owners.Repos)
			}
		})
	}
}

func TestLabelFor(t *testing.T) {
	testcases := []struct {
		name        string
		label       *TiCommunityLabel
		org         string
		repo        string
		expectEmpty *TiCommunityLabel
	}{
		{
			name: "Full name",
			label: &TiCommunityLabel{
				Repos:    []string{"ti-community-infra/test-dev"},
				Prefixes: []string{"status"},
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name: "Only org",
			label: &TiCommunityLabel{
				Repos:    []string{"ti-community-infra"},
				Prefixes: []string{"status"},
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name:        "Can not find",
			label:       &TiCommunityLabel{},
			org:         "ti-community-infra1",
			repo:        "test-dev1",
			expectEmpty: &TiCommunityLabel{},
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			config := Configuration{TiCommunityLabel: []TiCommunityLabel{
				*tc.label,
			}}

			label := config.LabelFor(tc.org, tc.repo)

			if tc.expectEmpty != nil {
				assert.DeepEqual(t, label, &TiCommunityLabel{})
			} else {
				assert.DeepEqual(t, label.Repos, tc.label.Repos)
			}
		})
	}
}

func TestAutoresponderFor(t *testing.T) {
	testcases := []struct {
		name          string
		autoresponder *TiCommunityAutoresponder
		org           string
		repo          string
		expectEmpty   *TiCommunityAutoresponder
	}{
		{
			name: "Full name",
			autoresponder: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra/test-dev"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name: "Only org",
			autoresponder: &TiCommunityAutoresponder{
				Repos: []string{"ti-community-infra"},
				AutoResponds: []AutoRespond{
					{
						Regex:   `(?mi)^/merge\s*$`,
						Message: "/run-all-test",
					},
				},
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name:          "Can not find",
			autoresponder: &TiCommunityAutoresponder{},
			org:           "ti-community-infra1",
			repo:          "test-dev1",
			expectEmpty:   &TiCommunityAutoresponder{},
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			config := Configuration{TiCommunityAutoresponder: []TiCommunityAutoresponder{
				*tc.autoresponder,
			}}

			autoresponder := config.AutoresponderFor(tc.org, tc.repo)

			if tc.expectEmpty != nil {
				assert.DeepEqual(t, autoresponder, &TiCommunityAutoresponder{})
			} else {
				assert.DeepEqual(t, autoresponder.Repos, tc.autoresponder.Repos)
			}
		})
	}
}

func TestBlunderbussFor(t *testing.T) {
	testcases := []struct {
		name        string
		blunderbuss *TiCommunityBlunderbuss
		org         string
		repo        string
		expectEmpty *TiCommunityBlunderbuss
	}{
		{
			name: "Full name",
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra/test-dev"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name: "Only org",
			blunderbuss: &TiCommunityBlunderbuss{
				Repos:              []string{"ti-community-infra"},
				MaxReviewerCount:   2,
				ExcludeReviewers:   []string{},
				PullOwnersEndpoint: "https://bots.tidb.io/ti-community-bot",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name:        "Can not find",
			blunderbuss: &TiCommunityBlunderbuss{},
			org:         "ti-community-infra1",
			repo:        "test-dev1",
			expectEmpty: &TiCommunityBlunderbuss{},
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			config := Configuration{TiCommunityBlunderbuss: []TiCommunityBlunderbuss{
				*tc.blunderbuss,
			}}

			blunderbuss := config.BlunderbussFor(tc.org, tc.repo)

			if tc.expectEmpty != nil {
				assert.DeepEqual(t, blunderbuss, &TiCommunityBlunderbuss{})
			} else {
				assert.DeepEqual(t, blunderbuss.Repos, tc.blunderbuss.Repos)
			}
		})
	}
}

func TestTarsFor(t *testing.T) {
	testcases := []struct {
		name        string
		tars        *TiCommunityTars
		org         string
		repo        string
		expectEmpty *TiCommunityTars
	}{
		{
			name: "Full name",
			tars: &TiCommunityTars{
				Repos:         []string{"ti-community-infra/test-dev"},
				Message:       "updated",
				OnlyWhenLabel: "trigger-update",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name: "Only org",
			tars: &TiCommunityTars{
				Repos:         []string{"ti-community-infra"},
				Message:       "updated",
				OnlyWhenLabel: "trigger-update",
			},
			org:  "ti-community-infra",
			repo: "test-dev",
		},
		{
			name:        "Can not find",
			tars:        &TiCommunityTars{},
			org:         "ti-community-infra1",
			repo:        "test-dev1",
			expectEmpty: &TiCommunityTars{},
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			config := Configuration{TiCommunityTars: []TiCommunityTars{
				*tc.tars,
			}}

			tars := config.TarsFor(tc.org, tc.repo)

			if tc.expectEmpty != nil {
				assert.DeepEqual(t, tars, &TiCommunityTars{})
			} else {
				assert.DeepEqual(t, tars.Repos, tc.tars.Repos)
			}
		})
	}
}

func TestSetBlunderbussDefaults(t *testing.T) {
	testcases := []struct {
		name                      string
		gracePeriodDuration       int
		expectGracePeriodDuration int
	}{
		{
			name:                      "default",
			gracePeriodDuration:       0,
			expectGracePeriodDuration: 5,
		},
		{
			name:                      "overwrite",
			gracePeriodDuration:       3,
			expectGracePeriodDuration: 3,
		},
	}

	for _, testcase := range testcases {
		tc := testcase
		t.Run(tc.name, func(t *testing.T) {
			c := &Configuration{
				TiCommunityBlunderbuss: []TiCommunityBlunderbuss{
					{
						GracePeriodDuration: tc.gracePeriodDuration,
					},
				},
			}

			c.setDefaults()

			for _, blunderbuss := range c.TiCommunityBlunderbuss {
				if blunderbuss.GracePeriodDuration != tc.expectGracePeriodDuration {
					t.Errorf("unexpected grace_period_duration: %v, expected: %v",
						blunderbuss.GracePeriodDuration, tc.expectGracePeriodDuration)
				}
			}
		})
	}
}
