ldcli written using [kong](https://github.com/alecthomas/kong) instead of cobra.

# functionality to duplicate

* [x] testing
* [x] required flags
* [x] get flag command
* [x] list flag command
* [x] config file for flags
* [x] env vars for flags
* [ ] event tracking - attempt
* [ ] event tracking - success result
* [ ] event tracking - help
* [ ] event tracking - error
* [ ] plain text vs. json output

```
go run . flags get --environment test --project default --flag my-new-flag
go run . flags list --project default
```
