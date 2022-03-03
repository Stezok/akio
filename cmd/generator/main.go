package main

import (
	"NFTProject/internal/generator/penguin"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	// for i := 1; i <= 15; i++ {
	// 	path := fmt.Sprintf("./assets/gendata/metadata/%d", i)
	// 	entry, _ := os.ReadDir(path)
	// 	for _, e := range entry {
	// 		p := fmt.Sprintf("%s/%s", path, e.Name())
	// 		file, _ := os.Open(p)
	// 		data, _ := ioutil.ReadAll(file)
	// 		file.Close()
	// 		var fr meta.FragmentMetadata
	// 		json.Unmarshal(data, &fr)
	// 		// fr.Slot = meta.Slots[i-1]
	// 		arr := strings.Split(fr.FileName, "/")
	// 		// n, _ := strconv.Atoi(arr[1])
	// 		fr.FileName = arr[2]
	// 		data, _ = json.Marshal(fr)
	// 		f, _ := os.Create(p)
	// 		f.Write(data)
	// 		f.Close()
	// 	}
	// }

	gen := penguin.NewPinguinGenerator("./assets/gendata")
	rand.Seed(time.Now().Unix())
	for i := 0; i < 50; i++ {
		path := fmt.Sprintf("./generated/pengu#%d.png", i)
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		err = gen.GenerateRandomSingle(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
}
