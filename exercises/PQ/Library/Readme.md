# GO KIT TEMPLATE

This repository contains a template for starting go kit app

### Running the REST API server

The web server runs on port `8082` by default. It can be changed under the `config` folder. There is a configuration 
file per environment. The property that needs to be changed is `HttpPort`.

To are two ways of running the server locally, **with** and **without** *Docker*.

#### Without Docker

First make sure you have Go installed. Then run the following command:
```
make run
```
*Note: Dependencies will be downloaded as part of the run.*
 
 #### With Docker
 First make sure you have Docker installed. Then run the following command:
 ```
make docker-run-local
```

### Tests

First make sure you have Docker installed as the tests are run with it. To run the tests run the following command:
```
make tests
```

## IDE Setup
 
Run `goimports` on save.
 
Run `golint` and `go vet` to check for errors.
 
You can find information in editor support for Go tools here: [https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins)