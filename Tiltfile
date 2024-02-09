k8s_yaml('kafka.yaml')

k8s_yaml('./producer/deployment.yaml')

docker_build('producer', './producer')

k8s_yaml('./consumer/deployment.yaml')

docker_build('consumer', './consumer')

local_resource(name='local-registry', cmd='docker stop local-registry;docker rm local-registry;docker run -d -p 6000:5000 --restart=always --name local-registry registry:2.7.0')
