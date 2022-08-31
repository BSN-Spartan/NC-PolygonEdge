package validate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/0xPolygon/polygon-edge/network/common"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/mr-tron/base58/base58"
)

type ValidateInfo struct {
	nodeId peer.ID
	//Nonce   string `json:"key"`
	Address string `json:"address"`
	Mac     string `json:"validate_info"`
}

func (v *ValidateInfo) Sign(priv crypto.PrivKey) error {

	err := v.padding()
	if err != nil {
		return err
	}

	id, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		return fmt.Errorf("private key error")
	}

	if id != v.nodeId {
		return fmt.Errorf("id diff")
	}

	sign, err := priv.Sign(v.dist())
	if err != nil {
		return fmt.Errorf("sign failed")
	}

	v.Mac = base58.Encode(sign)
	return nil
}

//func (v *ValidateInfo) getNonce() []byte {
//	if v.Nonce != "" {
//		nonec, err := base58.Decode(v.Nonce)
//		if err == nil {
//			return nonec
//		}
//	}
//
//	nonce := GetRandomBytes(32)
//	v.Nonce = base58.Encode(nonce)
//
//	return nonce
//}

func (v *ValidateInfo) dist() []byte {
	data := []byte(v.Address)

	//data = append(data, v.getNonce()...)

	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func (v *ValidateInfo) padding() error {
	addr, err := common.StringToAddrInfo(v.Address)
	if err != nil {
		return fmt.Errorf("address format error")
	}
	v.nodeId = addr.ID
	return nil
}

func (v *ValidateInfo) Validate() (bool, error) {
	err := v.padding()
	if err != nil {
		return false, err
	}

	pubk, err := v.nodeId.ExtractPublicKey()
	if err != nil {
		return false, fmt.Errorf("node id format error")
	}
	mac, err := base58.Decode(v.Mac)
	if err != nil {
		return false, fmt.Errorf("decode failed")
	}

	return pubk.Verify(v.dist(), mac)
}

func (v *ValidateInfo) ToJson() string {
	jb, _ := json.Marshal(v)
	return string(jb)
}

// GetRandomBytes returns len random looking bytes
func GetRandomBytes(len int) []byte {
	key := make([]byte, len)

	// TODO: rand could fill less bytes then len
	_, err := rand.Read(key)
	if err != nil {
		key = []byte("nonce")
	}
	return key
}
