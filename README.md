Install docker in linux distro: https://docs.docker.com/engine/install/ubuntu/

Install k3s and create a local kubernetes cluster in Linux:
```
# curl -sfL https://get.k3s.io | sh -
k3d cluster create demo-cluster
```

Install Helm and Postgres (run on port 31432) (OPTIONAL):
```
# curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
# helm install stockvn-db bitnami/postgresql --set service.type=NodePort --set service.nodePort=31432 --set auth.postgresPassword=[YOUR_PASSWORD] --set primary.service.ports.postgresql=5432
# kubectl patch svc stockvn-db-postgresql -p '{"spec": {"type": "NodePort", "ports": [{"port": 5432, "targetPort": 5432, "nodePort": 31432}]}}'
```

Init and run the Jenkins container:

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
