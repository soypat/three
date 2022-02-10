# Examples

## Building examples
1. First step is to run `gopherjs build` inside each directory (each directory is a self contained project).

2. Once built, you can try serving the html+javascript via a server. If you built the project you can use `gopherjs`'s built
in server:
```shell
gopherjs serve
```
or any other http server in this folder and head over to the corresponding link (port :8080)
* [basic example](https://localhost:8080/basic/three_example.html)

## How to build `basic` example and serve it
```shell
cd basic # from examples, switch directory to basic
gopherjs build
cd .. # go back to examples directory
gopherjs serve
```
[Default gopherjs server landing page is http://localhost:8080](http://localhost:8080), you can navigate folders from here.