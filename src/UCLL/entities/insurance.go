package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Owner struct {
	Name    string
	Surname string
	OwnerID string
}

type Owners []Owner

type Accident struct {
	Description string
	Date        string
}

type Accidents []Accident

type CarInsurance struct {
	CarID        string
	OwnerList    Owners
	AccidentList Accidents
}

//Add Owner to a car
func (o *Owner) AddOwner(stub shim.ChaincodeStubInterface, args []string) error {
	carID := args[0]
	o.Name = args[1]
	o.Surname = args[2]
	o.OwnerID = args[3]

	var ci CarInsurance
	ciJson, _ := stub.GetState("ci-" + carID)
	if ciJson == nil {
		ci.CarID = carID
	} else {
		err := json.Unmarshal(ciJson, &ci)
		if err != nil {
			return errors.New(": Error in unmarshaling JSON")
		}
	}
	ci.OwnerList = append(ci.OwnerList, *o)
	ciJsonIndent, _ := json.MarshalIndent(ci, "", "  ")
	err := stub.PutState("ci-"+carID, ciJsonIndent)
	if err != nil {
		return errors.New("AddOwner: Unable to PutState")
	}
	return nil
}

//Add Accident to a car
func (a *Accident) AddAccident(stub shim.ChaincodeStubInterface, args []string) error {
	carID := args[0]
	a.Description = args[1]
	a.Date = args[2]

	var ci CarInsurance
	ciJson, _ := stub.GetState("ci-" + carID)
	if ciJson == nil {
		ci.CarID = carID
	} else {
		err := json.Unmarshal(ciJson, &ci)
		if err != nil {
			return errors.New(": Error in unmarshaling JSON")
		}
	}
	ci.AccidentList = append(ci.AccidentList, *a)
	ciJsonIndent, _ := json.MarshalIndent(ci, "", "  ")
	err := stub.PutState("ci-"+carID, ciJsonIndent)
	if err != nil {
		return errors.New("AddAccident: Unable to PutState")
	}
	return nil
}

//Getting insurance list for a car
func (ci *CarInsurance) GetCarInsuranceList(stub shim.ChaincodeStubInterface, carID string) ([]byte, error) {
	ciJsonIndent, err := stub.GetState("ci-" + carID)
	if err != nil {
		return nil, err
	}
	if ciJsonIndent == nil {
		ciJsonIndent = []byte("{\"Error\":\"No insurance track available for carID " + carID + "\"}")
	}
	fmt.Println("GetCarJSON returned:", string(ciJsonIndent))
	return ciJsonIndent, nil
}
