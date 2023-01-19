EN / [ES](./README_ES.md)

# Template Gin Golang

## Installation
To install this project it is necessary to have docker installed.
Once docker is installed, we can start the installation.

We run this command to build the dockerfile, it is important to 
```
docker build . -t template-go
```

Then we proceed to execute the image with:
```
docker run -i -t -p 80:8080 template-go
```

The name of the image can be changed, that is to say, instead of putting `template-go`, we can put a different name, what is important is to put a correct syntax (In small letters and with hyphens instead of spaces).
we can put a different name, what is important is to put a correct syntax (in lowercase and with hyphens instead of spaces)

`IMPORTANT: In case of modification of the version in Go.mod, it is important to change it also in the dockerfile. In case it is not changed, it may cause errors.`
