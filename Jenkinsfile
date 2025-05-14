// pipeline {
//     agent any
//     environment {
//         DOCKERHUB_CREDENTIALS = credentials('dockerhub-credentials')
//         IMAGE_NAME = 'yourusername/go-rest-api'
//         IMAGE_TAG = "${env.BUILD_NUMBER}"
//         // EMAIL_RECIPIENTS = 'your.email@gmail.com' // 알림 수신자
//     }
//     stages {
//         stage('Checkout') {
//             steps {
//                 // 사용자 github 레페지토리 경로 지정 
//                 git branch: 'main', url: 'https://github.com/your-repo/go-rest-api.git'
//             }
//         }
//         stage('Build Docker Image') {
//             steps {
//                 sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG} ."
//             }
//         }
//         stage('Push to Docker Hub') {
//             steps {
//                 sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
//                 sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}"
//             }
//         }
//     }
//     post {
//         always {
//             // 빌드 후 로컬 이미지 정리
//             sh "docker rmi ${IMAGE_NAME}:${IMAGE_TAG} || true"
//         }
//         success {
//             // 성공 시 Gmail 알림
//             // emailext (
//             //     to: "${EMAIL_RECIPIENTS}",
//             //     subject: "SUCCESS: Job ${env.JOB_NAME} [${env.BUILD_NUMBER}]",
//             //     body: """
//             //     <p>Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' succeeded.</p>
//             //     <p>Image pushed: ${IMAGE_NAME}:${IMAGE_TAG}</p>
//             //     <p>Check console output at <a href='${env.BUILD_URL}'>${env.BUILD_URL}</a></p>
//             //     """,
//             //     mimeType: 'text/html'
//             // )
//         }
//         failure {
//             // 실패 시 Gmail 알림
//             // emailext (
//             //     to: "${EMAIL_RECIPIENTS}",
//             //     subject: "FAILURE: Job ${env.JOB_NAME} [${env.BUILD_NUMBER}]",
//             //     body: """
//             //     <p>Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' failed.</p>
//             //     <p>Check console output at <a href='${env.BUILD_URL}'>${env.BUILD_URL}</a></p>
//             //     """,
//             //     mimeType: 'text/html'
//             // )
//         }
//     }
// }
pipeline {
    agent any
    environment {
        IMAGE_NAME = 'rlaehgns78/go-rest-api'
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