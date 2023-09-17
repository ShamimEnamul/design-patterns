```go
// prototype interface

type Inode interface {
	print(string)
	clone() Inode
}
```
this portion will define the behavior

```go
// Concrete prototype :: File

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation, f.name)
}
func (f *File) clone() Inode {
	return &File{
		name: f.name + "_clone",
	}
}

```
File is concrete file of Inode type
```go
// Concrete prototype :: Folder

type Folder struct {
	childreen []Inode
	name      string
}

func (f *Folder) print(indentation string) {
	fmt.Printf(indentation, f.name)
	for _, i := range f.childreen {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildreen []Inode
	for _, i := range f.childreen {
		copyItem := i.clone()
		tempChildreen = append(tempChildreen, copyItem)
	}
	cloneFolder.childreen = tempChildreen
	return cloneFolder
}

```
Folder is concrete file of Inode type
```go
// client
func main() {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	folder1 := &Folder{
		name:      "folder1",
		childreen: []Inode{file1},
	}

	folder2 := &Folder{
		name:      "folder2",
		childreen: []Inode{folder1, file2, file3},
	}

	folder2.print("-")
	clone := folder2.clone()

	clone.print("-")
}

```