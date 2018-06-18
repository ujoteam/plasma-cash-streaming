package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ChildChainService child client to reference server
type ChildChainService struct {
	url string
}

func (c *ChildChainService) CurrentBlock() (error, *Block) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/block", c.url), nil)
	if err != nil {
		fmt.Print(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("CurrentBlock Body:", string(body))
	block := Block{blockId: string(body)}
	return nil, &block
}

func (c *ChildChainService) BlockNumber() int64 {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocknumber", c.url), nil)
	if err != nil {
		fmt.Print(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("BlockNumber Body:", string(body))

	i, _ := strconv.Atoi(string(body))
	return int64(i)
}

func (c *ChildChainService) Block(blknum int64) (error, *Block) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/block/%d", c.url, blknum), nil)
	if err != nil {
		fmt.Print(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("Block Body:", string(body))
	block := Block{blockId: string(body)}
	return nil, &block
}

func (c *ChildChainService) Proof(blknum int64, uid int64) (error, *Proof) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/proof?blknum=%d&uid=%d", c.url, blknum, uid), nil)
	if err != nil {
		fmt.Print(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	proof := Proof{proofstring: string(body)}
	return nil, &proof
}

func (c *ChildChainService) SubmitBlock(block *Block) error {

	data := url.Values{}
	data.Set("block", block.blockId)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/submit_block", c.url), strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	return err
}

func (c *ChildChainService) SendTransaction() error {

	return nil
}

func NewChildChainService(url string) ChainServiceClient {
	return &ChildChainService{url}
}
