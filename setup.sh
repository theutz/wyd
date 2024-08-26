#!/usr/bin/env bash

if ! command -v asdf 2>/dev/null 1>/dev/null; then
  echo "Please install the asdf version manager. https://asdf-vm.com"
  exit 1
fi

asdf plugin add golang
asdf plugin add just https://github.com/olofvndrhr/asdf-just.git
asdf plugin add sqlc https://github.com/dylanrayboss/asdf-sqlc.git
asdf plugin add goose https://github.com/samhvw8/asdf-goose.git

asdf install

echo
echo "Setup complete!"
