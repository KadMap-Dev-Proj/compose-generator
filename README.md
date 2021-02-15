# Compose Generator
![Build passing](https://github.com/marcauberer/compose-generator/workflows/Go%20CI/badge.svg)
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/marcauberer/compose-generator?include_prereleases)
[![Go Report](https://goreportcard.com/badge/github.com/marcauberer/compose-generator)](https://goreportcard.com/report/github.com/marcauberer/compose-generator)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

## Usage
You can use the Compose Generator CLI by directly installing it on your Docker host system or by generating your compose file with the Compose Generator Docker container.

### Install Compose Generator CLI
**Linux (Debian based distributions)**<br>
```sh
$ sudo apt-get update
$ sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common lsb-release
$ sudo add-apt-repository "deb [arch=amd64] https://repo.chillibits.com/artifactory/compose-generator/linux/$(lsb_release -is) stable"
$ sudo apt-get install compose-generator
```

**Linux (RPM based distributions)**<br>
```sh
$ sudo apt-get update
$
$ sudo apt-get install compose-generator
```

**Windows**
Compose Generator gets distributed for Windows via the new Windows package manager called [winget](https://github.com/microsoft/winget-cli). In the future, winget will be available for download in the Microsoft Store. Currently, the easiest way to install winget is, to download it manually from GitHub. Visit the [installation instruction](https://github.com/microsoft/winget-cli#installing-the-client) from Microsoft.

As soon as the Windows package manager is installed on your Windows machine, you can open powershell and execute this installation command:

```sh
$ winget install ChilliBits.ComposeGenerator
```

After installing Compose Generator, you should restart your powershell instance to make it reload the available commands.

### Generate compose file on the fly with Docker container
*Note: This command does not work with Windows CMD command line. Please use PowerShell.*

```sh
$ docker run --rm -it -v ${pwd}:/compose-generator/out chillibits/compose-generator
```

## Supported host systems
There are also downloadable packages available for all supported platforms:

| **Platform**                | **x86_64 / amd64**                                                                     | **i386**                                                                             | **armv5**                                                                              | **armv6**                                                                              | **armv7**                                                                              | **arm64**                                                                              |
|-----------------------------|----------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------|
| **Darwin / MacOS (tar.gz)** | [download](../../releases/download/0.1.0/compose-generator_0.1.0_darwin_amd64.tar.gz)  | -                                                                                    | -                                                                                      | -                                                                                      | -                                                                                      | -                                                                                      |
| **FreeBSD (tag.gz)**        | [download](../../releases/download/0.1.0/compose-generator_0.1.0_freebsd_amd64.tar.gz) | [download](../../releases/download/0.1.0/compose-generator_0.1.0_freebsd_386.tar.gz) | [download](../../releases/download/0.1.0/compose-generator_0.1.0_freebsd_armv5.tar.gz) | [download](../../releases/download/0.1.0/compose-generator_0.1.0_freebsd_armv6.tar.gz) | [download](../../releases/download/0.1.0/compose-generator_0.1.0_freebsd_armv7.tar.gz) | [download](../../releases/download/0.1.0/compose-generator_0.1.0_freebsd_arm64.tar.gz) |
| **Alpine (apk)**            | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_amd64.apk)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_386.apk)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv5.apk)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv6.apk)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv7.apk)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_arm64.apk)      |
| **CentOS (rpm)**            | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_amd64.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_386.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv5.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv6.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv7.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_arm64.rpm)      |
| **Debian (deb)**            | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_amd64.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_386.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv5.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv6.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv7.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_arm64.deb)      |
| **Fedora (rpm)**            | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_amd64.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_386.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv5.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv6.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv7.rpm)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_arm64.rpm)      |
| **Raspbian (deb)**          | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_amd64.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_386.deb)      | -                                                                                      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv6.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv7.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_arm64.deb)      |
| **Ubuntu (deb)**            | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_amd64.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_386.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv5.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv6.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_armv7.deb)      | [download](../../releases/download/0.1.0/compose-generator_0.1.0_linux_arm64.deb)      |
| **Windows Installer (exe)** | [download](../../releases/download/0.1.0/ComposeGenerator_0.1.0_x64_Setup.exe)         | [download](../../releases/download/0.1.0/ComposeGenerator_0.1.0_x86_Setup.exe)       | -                                                                                      | -                                                                                      | -                                                                                      | -                                                                                      |
| **Windows Portable (zip)**  | [download](../../releases/download/0.1.0/compose-generator_0.1.0_windows_amd64.zip)    | [download](../../releases/download/0.1.0/compose-generator_0.1.0_windows_386.zip)    | -                                                                                      | -                                                                                      | -                                                                                      | -                                                                                      |

## Contribute to the project
If you want to contribute to this project, please feel free to send us a pull request.

© Marc Auberer 2021