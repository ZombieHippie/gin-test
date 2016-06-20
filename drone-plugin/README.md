Sample

In this directory,

make a `./settings.json` file for your own image, and then execute:
```bash
docker build --rm -t colelawr/drone-dryclean .
```
> (don't forget that last period indicating this directory)

```yaml

host: cov.dryclean.io
auth: super-secret-auth
loadersDir: ./loaders
files:
    -
        path: ./coverage/cobertura.xml
        label: Code Coverage
        loader: "dc-coverage"
    -
        path: ./report/tslint.txt
        label: Typescript Linting
        loader: "dc-lint?format=tslint-prose"

    -
        path: ./report/mocha.xml
        label: Mocha Tests
        loader: "junit"
  ```
