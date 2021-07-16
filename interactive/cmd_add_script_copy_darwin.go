package interactive

///copyFileToEnvironmentDirectory links the file
func copyFileToEnvironmentDirectory(src, dst string) error {
	return copyFile(src, dst)
}
