//Warning:

//This is for a tech demo and not intended to be used in any situation!
//Because of this, there will be no actual CLI functionally. Any input to this applcation assumes it is correct.
//Error handling is just reporting.
//You hath been warned.

//--

//Checkpoint the container:
//https://github.com/moby/moby/blob/master/libcontainerd/client_daemon.go#L558

//Restore the container:
//https://github.com/moby/moby/blob/master/libcontainerd/client_daemon.go#L247
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//Arguments:
//	1) Container ID to checkpoint
//	2) Name of the checkpoint
func main() {

	client, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
	}

	err = client.ContainerStart(context.Background(), "f7787ad1ffe7", types.ContainerStartOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(20 * time.Second)

	//Create the checkpoint
	err = client.CheckpointCreate(context.Background(), "f7787ad1ffe7", types.CheckpointCreateOptions{
		CheckpointID: "cp1",
		Exit:         true,
	})
	if err != nil {
		fmt.Println(err)
		erra := client.ContainerStop(context.Background(), "f7787ad1ffe7", nil)
		if erra != nil {
			fmt.Println(erra)
			return
		}
		return
	}
}
