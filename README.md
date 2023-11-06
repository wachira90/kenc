# kenc
Kubernetes encrypt yaml file 

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

## Test

encrypt

```
C:\test\kenc>kenc encode test.yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
type: Opaque
data:
  access_key_cert: WlpaWlpYWFhYWENDQ0NDVlZWVlY=
  access_key_secret: VlZWVlZDQ0NDQ1hYWFhYWlpaWlo=
  secret_key_cert: RU5hV3hvUWo0Y3RMNW81ZFZNaTlMdHl3RDJ1eW9ERWRGNUplZ2dGSA==
  secret_key_secret: ejJiQjQ4UDdOVXFNb2ZmSjNUVWZqWHI4VFR0YURXZDRqVFlEdVFDcw==
```

decrypt

```
C:\test\kenc>kenc decode enc-test.yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
type: Opaque
data:
  access_key_cert: ZZZZZXXXXXCCCCCVVVVV
  access_key_secret: VVVVVCCCCCXXXXXZZZZZ
  secret_key_cert: ENaWxoQj4ctL5o5dVMi9LtywD2uyoDEdF5JeggFH
  secret_key_secret: z2bB48P7NUqMoffJ3TUfjXr8TTtaDWd4jTYDuQCs
```
