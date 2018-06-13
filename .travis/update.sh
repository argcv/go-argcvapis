#!/usr/bin/env bash

# if [[ $(git status --porcelain --untracked-files=no ) ]]; then # Ignore untracked files
if [[ $(git status --porcelain) ]]; then
  # Changes
  echo "Changed: make an auto commit"
  git add . -A
  git commit -m "Auto Commit: $(date '+%Y-%m-%d %H:%M:%S %Z')"
  TMP_BRANCH="tmp$(date '+%Y-%m-%d-%H-%M-%S-%Z')"
  git branch $TMP_BRANCH
  git checkout master
  git merge $TMP_BRANCH
  git push https://${GH_TOKEN}@github.com/argcv/go-argcvapis.git master
  echo "Updated..."
else
  echo "Not changed: skip..."
fi

