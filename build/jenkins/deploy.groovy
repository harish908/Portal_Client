pipeline{
    agent{
        kubernetes{
            yamlFile 'build/jenkins/pod.yaml'
        }
    }
    parameters{
        string(name: 'DOMAIN', defalultValue: 'portal', description: 'Portal Domain')
        string(name: 'SERVICE', defalultValue: 'portal-client', description: 'Portal Client Service')
        choice(name: "ECR_ACCOUNT_ID", choices: "859114173848", description: "ECR Account Id")
        choice(name: "ECR_REGION", choices: "ap-south-1", description: "ECR Region")
        gitParameter name: 'DEPLOY_BRANCH', description: 'Select a branch to deploy', type: 'PT_BRANCH', defalultValue: 'origin/master', listSize: '0', selectedValue: 'DEFAULT'
    }
    environment{
        DEPLOYMENT_ROLE         = "JenkinsDeploymentRole"
        DEPLOY_TARGET_ACCOUNT   = '859114173848'
    }
    stages{
        stage('checkout'){
            steps{
                script{
                    echo "DEPLOY BRANCH: ${params.DEPLOY_BRANCH}"
                    sh(script: "git checkout ${params.DEPLOY_BRANCH}")
                }
            }
        }

        stage('env'){
            steps{
                script{
                    env.REVISION = sh(script: "git rev-parse -short HEAD", returnStdout: true).trim()
                    env.ECR_PORTAL_IMAGE = "${ECR_ACCOUNT_ID}.dkr.ecr.${ECR_REGION}.amazonaws.com/${DOMAIN}/${SERVICE}:${REVISION}.${BUILD_NUMBER}"
                }
            }
        }

        stage('image'){
            steps{
                container('kaniko'){
                    withAWS(role: "${DEPLOYMENT_ROLE}", roleAccount: "${ECR_ACCOUNT_ID}", region: "${ECR_REGION}")
                    sh "/kaniko/executor -f Dockerfile -c `pwd` --skip-tls-verify --cache=false --destination=${ECR_PORTAL_IMAGE}"
                }
            }
        }
    }
}