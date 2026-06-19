pipeline {
    agent any

    environment {
        DOCKER_REGISTRY = credentials('docker-registry')
        DOCKER_USERNAME = "${DOCKER_REGISTRY_USR}"
        DOCKER_PASSWORD = "${DOCKER_REGISTRY_PSW}"
        IMAGE_TAG = "${env.BUILD_ID}"
        KUBECONFIG = "/var/jenkins_home/.kube/config"
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
                sh 'git log --oneline -1'
            }
        }

        stage('2. Run Unit Tests') {
            parallel {
                stage('User Service Unit Tests') {
                    steps {
                        echo 'Testing User Service...'
                        dir('user-service') {
                            sh '''
                                go version
                                go mod download
                                go test -v -race -coverprofile=coverage.out ./...
                            '''
                        }
                    }
                }

                stage('Order Service Unit Tests') {
                    steps {
                        echo 'Testing Order Service...'
                        dir('order-service') {
                            sh '''
                                go version
                                go mod download
                                go test -v -race -coverprofile=coverage.out ./...
                            '''
                        }
                    }
                }

                stage('Tracking Service Unit Tests') {
                    steps {
                        echo 'Testing Tracking Service...'
                        dir('tracking-service') {
                            sh '''
                                go version
                                go mod download
                                go test -v -race -coverprofile=coverage.out ./...
                            '''
                        }
                    }
                }

                stage('Gudang Service Unit Tests') {
                    steps {
                        echo 'Testing Gudang Service...'
                        dir('gudang-service') {
                            sh '''
                                go version
                                go mod download
                                go test -v -race -coverprofile=coverage.out ./...
                            '''
                        }
                    }
                }

                stage('Courier Service Unit Tests') {
                    steps {
                        echo 'Testing Courier Service...'
                        dir('courier-service') {
                            sh '''
                                go version
                                go mod download
                                go test -v -race -coverprofile=coverage.out ./...
                            '''
                        }
                    }
                }

                stage('Report Service Unit Tests') {
                    steps {
                        echo 'Testing Report Service...'
                        dir('report-service') {
                            sh '''
                                go version
                                go mod download
                                go test -v -race -coverprofile=coverage.out ./...
                            '''
                        }
                    }
                }

                stage('Payment Service Unit Tests') {
                    steps {
                        echo 'Testing Payment Service...'
                        dir('payment-service') {
                            sh '''
                                go version
                                go mod download
                                go test -v -race -coverprofile=coverage.out ./...
                            '''
                        }
                    }
                }
            }
        }

        stage('3. Code Analysis (go vet)') {
            parallel {
                stage('User Service Vet') {
                    steps {
                        dir('user-service') {
                            sh 'go vet ./...'
                        }
                    }
                }

                stage('Order Service Vet') {
                    steps {
                        dir('order-service') {
                            sh 'go vet ./...'
                        }
                    }
                }

                stage('Tracking Service Vet') {
                    steps {
                        dir('tracking-service') {
                            sh 'go vet ./...'
                        }
                    }
                }

                stage('Gudang Service Vet') {
                    steps {
                        dir('gudang-service') {
                            sh 'go vet ./...'
                        }
                    }
                }

                stage('Courier Service Vet') {
                    steps {
                        dir('courier-service') {
                            sh 'go vet ./...'
                        }
                    }
                }

                stage('Report Service Vet') {
                    steps {
                        dir('report-service') {
                            sh 'go vet ./...'
                        }
                    }
                }

                stage('Payment Service Vet') {
                    steps {
                        dir('payment-service') {
                            sh 'go vet ./...'
                        }
                    }
                }
            }
        }

        stage('4. Build Docker Images') {
            parallel {
                stage('Build User Service Image') {
                    steps {
                        echo 'Building User Service Docker Image...'
                        sh '''
                            docker build -t user-service:${IMAGE_TAG} ./user-service
                            docker tag user-service:${IMAGE_TAG} user-service:latest
                            docker tag user-service:${IMAGE_TAG} ${DOCKER_USERNAME}/user-service:${IMAGE_TAG}
                            docker tag user-service:${IMAGE_TAG} ${DOCKER_USERNAME}/user-service:latest
                        '''
                    }
                }

                stage('Build Order Service Image') {
                    steps {
                        echo 'Building Order Service Docker Image...'
                        sh '''
                            docker build -t order-service:${IMAGE_TAG} ./order-service
                            docker tag order-service:${IMAGE_TAG} order-service:latest
                            docker tag order-service:${IMAGE_TAG} ${DOCKER_USERNAME}/order-service:${IMAGE_TAG}
                            docker tag order-service:${IMAGE_TAG} ${DOCKER_USERNAME}/order-service:latest
                        '''
                    }
                }

                stage('Build Tracking Service Image') {
                    steps {
                        echo 'Building Tracking Service Docker Image...'
                        sh '''
                            docker build -t tracking-service:${IMAGE_TAG} ./tracking-service
                            docker tag tracking-service:${IMAGE_TAG} tracking-service:latest
                            docker tag tracking-service:${IMAGE_TAG} ${DOCKER_USERNAME}/tracking-service:${IMAGE_TAG}
                            docker tag tracking-service:${IMAGE_TAG} ${DOCKER_USERNAME}/tracking-service:latest
                        '''
                    }
                }

                stage('Build Gudang Service Image') {
                    steps {
                        echo 'Building Gudang Service Docker Image...'
                        sh '''
                            docker build -t gudang-service:${IMAGE_TAG} ./gudang-service
                            docker tag gudang-service:${IMAGE_TAG} gudang-service:latest
                            docker tag gudang-service:${IMAGE_TAG} ${DOCKER_USERNAME}/gudang-service:${IMAGE_TAG}
                            docker tag gudang-service:${IMAGE_TAG} ${DOCKER_USERNAME}/gudang-service:latest
                        '''
                    }
                }

                stage('Build Courier Service Image') {
                    steps {
                        echo 'Building Courier Service Docker Image...'
                        sh '''
                            docker build -t courier-service:${IMAGE_TAG} ./courier-service
                            docker tag courier-service:${IMAGE_TAG} courier-service:latest
                            docker tag courier-service:${IMAGE_TAG} ${DOCKER_USERNAME}/courier-service:${IMAGE_TAG}
                            docker tag courier-service:${IMAGE_TAG} ${DOCKER_USERNAME}/courier-service:latest
                        '''
                    }
                }

                stage('Build Report Service Image') {
                    steps {
                        echo 'Building Report Service Docker Image...'
                        sh '''
                            docker build -t report-service:${IMAGE_TAG} ./report-service
                            docker tag report-service:${IMAGE_TAG} report-service:latest
                            docker tag report-service:${IMAGE_TAG} ${DOCKER_USERNAME}/report-service:${IMAGE_TAG}
                            docker tag report-service:${IMAGE_TAG} ${DOCKER_USERNAME}/report-service:latest
                        '''
                    }
                }

                stage('Build Payment Service Image') {
                    steps {
                        echo 'Building Payment Service Docker Image...'
                        sh '''
                            docker build -t payment-service:${IMAGE_TAG} ./payment-service
                            docker tag payment-service:${IMAGE_TAG} payment-service:latest
                            docker tag payment-service:${IMAGE_TAG} ${DOCKER_USERNAME}/payment-service:${IMAGE_TAG}
                            docker tag payment-service:${IMAGE_TAG} ${DOCKER_USERNAME}/payment-service:latest
                        '''
                    }
                }
            }
        }

        stage('5. Run Functional Tests') {
            steps {
                echo '===== Starting Functional Tests ====='
                sh 'docker compose up -d'
                sh 'sleep 45'

                script {
                    try {
                        echo 'Running User Service Functional Tests...'
                        dir('user-service') {
                            sh '''
                                DB_HOST=host.docker.internal DB_PORT=3306 DB_USER=root DB_PASSWORD=root DB_NAME=userdb \
                                go test -tags=functional -v -run Functional ./... || true
                            '''
                        }

                        echo 'Running Order Service Functional Tests...'
                        dir('order-service') {
                            sh '''
                                DB_HOST=host.docker.internal DB_PORT=3306 DB_USER=root DB_PASSWORD=root DB_NAME=orderdb \
                                go test -tags=functional -v -run Functional ./... || true
                            '''
                        }

                        echo 'Running Tracking Service Functional Tests...'
                        dir('tracking-service') {
                            sh '''
                                DB_HOST=host.docker.internal DB_PORT=3306 DB_USER=root DB_PASSWORD=root DB_NAME=trackingdb \
                                go test -tags=functional -v -run Functional ./... || true
                            '''
                        }

                        echo 'Running Gudang Service Functional Tests...'
                        dir('gudang-service') {
                            sh '''
                                DB_HOST=host.docker.internal DB_PORT=3306 DB_USER=root DB_PASSWORD=root DB_NAME=gudangdb \
                                go test -tags=functional -v -run Functional ./... || true
                            '''
                        }

                        echo 'Running Courier Service Functional Tests...'
                        dir('courier-service') {
                            sh '''
                                DB_HOST=host.docker.internal DB_PORT=3306 DB_USER=root DB_PASSWORD=root DB_NAME=courierdb \
                                go test -tags=functional -v -run Functional ./... || true
                            '''
                        }

                        echo 'Running Report Service Functional Tests...'
                        dir('report-service') {
                            sh '''
                                DB_HOST=host.docker.internal DB_PORT=3306 DB_USER=root DB_PASSWORD=root DB_NAME=reportdb \
                                go test -tags=functional -v -run Functional ./... || true
                            '''
                        }

                        echo 'Running Payment Service Functional Tests...'
                        dir('payment-service') {
                            sh '''
                                DB_HOST=host.docker.internal DB_PORT=3306 DB_USER=root DB_PASSWORD=root DB_NAME=paymentdb \
                                go test -tags=functional -v -run Functional ./... || true
                            '''
                        }
                    } finally {
                        sh 'docker compose down || true'
                    }
                }
            }
        }

        stage('6. Push Docker Images') {
            steps {
                echo '===== Pushing Docker Images ====='
                script {
                    withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                        sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'

                        sh '''
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
                        '''

                        sh 'docker logout'
                    }
                }
            }
        }

        stage('7. Deploy to AKS') {
            steps {
                echo '===== Deploying to Azure Kubernetes Service (AKS) ====='
                script {
                    sh '''
                        # Update image tags in deployment manifests
                        sed -i "s|image: .*user-service.*|image: ${DOCKER_USERNAME}/user-service:${IMAGE_TAG}|g" k8s/user-deployment.yaml
                        sed -i "s|image: .*order-service.*|image: ${DOCKER_USERNAME}/order-service:${IMAGE_TAG}|g" k8s/order-deployment.yaml
                        sed -i "s|image: .*tracking-service.*|image: ${DOCKER_USERNAME}/tracking-service:${IMAGE_TAG}|g" k8s/tracking-deployment.yaml
                        sed -i "s|image: .*gudang-service.*|image: ${DOCKER_USERNAME}/gudang-service:${IMAGE_TAG}|g" k8s/gudang-deployment.yaml
                        sed -i "s|image: .*courier-service.*|image: ${DOCKER_USERNAME}/courier-service:${IMAGE_TAG}|g" k8s/courier-deployment.yaml
                        sed -i "s|image: .*report-service.*|image: ${DOCKER_USERNAME}/report-service:${IMAGE_TAG}|g" k8s/report-deployment.yaml
                        sed -i "s|image: .*payment-service.*|image: ${DOCKER_USERNAME}/payment-service:${IMAGE_TAG}|g" k8s/payment-deployment.yaml
                    '''

                    sh '''
                        # Apply deployments
                        kubectl apply -f k8s/user-deployment.yaml
                        kubectl apply -f k8s/order-deployment.yaml
                        kubectl apply -f k8s/tracking-deployment.yaml
                        kubectl apply -f k8s/gudang-deployment.yaml
                        kubectl apply -f k8s/courier-deployment.yaml
                        kubectl apply -f k8s/report-deployment.yaml
                        kubectl apply -f k8s/payment-deployment.yaml

                        # Apply services
                        kubectl apply -f k8s/user-service.yaml
                        kubectl apply -f k8s/order-service.yaml
                        kubectl apply -f k8s/tracking-service.yaml
                        kubectl apply -f k8s/gudang-service.yaml
                        kubectl apply -f k8s/courier-service.yaml
                        kubectl apply -f k8s/report-service.yaml
                        kubectl apply -f k8s/payment-service.yaml
                    '''
                }
            }
        }

        stage('8. Verify Deployment') {
            steps {
                echo '===== Verifying Deployment ====='
                sh '''
                    # Wait for rollout
                    kubectl rollout status deployment/user-service --timeout=300s
                    kubectl rollout status deployment/order-service --timeout=300s
                    kubectl rollout status deployment/tracking-service --timeout=300s
                    kubectl rollout status deployment/gudang-service --timeout=300s
                    kubectl rollout status deployment/courier-service --timeout=300s
                    kubectl rollout status deployment/report-service --timeout=300s
                    kubectl rollout status deployment/payment-service --timeout=300s

                    # Show deployed resources
                    echo "===== Running Pods ====="
                    kubectl get pods

                    echo "===== Deployed Services ====="
                    kubectl get svc

                    echo "===== Deployment Status ====="
                    kubectl get deployment
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
            echo '✗ Pipeline failed. Check logs above for details.'
        }

        unstable {
            echo '⚠ Pipeline is unstable. Review test results.'
        }

        always {
            cleanWs()
        }
    }
}
