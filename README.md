# sepamin
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub license](https://img.shields.io/github/license/FerdinaKusumah/sepamin.svg)](https://github.com/FerdinaKusumah/sepamin/blob/master/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/FerdinaKusumah/sepamin.svg)](https://GitHub.com/FerdinaKusumah/sepamin/issues/)
[![GitHub issues-closed](https://img.shields.io/github/issues-closed/FerdinaKusumah/sepamin.svg)](https://GitHub.com/FerdinaKusumah/sepamin/issues?q=is%3Aissue+is%3Aclosed)
[![GitHub pull-requests](https://img.shields.io/github/issues-pr/Naereen/StrapDown.js.svg)](https://GitHub.com/Naereen/StrapDown.js/pull/)

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Resource</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- GETTING STARTED -->
## Getting Started

This is a new variant of email spam, enjoy email spam to multiple targets at once.

### Prerequisites

Need go `1.12+` or later.

## Installation

### from Source

If you have go1.15+ compiler installed and configured:

```bash
▶ GO111MODULE=on go get github.com/FerdinaKusumah/sepamin
```

### from GitHub

```bash
▶ git clone https://github.com/FerdinaKusumah/sepamin
▶ cd sepamin
▶ go build .
▶ (sudo) mv sepamin /usr/local/bin
```


<!-- USAGE EXAMPLES -->
## Usage

| **Flag**          	| **Description**                                                 	                                |
|-------------------	|-----------------------------------------------------------------------------------------------	|
| -u, --user         	| Email account                 	                                                                |
| -f, --from         	| Sender user                 	                                                                    |
| -pwd, --password      | Password account                 	                                                            |
| -t, --target         	| Email target ex: hahahehe@gmail.com to multiple sender ex: hahahehe@gmail.com, cobacoba@gmail.com|
| -c, --count         	| How much you wanna spam email                 	                                                |
| -h, --host         	| SMTP Host, default use google smtp host                 	                                        |
| -p, --port         	| SMTP Port, default use google smtp port                 	                                        |
| -v, --verbose         | Display debug message                 	                                                        |
| -h, --help        	| Display its helps                                               	                                |


### Examples

#### Single Target

```bash
▶ sepamin -user="example@gmail.com" -pwd "cobatebakhayoo" -t "target@gmail.com" -c 2 -v true
```

#### Multiple Target

```bash
▶ sepamin -user="example@gmail.com" -pwd "cobatebakhayoo" -t "target1@gmail.com, target2@gmail.com, target3@gmail.com" -c 2
```

### Library

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/FerdinaKusumah/sepamin)

## TODOs

- [ ] Provide more email
- [ ] Sent email with image body

## Help & Bugs

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-blue.svg)](https://github.com/FerdinaKusumah/sepamin/issues)

If you are still confused or found a bug, please [open the issue](https://github.com/FerdinaKusumah/sepamin/issues). All bug reports are appreciated, some features have not been tested yet due to lack of free time.

## License

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

**sepamin** released under MIT. See `LICENSE` for more details.
