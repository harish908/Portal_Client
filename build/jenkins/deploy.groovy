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
        // stage('checkout'){
        //     agent{
        //         docker { image 'alpine/git:latest' }
        //     }
        //     steps{
        //         script{
        //             echo "DEPLOY BRANCH: ${params.DEPLOY_BRANCH}"
        //             sh(script: "git checkout ${params.DEPLOY_BRANCH}")
        //         }
        //     }
        // }

        // stage('env'){
        //     agent{
        //         docker { image 'alpine/git:latest' }
        //     }
        //     steps{
        //         script{
        //             env.REVISION = sh(script: "git rev-parse -short HEAD", returnStdout: true).trim()
        //             env.ECR_PORTAL_IMAGE = "${ECR_ACCOUNT_ID}.dkr.ecr.${ECR_REGION}.amazonaws.com/${DOMAIN}/${SERVICE}:${REVISION}.${BUILD_NUMBER}"
        //         }
        //     }
        // }

        stage('image'){
            // agent{
            //     docker{ image 'gcr.io/kaniko-project/executor:debug' }
            // }
            // steps{
            //     script{
            //         env.ECR_PORTAL_IMAGE = "${ECR_ACCOUNT_ID}.dkr.ecr.${ECR_REGION}.amazonaws.com/${DOMAIN}/${SERVICE}:123456.${BUILD_NUMBER}"
            //     }
            //     withAWS(role: "${DEPLOYMENT_ROLE}", roleAccount: "${ECR_ACCOUNT_ID}", region: "${ECR_REGION}"){
            //         sh "/kaniko/executor -f Dockerfile -c `pwd` --skip-tls-verify --cache=true --destination=${ECR_PORTAL_IMAGE}"
            //     }
            // }  
            agent{
                docker{ image 'node:16-alpine' }
            }
            steps{
                sh 'node --version'
            }    
        }
    }
}