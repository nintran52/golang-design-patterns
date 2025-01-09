package main

import "fmt"

// inode.go: Prototype interface

// Step 1: Create a prototype interface and declare the clone method.
type Inode interface {
	print(indentation string)
	clone() Inode
}

// file.go: Concrete prototype

// Step 2: Define a concrete prototype for File.
type File struct {
	name string
}

// Implement the print method to display file details.
func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

// Implement the clone method to create a copy of the File object.
func (f *File) clone() Inode {
	// Step 3: Clone creates a new File with "_clone" appended to its name.
	return &File{name: f.name + "_clone"}
}

// folder.go: Concrete prototype

// Step 2: Define a concrete prototype for Folder.
type Folder struct {
	children []Inode
	name     string
}

// Implement the print method to display folder hierarchy.
func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

// Implement the clone method to create a copy of the Folder object.
func (f *Folder) clone() Inode {
	// Step 3: Clone creates a new Folder and clones all its children recursively.
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

// main.go: Client code

func main() {
	// Step 5: Client creates instances of File and Folder and interacts with them using the prototype interface.
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}

	// Print the original folder hierarchy.
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	// Clone the folder and print the hierarchy of the clone.
	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}

/*
Step-by-Step Implementation:

1. Created a prototype interface (Inode) with the clone method.
2. Defined concrete prototypes (File and Folder) and implemented the clone and print methods.
3. Ensured the clone methods create new instances with "_clone" appended to their names.
4. Folder.clone recursively clones its children to maintain hierarchy.
5. Client code demonstrates creating objects, printing hierarchies, and cloning.

Execution Result:
-----------------
Printing hierarchy for Folder2
  Folder2
    Folder1
        File1
    File2
    File3

Printing hierarchy for clone Folder
  Folder2_clone
    Folder1_clone
        File1_clone
    File2_clone
    File3_clone
*/
