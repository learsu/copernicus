package script

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/copernet/copernicus/util"
)

func TestPublicKeyToAddress(t *testing.T) {
	publicKey := "03a34b99f22c790c4e36b2b3c2c35a36db06226e41c692fc82b8b56ac1c540c5bd"
	bytes, err := hex.DecodeString(publicKey)
	if err != nil {
		t.Fatal(err)
		return
	}
	address, err := AddressFromPublicKey(bytes)
	if err != nil {
		t.Fatal(err)
		return
	}
	hash160 := make([]byte, 20)
	copy(hash160[:], address.hash160[:])
	hash160Hex := hex.EncodeToString(hash160)
	if hash160Hex != "9a1c78a507689f6f54b847ad1cef1e614ee23f1e" {
		t.Errorf("hash160Hex is wrong 9a1c78a507689f6f54b847ad1cef1e614ee23f1e  --  %s", hash160Hex)
		return
	}
	if address.addressStr != "1F3sAm6ZtwLAUnj7d38pGFxtP3RVEvtsbV" {
		t.Errorf("address is wrong 1F3sAm6ZtwLAUnj7d38pGFxtP3RVEvtsbV  --  %s", address.addressStr)
		return
	}
}

func TestHash160Address(t *testing.T) {
	hash160, err := hex.DecodeString("0000000000000000000000000000000000000000")
	if hex.EncodeToString(hash160) != "0000000000000000000000000000000000000000" {
		t.Error(err)
		return
	}
	address, err := Hash160ToAddressStr(hash160, PublicKeyToAddress)
	if err != nil {
		t.Error(err)
		return
	}
	if address != "1111111111111111111114oLvT2" {
		t.Error("address is worng ,", address)
		return
	}
}

func TestAddressMatch(t *testing.T) {
	for v := 0; v < 100000; v++ {
		x := fmt.Sprintf("%02x", v)
		len1 := len(x)
		for i := len1; i <= 6; i++ {
			x = fmt.Sprintf("0%s", x)
		}
		result := fmt.Sprintf("000000000000000000000000000000000%s", x) //16进制 length=32

		hash160, err := hex.DecodeString(result)
		if err != nil {
			t.Error(err) // encoding/hex: odd length hex string
		}
		address, err1 := Hash160ToAddressStr(hash160, PublicKeyToAddress)
		if err1 != nil {
			t.Error(err1) //hash160 length 0 not 20
			return
		}

		if strings.Contains(address, "bwhc") {
			fmt.Printf("%v=====%v\n", result, address)
		}

		if strings.Contains(address, "BWHC") {
			fmt.Printf("%v=====%v\n", result, address)
		}
	}
}

func TestAdd(t *testing.T) {
	hash160, err := hex.DecodeString("000000000000000000000000000000000001210f")
	if err != nil {
		t.Error(err) // encoding/hex: odd length hex string
	}
	address, err1 := Hash160ToAddressStr(hash160, PublicKeyToAddress)
	if err1 != nil { //
		t.Error(err1) //hash160 length 0 not 20
		return
	}

	fmt.Println(address)
}

func TestHash160ToAddress(t *testing.T) {
	data, err := hex.DecodeString("0014a4b4ca48de0b3fffc15404a1acdc8dbaae226955")
	if err != nil {
		t.Error(err)
		return
	}
	hash160 := util.Hash160(data)
	if hex.EncodeToString(hash160) != "2928f43af18d2d60e8a843540d8086b305341339" {
		t.Error(err)
		return
	}
	address, err := Hash160ToAddressStr(hash160, ScriptToAddress)
	if err != nil {
		t.Error(err)
		return
	}
	if address != "35SegwitPieWKVHieXd97mnurNi8o6CM73" {
		t.Error("address is worng ,", address)
		return
	}
}

func TestPrivateKeyToAddress(t *testing.T) {
	address, err := AddressFromPrivateKey("5KYZdUEo39z3FPrtuX2QbbwGnNP5zTd7yyr2SC1j299sBCnWjss")
	if err != nil {
		t.Error(err)
	}
	if address.addressStr != "1HZwkjkeaoZfTSaJxDw6aKkxp45agDiEzN" {
		t.Errorf("address (%s) is error", address.addressStr)
	}

	address, err = AddressFromPrivateKey("L4rK1yDtCWekvXuE6oXD9jCYfFNV2cWRpVuPLBcCU2z8TrisoyY1")
	if err != nil {
		t.Error(err)
	}
	if address.addressStr != "1F3sAm6ZtwLAUnj7d38pGFxtP3RVEvtsbV" {
		t.Errorf("address (%s) is error", address.addressStr)
	}

}

func TestPrivateKeyFromHex(t *testing.T) {

	//str1 := "3714c34e68f8481d"
	//str2 := "9e3647445d5ca65e"
	//str3 := "9d150ddb24d2182a"
	//str4 := "6ac12143f1293835"
	//pri := core.PrivateKeyFromBytes(hexToBytes(str1 + str2 + str3 + str4))
	//if pri == nil {
	//	t.Error("pri is nil")
	//}
	//pub := pri.PubKey()
	//pub.Compressed = true
	//address, err := AddressFromPublicKey(pub.ToBytes())
	//if err != nil {
	//	t.Error(err.Error())
	//}
}
