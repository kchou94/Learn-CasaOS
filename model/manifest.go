package model

type PathMap struct {
	ContainerPath string `json:"container"`
	Path          string `json:"host"`
	Type          string `json:"type"`
	Desc          string `json:"desc"`
}
