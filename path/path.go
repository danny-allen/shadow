package path

import (
	"os"
	"io"
)

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

// Copy file from source to destination.
func CopyFile(src string, dst string) error {

	// Open the source.
	in, err := os.Open(src)

	// Check for error.
	if (err != nil) { return err }

	// Defer the closing of source.
	defer in.Close()

	// Create the file specified in destination.
	out, err := os.Create(dst)

	// Check for error.
	if (err != nil) { return err }

	// Defer closing of the destination.
	defer out.Close()

	// Copy the contents from source to destination.
	_, err = io.Copy(out, in)

	// Close the destination file.
	cerr := out.Close()

	// Check for error.
	if (err != nil) { return err }

	// Return close error if any.
	return cerr
}