package lvm

type LogicalVolume struct {
	Name        string
	DeviceClass string
	SizeGb      uint64
	Tags        []string
}

type NodeStats struct {
	DeviceClasses []*DeviceClassStats
	Default *DeviceClassStats
}

type VolumeStats struct {
	TotalBytes uint64
	UsedBytes  uint64
}

type DeviceClassStats struct {
	VolumeStats
	DeviceClass string
}

type Client interface {
	GetLVList(deviceClass string) ([]*LogicalVolume, error)
	CreateLV(name, deviceClass string, size uint64, tags []string) (*LogicalVolume, error)
	RemoveLV(name, deviceClass string) error
	ResizeLV(name, deviceClass string, size uint64) error

	GetPath(name, deviceClass string) string

	VolumeStats(name, deviceClass string) (*VolumeStats, error)
	NodeStats() (*NodeStats, error)
}

func New(path string) (Client, error) {
	return &btrfs{poolPath: path}, nil
}
