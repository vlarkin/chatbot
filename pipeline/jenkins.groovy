pipeline {

    agent any

    parameters {
        choice(name: "OS", choices: ["linux", "macos"], description: "Pick OS")
        choice(name: "ARCH", choices: ["amd64", "arm64"], description: "Pick ARCH")
    }

    environment {
        REPO = "https://github.com/vlarkin/chatbot"
        DOCKER_REGISTRY = "https://ghcr.io/vlarkin"
        BRANCH = "master"
        GITHUB_TOKEN = credentials("github-token")
        TARGETARCH = "${params.ARCH}"
        TARGETOS = "${params.OS}"
    }

    stages {
        stage("clone") {
            steps {
                echo "Clone a repository"
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }

        stage("tests") {
            steps {
                echo "Run tests"
                sh "make tests"
            }
        }

        stage("registry") {
            steps {
                echo "Login to Container Registry"
                sh "echo $GITHUB_TOKEN_PSW | docker login $DOCKER_REGISTRY -u $GITHUB_TOKEN_USR --password-stdin"
            }
        }
        
        stage("dockerize") {
            steps {
                echo "Build and push a docker image"
                sh "make dockerize"
            }
        }
        stage("cleanup") {
            steps {
                echo "Cleanup the environment"
                sh "make clean"
                sh "docker system prune -a --volumes --force"
            }
        }
    }

}
