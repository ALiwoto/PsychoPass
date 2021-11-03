package shellUtils

const (
	ShellToUse = "bash"
	gitCmd     = "git rev-parse --short HEAD &&" +
		" git rev-parse --verify HEAD &&" +
		" git fetch && " +
		"git rev-list --left-right --count origin/master...master"
)
