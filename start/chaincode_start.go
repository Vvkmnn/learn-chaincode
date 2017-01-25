/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

// Dependencies =================================================

/*
fmt - contains Println for debugging/logging.
errors - standard go error format.
github.com/hyperledger/fabric/core/chaincode/shim - contains the definition for the chaincode interface and the chaincode stub, which you will need to interact with the ledger.
*/
import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// =================================================================
// Main
// ==================================================================

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

/*
Init is called when you first deploy your chaincode. As the name implies, this function should be used to do any initialization your chaincode needs.

In our example, we use Init to configure the initial state of a single key/value pair on the ledger.

In your chaincode_start.go file, change the Init function so that it stores the first element in the args argument to the key "hello_world".
*/

// func (t *SimpleChaincode) -
// Function:
// Init(, function string,
// Input:
// - stub: Take in a stub of type shim.ChaincodeStubInterface (from Hyperledger fabric)
// - function string: Take a function name as string
// - args []string: An array of args as type string
// Output:
// - []byte error: An array of bytes as type error

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 { // If insufficient arguments
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}
	fmt.Println("invoke did not find func: " + function) //error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "dummy_query" { //read a variable
		fmt.Println("hi there " + function) //error
		return nil, nil
	}
	fmt.Println("query did not find func: " + function) //error

	return nil, errors.New("Received unknown function query: " + function)
}
