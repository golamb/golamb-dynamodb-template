# Go Lambda Simple Template for GoLambda

## Init project

```sh
golamb init simple my-project
```

## Require

* [dep](https://github.com/golang/dep)
* [aws-lambda-go](https://github.com/aws/aws-lambda-go/)

## For developers on Windows
```bash
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```


## First Step

* Create function in AWS Lambda
* Choose Runtime to Go 1.x
* Create role or choose existing role

## Atfer create

* edit Handler to "main"
* save and deploy

## Dev

```sh
dep ensure
```

### But now not yet run lambda in locally. My way just normal run func Handle in main
