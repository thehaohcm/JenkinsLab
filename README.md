Init and run the Jenkins by docker:

```# sudo docker compose up -d```

If you encounter a "permission denied" error, run the following command:

```# chmod 777 -R /var/jenkins_home```

And then run the 1st cmd again


Jenkins DSL Example:
```
job('DSL-Example') {
    scm {
        git('https://github.com/thehaohcm/JenkinsLab.git')
    }
    triggers {
        scm('H/15 * * * *')
    }
    steps {
        golangProject('...')
    }
}
```

Jenkinsfile Pipeline Example:
```
pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        // send slack
        slackSend color: "#439FE0", message: "Build Started"

        // build image
        sh 'docker build -f project/Dockerfile -t thehaohcm/docker-gs-ping:latest ./project/'

        // push image
        sh 'docker image push thehaohcm/docker-gs-ping:latest'
      }
    }
    stage('Deploy') {
      steps {
        sh 'docker image pull thehaohcm/docker-gs-ping:latest'

      // stop an existing container
        sh '''
          if [ "$( docker container inspect -f '{{.State.Running}}' docker-gs-ping )" = "true" ]; then
            docker stop docker-gs-ping
          fi
          '''

      // start a container
        sh 'docker run -d --rm --name docker-gs-ping -p 8081:8080 thehaohcm/docker-gs-ping:latest'
      }
    }
  }

  post{
    success{
      echo "Cool :)"
    }
  }
}
```
