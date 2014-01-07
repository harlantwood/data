package data

import (
	"io"
	"io/ioutil"
	"launchpad.net/goyaml"
	"os"
	"path"
)

type SerializedFile struct {
	Path   string      "-"
	Format interface{} "-"
}

func (f *SerializedFile) Marshal() ([]byte, error) {
	dOut("Marshalling %s\n", f.Path)
	return goyaml.Marshal(f.Format)
}

func (f *SerializedFile) Unmarshal(buf []byte) error {
	err := goyaml.Unmarshal(buf, f.Format)
	if err != nil {
		return err
	}

	dOut("Unmarshalling %s\n", f.Path)
	return nil
}

func (f *SerializedFile) Write(w io.Writer) error {
	buf, err := f.Marshal()
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	return err
}

func (f *SerializedFile) Read(r io.Reader) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return f.Unmarshal(buf)
}

func (f *SerializedFile) WriteFile() error {
	buf, err := f.Marshal()
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(f.Path), 0777)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(f.Path, buf, 0666)
}

func (f *SerializedFile) ReadFile() error {
	buf, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}

	return f.Unmarshal(buf)
}
