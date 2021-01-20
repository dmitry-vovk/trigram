<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-96%25-brightgreen.svg?longCache=true&style=flat)</a>

# Trigram text generator

## App structure

The app consists of several packages:
  * trigrammer: core app code with public API
    * storage: stores trigrams
    * trigram: common types used across the app


## Requirements, assumptions, implementation details

Endpoint `GET /generate` has compulsory parameters `word1` & `word2`, so the URL for a request will look like `/generate?word1=it&word2=was`.

An effort was made to reduce memory consumption when storing trigrams.
Every word is stored as unique string, and indexed using uint32 numbers.
The type was selected for capacity enough to have a unique ID for each english word (up to 4,294,967,295 words),
but small enough to not waste memory unnecessary.

Further improvement can be made around words splitting/normalisation.
For example, the app could strip quotes in cases like this:
```
  “My dear Mr. Bennet,” said his lady to him one day, “have you
  heard that Netherfield Park is let at last?”
```
The first and the second trigram look like

  - [`“My` `dear` `Mr.`]
  - [`dear` `Mr.` `Bennet,”`]

whereas for better text generation quotes better be stripped:

  - [`My` `dear` `Mr.`]
  - [`dear` `Mr.` `Bennet,`]

Punctuation and register can be handled so that seed words `mr` `bennet` would match a trigram with `Mr.` `Bennet,` words.

### Testing

```shell
go test -race ./...
````

## Running

### Standalone app

```shell
$ go build main.go
$ ./main
```

### Containerized app

```shell
$ docker build . -t trigram
$ docker run -p 8080:8080 trigram
```

### Useful scripts

Folder `scripts` contains some helpful scripts:
  * `learn.sh` can be used to feed in a text file
  * `generate.sh` makes request to generate random txt
  * `pprof.sh` can be used to generate performance profile charts (can be found in `scripts/performance_reports` folder)
