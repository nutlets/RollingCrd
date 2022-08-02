node('slave') {
    container('jnlp-kubectl') {

            stage('Clone Stage') {
                sh 'curl "http://p.nju.edu.cn/portal_io/login" --data "username=201250117&password=hanghang5214.." '
                git branch: 'master', url: "https://github.com/nutlets/RollingCrd.git"
            }
            stage('Install Stage') {
                echo 'Install make'
                sh '''
                yum -y install make
                make -v
               '''
               echo 'install golang'
                sh '''
                yum install -y epel-release
                yum -y install golang
                go version
                echo $GOPATH
                '''
            }

            stage('Image Build Stage') {
                sh 'docker build -f Dockerfile -t rollingupdatecrd:${BUILD_ID} .'
                sh 'docker tag rollingupdatecrd:${BUILD_ID} harbor.edu.cn/nju17/rollingupdatecrd:${BUILD_ID}'
                sh 'docker login harbor.edu.cn -u nju17 -p nju172022'
                sh 'docker push harbor.edu.cn/nju17/rollingupdatecrd:${BUILD_ID}'
            }

            stage('Make Install And Deploy Stage'){


                sh 'go env -w GO111MODULE=auto'
                sh 'go env -w GOPROXY=https://goproxy.cn'
                sh 'chmod +x bin/controller-gen'
                sh 'chmod +x bin/kustomize'
                // sh 'make install'
                sh 'make manifests'
                sh 'bin/kustomize build config/crd | kubectl apply -f -'

                sh 'make manifests'
                sh 'cd config/manager && ls && ./../../bin/kustomize edit set image controller=harbor.nju.edu.cn/nju17/rollingupdatecrd:${BUILD_ID}'
                sh 'ls'
                sh 'bin/kustomize build config/default | kubectl apply -f -'

                sh 'kubectl apply -f config/samples/demo_v1_rollingupdatecrd.yaml -n nju17'
            }


    }
}