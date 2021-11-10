package accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/daniel1sender/Desafio-API/pkg/gateways/store"
)

func (ar AccountRepository) CheckCPF(cpf string) error {
	/* 	var responseReason server_http.Error
	   	json.Unmarshal(newResponse.Body.Bytes(), &responseReason) */

	file, err := ioutil.ReadFile(ar.storage.Name())
	if err != nil {
		return nil
	}
	var response []AccountResponse
	_ = json.NewDecoder(bytes.NewBuffer(file)).Decode(&response)
	fmt.Print(response)
	for _, value := range response {
		if value.CPF == cpf {
			return store.ErrExistingCPF
		}
	}

	//json.Unmarshal(file, &response)

	//fmt.Println(response)
	/* 	file, err := ioutil.ReadFile(ar.storage.Name())
	   	if err != nil {
	   		return nil
	   	}
	   	for _, value := range string(file) {
	   		if string(value) == cpf {
	   			return store.ErrExistingCPF
	   		}
	   	} */
	return nil
}
