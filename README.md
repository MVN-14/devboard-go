# Devboard
Track dev projects from the terminal with ease with this _customizable_ dev-dashboard cli

![LangBadge](https://img.shields.io/github/languages/top/MVN-14/devboard-go?style=for-the-badge)


## Demo
This is a rawdog example of interacting with the cli, although it was really written to allow for easy ui implementations or scripted automation.

![demo](https://vhs.charm.sh/vhs-5xd22qvwJz5fyC0hNWTvgO.gif)

## Installation

``` bash
git clone https://github.com/MVN-14/devboard-go
cd devboard-go
go install devboard.go
```

## Usage
### Flags and Environment Variables
Devboard allows for different flags to customize program behaviour.<br>
Devboard evaluates variables in the following order and stops when one exists:<br>
1. **Command line flag. - eg. (--dbpath=test.db)**
2. **Environment variable. - eg. (DEVBOARD_DBPATH=test.db)**
3. **Config value. - eg. (`dbpath: test.db` in config.yaml)**

|Cmd|Flag(s), Variable|Default|Decription|
|---|--------------|-------|----------|
|**all**|`--verbose,-v, DEVBOARD_VERBOSE`|false|Verbose output|
|**all**|`--dbpath, DEVBOARD_DBPATH`|$HOME/.devboard/devboard.db|Path to **_existing directory_** to open or create db|
|**all**|`-c,--config DEVBOARD_CONFIG`|$HOME/.devboard/config.yaml|Config file location (default is $HOME/.devboard/config.yaml)<br>**Config file must be yaml format** [See config](#config)
|**list**|`-d,--deleted DEVBOARD_DELETED`|false|Show deleted projects in list|
|**open**|`--command DEVBOARD_COMMAND`| `echo "no command present in project, environment, or config"`|Command to execute to open project if command field is not set for specified project|


### Add Project
```
devboard add <project>
```
[see project arg](#project)

### Update Project
```
devboard update <project>
```
[see project arg](#project)

### Open Project
```
devboard open <id>
```
[see id arg](#id)


### Delete Project
```
devboard delete <id>
```
[see id arg](#id)

### List Projects
```
devboard list
```

## Args
### Project
|Field|Type|Required|Description|
|-|-|-|-|
|command|string|No|Command to open the project|
|id|int|_only on update_|The ID of the project to update|
|name|string|Yes|Project name|
|path|string|Yes|Path to project directory|

### Id  
Positive integer represending the ID field of the project being updated

![LangBadge](https://img.shields.io/github/languages/top/MVN-14/devboard-go?style=for-the-badge)


## Demo
This is a rawdog example of interacting with the cli, although it was really written to allow for easy ui implementations or scripted automation.

![demo](https://vhs.charm.sh/vhs-5xd22qvwJz5fyC0hNWTvgO.gif)

## Installation

``` bash
git clone https://github.com/MVN-14/devboard-go
cd devboard-go
go install devboard.go
```

## Usage
### Flags and Environment Variables
Devboard allows for different flags to customize program behaviour.<br>
Devboard evaluates variables in the following order and stops when one exists:<br>
1. **Command line flag. - eg. (--dbpath=test.db)**
2. **Environment variable. - eg. (DEVBOARD_DBPATH=test.db)**
3. **Config value. - eg. (`dbpath: test.db` in config.yaml)**

|Cmd|Flag(s), Variable|Default|Decription|
|---|--------------|-------|----------|
|**all**|`--verbose,-v, DEVBOARD_VERBOSE`|false|Verbose output|
|**all**|`--dbpath, DEVBOARD_DBPATH`|$HOME/.devboard/devboard.db|Path to **_existing directory_** to open or create db|
|**all**|`-c,--config DEVBOARD_CONFIG`|$HOME/.devboard/config.yaml|Config file location (default is $HOME/.devboard/config.yaml)<br>**Config file must be yaml format** [See config](#config)
|**list**|`-d,--deleted DEVBOARD_DELETED`|false|Show deleted projects in list|
|**open**|`--command DEVBOARD_COMMAND`| `echo "no command present in project, environment, or config"`|Command to execute to open project if command field is not set for specified project|


### Add Project
```
devboard add <project>
```
[see project arg](#project)

### Update Project
```
devboard update <project>
```
[see project arg](#project)

### Open Project
```
devboard open <id>
```
[see id arg](#id)


### Delete Project
```
devboard delete <id>
```
[see id arg](#id)

### List Projects
```
devboard list
```

## Args
|Field|Type|Required|Description|
|-|-|-|-|
|command|string|No|Command to open the project|
|id|int|_only on update_|The ID of the project to update|
|name|string|Yes|Project name|
|path|string|Yes|Path to project directory|

Positive integer represending the ID field of the project being updated
