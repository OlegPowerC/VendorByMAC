package VendorByMAC

import (
	"testing"
)

func TestMacDB_Init(t *testing.T) {
	MACExeample1 := [6]byte{0x00, 0x23, 0x9c, 0x00, 0x02, 0x07}
	MACExeamplePrefix281 := [6]byte{0x70, 0xb3, 0xd5, 0x15, 0x50, 0x01}
	var TestedStruct MacDB

	InitErr := TestedStruct.Init("oui2.txt", "oui28.txt")
	if InitErr != nil {
		t.Errorf("Error in init: %s", InitErr)
	}

	if !TestedStruct.Initialized {
		if InitErr != nil {
			t.Errorf("Error in init")
		}
	}

	Er, Vn := TestedStruct.GetVendor(MACExeample1)
	if Er != nil {
		t.Errorf("Error in init: %s", Er)
	}
	if Vn != "Juniper Networks" {
		t.Errorf("Expected vendor 'Juniper Networks', but got: %s", Vn)
	}

	Er, Vn = TestedStruct.GetVendor(MACExeamplePrefix281)
	if Er != nil {
		t.Errorf("Error in init: %s", Er)
	}
	if Vn != "Sanwa New Tec Co.,Ltd" {
		t.Errorf("Expected vendor 'Sanwa New Tec Co.,Ltd', but got: %s", Vn)
	}
}
