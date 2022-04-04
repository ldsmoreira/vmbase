package main

import (
	"modalchemy-virt-plataform/internal/virtcontroller"

	"github.com/google/uuid"
)

func main() {
	testvm := virtcontroller.Vm{
		XmlTemplatePath: "/home/lucas/Desktop/modalch-virt-plataform/resources/templates/imagesconfig/ubuntu-lcp.xml",
		Name:            "Test",
		UUID:            uuid.NewString(),
		BlockFile:       "/home/lucas/Desktop/modalch-virt-plataform/ubuntu-lcp.qcow2",
	}

	testvm.CreateVm()
}
