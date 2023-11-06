# kenc
kubernetes encrypt 

## install

```
go get -u github.com/wachira90/kenc

go install github.com/wachira90/kenc
```


## To encode from string to base64

```
kenc encode filename.yaml
```

## By default the output of that YAML file will be printed to stdout. You can use -o to write the output to another file or -s to save output to the save file.

```
kenc encode filename.yaml -s

kenc encode filename.yaml -o filename.yaml
```

## To decode from base64 to string. The flag -o and -s can also be used here.

```
kenc decode filename.yaml
```
