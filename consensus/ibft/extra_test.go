package ibft

import (
	"fmt"
	"github.com/0xPolygon/polygon-edge/helper/hex"
	"reflect"
	"testing"

	"github.com/0xPolygon/polygon-edge/types"
)

func TestExtraEncoding(t *testing.T) {
	seal1 := types.StringToHash("1").Bytes()

	cases := []struct {
		extra []byte
		data  *IstanbulExtra
	}{
		{
			data: &IstanbulExtra{
				Validators: []types.Address{
					types.StringToAddress("1"),
				},
				WhiteAccountValidators: []types.Address{
					types.StringToAddress("2"),
				},
				ProposerSeal: seal1,
				CommittedSeal: [][]byte{
					seal1,
				},
			},
		},
	}

	for _, c := range cases {
		data := c.data.MarshalRLPTo(nil)

		ii := &IstanbulExtra{}
		if err := ii.UnmarshalRLP(data); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(c.data, ii) {
			t.Fatal("bad")
		}
	}
}

func TestName(t *testing.T) {
	d := "0x0000000000000000000000000000000000000000000000000000000000000000d9d5944cffee4d09a9afe0928a876d6782d3feb2e25dd5c080c0"
	bys, err := hex.DecodeHex(d)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(bys))
	data := bys[IstanbulExtraVanity:]
	extra := &IstanbulExtra{}

	if err := extra.UnmarshalRLP(data); err != nil {
		t.Fatal(err)
	}

}
