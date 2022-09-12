# Redirect urls
The goal is to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs we would look for any incoming web requests with the path /dogs and redirect them.

Redirects rules can be stored in
- YAML (default: config.yaml)
- JSON

Original task: https://courses.calhoun.io/lessons/les_goph_04

## Redirect configuration
The rules for redirecting requests are located in the assets directory.

Depending on the input choice: yaml or json, the file has /assets/config. there will be a corresponding extension - /assets/config.yaml.

The default input method is the yaml file **_/assets/config.yaml_**.

The files contain data according to the following structure:

```
paths: [ path1{path, url}, ... ]
```

## Run options
The file main.go is located in **_/cmd_**.

### Build command:
```
go to the build -o redirect.
```

### Help command:
```
./redirect --help
```

### Launch Command
```
./redirect --input yaml --filepath "../assets/config.yaml"
```

This command will output:
```
start decode the yaml file...
the yaml file has been decoded...

config
   paths:
        - path: /go
          url: https://go.dev/
        - path: /go_pkg
          url: https://pkg.go.dev/
        - path: /path3
          url: http://localhost:8080/redirections/path3

init a new http handler...
start http server at 8080...
```
