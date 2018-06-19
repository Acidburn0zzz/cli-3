package constraints

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/pkg/projectfile"
	"github.com/ActiveState/sysinfo"
	"github.com/stretchr/testify/assert"
)

func TestPlatformConstraints(t *testing.T) {
	root, _ := environment.GetRootPath()
	project, err := projectfile.Parse(filepath.Join(root, "test", constants.ConfigFileName))
	project.Persist()
	assert.Nil(t, err, "There was no error parsing the config file")

	assert.True(t, platformIsConstrained("Windows10Label"))
}

func TestEnvironmentConstraints(t *testing.T) {
	os.Setenv(constants.EnvironmentEnvVarName, "dev")
	assert.False(t, environmentIsConstrained("dev"), "The current environment is in 'dev'")
	assert.False(t, environmentIsConstrained("dev,qa"), "The current environment is in 'dev,qa'")
	assert.False(t, environmentIsConstrained("qa,dev,prod"), "The current environment is in 'dev,qa,prod'")
	assert.True(t, environmentIsConstrained("qa"), "The current environment is not in 'qa'")
	assert.True(t, environmentIsConstrained("qa,devops"), "The current environment is not in 'qa,devops'")
}

func TestMatchConstraint(t *testing.T) {
	root, _ := environment.GetRootPath()
	project, err := projectfile.Parse(filepath.Join(root, "test", constants.ConfigFileName))
	project.Persist()
	assert.Nil(t, err, "There was no error parsing the config file")

	constraint := projectfile.Constraint{"Windows10Label", "dev"}
	assert.True(t, IsConstrained(constraint))
}

func TestOsMatches(t *testing.T) {
	for _, name := range osNames {
		osOverride = name
		assert.True(t, osMatches(name), "OS matches with override")
	}
	osOverride = "" // reset
}

func TestOsVersionMatches(t *testing.T) {
	// Linux tests.
	osVersionOverride = "4.10.0 Ubuntu 16.04.3 LTS"
	assert.False(t, osVersionMatches("4.10.1"), "Newer kernel required")
	assert.False(t, osVersionMatches("4.11"), "Newer kernel required")
	assert.False(t, osVersionMatches("5"), "Newer kernel required")
	assert.True(t, osVersionMatches("4.10.0"), "Kernel matches")
	assert.True(t, osVersionMatches("4.10"), "Kernel matches")
	assert.True(t, osVersionMatches("4.09.1"), "Older kernel is okay")
	assert.True(t, osVersionMatches("4.09"), "Older kernel is okay")
	assert.True(t, osVersionMatches("4"), "Older kernel is okay")

	// Windows tests.
	osVersionOverride = "6.1.999 Windows 7"
	assert.False(t, osVersionMatches("6.2.0"), "Windows 8 required")
	assert.False(t, osVersionMatches("6.2"), "Windows 8 required")
	assert.False(t, osVersionMatches("10"), "Windows 10 required")
	assert.True(t, osVersionMatches("6.1.0"), "Windows 7 is okay")
	assert.True(t, osVersionMatches("6.0"), "Windows Vista is okay")

	// macOS tests.
	osVersionOverride = "10.6.2 Mac OS X"
	assert.False(t, osVersionMatches("10.7.0"), "Lion required")
	assert.False(t, osVersionMatches("10.7"), "Lion required")
	assert.False(t, osVersionMatches("10.10"), "Mavericks required")
	assert.True(t, osVersionMatches("10.5.0"), "Leopard is okay")
	assert.True(t, osVersionMatches("10.4"), "Tiger is okay")

	osVersionOverride = "" // reset
}

func TestArchMatches(t *testing.T) {
	for _, name := range archNames {
		archOverride = name
		assert.True(t, archMatches(name), "Architecture matches with override")
	}
	archOverride = "" // reset
}

func TestLibcMatches(t *testing.T) {
	// Linux tests.
	libcOverride = "glibc 2.23"
	assert.False(t, libcMatches("glibc 2.24"), "Newer glibc required")
	assert.False(t, libcMatches("glibc 3.0"), "Newer glibc required")
	assert.True(t, libcMatches("glibc 2.23"), "glibc matches")
	assert.True(t, libcMatches("glibc 2.22"), "Older glibc is okay")
	assert.True(t, libcMatches("glibc 1.0"), "Older glibc is okay")
	assert.False(t, libcMatches("musl 2.23"), "Non-glibc (musl) is not okay")
	assert.False(t, libcMatches("musl 2"), "Non-glibc (musl) is not okay")
	assert.True(t, libcMatches("GLIBC 2.23"), "Case-insensitive matching")

	// Windows tests.
	libcOverride = "msvcrt 7.0"
	assert.False(t, libcMatches("msvcrt 8.0"), "Newer msvcrt required")
	assert.True(t, libcMatches("msvcrt 7.0"), "msvcrt matches")
	assert.True(t, libcMatches("msvcrt 6.0"), "Older msvcrt is okay")
	assert.False(t, libcMatches("glibc 2.23"), "Non-msvcrt (glibc) is not okay")
	assert.True(t, libcMatches("MSVCRT 7.0"), "Case-insensitive matching")

	// macOS tests.
	libcOverride = "bsdlibc 3.2"
	assert.False(t, libcMatches("bsdlibc 3.4"), "Newer bsdlibc required")
	assert.False(t, libcMatches("bsdlibc 4.0"), "Newer bsdlibc required")
	assert.True(t, libcMatches("bsdlibc 3.2"), "bsdlibc matches")
	assert.True(t, libcMatches("bsdlibc 3.0"), "Older bsdlibc is okay")
	assert.True(t, libcMatches("bsdlibc 2.0"), "Older bsdlibc is okay")
	assert.True(t, libcMatches("BSDlibc 3.2"), "Case-insensitive matching")

	libcOverride = "" // reset
}

func TestCompilerMatches(t *testing.T) {
	// Linux tests.
	compilerOverride = "gcc 5.2"
	assert.False(t, compilerMatches("gcc 5.4"), "Newer GCC required")
	assert.False(t, compilerMatches("gcc 6"), "Newer GCC required")
	assert.True(t, compilerMatches("gcc 5.2"), "GCC matches")
	assert.True(t, compilerMatches("gcc 5"), "Older GCC is okay")
	assert.True(t, compilerMatches("gcc 4"), "Older GCC is okay")
	assert.False(t, compilerMatches("clang 3.4"), "Non-GCC (Clang) is not okay")
	assert.True(t, compilerMatches("GCC 5.2"), "Case-insensitive matching")

	// Windows tests.
	compilerOverride = "cl 17.00"
	assert.False(t, compilerMatches("cl 19.00"), "Newer cl required")
	assert.False(t, compilerMatches("cl 19"), "Newer cl required")
	assert.True(t, compilerMatches("cl 17.00"), "cl matches")
	assert.True(t, compilerMatches("cl 17"), "cl matches")
	assert.True(t, compilerMatches("cl 15.00"), "Older cl is okay")
	assert.True(t, compilerMatches("cl 15"), "Older cl is okay")
	assert.False(t, compilerMatches("mingw 5.4"), "Non-cl (MinGW) is not okay")
	assert.True(t, compilerMatches("CL 17"), "Case-insensitive matching")

	// macOS tests.
	compilerOverride = "clang 6.0"
	assert.False(t, compilerMatches("clang 7.0"), "Newer clang required")
	assert.False(t, compilerMatches("clang 7"), "Newer clang required")
	assert.True(t, compilerMatches("clang 6.0"), "clang matches")
	assert.True(t, compilerMatches("clang 6"), "clang matches")
	assert.True(t, compilerMatches("clang 4"), "Older clang is okay")
	assert.True(t, compilerMatches("clang 3.4"), "Older clang is okay")
	assert.True(t, compilerMatches("CLANG 6"), "Case-insensitive matching")

	compilerOverride = "" // reset
}

// This test is not for constraints, but verifies that sysinfo is working
// correctly in a Linux development environment such that constraints will have
// an effect.
func TestSysinfoLinuxEnv(t *testing.T) {
	if sysinfo.OS() != sysinfo.Linux || os.Getenv("CIRCLECI") != "" {
		return // skip
	}
	assert.Equal(t, sysinfo.Linux, sysinfo.OS(), "Linux is the OS")
	version, err := sysinfo.OSVersion()
	assert.NoError(t, err, "No errors detecting OS version")
	assert.True(t, version.Major > 0, "Determined kernel version")
	assert.NotEqual(t, sysinfo.UnknownArch, sysinfo.Architecture(), "Architecture was recognized")
	libc, err := sysinfo.Libc()
	assert.NoError(t, err, "No errors detecting a Libc")
	assert.NotEqual(t, sysinfo.UnknownLibc, libc.Name, "Libc name was recognized")
	assert.True(t, libc.Major > 0, "Determined Libc version")
	compilers, err := sysinfo.Compilers()
	assert.NoError(t, err, "No errors detecting a compiler")
	for _, compiler := range compilers {
		assert.True(t, compiler.Major > 0, "Determined compiler version")
	}
}

// This test is not for constraints, but verifies that sysinfo is working
// correctly in a Windows development environment such that constraints will
// have an effect.
func TestSysinfoWindowsEnv(t *testing.T) {
	if sysinfo.OS() != sysinfo.Windows || os.Getenv("CIRCLECI") != "" {
		return // skip
	}
	assert.Equal(t, sysinfo.Windows, sysinfo.OS(), "Windows is the OS")
	version, err := sysinfo.OSVersion()
	assert.NoError(t, err, "No errors detecting OS version")
	assert.True(t, version.Major > 0, "Determined OS version")
	assert.NotEqual(t, sysinfo.UnknownArch, sysinfo.Architecture(), "Architecture was recognized")
	libc, err := sysinfo.Libc()
	assert.NoError(t, err, "No errors detecting a Libc")
	assert.NotEqual(t, sysinfo.UnknownLibc, libc.Name, "Libc name was recognized")
	assert.True(t, libc.Major > 0, "Determined Libc version")
	compilers, err := sysinfo.Compilers()
	assert.NoError(t, err, "No errors detecting a compiler")
	for _, compiler := range compilers {
		assert.True(t, compiler.Major > 0, "Determined compiler version")
	}
}

// This test is not for constraints, but verifies that sysinfo is working
// correctly in a macOS development environment such that constraints will have
// an effect.
func TestSysinfoMacOSEnv(t *testing.T) {
	if sysinfo.OS() != sysinfo.Mac || os.Getenv("CIRCLECI") != "" {
		return // skip
	}
	assert.Equal(t, sysinfo.Mac, sysinfo.OS(), "macOS is the OS")
	version, err := sysinfo.OSVersion()
	assert.NoError(t, err, "No errors detecting OS version")
	assert.True(t, version.Major > 0, "Determined OS version")
	assert.NotEqual(t, sysinfo.UnknownArch, sysinfo.Architecture(), "Architecture was recognized")
	libc, err := sysinfo.Libc()
	assert.NoError(t, err, "No errors detecting a Libc")
	assert.NotEqual(t, sysinfo.UnknownLibc, libc.Name, "Libc name was recognized")
	assert.True(t, libc.Major > 0, "Determined Libc version")
	compilers, err := sysinfo.Compilers()
	assert.NoError(t, err, "No errors detecting a compiler")
	for _, compiler := range compilers {
		assert.True(t, compiler.Major > 0, "Determined compiler version")
	}
}
