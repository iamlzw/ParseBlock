package main

import (
	"fmt"
	"github.com/hyperledger/fabric/common/tools/protolator"
	"github.com/hyperledger/fabric/protos/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)
func doInspectBlock(inspectBlock string) error {
	fmt.Println("Inspecting block")
	data, err := ioutil.ReadFile(inspectBlock)
	if err != nil {
		return fmt.Errorf("Could not read block %s", inspectBlock)
	}

	fmt.Println("Parsing genesis block")
	block, err := utils.UnmarshalBlock(data)
	if err != nil {
		return fmt.Errorf("error unmarshaling to block: %s", err)
	}
	err = protolator.DeepMarshalJSON(os.Stdout, block)
	if err != nil {
		return fmt.Errorf("malformed block contents: %s", err)
	}

	fmt.Println(data)
	return nil
}


func main(){
	blockDest := filepath.Join("mychannel_6.block")
	doInspectBlock(blockDest)
}
