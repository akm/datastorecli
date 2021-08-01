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
