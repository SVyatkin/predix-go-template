[![Go](https://github.com/SVyatkin/predix-go-template/workflows/Go/badge.svg)](https://github.com/SVyatkin/predix-go-template/actions)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=SVyatkin_predix-go-template&metric=alert_status)](https://sonarcloud.io/dashboard?id=SVyatkin_predix-go-template)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=SVyatkin_predix-go-template&metric=bugs)](https://sonarcloud.io/dashboard?id=SVyatkin_predix-go-template)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=SVyatkin_predix-go-template&metric=coverage)](https://sonarcloud.io/dashboard?id=SVyatkin_predix-go-template)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=SVyatkin_predix-go-template&metric=ncloc)](https://sonarcloud.io/dashboard?id=SVyatkin_predix-go-template)

# predix-go-template
Predix Golang Microservice Template

Please use  <a href=https://golang.org/doc/install?download=go1.5.windows-amd64.msi2> a link </a> to install Go and <a href=https://github.com/tools/godep/blob/master/Readme.md> another link </a> to install Godep

- $ git clone predix-go-template repository
- $ cd predix-go-template
- $ godep save
- $ go test
- $ cf push
- $ curl https://predix-go-template.run.aws-usw02-pr.ice.predix.io
    <div >Predix Go Microservice Template</div><a href='/env' target='_blank'>Environment variables</a>
- $ curl https://predix-go-template.run.aws-usw02-pr.ice.predix.io/env
    returns environment variables
