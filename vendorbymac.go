package VendorByMAC

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	MacUtils "github.com/OlegPowerC/macaddress_utils"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type OUIprefix28_36 struct {
	OUIfirst6byteStr string
	NextByteMin      uint32
	NextByteMax      uint32
	Vendor           string
}

type MacDB struct {
	OUIMap         map[string]string
	Prefix28MacMap map[string][]OUIprefix28_36
	Initialized    bool
}

func (MacData *MacDB) Init(FileFull string, FilePrefix28 string) error {
	MacData.Initialized = false
	file, err := os.Open(FileFull)
	if err != nil {
		return err
	}
	defer file.Close()

	ouimap := make(map[string]string)
	OUIv28prefix := make(map[string][]OUIprefix28_36)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str1 := scanner.Text()
		splitted := strings.Split(str1, ":")
		if len(splitted) == 2 {
			oui_map_key := strings.Trim(splitted[0], "\"")
			oui_map_value := strings.Trim(splitted[1], "\"")
			if len(oui_map_key) > 1 && len(oui_map_value) > 1 {
				ouimap[strings.ToLower(oui_map_key)] = oui_map_value
			}
		}
	}

	MacData.OUIMap = ouimap

	file28, err28 := ioutil.ReadFile(FilePrefix28)
	if err28 != nil {
		return err28
	}

	stringt28 := string(file28)
	OUIp28 := strings.Split(stringt28, "\r\n\r\n")

	for _, OUI28val := range OUIp28 {
		OUI28valPs := strings.Split(OUI28val, "\r\n")
		ValidF6b := false
		F6BOUIstrs := ""
		for vin1, vval1 := range OUI28valPs {
			if vin1 == 0 {
				if len(vval1) > 8 {
					vb, _ := regexp.MatchString(`^[0-9A-F]{2}-[0-9A-F]{2}-[0-9A-F{2}]`, vval1)
					F6BOUIstrs = strings.ToLower(fmt.Sprintf("%s%s%s", vval1[0:2], vval1[3:5], vval1[6:8]))
					ValidF6b = vb
				}
			}
			if vin1 == 1 && ValidF6b {
				vb2, _ := regexp.MatchString(`^[0-9A-F]{6}-[0-9A-F]{6}`, vval1)
				if vb2 {
					MinStr := fmt.Sprintf(vval1[0:6])
					MaxStr := fmt.Sprintf(vval1[7:13])
					hexstringtobytes, hexstringtobyteserr := hex.DecodeString(MinStr)
					hexstringtobytes2, hexstringtobyteserr2 := hex.DecodeString(MaxStr)
					if hexstringtobyteserr == nil && hexstringtobyteserr2 == nil {
						vendor := ""
						bs1 := make([]byte, 1)
						bs1 = append(bs1, hexstringtobytes...)
						minuint32 := binary.BigEndian.Uint32(bs1)
						bs2 := make([]byte, 1)
						bs2 = append(bs2, hexstringtobytes2...)
						maxuint32 := binary.BigEndian.Uint32(bs2)
						findvendorsbytabs := strings.Split(vval1, "\t\t")
						if len(findvendorsbytabs) == 2 {
							vendor = findvendorsbytabs[1]
						}
						if _, ok := OUIv28prefix[F6BOUIstrs]; ok {
							Oldv := OUIv28prefix[F6BOUIstrs]
							Oldv = append(Oldv, OUIprefix28_36{F6BOUIstrs, minuint32, maxuint32, vendor})
							OUIv28prefix[F6BOUIstrs] = Oldv

						} else {
							Oldv := make([]OUIprefix28_36, 0)
							Oldv = append(Oldv, OUIprefix28_36{F6BOUIstrs, minuint32, maxuint32, vendor})
							OUIv28prefix[F6BOUIstrs] = Oldv
						}
					}
				}
			}
		}
	}

	MacData.Prefix28MacMap = OUIv28prefix
	MacData.Initialized = true
	return nil
}

func (MacData *MacDB) GetVendor(MacAddr [6]byte) (err error, vendor string) {
	Ven := ""
	if !MacData.Initialized {
		return errors.New("MAC DB not initialized"), Ven
	}
	_, macst := MacUtils.SNMPMACfrom6bytestoHexString(MacAddr, MacUtils.MACFORMAT_XXXXdotXXXXdotXXXX)
	MAC6string := macst[0:4] + macst[5:7]

	//Поиск вендора в длинных префиксах
	venlongprefix, venlogprefixexsist := MacData.Prefix28MacMap[MAC6string]
	if venlogprefixexsist {
		lastmacaddrsbytes := make([]byte, 1)
		lastmacaddrsbytes = append(lastmacaddrsbytes, MacAddr[3], MacAddr[4], MacAddr[5])
		MacLastDigit := binary.BigEndian.Uint32(lastmacaddrsbytes)
		for _, MacRange := range venlongprefix {
			if MacLastDigit >= MacRange.NextByteMin && MacLastDigit <= MacRange.NextByteMax {
				Ven = MacRange.Vendor
				break
			}
		}
	} else {
		VendorFromFull, VendorFromFullKey := MacData.OUIMap[MAC6string]
		if VendorFromFullKey {
			Ven = VendorFromFull
		}
	}
	return nil, Ven
}
