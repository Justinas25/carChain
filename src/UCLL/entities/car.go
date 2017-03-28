package entities

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"strings"
	"time"
)

type Part struct {
	Name          string
	Description   string
	LocationOnCar string
	DateAdded     time.Time
	DateInspected time.Time
	Brand         string
}

type Specific struct {
	Name        string
	Description string
	Value       string
	Timestamp   time.Time
}
type SpecificsArray []Specific

type PartsArray []Part

//Car entity
type Car struct {
	CarID     string
	Brand     string
	Type      string
	Engine    string
	Year      int
	Parts     PartsArray
	Specifics SpecificsArray
}

type Cars []Car

func (s SpecificsArray) sampleSpecifics() {

	s = make([]Specific, 0, 5)

	s[0].Name = "Color"
	s[0].Description = "First Specific in the list."
	s[0].Value = "Red"
	s[0].Timestamp = time.Now()

	s[1].Name = "Power"
	s[1].Description = "Horse power!"
	s[1].Value = "144 HP"
	s[1].Timestamp = time.Now()

}

func (p PartsArray) sampleParts() {
	p = make([]Part, 0, 5)
	p[0].Brand = "Mercedes"
	p[0].DateAdded = time.Date(2010, time.January, 1, 0, 0, 0, 0, nil)
	p[0].DateInspected = time.Now()
	p[0].Description = "Left mirror"
	p[0].LocationOnCar = "Outside, left"
	p[0].Name = "Mirror"

}

//Insert an new car in the database + update index (CarId Separated by ,)
func (c *Car) CreateCar(stub shim.ChaincodeStubInterface, args []string) error {
	c.CarID = args[0]
	c.Brand = args[1]
	c.Type = args[2]
	c.Engine = args[3]
	c.Year, _ = strconv.Atoi(args[4])

	cJsonIndent, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println("CreateCar:", string(cJsonIndent))
	err := stub.PutState(c.CarID, cJsonIndent)
	if err != nil {
		return err
	}
	//Update Car index
	idxCarsByte, _ := stub.GetState("idx_Cars")
	if idxCarsByte == nil {
		err := stub.PutState("idx_Cars", []byte(args[0]))
		if err != nil {
			return err
		}
		return nil
	} else {
		idxCarsByte = []byte(string(idxCarsByte) + "," + args[0])
		err := stub.PutState("idx_Cars", idxCarsByte)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

//Get a car based on its ID
func (c *Car) GetCar(stub shim.ChaincodeStubInterface, carID string) ([]byte, error) {
	cJsonIndent, err := stub.GetState(carID)
	if err != nil {
		return nil, err
	}
	if cJsonIndent == nil {
		cJsonIndent = []byte("{\"Error\":\"Car with ID" + carID + " not found\"}")
	}
	fmt.Println("GetCarJSON returned:", string(cJsonIndent))
	return cJsonIndent, nil
}

//List all cars in the database
func (cs *Cars) ListCars(stub shim.ChaincodeStubInterface) ([]byte, error) {
	idxCarsByte, _ := stub.GetState("idx_Cars")
	carIDs := strings.Split(string(idxCarsByte), ",")
	carList := "{\"Cars\":"
	for i, carID := range carIDs {
		if i != 0 {
			carList = carList + ","
		}
		cJsonIndent, _ := stub.GetState(carID)
		carList = carList + string(cJsonIndent)
	}
	carList = carList + "\n}"
	return []byte(carList), nil
}

//Load Cars sample data
func (cs *Cars) LoadSample(stub shim.ChaincodeStubInterface) string {
	var c Car

	argslist := make([][]string, 6)
	argslist[0] = []string{"0001", "Renault", "Megane", "1600D", "2012"}
	argslist[1] = []string{"0002", "Mercedes", "C-Class", "220", "2014"}
	argslist[2] = []string{"0003", "Ford", "Focus", "1.8 16V", "2005"}
	argslist[3] = []string{"0004", "Renault", "Clio", "1200cc", "2014"}
	argslist[4] = []string{"0005", "Opel", "Astra", "1.9CDTI", "2011"}
	argslist[5] = []string{"0006", "Opel", "Astra", "2.0", "2010"}

	for _, args := range argslist {
		c.CreateCar(stub, args)
	}
	return "Load Car samples: 6 inserted"
}

