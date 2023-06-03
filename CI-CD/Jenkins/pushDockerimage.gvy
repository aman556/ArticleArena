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

    // stage('Push image') {
    //     withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'dockerhub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
	// 		usr = USERNAME
	// 		pswd = PASSWORD
	// 	}
    //     docker.withRegistry('https://hub.docker.com/', 'dockerhub') {
    //         sh "docker login -u ${USERNAME} -p ${PASSWORD}"
    //         app.push("${env.BUILD_NUMBER}")
    //     }
    // }
    
    stage('Push image') {
        /* Finally, we'll push the image with two tags:
        * First, the incremental build number from Jenkins
        * Second, the 'latest' tag. */
        withCredentials([usernamePassword( credentialsId: 'dockerhub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {

            docker.withRegistry('', 'dockerhub') {
                sh "docker login -u ${USERNAME} -p ${PASSWORD}"
                myImage.push("${env.BUILD_NUMBER}")
                myImage.push("latest")
            }
        }
    }
    stage('Trigger ManifestUpdate') {
        echo "triggering updatemanifestjob"
        build job: 'updatearticlearenamanifest', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER)]
    }
}
