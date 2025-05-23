package version

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	hashicorpVersion "github.com/hashicorp/go-version"
)

func HighestReleaseType(numbers []ReleaseType) ReleaseType {
	highestType := numbers[0]
	for _, value := range numbers {
		if value > highestType {
			highestType = value
		}
	}
	return highestType
}

// IsValidSemver checks that the given input is a valid semver version
func IsValidSemver(input string) bool {
	_, versionErr := semver.NewVersion(input)
	return versionErr == nil
}

// IsStable checks if the input is a stable semver version
func IsStable(input string) bool {
	ver, verErr := semver.NewVersion(input)
	if verErr != nil {
		return false
	}

	// no prereleases
	if len(ver.Prerelease()) > 0 {
		return false
	}

	return true
}

// Format formats the version string into a Semantic Versioning (SemVer) string format.
// If the input string is not a valid version, an error is returned.
// The function returns the formatted version string without a v prefix.
func Format(input string) (string, error) {
	ver, verErr := semver.NewVersion(input)
	if verErr != nil {
		return "", fmt.Errorf("malformed version: %s", input)
	}

	return ver.String(), nil
}

// Compare compares two versions.
// It returns:
// - 0 if both versions are equal
// - -1 if `left` version is older than `right` version
// - 1 if `left` version is newer than `right` version
// Note: `Compare` only compares the major, minor, and patch version numbers.
// It does not support comparison of build numbers or pre-release labels such as alpha, beta, etc.
func Compare(left string, right string) (int, error) {
	leftVer, leftVerErr := semver.NewVersion(left)
	if leftVerErr != nil {
		return 0, fmt.Errorf("failed to compare versions. left version is invalid: %s", leftVerErr.Error())
	}

	rightVer, rightVerErr := semver.NewVersion(right)
	if rightVerErr != nil {
		return 0, fmt.Errorf("failed to compare versions. right version is invalid: %s", rightVerErr.Error())
	}

	return leftVer.Compare(rightVer), nil
}

// FulfillsConstraint checks if the given `version` fulfills the `constraint`.
// It returns `true` if `version` satisfies the constraints specified in `constraint`, and `false` otherwise.
// `constraint` should be a string that follows the format described in the semver 2.0.0 specification.
// Example constraint strings: ">=1.2.3", "^1.2.3", "2.0.0".
func FulfillsConstraint(version string, constraint string) bool {
	ver, vErr := semver.NewVersion(version)
	if vErr != nil {
		return false
	}

	// verify constraint
	constraints, constraintsErr := semver.NewConstraint(constraint)
	if constraintsErr != nil {
		return false
	}
	if constraints.Check(ver) {
		return true
	}

	return false
}

// Bump bumps a version component up by one.
// The releaseType parameter specifies the component to be bumped:
// ReleaseMajor bumps the major version number,
// ReleaseMinor bumps the minor version number,
// ReleasePatch bumps the patch version number.
// The returned string is the new version and the error is nil if bumping was successful.
func Bump(version string, releaseType ReleaseType) (string, error) {
	v, vErr := hashicorpVersion.NewSemver(version)
	if vErr != nil {
		return "", vErr
	}

	segments := v.Segments()
	if releaseType == ReleaseMajor {
		segments[0]++
		segments[1] = 0
		segments[2] = 0
	} else if releaseType == ReleaseMinor {
		segments[1]++
		segments[2] = 0
	} else if releaseType == ReleasePatch {
		segments[2]++
	} else {
		return "", errors.New("can't patch releaseType " + string(releaseType))
	}

	var buf bytes.Buffer
	_, _ = fmt.Fprintf(&buf, "%d.%d.%d", segments[0], segments[1], segments[2])
	if v.Prerelease() != "" {
		_, _ = fmt.Fprintf(&buf, "-%s", v.Prerelease())
	}
	if v.Metadata() != "" {
		_, _ = fmt.Fprintf(&buf, "+%s", v.Metadata())
	}

	return buf.String(), nil
}
