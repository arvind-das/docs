#!/bin/bash

set -euxo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")/../.."

echo "test building Go SDK examples"
cd code-examples/sdk/go
go build ./...
