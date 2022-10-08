package VendorByMAC

import (
	"testing"
)

func TestMacDB_Init(t *testing.T) {
	MACExeample1 := [6]byte{0x00, 0x23, 0x9c, 0x00, 0x02, 0x12}
	MACExeamplePrefix281 := [6]byte{0x70, 0xb3, 0xd5, 0x15, 0x50, 0x031}
	MACExeample3 := [6]byte{0x00, 0x19, 0xb9, 0x85, 0x8b, 0x0d}
	MacHPnew := [6]byte{0x6c, 0x02, 0xe0, 0x0f, 0xa8, 0x74}
	MacN2 := [6]byte{0x00, 0xa2, 0x3c, 0xa0, 0x21, 0x07}
	var TestedStruct MacDB

	InitErr := TestedStruct.Init("oui2.txt", "oui.txt", "oui36.txt")
	if InitErr != nil {
		t.Errorf("Error in init: %s", InitErr)
	} else {
		if !TestedStruct.Initialized {
			if InitErr != nil {
				t.Errorf("Error in init")
			}
		} else {
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

			Er, Vn = TestedStruct.GetVendor(MACExeample3)
			if Er != nil {
				t.Errorf("Error in init: %s", Er)
			}
			if Vn != "Dell Inc." {
				t.Errorf("Expected vendor 'Juniper Networks', but got: %s", Vn)
			}

			Er, Vn = TestedStruct.GetVendor(MacHPnew)
			if Er != nil {
				t.Errorf("Error in init: %s", Er)
			}
			if Vn != "HP Inc." {
				t.Errorf("Expected vendor 'HP Inc.', but got: %s", Vn)
			}

			Er, Vn = TestedStruct.GetVendor(MacN2)
			if Er != nil {
				t.Errorf("Error in init: %s", Er)
			}
			if Vn != "" {
				t.Errorf("Expected vendor '', but got: %s", Vn)
			}
		}
	}

}
