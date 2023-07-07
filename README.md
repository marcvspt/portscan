# TCP Port Scanner
Port scanner for TCP IPv4 and IPv6 programming in golang.

## Set up
Build from source code, the golang script name **portscan.go**.

```bash
git clone https://github.com/marcvspt/portscan
cd portscan
go build -ldflags "-s -w" -o portscan portscan.go
```

`-ldflags "-s -w"` reduce the size of the binary.

### Optional
Reduce a little bit more the size of the binary.

```bash
upx brute portscan
```

## Usage
`./portscan -h 1.1.1.1 -p 8080 [options]`
|**Argument**|**Description**|**Example**|
|-|-|-|
|`-help`|Show help panel|`-help`|
|`-attempts` *int*|Number of connection attempts for each port|`-attempts 4`|
|`-h` *string*|Host to scan, available at IPv4 and IPv6|`-h example.com`, `-h 8.8.4.4`, `-h 2001:4860:4860::8888`|
|`-p` *string*|Port or range of ports to be scanned|`-p 22`, `-p 1-1024`|
|`-timeout` *int*|Response timeout in milliseconds for each port|`-timeout 2198`|



## Download the binary
You can download the compiled binary in the [releases](https://github.com/marcvspt/portscan/releases/tag/v1.0) page.
