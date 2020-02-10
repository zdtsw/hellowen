# Go program: hellowen
# Description: a simple demo project to show how CI/CD works

Check
>go fmt
>go vet
>golint

Build
>go build

Run
>./hellowen

Test
>go test



# Due to current v0.2 image from jenkinsci does only support mvn build, 
# User for image do not grant sudo to install go package.
# Change projec to a even more simple solution: bash kinda hello Wen
Test Positive result
> ./test.sh
hello Wen
Test pass!

Test Negative result
>./test.sh
exit 1


![Jenkins single-shot master](https://github.com/zdtsw/hellowen/workflows/Jenkins%20single-shot%20master/badge.svg?branch=master)
