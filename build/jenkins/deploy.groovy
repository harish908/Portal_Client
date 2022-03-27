pipeline{
    agent none
    parameters{
        string(name: "DOMAIN", defaultValue: "portal", description: "Portal Domain")
        string(name: "SERVICE", defaultValue: "portal-client", description: "Portal Client Service")
        choice(name: "ECR_ACCOUNT_ID", choices: "859114173848", description: "ECR Account Id")
        choice(name: "ECR_REGION", choices: "ap-south-1", description: "ECR Region")
        gitParameter name: 'DEPLOY_BRANCH', description: 'Select a branch to deploy', type: 'PT_BRANCH', defaultValue: 'origin/master', listSize: '0', selectedValue: 'DEFAULT'
    }
    environment{
        DEPLOYMENT_ROLE         = "JenkinsDeploymentRole"
        DEPLOY_TARGET_ACCOUNT   = '859114173848'
    }
    stages{
        stage('checkout'){
            agent{
                docker{ 
                    image 'alpine/git:latest' 
                    args '--entrypoint='                                // keep container alive, smiliar to cat command 
                }
            }
            steps{
                script{
                    echo "DEPLOY BRANCH: ${params.DEPLOY_BRANCH}"
                    sh(script: "git checkout ${params.DEPLOY_BRANCH}")
                    env.REVISION = sh(script: "git rev-parse -short HEAD", returnStdout: true).trim()
                    env.ECR_PORTAL_IMAGE = "${ECR_ACCOUNT_ID}.dkr.ecr.${ECR_REGION}.amazonaws.com/${DOMAIN}:latest"
                }
            }
        }

        stage('image'){
            agent{
                docker{ 
                    image 'gcr.io/kaniko-project/executor:debug'        // use debug version to keep container alive
                    args '--user 0 --entrypoint='                       // run as root user
                }
            }
            steps{
                // withAWS(role: "${DEPLOYMENT_ROLE}", roleAccount: "${ECR_ACCOUNT_ID}", region: "${ECR_REGION}"){
                //     withCredentials([aws(credentialsId: "$GITHUB_CREDENTIALS", variable: "GITHUB_TOKEN")]) {
                //         sh "/kaniko/executor -f Dockerfile -c `pwd` --skip-tls-verify --cache=true --destination=${ECR_PORTAL_IMAGE}"
                //     }
                // }
                withCredentials([aws(accessKeyVariable:'AWS_ACCESS_KEY_ID', credentialsId:'harish-aws-creds', secretKeyVariable:'AWS_SECRET_ACCESS_KEY')]){
                    sh "/kaniko/executor -f Dockerfile -c `pwd` --skip-tls-verify --cache=true --destination=${ECR_PORTAL_IMAGE}"
                }
            }
        }
    }
}