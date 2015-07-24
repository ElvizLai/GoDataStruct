package cipher
import "fmt"

type Encipher func(plaintext string) []byte

//返回一个函数
func GenEncryptionFunc(encrypt Encipher) func(string) (ciphertext string) {
	return func(plaintext string) string {
		return fmt.Sprintf("%x", encrypt(plaintext))
	}
}


