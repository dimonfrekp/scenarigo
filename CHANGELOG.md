# CHANGELOG

<a name="v0.16.2"></a>
## [v0.16.2] - 2023-11-02
### Bug Fixes
- **template:** don't convert int to string

<a name="v0.16.1"></a>
## [v0.16.1] - 2023-11-01
### Features
- **grpc:** dump invalid utf8 strings as hex

<a name="v0.16.0"></a>
## [v0.16.0] - 2023-10-24
### Bug Fixes
- don't panic if the protocol is empty

### Features
- add continueOnError to prevent failure due to step errors
- add if field for controlling step running
- allow to access results of each step
- enable to assert by template string expressions
- cancel request contexts after each step
- **assert:** enable to pass custom equalers
- **config:** add global variables
- **template:** add size() function
- **template:** allow to call values having Call method as a function
- **template:** allow '$' identifier

### BREAKING CHANGE

assert.Build function requires a context.Context value as an argument to avoid a goroutine leak.

<a name="v0.15.1"></a>
## [v0.15.1] - 2023-09-15
### Code Refactoring
- remove workaround

<a name="v0.15.0"></a>
## [v0.15.0] - 2023-09-06
### Bug Fixes
- **plugin:** setup plugins in the order in which they are registered
- **schema:** add a workaround to avoid failing to load scenarios
- **template:** check overflow
- **template:** evaluate only an expression that matched the condition

### Code Refactoring
- add OrderedMap

### Features
- add dump sub-command
- add ytt integration
- add input config
- **grpc:** contain response status in log
- **template:** add time and duration type
- **template:** add bytes type

<a name="v0.14.2"></a>
## [v0.14.2] - 2023-03-03
- bump up the version of dependent modules

<a name="v0.14.1"></a>
## [v0.14.1] - 2023-02-27
### Features
- **schema:** add Comments field

<a name="v0.14.0"></a>
## [v0.14.0] - 2023-02-20
### Bug Fixes
- pass bound variables to the next step
- fix to filter correctly even if / is included in subtest names
- filter test by -run flag of go test
- **plugin:** make RegisterSetup() not cause an error if called in tests

### Code Refactoring
- **reporter:** change FromT implementation

### Features
- change retry unit from request to entire step
- **http:** add Accept-Encoding header by default
- **http:** enable decoding of response bodies with character encodings other than utf-8
- **http:** add text/html unmarshaler

<a name="v0.13.2"></a>
## [v0.13.2] - 2022-12-16
- bump up the version of dependent modules

<a name="v0.13.1"></a>
## [v0.13.1] - 2022-12-15
- bump up the version of dependent modules

<a name="v0.13.0"></a>
## [v0.13.0] - 2022-12-08
### Bug Fixes
- enable to specify report paths by absolute path
- fix generate CREDITS workflow

### Features
- enable to read config from stdin
- enable to marshal schema.Config to YAML
- **errors:** change error message format

<a name="v0.12.8"></a>
## [v0.12.8] - 2022-10-18
### Bug Fixes
- don't bind vars if included scenario failed

<a name="v0.12.7"></a>
## [v0.12.7] - 2022-09-27
### Features
- **template:** enable to call methods

<a name="v0.12.6"></a>
## [v0.12.6] - 2022-09-13
### Features
- enable to specify step timeout
- **grpc:** enable to use template in error details
- **http:** make method name case-insensitive

<a name="v0.12.5"></a>
## [v0.12.5] - 2022-08-22
### Bug Fixes
- **plugin:** go mod tidy with -compat option

<a name="v0.12.4"></a>
## [v0.12.4] - 2022-07-25
### Bug Fixes
- **plugin:** enable to replace modules to local paths
- **plugin:** keep replace directives

<a name="v0.12.3"></a>
## [v0.12.3] - 2022-07-21
### Bug Fixes
- **plugin:** remove plugin modules from the cache
- **plugin:** check remote module source versions

<a name="v0.12.2"></a>
## [v0.12.2] - 2022-07-20
### Bug Fixes
- **mock:** fix nil error bug
- **plugin:** force all plugins to use the same version of package

### Code Refactoring
- fix maintidx error
- fix cyclop error

<a name="v0.12.1"></a>
## [v0.12.1] - 2022-06-26
### Bug Fixes
- **release:** reduce target Go versions

<a name="v0.12.0"></a>
## [v0.12.0] - 2022-06-13
### Bug Fixes
- **plugin:** suppress unnecessary plugin build logs
- **plugin:** don't use "main" as module name

### Features
- **template:** allow functions to return an error

<a name="v0.11.2"></a>
## [v0.11.2] - 2022-04-26
### Bug Fixes
- **plugin:** allow specifying sub directories of remote modules as src

<a name="v0.11.1"></a>
## [v0.11.1] - 2022-04-18
### Bug Fixes
- print error if fail to open plugin
- **doc:** setup field was deprecated

<a name="v0.11.0"></a>
## [v0.11.0] - 2022-04-15
### Bug Fixes
- **plugin:** fix issue with plugin build failure in Go1.18

### Features
- enable to marshal scenarios into YAML
- **mock:** enable to assert request
- **template:** allow writing left arrow function call in map syntax
- **template:** enable to use template in map keys
- **template:** enable to escape { by \

<a name="v0.10.0"></a>
## [v0.10.0] - 2022-01-31
### Bug Fixes
- update the go directive of go.mod
- **plugin:** use the same module version as scenarigo for building plugins

### BREAKING CHANGE

This package requires Go 1.17 or later.

<a name="v0.9.0"></a>
## [v0.9.0] - 2021-12-03
### Bug Fixes
- **errors:** Errors returns nil if no errors

### Code Refactoring
- use yaml.PathBuilder to specify the pos

### Features
- add setup feature
- add "scenarigo plugin list" command
- add "scenarigo config validate" command
- add plugin sub-command
- **plugin:** enable registration of setup functions to be executed for each scenario
- **plugin:** enable to build plugin from remote "go gettable" src
- **template:** add bool literals

<a name="v0.8.1"></a>
## [v0.8.1] - 2021-09-27
### Bug Fixes
- add workaround to avoid the bug of Go 1.17

### Code Refactoring
- export functions

### Features
- list command refers to the configuration file
- remove blank lines from logs

### BREAKING CHANGE

"file" and "verbose" options are removed from the list sub-command.

<a name="v0.8.0"></a>
## [v0.8.0] - 2021-09-08
### Bug Fixes
- enable CGO on release build
- **query:** do not extract by the inline field name
- **template:** fix a bug by nil struct field
- **template:** marshal variables to YAML in LAF arguments
- **template:** keep the original memory address
- **template:** marshal LAF arguments with indent

### Features
- enable cross compile with CGO
- **grpc:** loose type checking for equaler
- **template:** execute templates of data
- **version:** get version from build info

<a name="v0.7.0"></a>
## [v0.7.0] - 2021-07-30
### Bug Fixes
- **assert:** fix the assertion operators
- **assert:** fix the logic to compare Go protobuf APIv2 messages
- **grpc:** rename body field to message
- **query:** don't access unexported field

### Code Refactoring
- don't use ioutil package

### Features
- change default configuration filename
- enable to set configurations by file
- add WithConfig option
- colorize outputs
- support NO_COLOR standard
- enable strictly check on request field
- use Go protobuf APIv2
- **assert:** enable to change the behavior of equal assertion
- **query:** allow accessing anonymous fields

### Performance Improvements
- reuse parsed AST node to print error tokens

### BREAKING CHANGE

This package requires Go 1.16 or later.

<a name="v0.6.3"></a>
## [v0.6.3] - 2021-04-08
### Bug Fixes
- enable to bind vars defined in the included scenario

<a name="v0.6.2"></a>
## [v0.6.2] - 2021-04-07
### Bug Fixes
- **plugin:** avoid the error caused by loading plugins concurrently ([#78](https://github.com/zoncoen/scenarigo/issues/78))

### Code Refactoring
- **assert:** remove query from arguments

### Features
- **assert:** add length assertion
- **assert:** add greaterThan/greaterThanOrEqual/lessThan/lessThanOrEqual ([#77](https://github.com/zoncoen/scenarigo/issues/77))
- **reporter:** enable to generate test report ([#83](https://github.com/zoncoen/scenarigo/issues/83))
- **reporter:** include the execution time of sub-tests ([#82](https://github.com/zoncoen/scenarigo/issues/82))

<a name="v0.6.1"></a>
## [v0.6.1] - 2021-01-14
### Bug Fixes
- **template:** don't convert invalid values to avoid panic

<a name="v0.6.0"></a>
## [v0.6.0] - 2021-01-12
### Bug Fixes
- **template:** enable to set to pointer values

### Features
- export RunScenario function
- add WithScenariosFromReader option
- allow template in header assertion
- **assert:** add regexp function
- **context:** add ScenarioFilePath

<a name="v0.5.1"></a>
## [v0.5.1] - 2020-10-23
### Bug Fixes
- **template:** restore funcs in args of left arrow function

### Features
- **assert:** add "and" function

<a name="v0.5.0"></a>
## [v0.5.0] - 2020-10-05
### Features
- **assert:** add "or" function
- **expect:** enable strict option when decoding yaml for expect to prevent field misplacement ([#59](https://github.com/zoncoen/scenarigo/issues/59))
- **grpc:** allow using a template as code and msg
- **http:** allow using a template as code

<a name="v0.4.0"></a>
## [v0.4.0] - 2020-09-02
### Bug Fixes
- register errdetails proto messages to unmarshal Any
- **expect:** use the default assertion if no expect ([#55](https://github.com/zoncoen/scenarigo/issues/55))
- **template:** avoid to panic ([#54](https://github.com/zoncoen/scenarigo/issues/54))

### Features
- **cmd:** add list sub-command ([#51](https://github.com/zoncoen/scenarigo/issues/51))

<a name="v0.3.3"></a>
## [v0.3.3] - 2020-06-17
### Bug Fixes
- **core:** add generated files to avoid the import error ([#41](https://github.com/zoncoen/scenarigo/issues/41))
- **deps:** update YAML library ( v1.7.12 => v1.7.15 ) ([#47](https://github.com/zoncoen/scenarigo/issues/47))
- **deps:** update YAML library ( v1.7.10 => v1.7.11 ) ([#42](https://github.com/zoncoen/scenarigo/issues/42))
- **deps:** update YAML library to fix a bug ( v1.7.9 => v1.7.10 ) ([#40](https://github.com/zoncoen/scenarigo/issues/40))
- **template:** fix processing for variadic arguments of function ([#48](https://github.com/zoncoen/scenarigo/issues/48))

<a name="v0.3.2"></a>
## [v0.3.2] - 2020-06-15
### Bug Fixes
- **deps:** update YAML library to fix a bug ( v1.7.8 => v1.7.9 ) ([#39](https://github.com/zoncoen/scenarigo/issues/39))

<a name="v0.3.1"></a>
## [v0.3.1] - 2020-06-12
### Bug Fixes
- **core:** fix ctx.Response() for http protocol ([#35](https://github.com/zoncoen/scenarigo/issues/35))
- **errors:** fix incorrect line number in YAML source ([#38](https://github.com/zoncoen/scenarigo/issues/38))

<a name="v0.3.0"></a>
## [v0.3.0] - 2020-06-11
### Features
- **core:** support to output error with YAML ([#33](https://github.com/zoncoen/scenarigo/issues/33))

<a name="v0.2.0"></a>
## [v0.2.0] - 2020-06-03
### Code Refactoring
- **core:** replace YAML libraries to goccy/go-yaml ([#31](https://github.com/zoncoen/scenarigo/issues/31))

### Features
- **core:** read YAML files only as scenarios ([#28](https://github.com/zoncoen/scenarigo/issues/28))
- **grpc:** enable to check header/trailer metadata of gRPC response ([#29](https://github.com/zoncoen/scenarigo/issues/29))
- **http:** enable to check HTTP response headers ([#30](https://github.com/zoncoen/scenarigo/issues/30))

### BREAKING CHANGE

change protocl.Protocol interface

<a name="v0.1.0"></a>
## v0.1.0 - 2020-05-17
- first release


[v0.16.2]: https://github.com/zoncoen/scenarigo/compare/v0.16.1...v0.16.2
[v0.16.1]: https://github.com/zoncoen/scenarigo/compare/v0.16.0...v0.16.1
[v0.16.0]: https://github.com/zoncoen/scenarigo/compare/v0.15.1...v0.16.0
[v0.15.1]: https://github.com/zoncoen/scenarigo/compare/v0.15.0...v0.15.1
[v0.15.0]: https://github.com/zoncoen/scenarigo/compare/v0.14.2...v0.15.0
[v0.14.2]: https://github.com/zoncoen/scenarigo/compare/v0.14.1...v0.14.2
[v0.14.1]: https://github.com/zoncoen/scenarigo/compare/v0.14.0...v0.14.1
[v0.14.0]: https://github.com/zoncoen/scenarigo/compare/v0.13.2...v0.14.0
[v0.13.2]: https://github.com/zoncoen/scenarigo/compare/v0.13.1...v0.13.2
[v0.13.1]: https://github.com/zoncoen/scenarigo/compare/v0.13.0...v0.13.1
[v0.13.0]: https://github.com/zoncoen/scenarigo/compare/v0.12.8...v0.13.0
[v0.12.8]: https://github.com/zoncoen/scenarigo/compare/v0.12.7...v0.12.8
[v0.12.7]: https://github.com/zoncoen/scenarigo/compare/v0.12.6...v0.12.7
[v0.12.6]: https://github.com/zoncoen/scenarigo/compare/v0.12.5...v0.12.6
[v0.12.5]: https://github.com/zoncoen/scenarigo/compare/v0.12.4...v0.12.5
[v0.12.4]: https://github.com/zoncoen/scenarigo/compare/v0.12.3...v0.12.4
[v0.12.3]: https://github.com/zoncoen/scenarigo/compare/v0.12.2...v0.12.3
[v0.12.2]: https://github.com/zoncoen/scenarigo/compare/v0.12.1...v0.12.2
[v0.12.1]: https://github.com/zoncoen/scenarigo/compare/v0.12.0...v0.12.1
[v0.12.0]: https://github.com/zoncoen/scenarigo/compare/v0.11.2...v0.12.0
[v0.11.2]: https://github.com/zoncoen/scenarigo/compare/v0.11.1...v0.11.2
[v0.11.1]: https://github.com/zoncoen/scenarigo/compare/v0.11.0...v0.11.1
[v0.11.0]: https://github.com/zoncoen/scenarigo/compare/v0.10.0...v0.11.0
[v0.10.0]: https://github.com/zoncoen/scenarigo/compare/v0.9.0...v0.10.0
[v0.9.0]: https://github.com/zoncoen/scenarigo/compare/v0.8.1...v0.9.0
[v0.8.1]: https://github.com/zoncoen/scenarigo/compare/v0.8.0...v0.8.1
[v0.8.0]: https://github.com/zoncoen/scenarigo/compare/v0.7.0...v0.8.0
[v0.7.0]: https://github.com/zoncoen/scenarigo/compare/v0.6.3...v0.7.0
[v0.6.3]: https://github.com/zoncoen/scenarigo/compare/v0.6.2...v0.6.3
[v0.6.2]: https://github.com/zoncoen/scenarigo/compare/v0.6.1...v0.6.2
[v0.6.1]: https://github.com/zoncoen/scenarigo/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/zoncoen/scenarigo/compare/v0.5.1...v0.6.0
[v0.5.1]: https://github.com/zoncoen/scenarigo/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/zoncoen/scenarigo/compare/v0.4.0...v0.5.0
[v0.4.0]: https://github.com/zoncoen/scenarigo/compare/v0.3.3...v0.4.0
[v0.3.3]: https://github.com/zoncoen/scenarigo/compare/v0.3.2...v0.3.3
[v0.3.2]: https://github.com/zoncoen/scenarigo/compare/v0.3.1...v0.3.2
[v0.3.1]: https://github.com/zoncoen/scenarigo/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/zoncoen/scenarigo/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/zoncoen/scenarigo/compare/v0.1.0...v0.2.0
