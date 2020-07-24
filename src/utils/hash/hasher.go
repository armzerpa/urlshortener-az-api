package hash

import (
	"time"

	"github.com/speps/go-hashids"
)

type HasherService interface {
	GetHashId() string
}

type hasher struct {
}

func (r *hasher) GetHashId() string {
	var id string
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	id, _ = h.Encode([]int{int(now.Unix())})
	return id
}

func GetHasher() HasherService {
	return &hasher{}
}
