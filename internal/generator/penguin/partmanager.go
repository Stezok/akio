package penguin

import (
	"NFTProject/internal/generator/penguin/rules"
	"NFTProject/internal/meta"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type PartManager struct {
	metadata []meta.FragmentMetadata
}

func (pm *PartManager) LoadMetadata(dataPath string) error {

	for i := 0; i < len(meta.Slots); i++ {
		path := fmt.Sprintf("%s/%d", dataPath, i+1)
		entry, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for _, e := range entry {
			if !e.IsDir() {
				filePath := fmt.Sprintf("%s/%s", path, e.Name())
				file, err := os.Open(filePath)
				if err != nil {
					return err
				}
				defer file.Close()

				buf, err := ioutil.ReadAll(file)
				if err != nil {
					return err
				}
				var fragmentMeta meta.FragmentMetadata
				err = json.Unmarshal(buf, &fragmentMeta)
				if err != nil {
					return err
				}

				if _, err := os.Stat("./assets/gendata/" + fragmentMeta.GetFileName()); err != nil {
					log.Print(err)
				}

				pm.metadata = append(pm.metadata, fragmentMeta)
			}
		}
	}

	return nil
}

type Filter struct {
	Rules rules.PinguinRules

	Slot     meta.Slot
	Color    meta.Color
	PromoTag meta.PromoTag
	Hash     int64
}

func (pm *PartManager) GetPartFiltred(filter Filter) *meta.FragmentMetadata {
	action := filter.Rules.Decide(filter.Slot)
	if action == rules.Continue {
		return nil
	}

	var filteredIndexes []int
	for i, frMeta := range pm.metadata {
		if filter.Slot != frMeta.Slot {
			continue
		}

		if filter.PromoTag != frMeta.PromoTag {
			continue
		}

		if filter.Color != meta.NoneColor && frMeta.Color != filter.Color && frMeta.Color != meta.NoneColor {
			continue
		}

		filteredIndexes = append(filteredIndexes, i)
	}
	if len(filteredIndexes) == 0 {
		return nil
	}

	i := rand.Intn(len(filteredIndexes))
	return &pm.metadata[filteredIndexes[i]]
}
