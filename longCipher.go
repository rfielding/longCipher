package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
)

func doCipherByName(inName, outName string, key []byte, iv [aes.BlockSize]byte) error {
	inFile, err := os.Open(inName)
	if err != nil {
		return err
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(outName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writeCipher, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	writeCipherStream := cipher.NewOFB(writeCipher, iv[:])
	if err != nil {
		return err
	}

	reader := &cipher.StreamReader{S: writeCipherStream, R: inFile}
	if _, err := io.Copy(outFile, reader); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	theFile := "kali-linux-1.1.0a-amd64.iso"
	key := []byte("asdfasdfasdfasdf")
	var iv [aes.BlockSize]byte
	var err error

	err = doCipherByName(theFile, theFile+".encrypted", key, iv)
	if err != nil {
		panic(err)
	}
	err = doCipherByName(theFile+".encrypted", theFile+".decrypted", key, iv)
	if err != nil {
		panic(err)
	}
}
