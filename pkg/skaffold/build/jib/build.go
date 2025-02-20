/*
Copyright 2019 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package jib

import (
	"context"
	"io"

	latest_v1 "github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest/v1"
)

// Build builds an artifact with Jib.
func (b *Builder) Build(ctx context.Context, out io.Writer, artifact *latest_v1.Artifact, tag string) (string, error) {
	t, err := DeterminePluginType(artifact.Workspace, artifact.JibArtifact)
	if err != nil {
		return "", err
	}

	switch t {
	case JibMaven:
		if b.pushImages {
			return b.buildJibMavenToRegistry(ctx, out, artifact.Workspace, artifact.JibArtifact, artifact.Dependencies, tag)
		}
		return b.buildJibMavenToDocker(ctx, out, artifact.Workspace, artifact.JibArtifact, artifact.Dependencies, tag)

	case JibGradle:
		if b.pushImages {
			return b.buildJibGradleToRegistry(ctx, out, artifact.Workspace, artifact.JibArtifact, artifact.Dependencies, tag)
		}
		return b.buildJibGradleToDocker(ctx, out, artifact.Workspace, artifact.JibArtifact, artifact.Dependencies, tag)

	default:
		return "", unknownPluginType(artifact.Workspace)
	}
}
