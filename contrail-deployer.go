package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func IsValidAction(actionList []string) bool {
	actionFound := false
	for _, action := range actionList {
		actionFound = false
		switch action {
		case
			"provision",
			"configure",
			"install",
			"all",
			"1",
			"2",
			"12",
			"23":
			actionFound = true
		}
	}
	return actionFound
}

func main() {
	instanceFile := flag.String("i", "instance.yaml", "Absolute path to instance.yaml")
	orchestrator := flag.String("o", "none", "openstack|kubernetes|none")
	privateKey := flag.String("privk", "", "Absolute path to private ssh key")
	publicKey := flag.String("pubk", "", "Absolute path to public ssh key")
	deployerImage := flag.String("di", "michaelhenkel/contrail-deployer", "Contrail Deployer Docker image name")
	cherryPick := flag.String("cp", "", "cherry pick id for ansible deployer (e.g. 90/42790/1)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "./contrail-deployer [OPTIONS] [ACTION]\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "ACTIONS:\n")
		fmt.Fprintf(os.Stderr, "  provision|1\n")
		fmt.Fprintf(os.Stderr, "       provisions instances\n")
		fmt.Fprintf(os.Stderr, "  configure|2\n")
		fmt.Fprintf(os.Stderr, "       configures instances\n")
		fmt.Fprintf(os.Stderr, "  install|3\n")
		fmt.Fprintf(os.Stderr, "       installs instances\n")
		fmt.Fprintf(os.Stderr, "  all\n")
		fmt.Fprintf(os.Stderr, "       provisions, configures and installs instances\n")
		fmt.Fprintf(os.Stderr, "  12\n")
		fmt.Fprintf(os.Stderr, "       provisions and configures instances\n")
		fmt.Fprintf(os.Stderr, "  23\n")
		fmt.Fprintf(os.Stderr, "       configures and installs instances\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "OPTIONS:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if !IsValidAction(flag.Args()) {
		fmt.Println("ERROR: action is missing/wrong")
		flag.Usage()
		return
	}
	var envList []string
	for _, action := range flag.Args() {
		envList = append(envList, "action="+action)
	}
	envList = append(envList, "orchestrator="+*orchestrator)
	if *cherryPick != "" {
		envList = append(envList, "CP="+*cherryPick)
	}
	var mountList []mount.Mount
	if *instanceFile != "" {
		if _, err := os.Stat(*instanceFile); os.IsNotExist(err) {
			fmt.Printf("ERROR: instance file %s is missing/wrong\n", *instanceFile)
			flag.Usage()
			return
		}
		instanceMount := mount.Mount{Type: mount.TypeBind, Source: *instanceFile, Target: "/instances.yaml"}
		mountList = append(mountList, instanceMount)
	} else {
		fmt.Println("ERROR: instance file is not specified")
		flag.Usage()
		return
	}
	if *privateKey != "" {
		privKeyMount := mount.Mount{Type: mount.TypeBind, Source: *privateKey, Target: "/id_rsa"}
		mountList = append(mountList, privKeyMount)
	}
	if *publicKey != "" {
		pubKeyMount := mount.Mount{Type: mount.TypeBind, Source: *publicKey, Target: "/id_rsa.pub"}
		mountList = append(mountList, pubKeyMount)
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if r, err := cli.ImagePull(ctx, *deployerImage, types.ImagePullOptions{}); err != nil {
		panic(err)
	} else {
		io.Copy(os.Stdout, r)
	}

	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: *deployerImage,
			//Cmd:   []string{"cat", "/instances.yaml"},
			Tty: true,
			Env: envList,
		},
		&container.HostConfig{
			Mounts: mountList,
		},
		nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		Follow:     true,
	})
	if err != nil {
		panic(err)
	} else {
		io.Copy(os.Stdout, out)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	if err := cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{}); err != nil {
		panic(err)
	}
}
