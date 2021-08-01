# datastorecli

## Install

Download released files from [Releases page](https://github.com/akm/datastorecli/releases).

Or

```
cd somwwhere/to/parent/dir
git clone https://github.com/akm/datastorecli.git
cd datastorecli
go install ./cmd/datastorecli
```

## Setup

Login for Application Default Credentials.

```
gcloud auth application-default login
```

## Usage

### Query

```
$ datastorecli query calcsvc --project-id=your-project
{"A":100,"B":23,"R":123}
{"A":1,"B":5,"R":6}
{"A":100,"B":24,"R":124}
{"A":100,"B":24,"R":124}
{"A":1,"B":5,"R":6}
{"A":100,"B":26,"R":126}
{"A":1,"B":5,"R":6}
{"A":100,"B":26,"R":126}
```

### Put

```
$ datastorecli put calcsvc '{"A":200,"B":36,"R":128}' --project-id=your-project
/calcsvc,5646874153320448
```

```
$ datastorecli put calcsvc 5646874153320448 '{"A":300,"B":36,"R":128}' --project-id=your-project
/calcsvc,5646874153320448
```

### Get

```
$ datastorecli get calcsvc 5646874153320448 --project-id=your-project
{"A":300,"B":36,"R":128}
```

### Delete

```
$ datastorecli delete calcsvc 5646874153320448 --project-id=ayour-project
$ echo $?
0
```

### Encode and Decode Key

#### --id

```
$ datastorecli key encode calcsvc --id 5646874153320448; echo
EhIKB2NhbGNzdmMQgICAwNX5gwo
```

```
$ datastorecli key decode EhIKB2NhbGNzdmMQgICAwNX5gwo; echo
/calcsvc,5646874153320448
```

#### --name

```
$ datastorecli key encode accounts --name user1@example.com; echo
Eh0KCGFjY291bnRzGhF1c2VyMUBleGFtcGxlLmNvbQ
```

```
$ datastorecli key decode Eh0KCGFjY291bnRzGhF1c2VyMUBleGFtcGxlLmNvbQ; echo
/accounts,user1@example.com
```

#### --namespace

```
$ datastorecli key encode accounts --name user1@example.com --namespace shop1; echo
CgciBXNob3AxEh0KCGFjY291bnRzGhF1c2VyMUBleGFtcGxlLmNvbQ
```

```
$ datastorecli key decode CgciBXNob3AxEh0KCGFjY291bnRzGhF1c2VyMUBleGFtcGxlLmNvbQ; echo
/accounts,user1@example.com (namespace:shop1)
```

#### --encoded-parent

```
$ datastorecli key encode account_logs --id 123456789 --encoded-parent Eh0KCGFjY291bnRzGhF1c2VyMUBleGFtcGxlLmNvbQ; echo
Eh0KCGFjY291bnRzGhF1c2VyMUBleGFtcGxlLmNvbRITCgxhY2NvdW50X2xvZ3MQlZrvOg
```

```
$ datastorecli key decode Eh0KCGFjY291bnRzGhF1c2VyMUBleGFtcGxlLmNvbRITCgxhY2NvdW50X2xvZ3MQlZrvOg; echo
/accounts,user1@example.com/account_logs,123456789
```
