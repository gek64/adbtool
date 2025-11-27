# ADBTool

## Features

- ADB Batch Tool

## Example

```sh
# authentication first
adb devices

# if you want to get a list of installed apps on your device
adb shell pm list packages | sort > app.txt

# uninstall all apps
adbtool uninstall -a

# uninstall all apps from a file
adbtool uninstall -f app.txt

# uninstall all apps for user 0 from a file
adbtool uninstall-user -u 0 -f app.txt
```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/adbtool https://github.com/unix755/adbtool/releases/latest/download/adbtool-linux-amd64
chmod +x /usr/local/bin/adbtool
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/unix755/adbtool.git
cd adbtool
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
