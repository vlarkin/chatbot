
# Using Gitleaks in a Pre-Commit Hook Script

Gitleaks is a SAST tool for detecting and preventing hardcoded secrets like passwords, api keys, and tokens in git repos.

## Installing a Pre-Commit Hook Script to Install and Run Gitleaks

Copy the pre-commit hook script to the `.git/hooks` directory and make it executable. Then enable Gitleaks using Git config options.

```
cp .githooks/pre-commit .git/hooks/pre-commit
chmod 0755 .git/hooks/pre-commit
git config --local gitleaks.enable "true"
```

## Verifying Gitleaks Installation

After installing the pre-commit hook script and enabling Gitleaks in Git config, try to commit your changes:

```
git add .
git commit -m "Test commit"
```
