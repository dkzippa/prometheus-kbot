pipeline {
    agent any
    parameters {
        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        choice(name: 'ARCH', choices: ['amd64', 'arm64', 'all'], description: 'Pick ARCH')
    }
    environment {
        REPO = 'https://github.com/dkzippa/prometheus-kbot'
        BRANCH = 'main'
    }
    stages {

        stage('Example') {
            steps {
                echo "Build for platform ${params.OS}"

                echo "Build for arch: ${params.ARCH}"

            }
        }

        stage(" clone") {
            steps {
            echo 'CLONE REPOSITORY'
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }
        
        stage("test") {
            steps {
                echo 'TEST EXECUTION STARTED'
                sh 'make test'
            }
        }
        
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh "make build TARGETOS=${params.OS} GO_ARCH=${params.ARCH}"
            }
        }

        stage("image") {
            steps {
                echo 'IMAGE EXECUTION STARTED'
                sh "make image TARGETOS=${params.OS} GO_ARCH=${params.ARCH}"
            }
        }

        stage("push") {
            steps {
                script {
                    docker.withRegistry('', 'dockerhub'){
                        sh 'make push'
                    }    
                }
                
            }
        }
    }
}
