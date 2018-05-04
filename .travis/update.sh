#!/usr/bin/env bash

# if [[ $(git status --porcelain --untracked-files=no ) ]]; then # Ignore untracked files
if [[ $(git status --porcelain) ]]; then
  # Changes
  echo "Changed: make an auto commit"
  git commit -m "Auto Commit: $(date '+%Y-%m-%d %H:%M:%S')"
  git push https://${GH_TOKEN}@github.com/argcv/go-argcvapis.git
  echo "Updated..."
else
  echo "Not changed: skip..."
fi

