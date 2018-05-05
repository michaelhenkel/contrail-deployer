#!/bin/bash
case $1 in
	mac)
		GOOS=darwin GOARCH=amd64 go build -o contrail-deployer_darwin contrail-deployer.go;;
	linux)
		GOOS=linux GOARCH=amd64 GOARM=7 go build -o contrail-deployer_linux contrail-deployer.go;;
	win)
		GOOS=windows GOARCH=386 go build -o contrail-deployer.exe contrail-deployer.go;;
	*)
		GOOS=linux GOARCH=amd64 GOARM=7 go build -o contrail-deployer contrail-deployer.go;;
esac

