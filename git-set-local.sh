#!/bin/sh

git config --local user.name "alex-worker"
git config --local user.email "alex-worker@no-mail.com"
git config --local core.sshCommand 'ssh -i ~/.ssh/ed_alex-worker -o IdentityAgent=none'
