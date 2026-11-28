# Built by deploy/update.sh from the repo root: docker build -f deploy/Dockerfile --build-arg VERSION=<git describe> .
# The build stage is task build:release, so the server runs the same code and tests as a release and
# produces the archives the app hands out for taking a character home (the archives stage).
FROM golang:1.27 AS build
RUN apt-get update && apt-get install -y --no-install-recommends zip \
	&& rm -rf /var/lib/apt/lists/* \
	&& go install github.com/go-task/task/v3/cmd/task@latest
WORKDIR /src
COPY . .
ARG VERSION
# scripts/build.sh runs git rev-parse, so .git stays in the context.
RUN VERSION=$VERSION /go/bin/task build:release

FROM scratch AS archives
COPY --from=build /src/dist/ /

# Last stage, so a plain docker build gives the app image.
FROM gcr.io/distroless/static:nonroot
# The release stage folder is named after the version, hence the glob. DATA_DIR comes from compose.yaml.
COPY --from=build /src/build/stage/*-linux-amd64/gm-session-manager /gm-session-manager
ENTRYPOINT ["/gm-session-manager"]
