package instruction_sets

import (
	"os/exec"
	os "os/exec"
	"strconv"

	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/configs"
)

func GetAppendBandwidthControl(dlSpeed uint32, upSpeed uint32, ipAddress string) []*os.Cmd {

	dlSpeedStr := strconv.FormatUint(uint64(dlSpeed), 10) + "bps"
	upSpeedStr := strconv.FormatUint(uint64(upSpeed), 10) + "bps"
	allCommand := []*os.Cmd{
		exec.Command("tcset", configs.SystemConfig.ZAuth.Interfaces.WAN, "--overwrite", "--delay", "1ms", "--loss",
			"0.01%", "--rate", dlSpeedStr, "--dst-network", ipAddress, "--direction", "incoming"),
		exec.Command("tcset", configs.SystemConfig.ZAuth.Interfaces.LAN, "--overwrite", "--delay", "1ms", "--loss",
			"0.01%", "--rate", upSpeedStr, "--src-network", ipAddress, "--direction", "incoming"),
	}

	return allCommand
}

func GetDeleteBandwidthControl(ipAddress string) []*os.Cmd {

	allCommand := []*os.Cmd{
		exec.Command("tcdel", configs.SystemConfig.ZAuth.Interfaces.WAN, "--dst-network", ipAddress, "--direction", "incoming"),
		exec.Command("tcdel", configs.SystemConfig.ZAuth.Interfaces.LAN, "--src-network", ipAddress, "--direction", "incoming"),
	}

	return allCommand
}
