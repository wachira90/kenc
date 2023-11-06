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

## use -s for overwrite or -o for output new file

```
kenc encode filename.yaml -s

kenc encode filename1.yaml -o filename2.yaml
```

## To decode from base64 to string

```
kenc decode filename.yaml
```

## Test

encode

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

decode

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
