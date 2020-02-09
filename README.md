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



# Due to current v2.0 image from jenkinsci does only support mvn build, and I do not have sudo to install go package.
# Change projec to a much simple solution: bash which support by ubuntu for sure
Test Positive result
> ./test.sh
hello Wen
Test pass!

Test Negative result
>./test.sh
exit 1
