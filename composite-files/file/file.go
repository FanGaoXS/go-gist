package file

import "fmt"

type Filename = string

type CompositeFile interface {
	Name() string
	Info() string
	Add(file CompositeFile)
	Remove(name Filename)
	Child(name Filename) CompositeFile
	Children() []CompositeFile
}

func NewFile(name string) *File {
	return &File{name: name}
}

type File struct {
	name Filename
}

func (f File) Name() string {
	return f.name
}

func (f File) Info() string {
	return fmt.Sprintf("name = %s, type= %s", f.name, "file")
}

func (f File) Add(_ CompositeFile) { return }

func (f File) Remove(_ Filename) { return }

func (f File) Child(_ Filename) CompositeFile { return nil }

func (f File) Children() []CompositeFile { return nil }

func NewFolder(name string) *Folder {
	return &Folder{
		name:     name,
		elements: make(map[Filename]CompositeFile),
	}
}

type Folder struct {
	name     Filename
	elements map[Filename]CompositeFile
}

func (f *Folder) Name() string {
	return f.name
}

func (f *Folder) Info() string {
	return fmt.Sprintf("name = %s, type= %s, len= %d", f.name, "folder", len(f.elements))
}

func (f *Folder) Add(file CompositeFile) {
	f.elements[file.Name()] = file
}

func (f *Folder) Remove(name Filename) {
	delete(f.elements, name)
}

func (f *Folder) Child(name Filename) CompositeFile {
	return f.elements[name]
}

func (f *Folder) Children() []CompositeFile {
	res := make([]CompositeFile, 0, len(f.elements))
	for _, cf := range f.elements {
		res = append(res, cf)
	}
	return res
}
