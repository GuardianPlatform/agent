package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"strconv"
	"time"
)

type ImageSummary []struct {

	// containers
	// Required: true
	Containers int64 `json:"containers"`

	// created
	// Required: true
	Created int64 `json:"created"`

	// Id
	// Required: true
	ID string `json:"id"`

	// labels
	// Required: true
	Labels map[string]string `json:"labels"`

	// parent Id
	// Required: true
	ParentID string `json:"parentId"`

	// repo digests
	// Required: true
	RepoDigests []string `json:"repoDigests"`

	// repo tags
	// Required: true
	RepoTags []string `json:"repoTags"`

	// shared size
	// Required: true
	SharedSize int64 `json:"sharedSize"`

	// size
	// Required: true
	Size int64 `json:"size"`

	// virtual size
	// Required: true
	VirtualSize int64 `json:"virtualSize"`
}

func GetImageList(ch chan interface{},cli *client.Client){
	ctx := context.Background()
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	ch <- images
}
func GetContainersList(ch chan interface{},cli *client.Client)  {
	ctx := context.Background()
	container, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	ch <- container
}

/*
*/
func CreateContainer(cli *client.Client) string {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.40"))
	exports := make(nat.PortSet, 10)
	portMap := make(nat.PortMap, 10)
	//tmp := make([]nat.PortBinding, 0, 10)

	for i := 0; i < 4; i++ {
		port, err := nat.NewPort("tcp", strconv.Itoa(80 + i))
		if err != nil {
			panic(err)
		}
		exports[port] = struct{}{}
		portBind := nat.PortBinding{HostPort: strconv.Itoa(80 + i)}
		fmt.Println(portBind)
		tmp := make([]nat.PortBinding, 0)
		tmp = append(tmp, portBind)
		portMap[port] = tmp
	}
	//fmt.Println(tmp)
	//port, err := nat.NewPort("tcp", "80")
	//if err != nil {
	//	panic(err)
	//}
	//exports[port] = struct{}{}
	fmt.Println(exports)
	fmt.Println(portMap)
	//port, err = nat.NewPort("tcp", "81")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(port)
	//exports[port] = struct{}{}
	config := &container.Config{Image: "nginx", ExposedPorts: exports}
	//portBind := nat.PortBinding{HostPort: "80"}
	//tmp := make([]nat.PortBinding, 0, 10)
	//tmp = append(tmp, portBind)
	//portBind = nat.PortBinding{HostPort: "81"}
	//tmp = append(tmp, portBind)
	//
	//portMap[port] = tmp

	hostConfig := &container.HostConfig{PortBindings: portMap}
	cjson,_ := json.Marshal(config)
	fmt.Println(string(cjson));
	pjson,_ := json.Marshal(hostConfig)
	fmt.Println(string(pjson));
	containerName := "hello313"
	body, err := cli.ContainerCreate(context.Background(), config, hostConfig, nil, containerName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %s\n", body.ID)
	return body.ID
}
// 启动
func StartContainer(containerID string, cli *client.Client) {
	err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err == nil {
		fmt.Println("容器", containerID, "启动成功")
	}
}
// 停止
func StopContainer(containerID string, cli *client.Client) {
	timeout := time.Second * 10
	err := cli.ContainerStop(context.Background(), containerID, &timeout)
	if err != nil {
	} else {
		fmt.Printf("容器%s已经被停止\n", containerID)
	}
}