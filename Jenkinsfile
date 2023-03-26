pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        echo 'build stage'
      }
    }

    stage('test') {
      steps {
        echo 'test stage'
      }
    }

    stage('deploy') {
      steps {
        sleep 30
        echo 'deploy stage'
      }
    }

  }
}