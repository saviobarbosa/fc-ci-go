package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invalido: Resultado %d. Esperado %d", total, 30)
	}
}