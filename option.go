package gdrive // nolint: golint

// Option can be used to pass optional Options to GDriver
type Option func(driver *GDriver) error

// RootDirectory sets the root directory for all operations
func RootDirectory(path string) Option {
	return func(driver *GDriver) error {
		_, err := driver.SetRootDirectory(path)
		return err
	}
}

// Sets the drive to act on
func RootDrive(drive string, path string) Option {
	return func(driver *GDriver) error {
		driver.SetDrive(drive)
		_, err := driver.SetRootDirectory(path)
		return err
	}
}
