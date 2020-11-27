package main

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/common/tools/protolator"
	"github.com/hyperledger/fabric/protos/utils"
	"io/ioutil"
	"path/filepath"
)
func doInspectBlock(inspectBlock string)  {
	fmt.Println("Inspecting block")
	data, err := ioutil.ReadFile(inspectBlock)
	if err != nil {
		fmt.Errorf("Could not read block %s", inspectBlock)
	}

	fmt.Println("Parsing genesis block")
	block, err := utils.UnmarshalBlock(data)
	if err != nil {
		fmt.Errorf("error unmarshaling to block: %s", err)
	}
	buf := new (bytes.Buffer)
	err = protolator.DeepMarshalJSON(buf, block)
	if err != nil {
		fmt.Errorf("malformed block contents: %s", err)
	}
	err = ioutil.WriteFile("mychannel_6.json",buf.Bytes(),0644)
	if err != nil{
		fmt.Println("write to file failure:",err)
	}
}


func main(){
	blockDest := filepath.Join("mychannel_6.block")
	doInspectBlock(blockDest)
}
