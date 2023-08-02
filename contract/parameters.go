package contract

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

const filePath = "configs/%s.yml"

type KeyValue struct {
	Key   string  `yaml:"key"`
	Value float64 `yaml:"value"`
}

type Parameters struct {
	Month   string `yaml:"month"`
	Savings struct {
		Bank     float64 `yaml:"bank"`
		Cashback float64 `yaml:"cashback"`
		FGTS     float64 `yaml:"fgts"`
	} `yaml:"savings"`
	Gains struct {
		Taxables    []KeyValue `yaml:"taxables"`
		NonTaxables []KeyValue `yaml:"non_taxables"`
	} `yaml:"gains"`
	Spents struct {
		Debits      []KeyValue `yaml:"debits"`
		Credits     []KeyValue `yaml:"credits"`
		CreditTotal float64    `yaml:"credit_total"`
	} `yaml:"spents"`
}

func LoadParameters() (*Parameters, error) {
	file := "parameters"

	fmt.Printf("loading config file from env=%s\n", file)

	f, err := os.Open(fmt.Sprintf(filePath, file))
	if err != nil {
		fmt.Printf("could not load config file from env=%s\n", file)
		return nil, err
	}
	defer f.Close()

	cfg := &Parameters{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf("could not parse config file from env=%s\n", file)
		return nil, err
	}

	return cfg, nil
}
