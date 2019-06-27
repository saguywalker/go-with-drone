pipeline{
    agent{
        docker{
            image 'golang:alpine'
        }
    }
    stages{
        stage('Format'){
            steps{
                sh 'go fmt $(go list ./... | grep -v /vendor/)'
                sh 'go vet $(ge list ./... | grep -v /vendor/)'
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
            }
        }
        stage('Run'){
            steps{
                sh 'go build'
                sh './go-with-drone=${GIT_COMMIT}'
            }
            post{
                success{
                    archiveArtifacts 'dist/go-with-drone'
                }
            }
        }
    }
}