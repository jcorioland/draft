package version

type Version struct {
	SemVer string
	GitCommit string
	GitTreeState string
}

func (v *Version) String() string {
	return v.SemVer
}

var (
	// Release is the current release of Prow.
	// Update this whenever making a new release.
	// The release is of the format Major.Minor.Patch[-Prerelease][+BuildMetadata]
	//
	// Increment major number for new feature additions and behavioral changes.
	// Increment minor number for bug fixes and performance enhancements.
	// Increment patch number for critical fixes to existing releases.
	Release = "v0.1.0"

	// BuildMetadata is extra build time data
	BuildMetadata = ""
	// GitCommit is the git sha1
	GitCommit = ""
	// GitTreeState is the state of the git tree
	GitTreeState = ""
)

// getVersion returns the semver string of the version
func getVersion() string {
	if BuildMetadata == "" {
		return Release
	}
	return Release + "+" + BuildMetadata
}

// GetVersion returns the semver interpretation of the version
func GetVersion() *Version {
	return &Version{
		SemVer:       getVersion(),
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
	}
}