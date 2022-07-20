# Fairu

Maintain your work directory with Fairu.

Fairu is a tool to help you keep your file system ordered, based on a simple
system of rules.

**WARNING:** This is an experimental tool, so use it with great care, I do not
make myself responsible for any loss of data caused by the use of this
utility.

## Installation

Execute the follow command: `go install github.com/anderson-lt/fairu/cmd/fairu@latest`.

## Get Started

Copy the following to the program configuration path:

```yaml
Rules:
	"Show Big Files":
		Size: 50GB
		Print: Big file $ShellPath

	"Files With Extension":
		Glob: "*.*"
		Print: Has extension $Extension the file $ShellPath

	"Rest Of Files":
		Print: File $ShellPath
```

Run the program without arguments the first time to see the configuration file
path.

Then execute the `fairu` command to see the program in action.

## State

This program is at a very early stage of development and therefore it is very
unstable.

Although most of the characteristics are not implemented; You can get more
information about the project, reading the documentation (in Spanish) of the
program at: `./Documentation/draft/Configuration.adoc`.
