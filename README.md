# Contrail Deployer

The Contrail Deployer tool installs contrail on a set of defined    
instances.    

# Requirements    

- docker client    
- internet access    

# Installation

## OSX

```
curl -L https://github.com/michaelhenkel/contrail-deployer/blob/master/mac/contrail-deployer-Darwin-x86_64\?raw\=true -o /usr/local/bin/contrail-deployer
chmod +x /usr/local/bin/contrail-deployer
```

## Linux
```
curl -L https://github.com/michaelhenkel/contrail-deployer/blob/master/linux/contrail-deployer-Linux-x86_64\?raw\=true -o /usr/local/bin/contrail-deployer
chmod +x /usr/local/bin/contrail-deployer
```

## Windows
Download https://github.com/michaelhenkel/contrail-deployer/raw/master/win/contrail-deployer-Windows-x86_64.exe

# Usage
```
contrail-deployer [OPTIONS] [ACTION]

ACTIONS:
  provision|1
       provisions instances
  configure|2
       configures instances
  install|3
       installs instances
  all
       provisions, configures and installs instances
  12
       provisions and configures instances
  23
       configures and installs instances

OPTIONS:
  -di string
        Contrail Deployer Docker image name (default "michaelhenkel/contrail-deployer")
  -i string
        Absolute Path to instance.yaml (default "instance.yaml")
  -o string
        openstack|kubernetes|none (default "none")
  -privk string
        Absolute path to private ssh key
  -pubk string
        Absolute path to public ssh key
```

# Examples

## KVM k8s
```
cat << EOF > instances.yaml
provider_config:
  kvm:
    image: CentOS-7-x86_64-GenericCloud-1802.qcow2.xz
    image_url: http://10.87.64.32/
    ssh_pwd: c0ntrail123
    ssh_user: root
    vcpu: 8
    vram: 24000
    vdisk: 100G
    subnet_prefix: 192.168.1.0
    subnet_netmask: 255.255.255.0
    gateway: 192.168.1.1
    nameserver: 10.84.5.100
    ntpserver: 192.168.1.1
    domainsuffix: local
instances:
  kvm1:
    provider: kvm
    host: 10.87.64.31
    bridge: br1
    ip: 192.168.1.100
    roles:
      config_database:
      config:
      control:
      analytics_database:
      analytics:
      webui:
      k8s_master:
      kubemanager:
  kvm2:
    provider: kvm
    host: 10.87.64.32
    bridge: br1
    ip: 192.168.1.101
    roles:
      config_database:
      config:
      control:
      analytics_database:
      analytics:
      webui:
      kubemanager:
  kvm3:
    provider: kvm
    host: 10.87.64.33
    bridge: br1
    ip: 192.168.1.102
    roles:
      config_database:
      config:
      control:
      analytics_database:
      analytics:
      webui:
      kubemanager:
  kvm4:
    provider: kvm
    host: 10.87.64.33
    bridge: br1
    ip: 192.168.1.104
    UPGRADE_KERNEL: true
    roles:
      vrouter:
      k8s_node:
  kvm5:
    provider: kvm
    host: 10.87.64.32
    bridge: br1
    ip: 192.168.1.105
    UPGRADE_KERNEL: true
    roles:
      vrouter:
      k8s_node:
global_configuration:
  CONTAINER_REGISTRY: opencontrailnightly
contrail_configuration:
  CONTRAIL_VERSION: latest
EOF

contrail-deployer -i `pwd`/instances.yaml -o kubernetes all
```

## AWS k8s
```
# You have to set the following vars!!!
EC2_ACCESS_KEY=YOUR_EC2_ACCESS_KEY
EC2_SECRET_KEY=YOUR_EC2_SECRET_KEY
AMI_ID=REGION_AMI_ID_FOR_CENTOS_7
REGION=YOUR_REGION
VPC_SUBNET_ID=YOUR_VPC_SUBNET_ID

cat << EOF > instances.yaml
provider_config:
  aws:
    ec2_access_key: ${EC2_ACCESS_KEY}
    ec2_secret_key: ${EC2_SECRET_KEY}
    ssh_public_key: /id_rsa.pub
    ssh_private_key: /id_rsa
    ssh_user: centos
    instance_type: t2.xlarge
    image: ${AMI_ID}
    region: ${REGION}
    vpc_subnet_id: ${VPC_SUBNET_ID}
    assign_public_ip: yes
    volume_size: 50
    key_pair: ansible-deployer
    ntpserver: 169.254.169.123
instances:
  aws1:
    provider: aws
    roles:
      config_database:
      config:
      control:
      analytics_database:
      analytics:
      webui:
      k8s_master:
      kubemanager:
  aws2:
    UPGRADE_KERNEL: true
    provider: aws
    roles:
      vrouter:
      k8s_node:
  aws3:
    UPGRADE_KERNEL: true
    provider: aws
    roles:
      vrouter:
      k8s_node:
contrail_configuration:
  CONTRAIL_VERSION: latest
global_configuration:
  CONTAINER_REGISTRY: opencontrailnightly
EOF

contrail-deployer -i `pwd`/instances.yaml -privk /root/.ssh/id_rsa -pubk /root/.ssh/id_rsa.pub -o kubernetes all
```

# Known issues

- currently the install action cannot be run separately from the configure action.    
  Whenever an action is ran separately, a new container is being created. The configure    
  action pulls in the kolla ansible playbooks, which are required for the install action.   
  For now, the action must be set to 23 (configure and install).    
