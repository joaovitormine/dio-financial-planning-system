package util

import (
	"testing"
)

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2019-02-12T10:10:10")

	if convertedTime.Year() != 2019 {
		t.Errorf("Espera que o ano seja 2019")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Espera que o mês seja 02")
	}

	if convertedTime.Day() != 12 {
		t.Errorf("Espera que o(s) dia(s) seja(m) 12")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Espera que a(s) hora(s) seja(m) 10")
	}

	if convertedTime.Minute() != 10 {
		t.Errorf("Espera que o(s) minuto(s) seja(m) 10")
	}

	if convertedTime.Second() != 10 {
		t.Errorf("Espera que o(s) segundo(s) seja(m) 10")
	}
}
