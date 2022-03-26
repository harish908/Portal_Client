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
        stage('Front'){
            agent{
                docker{ image 'node:16-alpine' }
            }
            steps{
                    sh 'node --version'
            }  
        }
    }
}