#!/usr/bin/env bash
set -euo pipefail

trap 'dirs -c' EXIT
shopt -s nullglob

SCRIPT_DIR="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
TESTDATA_DIR="$SCRIPT_DIR/../../internal/generate/test_examples"
ROOT_DIR="$SCRIPT_DIR/../.."

if [[ ! -d "$TESTDATA_DIR" ]]; then
  echo "ERROR: Testdata directory does not exist: $TESTDATA_DIR" >&2
  exit 1
fi

BIN_DIR="$ROOT_DIR/_build_tmp"
TAG_BIN="$BIN_DIR/terragrunt-atlantis-config"
trap 'rm -rf "$BIN_DIR"' EXIT

pushd "$ROOT_DIR" >/dev/null
rm -rf "$BIN_DIR"
mkdir -p "$BIN_DIR"
go build -o "$TAG_BIN" ./cmd/generate
popd >/dev/null

pushd "$TESTDATA_DIR" >/dev/null

echo "Reset atlantis.yaml files and terragrunt stacks in testdata case directories..."
# For each immediate child of testdata, check base/target and remove atlantis.yaml
for case_dir in *; do
  [[ -d "$case_dir" ]] || continue
  echo "==> Case: $case_dir"

  pushd "$case_dir" >/dev/null

  find . -type d -name '.terragrunt-stack' -prune -exec rm -rf {} +
  terragrunt stack generate --non-interactive --log-disable=true
  popd  >/dev/null
done

git add .

popd  >/dev/null

echo
echo "Done resetting \`internal/generate\` test data."
