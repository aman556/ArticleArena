#!groovy

pipeline {
	agent none  stages {
  	stage('Maven Install') {
    	agent {
      	docker {
        	image 'maven:3.5.0'
        }
      }
      steps {
      	sh 'mvn clean install'
      }
    }
    stage('Docker Build') {
    	agent any
      steps {
      	sh "docker build -t aman55/articlearena:${DOCKERTAG} backend"
      }
    }
    stage('Docker Push') {
    	agent any
      steps {
      	withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
        	sh "docker login -u ${env.dockerHubUser} -p ${env.dockerHubPassword}"
          sh 'docker push shanem/spring-petclinic:latest'
        }
      }
    }

    stage('Trigger ManifestUpdate') {
            echo "triggering updatemanifestjob"
            build job: 'updatearticlearenamanifest', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER)]
    }
  }
}

// node {
//     def app

//     stage('Clone repository') {
//         checkout scm
//     }

//     stage('Build image') {   
//         app = docker.build("aman55/articlearena","-f backend/Dockerfile .")
//     }

//     stage('Push image') {
//         docker.withRegistry('https://hub.docker.com/repository/docker/aman55/articlearena', 'dockerhub') {
//             app.push("${env.BUILD_NUMBER}")
//         }
//     }
    
//     stage('Trigger ManifestUpdate') {
//                 echo "triggering updatemanifestjob"
//                 build job: 'updatearticlearenamanifest', parameters: [string(name: 'DOCKERTAG', value: env.BUILD_NUMBER)]
//     }
// }


