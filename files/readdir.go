package files

import (
	"io/ioutil"
	"os"
)

func ReadFiles(srcDir string) (map[string]string, error) {
	result := make(map[string]string, 100)

	fileInfos, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}

	for _, file := range fileInfos {
		content, err := ioutil.ReadFile(srcDir + string(os.PathSeparator) + file.Name())
		if err != nil {
			return nil, err
		}

		result[file.Name()] = string(content)
	}
	return result, nil
}
