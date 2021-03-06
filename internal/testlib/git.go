package testlib

import (
	"testing"
	"time"

	"github.com/goreleaser/goreleaser/internal/git"
	"github.com/stretchr/testify/assert"
)

// GitInit inits a new git project.
func GitInit(t *testing.T) {
	out, err := fakeGit("init")
	assert.NoError(t, err)
	assert.Contains(t, out, "Initialized empty Git repository")
	assert.NoError(t, err)
}

// GitRemoteAdd adds the given url as remote.
func GitRemoteAdd(t *testing.T, url string) {
	out, err := fakeGit("remote", "add", "origin", url)
	assert.NoError(t, err)
	assert.Empty(t, out)
}

// GitCommit creates a git commits.
func GitCommit(t *testing.T, msg string) {
	GitCommitWithDate(t, msg, time.Time{})
}

// GitCommitWithDate creates a git commit with a commit date.
func GitCommitWithDate(t *testing.T, msg string, commitDate time.Time) {
	env := (map[string]string)(nil)
	if !commitDate.IsZero() {
		env = map[string]string{
			"GIT_COMMITTER_DATE": commitDate.Format(time.RFC1123Z),
		}
	}
	out, err := fakeGitEnv(env, "commit", "--allow-empty", "-m", msg)
	assert.NoError(t, err)
	assert.Contains(t, out, "master", msg)
}

// GitTag creates a git tag.
func GitTag(t *testing.T, tag string) {
	out, err := fakeGit("tag", tag)
	assert.NoError(t, err)
	assert.Empty(t, out)
}

// GitAdd adds all files to stage.
func GitAdd(t *testing.T) {
	out, err := fakeGit("add", "-A")
	assert.NoError(t, err)
	assert.Empty(t, out)
}

func fakeGitEnv(env map[string]string, args ...string) (string, error) {
	var allArgs = []string{
		"-c", "user.name='GoReleaser'",
		"-c", "user.email='test@goreleaser.github.com'",
		"-c", "commit.gpgSign=false",
		"-c", "log.showSignature=false",
	}
	allArgs = append(allArgs, args...)
	return git.RunEnv(env, allArgs...)
}

func fakeGit(args ...string) (string, error) {
	return fakeGitEnv(nil, args...)
}

// GitCheckoutBranch allows us to change the active branch that we're using.
func GitCheckoutBranch(t *testing.T, tag string) {
	out, err := fakeGit("checkout", "-b", tag)
	assert.NoError(t, err)
	assert.Contains(t, out, tag)
}
