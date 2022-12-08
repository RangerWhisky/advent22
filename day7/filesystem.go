package day7

type Filesystem struct {
	directoryList map[string]Directory
	pwd           []string
}

func InitialiseFilesystem() Filesystem {
	fs_map := make(map[string]Directory)

	fs := Filesystem{
		directoryList: fs_map,
		pwd:           []string{},
	}

	return fs
}

func PrintWorkingDirectory(fs *Filesystem) string {
	return "/"
}

func ChangeDir(fs *Filesystem, newdir string) {

}
