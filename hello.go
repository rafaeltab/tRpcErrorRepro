package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("Hello, world.")
	bytes, err := getFixture("pnpm-lock.yaml")
	if err != nil {
		panic(err)
	}

	_, err2 := DecodePnpmLockfile(bytes)
	if err2 != nil {
		panic(err2)
	}
}

func getFixture(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func DecodePnpmLockfile(contents []byte) (*PnpmLockfile, error) {
	var lockfile PnpmLockfile
	if err := yaml.Unmarshal(contents, &lockfile); err != nil {
		panic(err)
	}

	return &lockfile, nil
}

type PnpmLockfile struct {
	Version                   float32                    `yaml:"lockfileVersion"`
	NeverBuiltDependencies    []string                   `yaml:"neverBuiltDependencies,omitempty"`
	OnlyBuiltDependencies     []string                   `yaml:"onlyBuiltDependencies,omitempty"`
	Overrides                 map[string]string          `yaml:"overrides,omitempty"`
	PackageExtensionsChecksum string                     `yaml:"packageExtensionsChecksum,omitempty"`
	PatchedDependencies       map[string]PatchFile       `yaml:"patchedDependencies,omitempty"`
	Importers                 map[string]ProjectSnapshot `yaml:"importers"`
	Packages                  map[string]PackageSnapshot `yaml:"packages,omitempty"`
	Time                      map[string]string          `yaml:"time,omitempty"`
}

// ProjectSnapshot Snapshot used to represent projects in the importers section
type ProjectSnapshot struct {
	Specifiers           map[string]string           `yaml:"specifiers"`
	Dependencies         map[string]string           `yaml:"dependencies,omitempty"`
	OptionalDependencies map[string]string           `yaml:"optionalDependencies,omitempty"`
	DevDependencies      map[string]string           `yaml:"devDependencies,omitempty"`
	DependenciesMeta     map[string]DependenciesMeta `yaml:"dependenciesMeta,omitempty"`
	PublishDirectory     string                      `yaml:"publishDirectory,omitempty"`
}

// PackageSnapshot Snapshot used to represent a package in the packages setion
type PackageSnapshot struct {
	Resolution PackageResolution `yaml:"resolution,flow"`
	ID         string            `yaml:"id,omitempty"`

	// only needed for packages that aren't in npm
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`

	Engines struct {
		Node string `yaml:"node"`
		NPM  string `yaml:"npm,omitempty"`
	} `yaml:"engines,omitempty,flow"`
	CPU  []string `yaml:"cpu,omitempty,flow"`
	Os   []string `yaml:"os,omitempty,flow"`
	LibC []string `yaml:"libc,omitempty"`

	Deprecated    string `yaml:"deprecated,omitempty"`
	HasBin        bool   `yaml:"hasBin,omitempty"`
	Prepare       bool   `yaml:"prepare,omitempty"`
	RequiresBuild bool   `yaml:"requiresBuild,omitempty"`

	BundledDependencies  []string          `yaml:"bundledDependencies,omitempty"`
	PeerDependencies     map[string]string `yaml:"peerDependencies,omitempty"`
	PeerDependenciesMeta map[string]struct {
		Optional bool `yaml:"optional"`
	} `yaml:"peerDependenciesMeta,omitempty"`

	Dependencies         map[string]string `yaml:"dependencies,omitempty"`
	OptionalDependencies map[string]string `yaml:"optionalDependencies,omitempty"`

	TransitivePeerDependencies []string `yaml:"transitivePeerDependencies,omitempty"`
	Dev                        bool     `yaml:"dev"`
	Optional                   bool     `yaml:"optional,omitempty"`
	Patched                    bool     `yaml:"patched,omitempty"`
}

// PackageResolution Various resolution strategies for packages
type PackageResolution struct {
	Type string `yaml:"type,omitempty"`
	// For npm or tarball
	Integrity string `yaml:"integrity,omitempty"`

	// For tarball
	Tarball string `yaml:"tarball,omitempty"`

	// For local directory
	Dir string `yaml:"directory,omitempty"`

	// For git repo
	Repo   string `yaml:"repo,omitempty"`
	Commit string `yaml:"commit,omitempty"`
}

// PatchFile represent a patch applied to a package
type PatchFile struct {
	Path string `yaml:"path"`
	Hash string `yaml:"hash"`
}

type DependenciesMeta struct {
	Injected bool   `yaml:"injected,omitempty"`
	Node     string `yaml:"node,omitempty"`
	Patch    string `yaml:"patch,omitempty"`
}
