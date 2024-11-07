package gdmisc

import "syscall"

func MakeDirHidden(dir *string) error {
	newDirPtr, err := syscall.UTF16PtrFromString(*dir)
	if err != nil {
		return err
	}

	// Set directory to hidden
	err = syscall.SetFileAttributes(newDirPtr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		return err
	}

	return nil
}
