Sample

```yaml

host: cov.dryclean.io
auth: super-secret-auth
loadersDir: ./loaders
files:
    -
        path: ./coverage/cobertura.xml
        label: cover/cobertura
        filename: cover.xml
        binary: false
    -
        path: ./report/lint.xml
        label: lint/codestyle
        filename: lint.xml
    -
        path: ./report/mocha.xml
        label: test/junit
        filename: test.xml
  ```