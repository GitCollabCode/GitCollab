void setBuildStatus(String message, String state) {
  step([
      $class: "GitHubCommitStatusSetter",
      reposSource: [$class: "ManuallyEnteredRepositorySource", url: env.GIT_URL],
      contextSource: [$class: "ManuallyEnteredCommitContextSource", context: "ci/jenkins/build-status"],
      errorHandlers: [[$class: "ChangingBuildStatusErrorHandler", result: "UNSTABLE"]],
      statusResultSource: [ $class: "ConditionalStatusResultSource", results: [[$class: "AnyBuildResult", message: message, state: state]] ]
  ]);
}

pipeline {
    agent any

    tools {
        go 'go1.19.1'
        nodejs 'node18.9.0'
    }

    environment {
        GO111MODULE = 'on'
        GOBIN = '/tmp/go-bin'
    }

    stages {
        stage('Setup') {
            steps {
                sh 'printenv'
                script {
                    if (env.GIT_BRANCH == "origin/main") {
                        setBuildStatus("Build pending", "PENDING");
                    }
                }
                echo 'Installing dependencies...'
                sh 'go version'
                sh 'go mod vendor'
                sh 'curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $WORKSPACE v1.49.0'
            }
        }

        stage('Go Build') {
            steps {
                echo 'Running Go build...'
                sh 'go build $WORKSPACE/cmd/gitcollab/main.go'
            }
        }

        stage('Go Test') {
            steps {
                echo 'Running Go tests...'
                sh 'go test $WORKSPACE/... -v'
            }
        }

        stage('Go Code Analysis') {
            // disable stage
            // when {
            //     expression { false }
            // }
            steps {
                echo 'Preforming Go Code Analysis...'
                sh '$WORKSPACE/golangci-lint --timeout=5m run'
            }
        }

        stage('npm Install') {
            steps {
                dir('web') {
                    echo 'Running npm install...'
                    sh 'npm install'
                }
            }
        }

        stage('npm Build') {
            steps {
                dir('web') {
                    echo 'Running npm build...'
                    sh 'npm run build'
                }
            }
        }

        stage('npm Test') {
            steps {
                dir('web') {
                    echo 'Running npm test...'
                    sh 'npm test'
                }
            }
        }

        stage('Update Live Deployment Server') {
            when {
                expression { env.GIT_BRANCH == "origin/main" }
            }
            steps {
                echo 'Updating live deployment server with new changes...'
                withCredentials([usernamePassword(credentialsId: 'mqtt-server', 
                                usernameVariable: 'USER', 
                                passwordVariable: 'PASSWORD')]) {
                    //Fix credential warn
                    sh "mosquitto_pub -h monkeymoment.duckdns.org -u $USER -P $PASSWORD -t \"dev-server\" -m \"update\""
                }
            }
        }
    } 

    post {
        always {
			cleanWs()
		}
        success {
            script {
                if (env.GIT_BRANCH == "origin/main") {
                    setBuildStatus("Build succeeded", "SUCCESS");
                }
            }
        }
        failure {
            script {
                if (env.GIT_BRANCH == "origin/main") {
                    setBuildStatus("Build failed", "FAILURE");
                }
            }
        }
    }
}
