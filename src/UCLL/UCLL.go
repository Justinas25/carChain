package main

import (
	"UCLL/entities"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//Chaincode is a blank struct to use with shim
type Chaincode struct {
}

//Main function to start chan code execution
func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Println("Error starting Chaincode: %s", err)
	}
}

//Init function is executed when chain code is deployed
func (t *Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

//Invoke is executed when data is stored and manipulated
func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case "createCar":
		var c entities.Car
		err := c.CreateCar(stub, args)
		if err != nil {
			return nil, err
		}
	case "loadSampleCars":
		var cs entities.Cars
		message := cs.LoadSample(stub)
		return []byte(message), nil
	case "addMaintenance":
		var m entities.Maintenance
		err := m.AddMaintenance(stub, args)
		if err != nil {
			return nil, err
		}
	case "loadMaintenanceSample":
		var cm entities.CarMaintenance
		message := cm.LoadMaintenanceSample(stub)
		return []byte(message), nil
	default:
		return nil, errors.New("Invoke: Received unknonw function name")
	}
	return nil, nil
}

//Query returns a result from the database
func (t *Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	switch function {
	case "getCar":
		var c entities.Car
		cJsonIndent, err := c.GetCar(stub, args[0])
		return cJsonIndent, err
	case "getCarMaintenceList":
		var cm entities.CarMaintenance
		cmJsonIndent, err := cm.GetCarMaintenceList(stub, args[0])
		return cmJsonIndent, err
	case "listCars":
		var cs entities.Cars
		csJsonIndent, err := cs.ListCars(stub)
		return csJsonIndent, err
	default:
		return nil, errors.New("Invoke: Received unknonw function name")
	}
	return nil, nil
}
