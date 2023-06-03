node {
    def app

    stage('Clone repository') {
      

        checkout scm
    }

    stage('Build image') {
       sh 'cd backend' 
       app = docker.build("aman55/articlearena")
    }

    stage('Test image') {
  

        app.inside {
            sh 'echo "Tests passed"'
        }
    }

    stage('Push image') {
        
        docker.withRegistry('https://hub.docker.com/repository/docker/aman55/articlearena', 'dockerhub') {
            app.push("${env.BUILD_NUMBER}")
        }
    }
    
    stage('Trigger ManifestUpdate') {
                echo "triggering updatemanifestjob"
                build job: 'updatemanifest', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER)]
        }
}
