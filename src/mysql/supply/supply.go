package supply

import (
	"fmt"
	"io"

	"github.com/cloudfoundry/libbuildpack"
)

type Stager interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/stager.go
	BuildDir() string
	DepDir() string
	DepsIdx() string
	DepsDir() string
	WriteProfileD(string, string) error
}

type Manifest interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/manifest.go
	AllDependencyVersions(string) []string
	DefaultVersion(string) (libbuildpack.Dependency, error)
}

type Installer interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/installer.go
	InstallDependency(libbuildpack.Dependency, string) error
	InstallOnlyVersion(string, string) error
}

type Command interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/command.go
	Execute(string, io.Writer, io.Writer, string, ...string) error
	Output(dir string, program string, args ...string) (string, error)
}

type Supplier struct {
	Manifest  Manifest
	Installer Installer
	Stager    Stager
	Command   Command
	Log       *libbuildpack.Logger
}

func (s *Supplier) Run() error {
	s.Log.BeginStep("Supplying css")
	fmt.Println("Hi I am the supplier and I am being run!")

	dep := libbuildpack.Dependency{Name: "css", Version: "0.0.1"}
	if err := s.Installer.InstallDependency(dep, s.Stager.DepDir()); err != nil {
		return err
	}

	if err := s.Stager.WriteProfileD("symlink.bat", "mklink /j c:\\Users\\vcap\\app\\Content\\override c:\\Users\\vcap\\deps\\0"); err != nil {
		fmt.Printf("Couldn't write profile.d: %s", err.Error())
		return err
	}
	return nil
}
