pipeline {
    agent any

    environment {
        IMAGE_TAG = "${env.BUILD_ID}"
        DOCKER_USERNAME = "hawadwi" 
    }

    options {
        buildDiscarder(logRotator(numToKeepStr: '10'))
        timestamps()
        timeout(time: 1, unit: 'HOURS')
    }

    stages {
        stage('1. Checkout Code') {
            steps {
                echo '===== Starting Checkout Stage ====='
                checkout scm
                bat 'git log --oneline -1'
            }
        }

        stage('2. Run Unit Tests') {
            steps {
                echo '===== Running Unit Tests Sequentially using *.go ====='
                script {
                    def services = ['user-service', 'order-service', 'tracking-service', 'gudang-service', 'courier-service', 'report-service', 'payment-service']
                    for (service in services) {
                        echo "--- Testing ${service} ---"
                        dir(service) {
                            bat '''
                            go version
                            go mod download
                            go test -v *.go -coverprofile=coverage.out
                            '''
                        }
                    }
                }
            }
        }

        stage('3. Code Analysis (go vet)') {
            steps {
                echo '===== Running Code Analysis Sequentially using *.go ====='
                script {
                    def services = ['user-service', 'order-service', 'tracking-service', 'gudang-service', 'courier-service', 'report-service', 'payment-service']
                    for (service in services) {
                        echo "--- Vetting ${service} ---"
                        dir(service) {
                            bat 'go vet *.go'
                        }
                    }
                }
            }
        }

        stage('4. Build Docker Images') {
            parallel {
                stage('Build User Service Image') {
                    steps { dir('user-service') { bat """
                        docker build -t user-service:${IMAGE_TAG} .
                        docker tag user-service:${IMAGE_TAG} user-service:latest
                        docker tag user-service:${IMAGE_TAG} ${DOCKER_USERNAME}/user-service:${IMAGE_TAG}
                        docker tag user-service:${IMAGE_TAG} ${DOCKER_USERNAME}/user-service:latest
                    """ } }
                }
                stage('Build Order Service Image') {
                    steps { dir('order-service') { bat """
                        docker build -t order-service:${IMAGE_TAG} .
                        docker tag order-service:${IMAGE_TAG} order-service:latest
                        docker tag order-service:${IMAGE_TAG} ${DOCKER_USERNAME}/order-service:${IMAGE_TAG}
                        docker tag order-service:${IMAGE_TAG} ${DOCKER_USERNAME}/order-service:latest
                    """ } }
                }
                stage('Build Tracking Service Image') {
                    steps { dir('tracking-service') { bat """
                        docker build -t tracking-service:${IMAGE_TAG} .
                        docker tag tracking-service:${IMAGE_TAG} tracking-service:latest
                        docker tag tracking-service:${IMAGE_TAG} ${DOCKER_USERNAME}/tracking-service:${IMAGE_TAG}
                        docker tag tracking-service:${IMAGE_TAG} ${DOCKER_USERNAME}/tracking-service:latest
                    """ } }
                }
                stage('Build Gudang Service Image') {
                    steps { dir('gudang-service') { bat """
                        docker build -t gudang-service:${IMAGE_TAG} .
                        docker tag gudang-service:${IMAGE_TAG} gudang-service:latest
                        docker tag gudang-service:${IMAGE_TAG} ${DOCKER_USERNAME}/gudang-service:${IMAGE_TAG}
                        docker tag gudang-service:${IMAGE_TAG} ${DOCKER_USERNAME}/gudang-service:latest
                    """ } }
                }
                stage('Build Courier Service Image') {
                    steps { dir('courier-service') { bat """
                        docker build -t courier-service:${IMAGE_TAG} .
                        docker tag courier-service:${IMAGE_TAG} courier-service:latest
                        docker tag courier-service:${IMAGE_TAG} ${DOCKER_USERNAME}/courier-service:${IMAGE_TAG}
                        docker tag courier-service:${IMAGE_TAG} ${DOCKER_USERNAME}/courier-service:latest
                    """ } }
                }
                stage('Build Report Service Image') {
                    steps { dir('report-service') { bat """
                        docker build -t report-service:${IMAGE_TAG} .
                        docker tag report-service:${IMAGE_TAG} report-service:latest
                        docker tag report-service:${IMAGE_TAG} ${DOCKER_USERNAME}/report-service:${IMAGE_TAG}
                        docker tag report-service:${IMAGE_TAG} ${DOCKER_USERNAME}/report-service:latest
                    """ } }
                }
                stage('Build Payment Service Image') {
                    steps { dir('payment-service') { bat """
                        docker build -t payment-service:${IMAGE_TAG} .
                        docker tag payment-service:${IMAGE_TAG} payment-service:latest
                        docker tag payment-service:${IMAGE_TAG} ${DOCKER_USERNAME}/payment-service:${IMAGE_TAG}
                        docker tag payment-service:${IMAGE_TAG} ${DOCKER_USERNAME}/payment-service:latest
                    """ } }
                }
            }
        }

        stage('5. Run Functional Tests') {
            steps {
                echo '===== Starting Functional Tests ====='
                bat 'docker compose up -d'
                bat 'timeout /t 45 >nul'
                
                script {
                    try {
                        def testServices = [
                            [dir: 'user-service', db: 'userdb'],
                            [dir: 'order-service', db: 'orderdb'],
                            [dir: 'tracking-service', db: 'trackingdb'],
                            [dir: 'gudang-service', db: 'gudangdb'],
                            [dir: 'courier-service', db: 'courierdb'],
                            [dir: 'report-service', db: 'reportdb'],
                            [dir: 'payment-service', db: 'paymentdb']
                        ]
                        
                        for (ts in testServices) {
                            dir(ts.dir) {
                                bat """
                                set DB_HOST=host.docker.internal
                                set DB_PORT=3306
                                set DB_USER=root
                                set DB_PASSWORD=root
                                set DB_NAME=${ts.db}
                                go test -tags=functional -v -run Functional *.go
                                """
                            }
                        }
                    } finally {
                        bat 'docker compose down'
                    }
                }
            }
        }

        stage('6. Push Docker Images') {
            steps {
                echo '===== Pushing Docker Images ====='
                script {
                    withCredentials([usernamePassword(
                            credentialsId: 'hawadwi',
                            usernameVariable: 'DOCKER_USER',
                            passwordVariable: 'DOCKER_PASS')]) {

                        bat 'echo %DOCKER_PASS% | docker login -u %DOCKER_USER% --password-stdin'

                        bat """
                        docker push ${DOCKER_USERNAME}/user-service:${IMAGE_TAG}
                        docker push ${DOCKER_USERNAME}/user-service:latest

                        docker push ${DOCKER_USERNAME}/order-service:${IMAGE_TAG}
                        docker push ${DOCKER_USERNAME}/order-service:latest

                        docker push ${DOCKER_USERNAME}/tracking-service:${IMAGE_TAG}
                        docker push ${DOCKER_USERNAME}/tracking-service:latest

                        docker push ${DOCKER_USERNAME}/gudang-service:${IMAGE_TAG}
                        docker push ${DOCKER_USERNAME}/gudang-service:latest

                        docker push ${DOCKER_USERNAME}/courier-service:${IMAGE_TAG}
                        docker push ${DOCKER_USERNAME}/courier-service:latest

                        docker push ${DOCKER_USERNAME}/report-service:${IMAGE_TAG}
                        docker push ${DOCKER_USERNAME}/report-service:latest

                        docker push ${DOCKER_USERNAME}/payment-service:${IMAGE_TAG}
                        docker push ${DOCKER_USERNAME}/payment-service:latest
                        """
                        bat 'docker logout'
                    }
                }
            }
        }

        stage('7. Deploy to AKS') {
            steps {
                echo '===== Deploying to AKS ====='
                bat """
                kubectl set image deployment/user-service user-service=${DOCKER_USERNAME}/user-service:${IMAGE_TAG}
                kubectl set image deployment/order-service order-service=${DOCKER_USERNAME}/order-service:${IMAGE_TAG}
                kubectl set image deployment/tracking-service tracking-service=${DOCKER_USERNAME}/tracking-service:${IMAGE_TAG}
                kubectl set image deployment/gudang-service gudang-service=${DOCKER_USERNAME}/gudang-service:${IMAGE_TAG}
                kubectl set image deployment/courier-service courier-service=${DOCKER_USERNAME}/courier-service:${IMAGE_TAG}
                kubectl set image deployment/report-service report-service=${DOCKER_USERNAME}/report-service:${IMAGE_TAG}
                kubectl set image deployment/payment-service payment-service=${DOCKER_USERNAME}/payment-service:${IMAGE_TAG}
                """
            }
        }

        stage('8. Verify Deployment') {
            steps {
                echo '===== Verify Deployment ====='
                bat '''
                kubectl rollout status deployment/user-service --timeout=300s
                kubectl rollout status deployment/order-service --timeout=300s
                kubectl rollout status deployment/tracking-service --timeout=300s
                kubectl rollout status deployment/gudang-service --timeout=300s
                kubectl rollout status deployment/courier-service --timeout=300s
                kubectl rollout status deployment/report-service --timeout=300s
                kubectl rollout status deployment/payment-service --timeout=300s

                kubectl get pods
                kubectl get svc
                kubectl get deployments
                '''
            }
        }
    }

    post {
        success {
            echo '✓ Pipeline executed successfully!'
            echo 'All microservices deployed to AKS!'
        }
        failure {
            echo '✗ Pipeline failed. Check logs above.'
        }
        always {
            cleanWs()
        }
    }
}
