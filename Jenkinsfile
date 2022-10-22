void setBuildStatus(String message, String state) {
  step([
      $class: "GitHubCommitStatusSetter",
      reposSource: [$class: "ManuallyEnteredRepositorySource", url: "https://github.com/GitCollabCode/GitCollab"],
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
                setBuildStatus("Build pending", "PENDING");
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go mod vendor'
                sh 'curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOBIN) v1.49.0'
            }
        }

        stage('Go Build') {
            steps {
                echo 'Building [Go]..'
                sh 'go build $WORKSPACE/cmd/gitcollab/main.go'
            }
        }

        stage('Go Test') {
            when {
                expression { false }
            }
            echo 'UNIT TEST EXECUTION STARTED'
            sh 'go test ./...'
        }

        stage('Go Code Analysis') {
            when {
                expression { false }
            }
            steps {
                echo 'Skipping, broken on jenkins please make sure to run golangci-lint locally!'
            }
            // steps {
            //     echo 'Preforming Code Analysis [Go]..'
            //     sh '$(go env GOBIN)/golangci-lint --timeout=5m run'
            // }
        }

        stage('npm Install') {
            steps {
                dir('web') {
                    sh 'npm install'
                }
            }
        }

        stage('npm Build') {
            steps {
                dir('web') {
                    sh 'npm run build'
                }
            }
        }

        stage('npm Test') {
            steps {
                dir('web') {
                    sh 'npm test'
                }
            }
        }

        stage('Update Live Deployment Server') {
            steps {
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
            setBuildStatus("Build succeeded", "SUCCESS");
        }
        failure {
            setBuildStatus("Build failed", "FAILURE");
        }
    }
}
