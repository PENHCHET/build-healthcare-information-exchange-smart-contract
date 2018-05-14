package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("main")

type SmartContract struct {
}

var bcFunctions = map[string]func(shim.ChaincodeStubInterface, []string) peer.Response{
	// Insurance Peer Smart Contract Functions
	"initInsurancePersonCheckHistoryLedger":                        initInsurancePersonCheckHistoryLedger,
	"createNewPersonCheckHistory":                                  createNewPersonCheckHistory,
	"getAllPersonCheckHistory":                                     getAllPersonCheckHistory,
	"getAllPersonCheckHistoryByPersonId":                           getAllPersonCheckHistoryByPersonId,
	"getAllPersonCheckHistoryByPersonIdAndCheckHospitalPermission": getAllPersonCheckHistoryByPersonIdAndCheckHospitalPermission,

	"createNewPersonRequestPermission": createNewPersonRequestPermission,
	"getPersonRequestPermission":       getPersonRequestPermission,
	"updatePersonRequestPermission":    updatePersonRequestPermission,

	// Hospital Peer Smart Contract Functions
	"initHospitalPatientMedicalHistory":                   initHospitalPatientMedicalHistory,
	"getAllPatientMedicalHistory":                         getAllPatientMedicalHistory,
	"addNewPatientMedicalHistory":                         addNewPatientMedicalHistory,
	"createNewPatientMedicalHistoryFromLocalDBByPersonId": createNewPatientMedicalHistoryFromLocalDBByPersonId,

	"getAllPatientByHospitalId": getAllPatientByHospitalID,

	"createNewPatientRequestPermission":                      createNewPatientRequestPermission,
	"getPatientRequestPermissionByHospitalIdAndPatientId":    getPatientRequestPermissionByHospitalIdAndPatientId,
	"updatePatientRequestPermissionByHospitalIdAndPatientId": updatePatientRequestPermissionByHospitalIdAndPatientId,
}

func (s *SmartContract) Init(APIStub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("========================================= Init Smart Contract =====================================")
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIStub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("========================================= Invoke Smart Contract ======================================")

	function, args := APIStub.GetFunctionAndParameters()

	// TODO: Get the create
	creator, err := APIStub.GetCreator()
	if err != nil {
		return shim.Error("Cannot get the Creator")
	}

	fmt.Println("Creator = ", creator)

	if function == "init" {
		return s.Init(APIStub)
	}

	bcFunc := bcFunctions[function]
	if bcFunc == nil {
		return shim.Error("Invalid invoke function.")
	}
	return bcFunc(APIStub, args)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
