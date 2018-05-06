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
        Path to instance.yaml (default "instance.yaml")
  -o string
        openstack|kubernetes|none (default "none")
  -privk string
        Path to private ssh key
  -pubk string
        Path to public ssh key
```

# Example

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
