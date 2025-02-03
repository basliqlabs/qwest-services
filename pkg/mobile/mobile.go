package main

import (
	"github.com/nyaruka/phonenumbers"
)

func IsValid(mobile string, region string) (bool, error) {
	num, err := phonenumbers.Parse("09399871525", "IR")
	if err != nil {
		return false, err
	}
	return phonenumbers.IsValidNumber(num), nil
}
