package path

import "os"

func Exists(path string) (bool, error) {

	// Get file info.
	if _, err := os.Stat(path); err != nil {

		// Check it exists.
		if os.IsNotExist(err) {

			// file does not exist
			return false, err
		} else {

			// other error
			return false, err
		}
	}

	return true, nil;
}