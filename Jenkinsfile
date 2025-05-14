pipeline {
    agent any
    environment {
        IMAGE_NAME = 'rlaehgns0714/go-rest-api'
        IMAGE_TAG = "${env.BUILD_NUMBER}"
    }
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main',
                    credentialsId: 'github-credentials',
                    url: 'https://github.com/Hand-Some-Guy/example-Sevice.git'
                sh 'ls -la' // 디버깅용
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    // Docker 이미지 빌드
                    def dockerImage = docker.build("${IMAGE_NAME}:${IMAGE_TAG}")
                }
            }
        }
        stage('Push to Docker Hub') {
            steps {
                script {
                    // Docker Hub 인증 및 푸시
                    docker.withRegistry('https://index.docker.io/v1/', 'dockerhub-credentials') {
                        docker.image("${IMAGE_NAME}:${IMAGE_TAG}").push()
                    }
                }
            }
        }
    }
    post {
        always {
            // 로컬 이미지 정리
            sh "docker rmi ${IMAGE_NAME}:${IMAGE_TAG} || true"
        }
        success {
            emailext (
                to: 'asd561298@gmail.com',
                subject: "SUCCESS: Job ${env.JOB_NAME} [${env.BUILD_NUMBER}]",
                body: "Image pushed: ${IMAGE_NAME}:${IMAGE_TAG}<br>Check: <a href='${env.BUILD_URL}'>${env.BUILD_URL}</a>",
                mimeType: 'text/html'
            )
        }
        failure {
            emailext (
                to: 'asd561298@gmail.com',
                subject: "FAILURE: Job ${env.JOB_NAME} [${env.BUILD_NUMBER}]",
                body: "Build failed. Check: <a href='${env.BUILD_URL}'>${env.BUILD_URL}</a>",
                mimeType: 'text/html'
            )
        }
    }
}