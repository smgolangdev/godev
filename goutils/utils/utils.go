package utils

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
	//"strconv"
)

//Basic promise implementation.
func Go(f func() error) chan error {
	ch := make(chan error, 1)

	go func() {
		ch <- f()
	}()
	return ch
}

////////////////////////////////////////
// Copy a file from src to destination.
////////////////////////////////////////

func CopyFile(src, dest string) (int64, error) {
	if src == dest {
		return 0, nil
	}
	sfile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer sfile.Close()
	//Check to see if the destination file exists, if so remove it
	//if it cannot be removed, then return error.:

	if err := os.Remove(dest); err != nil && os.IsExist(err) {
		return 0, err
	}
	df, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer df.Close()
	return io.Copy(df, sfile)

}

/////////////////////////////////////////////////////
//Generates a random number which is not an integer
////////////////////////////////////////////////////
func GenerateId() string {
	id := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, id); err != nil {
		panic(err)
	}
	value := hex.EncodeToString(id)

	return value
}
