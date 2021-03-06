package api

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/afero"
	"os"
	"runtime"
	"testing"
)

func TestFsAPI_ReadFile(t *testing.T) {
	fileData := "My Test File"
	memFs := afero.NewMemMapFs()
	afero.WriteFile(memFs, "test.txt", []byte(fileData), os.ModePerm)
	fsAPI := NewFsAPI(memFs)
	Convey("Check if FsAPI reads files without errors.", t, func() {
		s, err := fsAPI.ReadFile("test.txt")
		So(err, ShouldBeNil)
		Convey("FsAPI should return the right data", func() {
			So(s, ShouldEqual, fileData)
		})
	})
}

func TestFsAPI_WriteFile(t *testing.T) {
	fileData := "My Test File"
	memFs := afero.NewMemMapFs()
	fsAPI := NewFsAPI(memFs)
	Convey("Check if FsAPI writes files without errors.", t, func() {
		err := fsAPI.WriteFile("test.txt", fileData, false)
		So(err, ShouldBeNil)
	})
}

func TestFsAPI_Mkdir(t *testing.T) {
	memFs := afero.NewMemMapFs()
	fsAPI := NewFsAPI(memFs)
	Convey("Check if FsAPI creates folder without any errors.", t, func() {
		err := fsAPI.Mkdir("test_folder")
		So(err, ShouldBeNil)
		Convey("Check if FsAPI creates folder", func() {
			b, _ := afero.Exists(memFs, "test_folder")
			So(b, ShouldBeTrue)
		})
	})
}
func TestFsAPI_MkdirAll(t *testing.T) {
	memFs := afero.NewMemMapFs()
	fsAPI := NewFsAPI(memFs)
	Convey("Check if FsAPI creates folders and inner folders without any errors.", t, func() {
		err := fsAPI.MkdirAll("test_folder/test_inner/inner_2")
		So(err, ShouldBeNil)
		Convey("Check if FsAPI creates inner folders", func() {
			b, _ := afero.Exists(memFs, "test_folder/test_inner/inner_2")
			So(b, ShouldBeTrue)
		})
	})
}

func TestFsAPI_FilePathSeparator(t *testing.T) {
	memFs := afero.NewMemMapFs()
	fsAPI := NewFsAPI(memFs)
	Convey("Check if fileseperator is right.", t, func() {
		switch runtime.GOOS {
		case "linux":
			So(fsAPI.FilePathSeparator(), ShouldEqual, "/")
		case "darwin":
			So(fsAPI.FilePathSeparator(), ShouldEqual, "/")
		case "windows":
			So(fsAPI.FilePathSeparator(), ShouldEqual, "\\")
		}
	})
}

func TestFsAPI_Exists(t *testing.T) {
	memFs := afero.NewMemMapFs()
	memFs.Mkdir("test_folder", os.ModePerm)
	fsAPI := NewFsAPI(memFs)
	Convey("Check if FsAPI checks for existing file/folder without errors.", t, func() {
		b, err := fsAPI.Exists("test_folder")
		So(err, ShouldBeNil)
		Convey("Check if FsAPI finds existing folder", func() {
			So(b, ShouldBeTrue)
		})
	})
}

func TestFsAPI_Walk(t *testing.T) {
	memFs := afero.NewMemMapFs()
	memFs.MkdirAll("test_folder/test1/test2", os.ModePerm)
	afero.WriteFile(memFs, "test_folder/test1/test1.txt", []byte("Hello"), os.ModePerm)
	afero.WriteFile(memFs, "test_folder/test_folder.txt", []byte("Hello"), os.ModePerm)
	fsAPI := NewFsAPI(memFs)
	Convey("Check if FsAPI walks through all files.", t, func() {
		shouldWalk := []string{
			"test_folder",
			"test_folder/test1",
			"test_folder/test1/test1.txt",
			"test_folder/test1/test2",
			"test_folder/test_folder.txt",
		}
		walk := []string{}
		fsAPI.Walk("test_folder", func(path string, info os.FileInfo, err error) error {
			walk = append(walk, path)
			return nil
		})
		So(testEq(walk, shouldWalk), ShouldBeTrue)
	})
}
func testEq(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
