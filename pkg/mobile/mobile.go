package mobile

import (
	"github.com/nyaruka/phonenumbers"
)

func IsValid(mobile string) (bool, error) {
	num, err := phonenumbers.Parse(mobile, "")

	if err != nil {
		return false, err
	}

	return phonenumbers.IsValidNumber(num), nil
}

func GetRegion(mobile string) (string, error) {
	num, err := phonenumbers.Parse(mobile, "")

	if err != nil {
		return "", err
	}

	regionNumber := phonenumbers.GetRegionCodeForNumber(num)
	return regionNumber, nil
}

func NormalizePhoneNumber(mobile string) (string, error) {
	num, err := phonenumbers.Parse(mobile, "")

	if err != nil {
		return "", err
	}

	return phonenumbers.Format(num, phonenumbers.E164), nil
}
