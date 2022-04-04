package configVM

type VirtualMachineModel struct {
	XmlTemplatePath string
	Name            string `json:"name"`
	UUID            string
	BlockFile       string
	BlockDeviceSize int
}
