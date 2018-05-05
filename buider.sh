#!/bin/bash
case $1 in
	mac)
		GOOS=darwin GOARCH=amd64 go build -o mac/contrail-deployer-Darwin-x86_64 contrail-deployer.go;;
	linux)
		GOOS=linux GOARCH=amd64 GOARM=7 go build -o linux/contrail-deployer-Linux-x86_64 contrail-deployer.go;;
	win)
		GOOS=windows GOARCH=amd64 go build -o win/contrail-deployer-Windows-x86_64.exe contrail-deployer.go;;
	*)
		GOOS=darwin GOARCH=amd64 go build -o mac/contrail-deployer-Darwin-x86_64 contrail-deployer.go
		GOOS=linux GOARCH=amd64 GOARM=7 go build -o linux/contrail-deployer-Linux-x86_64 contrail-deployer.go
		GOOS=windows GOARCH=amd64 go build -o win/contrail-deployer-Windows-x86_64.exe contrail-deployer.go;;
esac

