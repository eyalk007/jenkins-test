pipeline {
    agent any
    
    environment {
        JF_URL = credentials('jfrog-url')
        JF_ACCESS_TOKEN = credentials('jfrog-token')
        JF_GIT_TOKEN = credentials('git-token')
        JF_GIT_PROVIDER = 'github'
    }
    
    stages {
        stage('Debug Environment') {
            steps {
                script {
                    sh '''
                        echo "=== Jenkins Environment Variables ==="
                        echo "JENKINS_URL: ${JENKINS_URL}"
                        echo "GIT_URL: ${GIT_URL}"
                        echo "GIT_BRANCH: ${GIT_BRANCH}"
                        echo "GIT_LOCAL_BRANCH: ${GIT_LOCAL_BRANCH}"
                        echo "GIT_COMMIT: ${GIT_COMMIT}"
                        echo "CHANGE_ID: ${CHANGE_ID}"
                        echo "BRANCH_NAME: ${BRANCH_NAME}"
                        echo "BUILD_NUMBER: ${BUILD_NUMBER}"
                    '''
                }
            }
        }
        
        stage('Build Go Application') {
            steps {
                sh '''
                    echo "=== Building Go Application ==="
                    go version
                    go mod download
                    go build -o app .
                    echo "Build completed successfully"
                '''
            }
        }
        
        stage('Frogbot Security Scan') {
            steps {
                script {
                    def frogbotCmd = env.CHANGE_ID ? 'scan-pull-request' : 'scan-repository'
                    
                    sh """
                        echo "=== Frogbot Auto-Detection Test ==="
                        echo "Command: ${frogbotCmd}"
                        
                        echo "=== Downloading Frogbot Test Binary ==="
                        curl -fL -u admin:password "https://z0newscafrogbotdemo2.jfrogdev.org/artifactory/local-frogbot-repo/frogbot/test-auto-detection/v1/frogbot-linux-amd64/frogbot" -o frogbot
                        chmod +x frogbot
                        
                        echo "=== Running Frogbot (auto-detection enabled) ==="
                        ./frogbot ${frogbotCmd}
                    """
                }
            }
        }
    }
    
    post {
        always {
            echo 'Pipeline completed'
        }
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}

