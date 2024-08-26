#!/usr/bin/env bash

set -euo pipefail

if ! command -v asdf 2>/dev/null 1>/dev/null; then
  echo "Please install the asdf version manager. https://asdf-vm.com"
  exit 1
fi

asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
asdf plugin add just https://github.com/olofvndrhr/asdf-just.git
asdf plugin add sqlc https://github.com/dylanrayboss/asdf-sqlc.git
asdf plugin add goose https://github.com/samhvw8/asdf-goose.git
asdf plugin add gum https://github.com/lwiechec/asdf-gum.git
asdf plugin add sqlite https://github.com/cLupus/asdf-sqlite.git

asdf install

gum format "# ✅ Setup complete!"
