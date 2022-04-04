package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"modalchemy-virt-plataform/helpers"
	"modalchemy-virt-plataform/internal/virtcontroller"
	"net/http"

	"github.com/google/uuid"
)

func HandleAPIResquest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET method requested"}`))
	case "POST":
		var body []byte
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(body)
		if err != nil {
			fmt.Println(err)
		}
		var vm virtcontroller.Vm
		w.WriteHeader(http.StatusCreated)
		err = json.Unmarshal(body, &vm)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(vm.GenerateMacAddress())

		newvolumefile := "/home/lucas/Desktop/modalch-virt-plataform/" + uuid.NewString() + ".qcow2"
		fmt.Println(newvolumefile)

		helpers.CopyFile("/home/lucas/Desktop/modalch-virt-plataform/ubuntu-lcp.qcow2", newvolumefile)

		vm.XmlTemplatePath = "/home/lucas/Desktop/modalch-virt-plataform/resources/templates/imagesconfig/ubuntu-lcp.xml"
		vm.UUID = uuid.NewString()
		vm.BlockFile = newvolumefile
		sshAddress, err := vm.CreateVm()
		if err != nil {
			fmt.Println(err)
		}
		resp := &CreateResponse{
			Name:       vm.Name,
			VMID:       vm.UUID,
			Status:     "Created",
			SSHAddress: sshAddress,
		}
		// fmt.Println(st)
		jsonresp, err := json.Marshal(resp)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(jsonresp)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}
