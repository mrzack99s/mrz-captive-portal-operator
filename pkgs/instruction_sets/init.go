package instruction_sets

import (
	"os/exec"
	os "os/exec"
	"strings"

	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/configs"
)

func GetInitCommand() []*os.Cmd {

	allCommand := []*os.Cmd{}
	prepareSystem := GetPrepareSystem()

	for _, cmd := range prepareSystem {
		allCommand = append(allCommand, cmd)
	}

	if len(configs.SystemConfig.ZAuth.Network.Routing) > 0 {
		for _, netAddr := range configs.SystemConfig.ZAuth.Network.Routing {
			prepareRouting := GetPrepareRouting(netAddr.NetAddress)
			for _, cmd := range prepareRouting {
				allCommand = append(allCommand, cmd)
			}
		}
	}

	if len(configs.SystemConfig.ZAuth.Network.EAPRouting) > 0 {
		for _, netAddr := range configs.SystemConfig.ZAuth.Network.EAPRouting {
			prepareRouting := GetPrepareEAPRouting(netAddr.NetAddress)
			for _, cmd := range prepareRouting {
				allCommand = append(allCommand, cmd)
			}
		}
	}

	if len(configs.SystemConfig.ZAuth.Network.BypassNetworkRouting) > 0 {
		for _, netAddr := range configs.SystemConfig.ZAuth.Network.BypassNetworkRouting {
			prepareRouting := GetPrepareBypassRouting(netAddr.NetAddress)
			for _, cmd := range prepareRouting {
				allCommand = append(allCommand, cmd)
			}
		}
	}

	return allCommand
}

func GetPrepareSystem() []*os.Cmd {

	allCommand := []*os.Cmd{
		exec.Command("iptables", "-F"),
		exec.Command("iptables", "-t", "nat", "-F"),
		exec.Command("./forwardIp"),
		exec.Command("iptables", "-A", "INPUT", "-i", configs.SystemConfig.ZAuth.Interfaces.WAN,
			"-m", "state", "--state", "ESTABLISHED,RELATED", "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-m", "conntrack",
			"--ctstate", "ESTABLISHED,RELATED", "-j", "ACCEPT"),
		exec.Command("tcdel", configs.SystemConfig.ZAuth.Interfaces.WAN, "--all"),
		exec.Command("tcdel", configs.SystemConfig.ZAuth.Interfaces.LAN, "--all"),
	}

	return allCommand
}

func GetPrepareRouting(netAddress string) []*os.Cmd {

	frontendMGMT := strings.Split(configs.SystemConfig.ZAuth.Network.FrontendMGMT, ":")
	backendMGMT := strings.Split(configs.SystemConfig.ZAuth.Network.BackendMGMT, ":")
	portalAPI := strings.Split(configs.SystemConfig.ZAuth.Network.PortalAPI, ":")
	switchesAPI := strings.Split(configs.SystemConfig.ZAuth.Network.SwitchesAPI, ":")
	switchesUI := strings.Split(configs.SystemConfig.ZAuth.Network.SwitchesUI, ":")
	webLogin := strings.Split(configs.SystemConfig.ZAuth.Network.WebLogin, ":")
	ipOperator := strings.Split(configs.SystemConfig.ZAuth.Network.IpOperator, ":")

	allCommand := []*os.Cmd{
		exec.Command("ip", "route", "add", netAddress, "via", configs.SystemConfig.ZAuth.Network.NextHop,
			"dev", configs.SystemConfig.ZAuth.Interfaces.LAN),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", "53", "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "udp", "--dport", "53", "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", webLogin[1], "-d",
			webLogin[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", frontendMGMT[1], "-d",
			frontendMGMT[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", backendMGMT[1], "-d",
			backendMGMT[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", portalAPI[1], "-d",
			portalAPI[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", switchesAPI[1], "-d",
			switchesAPI[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", ipOperator[1], "-d",
			ipOperator[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", switchesUI[1], "-d",
			switchesUI[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-s", netAddress, "-p", "tcp", "--dport", "80",
			"-j", "DNAT", "--to-destination", webLogin[0]+":"+webLogin[1]),
		exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-s", netAddress, "-p", "tcp", "--dport", "443",
			"-j", "DNAT", "--to-destination", webLogin[0]+":"+webLogin[1]),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-j", "DROP"),
		exec.Command("iptables", "-A", "FORWARD", "-o", configs.SystemConfig.ZAuth.Interfaces.WAN,
			"-i", configs.SystemConfig.ZAuth.Interfaces.LAN, "-s", netAddress, "-m", "conntrack",
			"--ctstate", "NEW", "-j", "ACCEPT"),
	}

	return allCommand
}

func GetPrepareEAPRouting(netAddress string) []*os.Cmd {

	frontendMGMT := strings.Split(configs.SystemConfig.ZAuth.Network.FrontendMGMT, ":")
	backendMGMT := strings.Split(configs.SystemConfig.ZAuth.Network.BackendMGMT, ":")
	portalAPI := strings.Split(configs.SystemConfig.ZAuth.Network.PortalAPI, ":")
	switchesAPI := strings.Split(configs.SystemConfig.ZAuth.Network.SwitchesAPI, ":")
	switchesUI := strings.Split(configs.SystemConfig.ZAuth.Network.SwitchesUI, ":")
	webEAPLogin := strings.Split(configs.SystemConfig.ZAuth.Network.WebEAPLogin, ":")

	allCommand := []*os.Cmd{
		exec.Command("ip", "route", "add", netAddress, "via", configs.SystemConfig.ZAuth.Network.NextHop,
			"dev", configs.SystemConfig.ZAuth.Interfaces.LAN),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", "53", "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "udp", "--dport", "53", "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", webEAPLogin[1], "-d",
			webEAPLogin[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", frontendMGMT[1], "-d",
			frontendMGMT[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", backendMGMT[1], "-d",
			backendMGMT[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", portalAPI[1], "-d",
			portalAPI[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", switchesAPI[1], "-d",
			switchesAPI[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-p", "tcp", "--dport", switchesUI[1], "-d",
			switchesUI[0], "-j", "ACCEPT"),
		exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-s", netAddress, "-p", "tcp", "--dport", "80",
			"-j", "DNAT", "--to-destination", webEAPLogin[0]+":"+webEAPLogin[1]),
		exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-s", netAddress, "-p", "tcp", "--dport", "443",
			"-j", "DNAT", "--to-destination", webEAPLogin[0]+":"+webEAPLogin[1]),
		exec.Command("iptables", "-A", "FORWARD", "-s", netAddress, "-j", "DROP"),
		exec.Command("iptables", "-A", "FORWARD", "-o", configs.SystemConfig.ZAuth.Interfaces.WAN,
			"-i", configs.SystemConfig.ZAuth.Interfaces.LAN, "-s", netAddress, "-m", "conntrack",
			"--ctstate", "NEW", "-j", "ACCEPT"),
	}

	return allCommand
}

func GetPrepareBypassRouting(netAddress string) []*os.Cmd {

	allCommand := []*os.Cmd{
		exec.Command("ip", "route", "add", netAddress, "via", configs.SystemConfig.ZAuth.Network.NextHop,
			"dev", configs.SystemConfig.ZAuth.Interfaces.LAN),
		exec.Command("iptables", "-I", "FORWARD", "1", "-s", netAddress, "-j", "ACCEPT"),
		exec.Command("iptables", "-A", "FORWARD", "-o", configs.SystemConfig.ZAuth.Interfaces.WAN,
			"-i", configs.SystemConfig.ZAuth.Interfaces.LAN, "-s", netAddress, "-m", "conntrack",
			"--ctstate", "NEW", "-j", "ACCEPT"),
	}

	return allCommand
}
