package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	HeaderSourceAccount  = "Source User"
	HeaderSourcePassword = "Source Password"
	HeaderRemoteAccount  = "Remote User"
	HeaderRemotePassword = "Remote Password"
)

func ReadCSV(path string) (Accounts, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	c, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	accounts := make(Accounts, 0, len(c)-1)
	for idx, row := range c {
		if idx == 0 {
			switch {
			case row[0] != HeaderSourceAccount:
				return nil, fmt.Errorf("first column in csv should be: %v", HeaderSourceAccount)
			case row[1] != HeaderSourcePassword:
				return nil, fmt.Errorf("second column in csv should be: %v", HeaderSourcePassword)
			case row[2] != HeaderRemoteAccount:
				return nil, fmt.Errorf("third column in csv should be: %v", HeaderRemoteAccount)
			case row[3] != HeaderRemotePassword:
				return nil, fmt.Errorf("fourth column in csv should be: %v", HeaderRemotePassword)
			}
			continue
		}
		accounts = append(accounts, NewSourceRemote(row, "Account", idx))

	}
	return accounts, nil
}
