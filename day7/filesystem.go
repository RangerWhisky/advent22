package day7

type Filesystem struct {
	directoryList map[string]Directory
	pwd           []byte
}

func InitialiseFilesystem() Filesystem {
	fs_map := make(map[string]Directory)

	fs := Filesystem{
		directoryList: fs_map,
		pwd:           []byte{},
	}

	return fs
}

func PrintWorkingDirectory(fs *Filesystem) byte {
	return '/'
}
