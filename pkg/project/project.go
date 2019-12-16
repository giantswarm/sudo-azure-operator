package project

var (
	bundleVersion = "0.0.1"
	description   = "The sudo-azure-operator does something."
	gitSHA        = "n/a"
	name          = "sudo-azure-operator"
	source        = "https://github.com/giantswarm/sudo-azure-operator"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
