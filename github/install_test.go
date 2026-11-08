package github

import "testing"

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
		got, err := pickLatestRelease(releases, "main", true, tc.installed)
		if err != nil || got.TagName != tc.want {
			t.Errorf("installed %s: got %q (err %v), want %s", tc.installed, got.TagName, err, tc.want)
		}
	}
}
