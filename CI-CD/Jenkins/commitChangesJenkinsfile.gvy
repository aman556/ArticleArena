node {
    def app

    stage('Clone repository') {
        checkout scm
    }

    stage('Update GIT') {
            script {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    withCredentials([usernamePassword(credentialsId: 'github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {
                        //def encodedPassword = URLEncoder.encode("$GIT_PASSWORD",'UTF-8')
                        sh "git config user.email amansharma14041998@gmail.com"
                        sh "git config user.name aman556"
                        //sh "git switch master"
                        sh "cat deploy/deploy-server/backend-server-deployment.yaml"
                        sh "sed -i 's+aman55/articlearena.*+aman55/articlearena:${DOCKERTAG}+g' deploy/deploy-server/backend-server-deployment.yaml"
                        sh "cat deploy/deploy-server/backend-server-deployment.yaml"
                        sh "git add ."
                        sh "git commit -m 'Done by Jenkins Job changemanifest: ${env.BUILD_NUMBER}'"
                        sh "git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/${GIT_USERNAME}/ArticleArena.git HEAD:CI-CD-Branch"
      }
    }
  }
}
}
