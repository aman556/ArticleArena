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
        withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'dockerhub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
			usr = USERNAME
			pswd = PASSWORD
		}
        docker.withRegistry('https://hub.docker.com/', 'dockerhub') {
            sh "docker login -u ${USERNAME} -p ${PASSWORD}"
            app.push("${env.BUILD_NUMBER}")
        }
    }
    
    stage('Trigger ManifestUpdate') {
        echo "triggering updatemanifestjob"
        build job: 'updatearticlearenamanifest', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER)]
    }
}
