# install

You may install `d2` through any of the following methods.

<!-- toc -->
- <a href="#installsh" id="toc-installsh">install.sh</a>
  - <a href="#security" id="toc-security">Security</a>
- <a href="#macos-homebrew" id="toc-macos-homebrew">macOS (Homebrew)</a>
- <a href="#standalone" id="toc-standalone">Standalone</a>
  - <a href="#manual" id="toc-manual">Manual</a>
  - <a href="#prefix" id="toc-prefix">PREFIX</a>
- <a href="#from-source" id="toc-from-source">From source</a>
  - <a href="#source-release" id="toc-source-release">Source Release</a>
- <a href="#windows" id="toc-windows">Windows</a>
  - <a href="#msys2" id="toc-msys2">MSYS2</a>
  - <a href="#wsl" id="toc-wsl">WSL</a>
- <a href="#coming-soon" id="toc-coming-soon">Coming soon</a>

## install.sh

The recommended and easiest way to install is with our install script, which will detect
the OS and architecture you're on and use the best method:

```sh
# With --dry-run the install script will print the commands it will use
# to install without actually installing so you know what it's going to do.
curl -fsSL https://d2lang.com/install.sh | sh -s -- --dry-run
# If things look good, install for real.
curl -fsSL https://d2lang.com/install.sh | sh -s --
```

For help on the terminal run, including the supported package managers and detection
methods:

```sh
curl -fsSL https://d2lang.com/install.sh | sh -s -- --help
```

### Security

The install script is not the most secure way to install d2. We recommend that if
possible, you use your OS's package manager directly or install from source with `go` as
described below.

But this does not mean the install script is insecure. There is no major flaw that
the install script is more vulnerable to than any other method of manual installation.
The most secure installation method involves a second independent entity, i.e your OS
package repos or Go's proxy server.

We're careful shell programmers and are aware of the many footguns of the Unix shell. Our
script was written carefully and with detail. For example, it is not vulnerable to partial
execution and the entire script runs with `set -eu` and very meticulous quoting.

It follows the XDG standards, installs `d2` properly into a Unix hierarchy path
(`/usr/local` unless `/usr/local` requires sudo in which case `~/.local` is used) and
allows for easy uninstall. You can easily adjust the used path with `--prefix`.

Some other niceties are that it'll tell you if you need to adjust `$PATH` or `$MANPATH` to
access `d2` and its manpages. It can also install
[TALA](https://github.com/terrastruct/tala) for you with `--tala`. You can also use it to
install a specific version of `d2` with `--version`. Run it with `--help` for more more
detailed docs on its various options and features.

If you're still concerned, remember you can run with `--dry-run` to avoid writing anything.

The install script does not yet verify any signature on the downloaded release
but that is coming soon. [#315](https://github.com/terrastruct/d2/issues/315)

## macOS (Homebrew)

If you're on macOS, you can install with `brew`.

```sh
brew tap terrastruct/d2
brew install d2
```

> The install script above does this automatically if you have `brew` installed and
> are running it on macOS.

## Standalone

We publish standalone release archives for every release on Github.

Here's a minimal example script that downloads a standalone release, extracts it into the
current directory and then installs it.
Adjust VERSION, OS, and ARCH as needed.

```sh
VERSION=v0.0.13 OS=macos ARCH=amd64 curl -fsSLO \
    "https://github.com/terrastruct/d2/releases/download/$VERSION/d2-$VERSION-$OS-$ARCH.tar.gz" \
    && tar -xzf "d2-$VERSION-$OS-$ARCH.tar.gz" \
    && make -sC "d2-$VERSION" install
```

To uninstall:

```sh
VERSION=v0.0.13 make -sC "d2-$VERSION" uninstall
```

### Manual

You can also manually download the `.tar.gz` release for your OS/ARCH combination and then
run the following inside the extracted directory to install:

```sh
make install
```

Run the following to uninstall:

```sh
make uninstall
```

### PREFIX

You can control the Unix hierarchy installation path with `PREFIX=`. For example:

```sh
# Install under ~/.local.
# Binaries will be at ~/.local/bin
# And manpages will be under ~/.local/share/man
# And supporting data like icons and fonts at ~/.local/share/d2
make install PREFIX=$HOME/.local
```

The install script places the standalone release into `$PREFIX/lib/d2/d2-<version>`
and we recommend doing the same with manually installed releases so that you
know where the release directory is for easy uninstall.

## From source

You can always install from source:

```sh
go install oss.terrastruct.com/d2@latest
```

### Source Release

To install a proper release from source clone the repository and then:

```sh
./ci/release/build.sh --install
# To uninstall:
# ./ci/release/build.sh --uninstall
```

Installing a real release will also install manpages and in the future other assets like
fonts and icons. Furthermore, when installing a non versioned commit, installing a proper
release will ensure that `d2 --version` works correctly by embedding the commit hash into
the `d2` binary.

## Windows

d2 builds and runs on Windows:

We have prebuilt standalone releases for Windows though they're structured in the same way
as our Unix releases. So after extracting a release, you'll have to manually put the d2
binary into your `$PATH` or add the `bin` directory of the release into your `$PATH`.

See https://www.wikihow.com/Change-the-PATH-Environment-Variable-on-Windows

Then you'll be able to call `d2` from the commandline in `cmd.exe` or `pwsh.exe`.

We intend to have a `.msi` release installer sometime soon that handles putting `d2` into
your `$PATH` for you.

### MSYS2

<img width="1680" alt="Screenshot 2022-12-06 at 2 55 27 AM" src="https://user-images.githubusercontent.com/10180857/205892927-6f3e116c-1c4a-440a-9972-82c306aa9779.png">

We recommend using [MSYS2](https://www.msys2.org/) or [Git
Bash](https://gitforwindows.org/#bash) (Git Bash is based on MSYS2) for an improved
terminal experience.

MSYS2 provides a unix style shell environment that is native to Windows (unlike
[Cygwin](https://www.cygwin.com/)). MSYS2 allows `install.sh` to work, enables proper
installation of our standalone releases via `make install` and makes the manpage
accessible via `man d2`.

The MSYS2 terminal also enables `d2` to display colors like in the above screenshot.

In addition, all of our development and CI scripts work under MSYS2 whereas they do not
under plain Windows.

### WSL

`d2` works perfectly under [WSL](https://learn.microsoft.com/en-us/windows/wsl/install)
aka Windows Subsystem for Linux if that's what you prefer. Installation is just like any
other Linux system.

## Coming soon

- Docker image
- rpm and deb packages
    - with repositories and standalone
- homebrew core
