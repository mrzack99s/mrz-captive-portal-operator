package structs

type SystemConfig struct {
	ZAuth struct {
		API struct {
			Port       string `yaml:"port"`
			Production bool   `yaml:"production"`
			ShareKey   string `yaml:"shareKey"`
		} `yaml:"api"`
		Interfaces struct {
			WAN string `yaml:"WAN"`
			LAN string `yaml:"LAN"`
		} `yaml:"interfaces"`
		Network struct {
			Routing []struct {
				NetAddress string `yaml:"netAddress"`
			} `yaml:"routing"`
			EAPRouting []struct {
				NetAddress string `yaml:"netAddress"`
			} `yaml:"eap_routing"`
			BypassNetworkRouting []struct {
				NetAddress string `yaml:"netAddress"`
			} `yaml:"bypassNetworkRouting"`
			NextHop            string `yaml:"nextHop"`
			BackendMGMT        string `yaml:"backendMGMT"`
			FrontendMGMT       string `yaml:"frontendMGMT"`
			PortalAPI          string `yaml:"portalAPI"`
			SwitchesAPI        string `yaml:"switchesAPI"`
			SwitchesUI         string `yaml:"switchesUI"`
			IpOperator         string `yaml:"ipOperator"`
			WebLogin           string `yaml:"webLogin"`
			WebEAPLogin        string `yaml:"webEAPLogin"`
			WirelessController string `yaml:"wirelessController"`
		} `yaml:"network"`
	} `yaml:"zauth"`
}
