package store

import (
	"github.com/kubernetes-incubator/cri-o/libkpod/info"
	"golang.org/x/net/context"
)

func init() {
	info.RegisterInfoGiver(storeInfo)
}

// top-level "store" info
func storeInfo(ctx context.Context) (string, map[string]interface{}, error) {
	store, err := getStore(c)
	if err != nil {
		return "store", nil, err
	}

	// lets say storage driver in use, number of images, number of containers
	i := map[string]interface{}{}
	i["GraphRoot"] = store.GraphRoot()
	i["GraphDriverName"] = store.GraphDriverName()
	if is, err := store.ImageStore(); err != nil {
		i["ImageStore"] = info.Err(err)
	} else {
		images, err := is.Images()
		if err != nil {
			i["ImageStore"] = info.Err(err)
		} else {
			i["ImageStore"] = map[string]interface{}{
				"number": len(images),
			}
		}
	}
	/* Oh this is in master on containers/storage, rebase later
	if is, err := store.ROImageStores(); err != nil {
		i["ROImageStore"] = info.Err(err)
	} else {
		images, err := is.Images()
		if err != nil {
			i["ROImageStore"] = info.Err(err)
		} else {
			i["ROImageStore"] = map[string]interface{}{
				"number": len(images),
			}
		}
	}
	*/
	if cs, err := store.ContainerStore(); err != nil {
		i["ContainerStore"] = info.Err(err)
	} else {
		containers, err := cs.Containers()
		if err != nil {
			i["ContainerStore"] = info.Err(err)
		} else {
			i["ContainerStore"] = map[string]interface{}{
				"number": len(containers),
			}
		}
	}
	return "store", i, nil
}
