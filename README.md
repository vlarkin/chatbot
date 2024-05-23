## chatbot

#### Yet Another Telegram Bot with very simple behavior 

How to build and execute it:
```
go get
go build -ldflags "-X="github.com/vlarkin/chatbot/cmd.appVersion=v1.0.0
export TELE_TOKEN="1234567890:JKwwWqp9n_SkIVN51eRKU78aB5f5O_nQuWd"
./chatbot start
```

Supported commands:
```
/start hello
/start joke
```

This bot was tested here https://t.me/RudeGnomeBot 
 

A digram for development workflow created in GitHub Actions: 
 
![Image](/images/workflow.png)

#### Install a pre commit hook script with gitleaks

Gitleaks is a SAST tool for detecting and preventing hardcoded secrets like passwords, api keys, and tokens in git repos. 
 
After cloning this repository to a local folder, navigate to that folder and run the following commands: 
```
cp .githooks/pre-commit .git/hooks/pre-commit
chmod 0755 .git/hooks/pre-commit
git config --local gitleaks.enable "true"
```

The installed pre-commit hook script automatically installs 'gitleaks' tool if it’s not already present on the system. 

You can disable Gitleaks in the pre-commit hook script with this command:
```
git config --local gitleaks.enable "false"
```

More details about the available git hook scripts can be found [here](/.githooks/README.md) 
