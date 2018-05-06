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
Download https://github.com/michaelhenkel/contrail-deployer/blob/master/win/contrail-deployer-Windows-x86_64.exe\?raw\=true

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
