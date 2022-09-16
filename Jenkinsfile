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
    stages {
        stage('Hello') {
            steps {
                setBuildStatus("Build pending", "PENDING");
                sh 'sleep 5'
                echo 'Hello World'
            }
        }
        stage('Update Live Deployment Server') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'mqtt-server', 
                                usernameVariable: 'USER', 
                                passwordVariable: 'PASSWORD')]) {
                    sh "mosquitto_pub -h monkeymoment.duckdns.org -u $USER -P $PASSWORD -t \"dev-server\" -m \"update\""
                }
            }
        }
    }

    post {
        success {
            setBuildStatus("Build succeeded", "SUCCESS");
        }
        failure {
            setBuildStatus("Build failed", "FAILURE");
        }
    }
}
