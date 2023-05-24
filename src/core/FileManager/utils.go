package FileManager

import "fmt"

func GetFile(r *FileManager, name string) (*[]byte, error) {
	fmt.Println("Get file: ", name)
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
