package containerruntime

// ContainerMount holds container volume mounts
type ContainerMount struct {
	MountType string
	Source    string
	Target    string
	Mode      MountMode
}

// EnvironmentProperty holds environment variables
type EnvironmentProperty struct {
	Name  string
	Value string
}

// ContainerPort holds container ports
type ContainerPort struct {
	Source int
	Target int
}

type MountMode string

const (
	WriteMode MountMode = "write"
	ReadMode  MountMode = "read"
)
