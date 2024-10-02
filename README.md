# gomod
gomod upgrades all go modules that you specifies

## Installation
```shell
go install github.com/sivchari/gomod@latest
```

## Usage

### Write your configuration file
```yaml
go_version: "1.23"
module_paths:
  - "$HOME/path/module1"
  - "$HOME/path/module2"
```

### Run gomod
```shell
gomod run
```

  
