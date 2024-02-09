k8s_yaml('kafka.yaml')

k8s_yaml('./producer/deployment.yaml')

docker_build('producer', './producer')

k8s_yaml('./consumer/deployment.yaml')

docker_build('consumer', './consumer')
