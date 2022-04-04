package virtcontroller

import (
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"os"
	"text/template"
	"time"

	"github.com/google/uuid"
	"libvirt.org/libvirt-go"
)

type Vm struct {
	XmlTemplatePath string
	Name            string `json:"name"`
	UUID            string
	BlockFile       string
	BlockDeviceSize int
	MacAddress      string
}

func (vm Vm) ConfigGen() (string, error) {

	var filename string = uuid.NewString()

	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	tmpl, err := template.ParseFiles(vm.XmlTemplatePath)
	if err != nil {
		log.Fatal("Parse: ", err)
		return "", err
	}

	fmt.Println(tmpl.Name())

	err = tmpl.ExecuteTemplate(file, tmpl.Name(), vm)

	if err != nil {
		log.Fatal("Execute: ", err)
		return "", err
	} else {
		fileContent, _ := os.ReadFile(filename)
		fileText := string(fileContent)
		os.Remove(filename)
		return fileText, nil
	}
}

func (vm Vm) CreateVm() (string, error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	vm.MacAddress = vm.GenerateMacAddress()
	fmt.Println(vm.MacAddress)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer conn.Close()

	xml, err := vm.ConfigGen()
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	dom, err := conn.DomainCreateXML(xml, 0)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	domInterfaces := new(libvirt.DomainInterfaceAddressesSource)

	name, err := dom.GetName()

	fmt.Printf("Inicializando maquina: %s\n", name)

	var interfaces []libvirt.DomainInterface
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		interfaces, err = dom.ListAllInterfaceAddresses(*domInterfaces)
		fmt.Println(interfaces)
	}

	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	sshAddress := interfaces[0].Addrs[0].Addr

	fmt.Printf("Maquina virtual habilitada para acesso em %s com usuario 'ubuntu-lcp'", sshAddress)

	return sshAddress, nil
}

func (vm Vm) GenerateMacAddress() string {
	buf := make([]byte, 6)
	var mac net.HardwareAddr

	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	buf[0] |= 2

	mac = append(mac, buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])

	return mac.String()
}
