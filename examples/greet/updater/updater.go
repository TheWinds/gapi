package updater

import (
	"bytes"
	"compress/gzip"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/zhiduoke/gapi/metadata"
	"github.com/zhiduoke/gapi/proto/pdparser"
	"io/ioutil"
	"sync"
)

type Updater struct {
	srvMD       map[string]*metadata.Metadata
	changedFunc func(md *metadata.Metadata)
	mu          sync.Mutex
}

func NewUpdater(f func(md *metadata.Metadata)) (*Updater, error) {
	updater := &Updater{changedFunc: f, srvMD: map[string]*metadata.Metadata{}}
	return updater, nil
}

func (u *Updater) SetChangedFn(f func(md *metadata.Metadata)) {
	u.mu.Lock()
	u.changedFunc = f
	u.mu.Unlock()
}

func (u *Updater) ReportChange(name string, data []byte) error {
	md, err := u.parseProtoFile(data)
	if err != nil {
		return err
	}
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.changedFunc == nil {
		return nil
	}
	u.srvMD[name] = md
	u.changedFunc(u.merge())
	return nil
}

func (u *Updater) merge() *metadata.Metadata {
	merged := new(metadata.Metadata)
	for _, md := range u.srvMD {
		merged.Routes = append(merged.Routes, md.Routes...)
	}
	return merged
}

func (u *Updater) unzipProtoFile(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return ioutil.ReadAll(reader)
}

func (u *Updater) parseProtoFile(zipData []byte) (*metadata.Metadata, error) {
	data, err := u.unzipProtoFile(zipData)
	if err != nil {
		return nil, err
	}
	var pd descriptor.FileDescriptorProto
	err = proto.Unmarshal(data, &pd)
	if err != nil {
		return nil, err
	}
	parser := pdparser.NewParser()
	if err = parser.AddFile(&pd); err != nil {
		return nil, err
	}
	parser.Resolve()
	routes, err := parser.CollectRoutes()
	if err != nil {
		return nil, err
	}
	return &metadata.Metadata{Routes: routes}, nil
}
