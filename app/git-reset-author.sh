#!/bin/sh

# Credits: http://stackoverflow.com/a/750191

git filter-branch -f --env-filter "
    GIT_AUTHOR_NAME='alex-worker'
    GIT_AUTHOR_EMAIL='alex-worker@email'
    GIT_COMMITTER_NAME='alex-worker'
    GIT_COMMITTER_EMAIL='new@email'
  " HEAD
