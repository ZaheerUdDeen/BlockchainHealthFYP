/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the participant structure, with 4 properties.  Structure tags are used by encoding/json library

type Patient struct {
	FirstName   string `json:"firstName"`
	SecondName  string `json:"secondName"`
	Age 	    string `json:"age"`
	Address     string `json:"address"`
	
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run this Smart Contract"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "registerPatient" {
		return s.registerPatient(APIstub, args)
	}else if function == "initLedger" {
		return s.initLedger(APIstub)
	}else if function == "queryPatient" {
		return s.queryPatient(APIstub)
	} 

	return shim.Error("Invalid Smart Contract function name.")
}
func (s *SmartContract) queryPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(patientAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	patient := []Patient{
		Patient{FirstName: "Zaheer", SecondName: "wasa", Age: "22", Address: "islamabad"},
		Patient{FirstName: "Khaliq", SecondName: "Ab", Age: "22", Address: "islamabad"},
		
	}

	i := 0
	for i < len(patient) {
		fmt.Println("i is ", i)
		patientAsBytes, _ := json.Marshal(patient[i])
		APIstub.PutState("patient"+strconv.Itoa(i), patientAsBytes)
		fmt.Println("Added", patient[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) registerPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var patient = Patient{FirstName: args[1], SecondName: args[2], Age: args[3], Address: args[4]}

	patientAsBytes, _ := json.Marshal(patient)
	APIstub.PutState(args[0], patientAsBytes)

	return shim.Success(nil)
}


// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
