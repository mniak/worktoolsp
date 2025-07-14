package main

import (
	"slices"
	"strings"

	"github.com/samber/lo"
)

type Environment string

func GetEnv(nickname string) Environment {
	nickname = strings.ToLower(strings.TrimSpace(nickname))
	switch nickname {
	case "abn", "abnamro":
		return ABN
	case "aus", "australia":
		return Australia
	case "bra02", "brazil2", "bra2":
		return Brazil2
	case "ind", "india":
		return India
	case "irl", "ireland", "irlanda":
		return Ireland
	case "itau", "itaú":
		return Itau
	case "nequi":
		return Nequi
	case "prod", "prod-mt", "mt", "multitenant":
		return Prod
	case "usa", "us":
		return USA
	case "ext", "sandbox", "dev", "development":
		return EXT
	default:
		return Environment(nickname)
	}
}

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
	nosuffix, _ := strings.CutSuffix(string(e), "-pci")
	return nosuffix
}

func (e Environment) PCI() string {
	if e == EXT {
		return e.NonPCI()
	}
	return string(e) + "-pci"
}

func (e Environment) IsMajor() bool {
	switch e.NonPCI() {
	case string(Prod), string(Itau), string(India):
		return true
	default:
		return false
	}
}

func (e Environment) IsProduction() bool {
	switch e.NonPCI() {
	case string(EXT):
		return false
	default:
		return true
	}
}

func (e Environment) Is(other Environment) bool {
	return e.NonPCI() == other.NonPCI()
}

func (e Environment) IsValid() bool {
	nosuffix, _ := strings.CutSuffix(string(e), "-pci")
	return slices.Contains(ProductionEnvironments(), Environment(nosuffix))
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
