package main

import (
	"slices"
	"strings"

	"github.com/samber/lo"
)

type Environment string

const (
	ABN       Environment = "abnamro-prod"
	Australia Environment = "aus-prod"
	Brazil2   Environment = "bra02-prod"
	India     Environment = "ind-prod"
	Ireland   Environment = "irl-prod"
	Itau      Environment = "itau"
	Nequi     Environment = "nequi-prod"
	Prod      Environment = "prod"
	USA       Environment = "usa-prod"
	EXT       Environment = "dev-ext"
)

func (e Environment) NonPCI() string {
	return string(e)
}

func (e Environment) PCI() string {
	if e == EXT {
		return e.NonPCI()
	}
	return string(e) + "-pci"
}

func (e Environment) IsMinor() bool {
	switch e {
	case Prod, Itau, India:
		return false
	default:
		return true
	}
}

func (e Environment) IsProduction() bool {
	switch e {
	case EXT:
		return false
	default:
		return true
	}
}

func (e Environment) IsValid() bool {
	nosuf, _ := strings.CutSuffix(string(e), "-pci")
	return slices.Contains(ProductionEnvironments(), Environment(nosuf))
}

func AllEnvironments() []Environment {
	return []Environment{
		ABN,
		Australia,
		Brazil2,
		India,
		Ireland,
		Itau,
		Nequi,
		Prod,
		USA,
		EXT,
	}
}

func ProductionEnvironments() []Environment {
	return lo.Filter(AllEnvironments(), func(e Environment, _ int) bool {
		return e.IsProduction()
	})
}

func MinorProductionEnvironments() []Environment {
	return lo.Filter(ProductionEnvironments(), func(e Environment, _ int) bool {
		return e.IsMinor()
	})
}

func MajorProductionEnvironments() []Environment {
	return lo.Filter(ProductionEnvironments(), func(e Environment, _ int) bool {
		return !e.IsMinor()
	})
}
