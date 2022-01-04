package model

type PortMap struct {
	ContainerPort string `json:"container"`
	CommendPort   string `json:"host"`
	Protocol      string `json:"protocol"`
	Desc          string `json:"desc"`
	Type          int    `json:"type"`
}

type PortArray []PortMap

type Env struct {
	Name  string `json:"container"`
	Value string `json:"host"`
	Desc  string `json:"desc"`
	Type  int    `json:"type"`
}

type EnvArray []Env

type PathMap struct {
	ContainerPath string `json:"container"`
	Path          string `json:"host"`
	Type          string `json:"type"`
	Desc          string `json:"desc"`
}

type PathArray []PathMap

type CustomizationPostData struct {
	Origin       string    `json:"origin"`
	NetworkModel string    `json:"network_model"`
	Index        string    `json:"index"`
	Icon         string    `json:"icon"`
	Image        string    `json:"image"`
	Envs         EnvArray  `json:"envs"`
	Ports        PortArray `json:"ports"`
	Volumes      PathArray `json:"volumes"`
	Devices      PathArray `json:"devices"`
	// Port         string    `json:"port,omitempty"`
	PortMap     string `json:"port_map"`
	CpuShares   int64  `json:"cpu_shares"`
	Memory      int64  `json:"memory"`
	Restart     string `json:"restart"`
	EnableUPNP  bool   `json:"enable_upnp"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Position    string `json:"position"`
}
