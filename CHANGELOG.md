## [1.1.0](https://github.com/maohieng/errs/compare/v1.0.3...v1.1.0) (2023-03-26)

### Changes
* `Error()` display format, for example
` svc.Create: persist.Create unable to create record, database error, code 13`
* Add support error stack msg
* Bring back support `string` args that was removed in `v1.0.2` 
to be used as stack msg

### New
* `Stack` via `Errors()` func that support json marshal. 
Generally be used to print error detail.

## [1.0.2](https://github.com/maohieng/errs/compare/v1.0.1...v1.0.2) (2023-03-25)

### Breaking changes
* Remove support `string` args of `New()` and `SNew()` 

### Performance improvement

* Improve `Error()` 49% performance and 44% allocation reduction ([950a7cc](https://github.com/maohieng/errs/commit/950a7cc9d653a3de576c98418821d193a36ff9e4))


## [1.0.1](https://github.com/maohieng/errs/compare/v1.0.1...v1.0.0) (2023-03-05)
