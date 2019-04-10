package game

import (
	// "fmt"
	// "bufio"
	"os"
)

type Level struct {
	Map[][]Tile
}

func LoadLevelFromFile(filename string) *Level {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scanner :=  bufio.NewScanner(file)
	// for ( scanner.Scan() ){
	// 	fmt.Println(scanner.Text())
	// }
	return nil
}