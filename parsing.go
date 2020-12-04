package main

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/common/tools/protolator"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)
//将Block解析为Json
func BlockToJson(blkPath string) {
	fmt.Println("Inspecting block")
	data, err := ioutil.ReadFile(blkPath)
	if err != nil {
		fmt.Errorf("Could not read block %s", blkPath)
	}

	fmt.Println("Parsing genesis block")
	block, err := utils.UnmarshalBlock(data)
	if err != nil {
		fmt.Errorf("error unmarshaling to block: %s", err)
	}

	outBytes := new(bytes.Buffer)
	err = protolator.DeepMarshalJSON(outBytes, block)
	if err != nil {
		fmt.Errorf("malformed block contents: %s", err)
	}
	err = ioutil.WriteFile("mychannel_6.json",outBytes.Bytes(),0644)
	if err != nil{
		fmt.Println("write to file failure:",err)
	}
}
//将Json解析为Block
func JsonToBlock(jsonPath string) {
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Errorf("Could not read block %s", jsonPath)
	}

	blockBytes := bytes.NewReader(data)

	block := common.Block{}

	err = protolator.DeepUnmarshalJSON(blockBytes,&block)

	if err != nil {
		fmt.Errorf("malformed block contents: %s", err)
	}
	err = protolator.DeepMarshalJSON(os.Stdout, &block)
	if err != nil {
		fmt.Errorf("malformed block contents: %s", err)
	}
}
func main(){
	blkPath := filepath.Join("mychannel_6.block")
	BlockToJson(blkPath)

	jsonPath := filepath.Join("mychannel_6.json")
	JsonToBlock(jsonPath)
}
