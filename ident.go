package gilmour

import (
	"fmt"
	"os"

	"github.com/pborman/uuid"
)

func getHostname() string {
	host, err := os.Hostname()
	if err != nil {
		host = "Host-Unknown"
	}
	return host
}

func getPid() int {
	return os.Getpid()
}

func getUUID() string {
	return uuid.New()
}

func makeIdent() string {
	hostname := getHostname()
	pid := getPid()
	uuid := getUUID()

	return fmt.Sprintf("%v-pid-%v-uuid-%v", hostname, pid, uuid)
}

func makeSenderId() string {
	return getUUID()
}
