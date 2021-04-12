package goonung

import "io/ioutil"

// Read : read a file and return bytes
func Read(path string, mode int) ([]byte, error) {
	var data []byte

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return data, err
	}

	return data, nil
}
