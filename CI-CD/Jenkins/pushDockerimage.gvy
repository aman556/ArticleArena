node {
    def app

    stage('Clone repository') {
        checkout scm
    }

    stage('Build image') {
        sh "ls" 
        sh "cd backend" 
        sh "ls"  
        app = docker.build("aman55/articlearena","backend")
    }

    stage('Push image') {
        docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
            app.push("${env.BUILD_NUMBER}")
        }
    }
    
    stage('Trigger ManifestUpdate') {
        echo "triggering updatemanifestjob"
        build job: 'updatearticlearenamanifest', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER)]
    }
}
