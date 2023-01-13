package jwt

import "github.com/Hooneats/Syeong_server/common/ciper"

type JWT struct {
	Salt string
	//AccessDuration  int
	//RefreshDuration int
}

func (j *JWT) DecryptSalt() error {
	salt, err := ciper.AESDecrypt(ciper.CipherBlock, j.Salt)
	if err != nil {
		return err
	}
	j.Salt = salt
	return nil
}
