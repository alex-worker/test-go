#!/bin/sh

# Credits: http://stackoverflow.com/a/750191

git remote remove origin
git remote add origin ssh://git@github.com/alex-worker/test-go
git push --set-upstream origin master
chmod 600 ~/.ssh/ed_alex-worker
