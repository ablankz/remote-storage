FROM golang:1.21 as dev

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
  && apt-get install -y --no-install-recommends curl \
  && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# Install aqua
ARG AQUA_INSTALLER_VERSION="3.0.0"
ARG AQUA_VERSION="2.25.1"
RUN curl -sSfL https://raw.githubusercontent.com/aquaproj/aqua-installer/v${AQUA_INSTALLER_VERSION}/aqua-installer \
  | bash -s -- -v v${AQUA_VERSION}

COPY ../aqua-policy.yaml /app/aqua-policy.yaml
COPY ../aqua-registry.yaml /app/aqua-registry.yaml
COPY ../aqua.yaml /app/aqua.yaml
COPY ../scripts/container-run.sh /init/run.sh

RUN chmod +x /init/run.sh

WORKDIR /app

ENV AQUA_GLOBAL_CONFIG=/app/aqua.yaml
ENV PATH="/root/.local/share/aquaproj-aqua/bin:$PATH"
ENV AQUA_POLICY_CONFIG=/app/aqua-policy.yaml

# Production build
FROM golang:1.21.8-alpine3.19 AS builder
WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/main.go

# Production image
FROM alpine:3.19 as prod
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 50054
CMD [ "/app/main" ]

