package file

import "testing"

func TestFile(t *testing.T) {
	root := NewFolder("root")
	println(root.Info())

	test := NewFolder("test")

	dev := NewFolder("dev")
	println(dev.Info())

	prod := NewFolder("prod")

	testFile := NewFile("测试文件")
	println(testFile.Info())

	root.Add(test)
	root.Add(dev)
	root.Add(prod)
	root.Add(testFile)

	println(root.Info())
	child := root.Child("prod")
	println(child.Info())

	files := root.Children()
	for _, f := range files {
		println(f.Info())
	}
}
