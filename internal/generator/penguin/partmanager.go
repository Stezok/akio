package penguin

import (
	"NFTProject/internal/generator/penguin/rules"
	"NFTProject/internal/meta"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type PartManager struct {
	dataPath string

	metadata []meta.FragmentMetadata
}

func (pm *PartManager) GetPartPath(part *meta.FragmentMetadata) string {
	return fmt.Sprintf("%s/%s", pm.dataPath, part.GetFileName())
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
	Rules    rules.PinguinRules
	Filename string

	Slot     meta.Slot
	Color    meta.Color
	PromoTag meta.PromoTag
	Hash     int64
}

func (pm *PartManager) GetBackgroundPartsList() (list []*meta.FragmentMetadata) {
	for i := range pm.metadata {
		if pm.metadata[i].Slot == meta.Back {
			list = append(list, &pm.metadata[i])
		}
	}
	return
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

		if filter.Filename != "" && filter.Filename != frMeta.FileName {
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

func (pm *PartManager) GetRandomPartReadCloser(slot meta.Slot) io.ReadCloser {
	part := pm.GetPartFiltred(Filter{
		Slot:  slot,
		Color: meta.NoneColor,
	})

	path := pm.GetPartPath(part)
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
	}
	return file
}

func (pm *PartManager) GetPartReadCloserByFilename(filename string) (io.ReadCloser, error) {
	for _, part := range pm.metadata {
		if part.FileName == filename {
			path := pm.GetPartPath(&part)
			file, err := os.Open(path)
			if err != nil {
				log.Print(err)
			}
			return file, nil
		}
	}

	return nil, errors.New("Part not found")
}
