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
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-credentials') // Docker Hub 자격 증명 ID
        IMAGE_NAME = 'rlaehgns78/go-rest-api' // Docker Hub 이미지 이름
        IMAGE_TAG = "${env.BUILD_NUMBER}" // 빌드 번호를 태그로 사용
    }
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main',
                    credentialsId: 'github-credentials',
                    url: 'https://github.com/Hand-Some-Guy/example-Sevice.git'
            }
        }
        stage('Build Docker Image') {
            steps {
                // Docker 이미지 빌드
                sh "docker build -t ${IMAGE_NAME}:${IMAGE_TAG} ."
            }
        }
        stage('Push to Docker Hub') {
            steps {
                // Docker Hub 로그인 및 푸시
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh "docker push ${IMAGE_NAME}:${IMAGE_TAG}"
            }
        }
    }
    post {
        always {
            // 빌드 후 로컬 이미지 정리
            sh "docker rmi ${IMAGE_NAME}:${IMAGE_TAG} || true"
        }
        success {
            echo "Docker 이미지 ${IMAGE_NAME}:${IMAGE_TAG}가 성공적으로 푸시되었습니다."
        }
        failure {
            echo "빌드 또는 푸시에 실패했습니다."
        }
    }
}