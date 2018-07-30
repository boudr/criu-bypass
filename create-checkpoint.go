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
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

//Arguments:
//	1) Container ID to checkpoint
//	2) Name of the checkpoint
func main() {
	containerID := os.Args[1]
	checkpointName := os.Args[2]

	client, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
	}

	//CreateCheckpoint(ctx context.Context, containerID, checkpointDir string, exit bool)
}
