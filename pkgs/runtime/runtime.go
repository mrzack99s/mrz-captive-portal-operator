package runtime

import (
	"fmt"
	os "os/exec"
	"sync"

	except "github.com/mrzack99s/mrz-captive-portal-operator/pkgs/exceptions"
)

func runProcess(allCommand []*os.Cmd) {

	for _, cmd := range allCommand {
		except.Block{
			Try: func() {
				err := cmd.Run()
				if err != nil {
					except.Throw("cmd.Run() failed with " + cmd.String())
				}
			},
			Catch: func(e except.Exception) {
				fmt.Printf("Caught %v\n", e)
			},
		}.Do()
	}

}

var (
	wg sync.WaitGroup
)

func Run(allCommand []*os.Cmd) {
	go runProcess(allCommand)
}
