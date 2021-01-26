package instruction_sets

import (
	"os/exec"
	os "os/exec"
)

func GetAllowNet(ipAddress string) []*os.Cmd {
	//iptables -t nat -I PREROUTING 1 -s ${SPEC_IP} -p tcp -m tcp --dport 80 -j ACCEPT
	allCommand := []*os.Cmd{
		exec.Command("iptables", "-t", "nat", "-I", "PREROUTING", "1", "-s", ipAddress, "-p", "tcp",
			"-m", "tcp", "--dport", "80", "-j", "ACCEPT"),
		exec.Command("iptables", "-t", "nat", "-I", "PREROUTING", "1", "-s", ipAddress, "-p", "tcp",
			"-m", "tcp", "--dport", "443", "-j", "ACCEPT"),
		exec.Command("iptables", "-I", "FORWARD", "1", "-s", ipAddress, "-j", "ACCEPT"),
	}

	return allCommand
}

func GetUnAllowNet(ipAddress string) []*os.Cmd {
	//iptables -t nat -I PREROUTING 1 -s ${SPEC_IP} -p tcp -m tcp --dport 80 -j ACCEPT
	allCommand := []*os.Cmd{
		exec.Command("iptables", "-t", "nat", "-D", "PREROUTING", "-s", ipAddress, "-p", "tcp",
			"-m", "tcp", "--dport", "80", "-j", "ACCEPT"),
		exec.Command("iptables", "-t", "nat", "-D", "PREROUTING", "-s", ipAddress, "-p", "tcp",
			"-m", "tcp", "--dport", "443", "-j", "ACCEPT"),
		exec.Command("iptables", "-D", "FORWARD", "-s", ipAddress, "-j", "ACCEPT"),
	}

	return allCommand
}
