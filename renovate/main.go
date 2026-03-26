package main

import (
	"context"
	"dagger/renovate/internal/dagger"
)

type Renovate struct{}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Renovate) RenovateScan(
	ctx context.Context,
	repository string,
	// +optional
	baseBranch string,
	renovateToken string,
	// +optional
	logLevel string,
	//TODO
) (string, error) {
	return dag.Container().
		From("renovate/renovate:43.91").
		WithSecretVariable("RENOVATE_REPOSITORIES", baseBranch).
		WithSecretVariable("RENOVATE_TOKEN", renovateToken).
		WithSecretVariable("RENOVATE_BASE_BRANCHES", baseBranch).
		WithEnvVariable("LOG_LEVEL", logLevel).
		WithExec([]string{"renovate"}).
		Stdout(ctx)
}