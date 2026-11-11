package github

import "testing"

func noSha(string) bool { return false }

func TestPickLatestReleaseStableKeepsNewerInstalledPrerelease(t *testing.T) {
	releases := []Release{
		{TagName: "v1.1.0-beta.2", TargetCommitish: "main", Prerelease: true},
		{TagName: "v1.1.0-beta.1", TargetCommitish: "main", Prerelease: true},
		{TagName: "v1.0.0", TargetCommitish: "main"},
	}
	for _, tc := range []struct{ installed, want string }{
		{"v1.1.0-beta.1", "v1.1.0-beta.1"}, // stay on the pre-release instead of downgrading
		{"v1.0.0", "v1.0.0"},               // never move onto a pre-release
		{"v0.9.0", "v1.0.0"},
	} {
		got, err := pickLatestRelease(releases, "main", true, tc.installed, noSha)
		if err != nil || got.TagName != tc.want {
			t.Errorf("installed %s: got %q (err %v), want %s", tc.installed, got.TagName, err, tc.want)
		}
	}
}

func TestPickLatestReleaseShaTarget(t *testing.T) {
	const sha = "af08737994d2ddf060c0ee40391dd4bbe754ca97"
	releases := []Release{
		{TagName: "v1.0.2", TargetCommitish: sha},
		{TagName: "v1.0.1", TargetCommitish: "main"},
	}
	onMain := func(s string) bool { return s == sha }
	for _, tc := range []struct {
		installed string
		onBranch  func(string) bool
		want      string
	}{
		{"v1.0.1", onMain, "v1.0.2"}, // SHA target on branch counts as the branch
		{"v1.0.1", noSha, "v1.0.1"},  // SHA not on branch: skipped
		{"v1.0.2", noSha, "v1.0.2"},  // installed release is never downgraded
	} {
		got, err := pickLatestRelease(releases, "main", true, tc.installed, tc.onBranch)
		if err != nil || got.TagName != tc.want {
			t.Errorf("installed %s: got %q (err %v), want %s", tc.installed, got.TagName, err, tc.want)
		}
	}
}
