pipeline{
    agent{
        docker{
            image 'golang:alpine'
        }
    }
    stages{
        stage('Format'){
            steps{
                sh 'go fmt'
                sh 'go vet'
            }
        }
        stage('Test'){
            steps{
                sh 'go test'
            }
        }
        stage('Build'){
            steps{
                sh 'go build'
                sh 'ls -R'
            }
        }
        stage('Run'){
            steps{
                sh 'go build'
                sh './go-with-drone -commit=${GIT_COMMIT}'
            }
            post{
                success{
                    archiveArtifacts 'go-with-drone'
                }
            }
        }
    }
}