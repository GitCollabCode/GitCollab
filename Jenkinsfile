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
    environment {
        /*
        define your command in variable
        */
        remoteCommands =
        """touch ~/test;
           echo "test" > ~/test"""
    }
    stages {
        stage('Hello') {
            steps {
                setBuildStatus("Build pending", "PENDING");
                sh 'sleep 5'
                echo 'Hello World'
                sshagent(['dev-server']) {
                    sh 'ssh -v -o StrictHostKeyChecking=no gitcollab@192.168.1.120 $remoteCommands'
                }
            }
        }

        //TODO: Add a final stage to ssh inside live deployment server, issue docker down, build and issue docker up
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
