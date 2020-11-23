# ```service_templated``` [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Download%20the%20brandnew%20Golang%20Service%20Template%20by%204thel00z&url=https://github.com/4thel00z/service_templated&hashtags=golang,go,service,template,architecture)


![service_templated-tests](https://github.com/4thel00z/service_templated/workflows/Test/badge.svg)
![service_templated-logo](https://github.com/4thel00z/service_templated/raw/assets/logo.png)

## What this project is about

This is [my](https://github.com/4thel00z) go http service template.
It sports features like:

- validation support (see debug module for example)
- jwt validation support
- module support (see debug module for example)


## How do I install it ?

To create a new project simply invoke this script, (make sure to use your own project name instead of `<project_name>` lol):

```
curl  --proto '=https' --tlsv1.2 -L -sSf https://shortly.fun/boilerplate | bash -s <project_name>
```

## How do I run it ?

After creating a new project like above you can simply run:

```
make run
```

or

```
just run
```

if you have [just](https://github.com/casey/just) installed.
Running `make help` will show you the rest of the targets.

## License

This project is licensed under the GPL-3 license.
