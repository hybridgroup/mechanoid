package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

func buildProject(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		return fmt.Errorf("target board required")
	}

	targetName := cCtx.Args().First()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	projectname := filepath.Base(wd)

	s := spinner.New(spinner.CharSets[5], 100*time.Millisecond, spinner.WithWriter(os.Stdout))
	s.Suffix = " Building application for " + targetName + " using interpreter " + cCtx.String("interpreter")
	s.FinalMSG = "Application built.\n"
	s.Start()
	defer s.Stop()

	intp := cCtx.String("interpreter")
	if intp == "wasman" {
		intp = "wasman nowazero"
	}

	if cCtx.Bool("debug") {
		intp += " debug"
	}

	args := []string{"build", "-o", projectname + ".uf2", "-size", "short", "-stack-size", "8kb", "-tags", intp, "-target", targetName}

	if len(cCtx.StringSlice("params")) > 0 {
		ldlags := ""
		for _, p := range cCtx.StringSlice("params") {
			ldlags += " -X " + p
		}
		args = append(args, "-ldflags", ldlags)
	}

	args = append(args, ".")

	var cmd = exec.Command("tinygo", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(&spinWriter{s, os.Stdout, false}, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(&spinWriter{s, os.Stderr, false}, &stderrBuf)

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s: %v\n", cmd.String(), err)
		os.Exit(1)
	}

	return nil
}

func buildModules(cCtx *cli.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if stat, err := os.Stat(filepath.Join(wd, "modules")); err == nil && stat.IsDir() {
		// build modules?
		dirs, err := os.ReadDir(filepath.Join(wd, "modules"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, dir := range dirs {
			f, err := dir.Info()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if !f.IsDir() {
				continue
			}
			// check if there is a go.mod file in the directory
			_, err = os.Stat(filepath.Join(wd, "modules", f.Name(), "go.mod"))
			if err == nil {
				if err := buildGoModule(filepath.Join(wd, "modules"), f.Name()); err != nil {
					return err
				}
				continue
			}

			// check if there is a Cargo.toml file in the directory
			_, err = os.Stat(filepath.Join(wd, "modules", f.Name(), "Cargo.toml"))
			if err == nil {
				if err := buildRustModule(filepath.Join(wd, "modules"), f.Name()); err != nil {
					return err
				}
				continue
			}

			// check if there is a build.zig file in the directory
			_, err = os.Stat(filepath.Join(wd, "modules", f.Name(), "build.zig"))
			if err == nil {
				if err := buildZigModule(filepath.Join(wd, "modules"), f.Name()); err != nil {
					return err
				}
				continue
			}

			// no go.mod or Cargo.toml file found
			fmt.Println("No module files found in", f.Name(), "skipping...")
		}
	}

	return nil
}

func buildGoModule(modulesPath, name string) error {
	s := spinner.New(spinner.CharSets[5], 100*time.Millisecond, spinner.WithWriter(os.Stdout))
	s.Suffix = " Building TinyGo module " + name
	s.FinalMSG = "Done.\n"
	s.Start()
	defer s.Stop()

	fmt.Println("Building TinyGo module", name)
	modulePath := filepath.Join(modulesPath, name)
	os.Chdir(modulePath)
	defer os.Chdir("../..")

	cmd := exec.Command("tinygo", "build", "-o", filepath.Join(modulesPath, name+".wasm"), "-target", name+".json", "-no-debug", "-size", "short", ".")

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(&spinWriter{s, os.Stdout, false}, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(&spinWriter{s, os.Stderr, false}, &stderrBuf)

	if err := cmd.Run(); err != nil {
		fmt.Printf("tinygo build error %s: %v\n", modulePath, err)
		os.Exit(1)
	}

	return nil
}

func buildRustModule(modulesPath, name string) error {
	s := spinner.New(spinner.CharSets[5], 100*time.Millisecond, spinner.WithWriter(os.Stdout))
	s.Suffix = " Building Rust module " + name
	s.FinalMSG = "Done.\n"
	s.Start()
	defer s.Stop()

	fmt.Println("Building Rust module", name)
	modulePath := filepath.Join(modulesPath, name)
	os.Chdir(modulePath)
	defer os.Chdir("../..")

	cmd := exec.Command("cargo", "build", "--target", "wasm32-unknown-unknown", "--release")

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(&spinWriter{s, os.Stdout, false}, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(&spinWriter{s, os.Stderr, false}, &stderrBuf)

	if err := cmd.Run(); err != nil {
		fmt.Printf("cargo build error %s: %v\n", modulePath, err)
		os.Exit(1)
	}

	if err := copyFile(filepath.Join(modulePath, "target", "wasm32-unknown-unknown", "release", name+".wasm"), filepath.Join(modulesPath, name+".wasm")); err != nil {
		fmt.Printf("copy file error %s: %v\n", modulePath, err)
		os.Exit(1)
	}

	return nil
}

func buildZigModule(modulesPath, name string) error {
	s := spinner.New(spinner.CharSets[5], 100*time.Millisecond, spinner.WithWriter(os.Stdout))
	s.Suffix = " Building Zig module " + name
	s.FinalMSG = "Done.\n"
	s.Start()
	defer s.Stop()

	fmt.Println("Building Zig module", name)
	modulePath := filepath.Join(modulesPath, name)
	os.Chdir(modulePath)
	defer os.Chdir("../..")

	cmd := exec.Command("zig", "build")

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(&spinWriter{s, os.Stdout, false}, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(&spinWriter{s, os.Stderr, false}, &stderrBuf)

	if err := cmd.Run(); err != nil {
		fmt.Printf("zig build error %s: %v\n", modulePath, err)
		os.Exit(1)
	}

	if err := copyFile(filepath.Join(modulePath, "zig-out", "lib", name+".wasm"), filepath.Join(modulesPath, name+".wasm")); err != nil {
		fmt.Printf("copy file error %s: %v\n", modulePath, err)
		os.Exit(1)
	}

	return nil
}

// copyFile copies the given file or directory from src to dst. It can copy over
// a possibly already existing file (but not directory) at the destination.
func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	st, err := source.Stat()
	if err != nil {
		return err
	}

	if st.IsDir() {
		err := os.Mkdir(dst, st.Mode().Perm())
		if err != nil {
			return err
		}
		names, err := source.Readdirnames(0)
		if err != nil {
			return err
		}
		for _, name := range names {
			err := copyFile(filepath.Join(src, name), filepath.Join(dst, name))
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		destination, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, st.Mode())
		if err != nil {
			return err
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		return err
	}
}
