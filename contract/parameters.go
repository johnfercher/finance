package contract

import (
	"finance/m/v2/domain/consts/month"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

const filePath = "configs/%s.yml"

type Transaction struct {
	Label      string        `yaml:"label"`
	Value      float64       `yaml:"value"`
	Recurrence []month.Month `yaml:"recurrence"`
}

type Parameters struct {
	YearCDIPercent float64     `yaml:"year_cdi_percent"`
	Month          month.Month `yaml:"month"`
	MonthsDuration int         `yaml:"months_duration"`
	Savings        struct {
		Bank     float64 `yaml:"bank"`
		Cashback float64 `yaml:"cashback"`
		FGTS     float64 `yaml:"fgts"`
	} `yaml:"savings"`
	Gains struct {
		Taxables    []Transaction `yaml:"taxables"`
		NonTaxables []Transaction `yaml:"non_taxables"`
	} `yaml:"gains"`
	Spents struct {
		Debits      []Transaction `yaml:"debits"`
		Credits     []Transaction `yaml:"credits"`
		CreditTotal float64       `yaml:"credit_total"`
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
