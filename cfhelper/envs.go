package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"
)

type Environment string

func ParseEnv(nickname string) (Environment, error) {
	nickname = strings.ToLower(strings.TrimSpace(nickname))
	switch nickname {
	case "abn", "abnamro":
		return ABN, nil
	case "aus", "australia":
		return Australia, nil
	case "bra02", "brazil2", "bra2":
		return Brazil2, nil
	case "ind", "india":
		return India, nil
	case "investec":
		return Investec, nil
	case "irl", "ireland", "irlanda":
		return Ireland, nil
	case "itau", "itaú":
		return Itau, nil
	case "nequi":
		return Nequi, nil
	case "prod", "prod-mt", "mt", "multitenant":
		return Prod, nil
	case "usa", "us":
		return USA, nil

	case "ext", "sandbox", "dev", "development":
		return EXT, nil

	case "integration", "integ":
		return Integration, nil
	default:
		result := Environment(nickname)
		if !result.IsValid() {
			return "", fmt.Errorf("unknown environment: %s", nickname)
		}
		return result, nil
	}
}

const (
	ABN       Environment = "abnamro-prod"
	Australia Environment = "aus-prod"
	Brazil2   Environment = "bra02-prod"
	India     Environment = "ind-prod"
	Investec  Environment = "investec-prd"
	Ireland   Environment = "irl-prod"
	Itau      Environment = "itau"
	Nequi     Environment = "nequi-prod"
	Prod      Environment = "prod"
	USA       Environment = "usa-prod"

	EXT         Environment = "dev-ext"
	Integration Environment = "integ"
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
	case string(EXT), string(Integration):
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
		Investec,
		Itau,
		Nequi,
		Prod,
		USA,
		EXT,
		Integration,
	}
}

func ProductionEnvironments() []Environment {
	return lo.Filter(AllEnvironments(), func(e Environment, _ int) bool {
		return e.IsProduction()
	})
}
