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
            steps {
                echo 'Skipping, no steps Go tests not setup!'
            }
            // Will need to get updated when some tests are added
            // steps {
			// 	withCredentials([usernamePassword(credentialsId: 'TEST_CREDENTIALS', usernameVariable: 'TEST_USERNAME', passwordVariable: 'TEST_PASSWORD'), string(credentialsId: 'KOPANO_SERVER_DEFAULT_URI', variable: 'KOPANO_SERVER_DEFAULT_URI')]) {
			// 		echo 'Testing..'
			// 		sh 'echo Kopano Server URI: \$KOPANO_SERVER_DEFAULT_URI'
			// 		sh 'echo Kopano Server Username: \$TEST_USERNAME'
			// 		sh 'go test -v -count=1 | tee tests.output'
			// 		sh 'PATH=$PATH:$GOBIN  go2xunit -fail -input tests.output -output tests.xml'
			// 	}
			// 	junit allowEmptyResults: true, testResults: 'tests.xml'
			// }
        }

        // stage('Go Code Analysis') {
        //     steps {
        //         echo 'Preforming Code Analysis [Go]..'
        //         sh 'PATH=$PATH:$GOBIN golangci-lint --timeout=5m run'
        //     }
        // }

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
