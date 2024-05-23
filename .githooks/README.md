
# Using Gitleaks in a Pre-Commit Hook Script

Gitleaks is a SAST tool for detecting and preventing hardcoded secrets like passwords, api keys, and tokens in git repos.

## Installing a Pre-Commit Hook Script

The `pre-commit` script can be installed in any repository's `.git/hooks` folder. 

Navigate to your repository folder and run the following commands: 
```
curl -fsL https://raw.githubusercontent.com/vlarkin/chatbot/master/.githooks/pre-commit > .git/hooks/pre-commit
chmod 0755 .git/hooks/pre-commit
git config --local gitleaks.enable "true"
```

The `pre-commit` script automatically installs `gitleaks` tool if it’s not already present on the system.

## Verifying Gitleaks Installation

After installing the `pre-commit` hook script in your repository, try committing changes:

```
git add .
git commit -m "Test commit"
```

## Disabling Gitleaks pre-commit checks 

You can disable Gitleaks in the pre-commit hook script with this command:
```
git config --local gitleaks.enable "false"
```
