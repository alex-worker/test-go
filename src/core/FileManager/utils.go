package FileManager

import (
	. "test-go/src/interfaces/IResourceSystem"
)

func GetFile(r *IResourceSystem, name string) (*[]byte, error) {
	res, err := (*r).GetResource(name)
	if err != nil {
		return nil, err
	}
	buf, err := res.GetContent()
	if err != nil {
		return nil, err
	}
	return buf, nil
}
