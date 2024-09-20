the program needs to be compiled into a binary
this command is used to compile and abselutely necessary
```
go build -o student/main student/main.go
```
example of full build.sh script
```
go build -o student/main student/main.go
docker build -q -t guessit .
docker run --rm --name audit -p 3000:3000 guessit
docker stop audit
docker rmi guessit
rm -rf student/main
```