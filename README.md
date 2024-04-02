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
def build() {
  sh '[cmd_1]'
  sh '[cmd_2]'
}

def deploy() {
  sh '[cmd_1]'
  sh '[cmd_2]'
}

pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        build()
      }
    }
    stage('Deploy') {
      steps {
        deploy()
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
