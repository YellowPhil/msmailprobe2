# msmailprobe 2.0

## IN DEVELOPMENT

A **BETTER** Office 365 and Exchange Enumeration

It is widely known that OWA (Outlook Webapp) is vulnerable to time-based user enumeration attacks. This tool leverages all known, and even some lesser-known services exposed by default Exchange installations to enumerate users. It also targets Office 365 for error-based user enumeration. 

Moreover this version of the script **fixes** some of the most common issues of the original.

## Motivation

Original `msmailprobe` had some arguable features (not to mention it was a single-file app) which this script aims to fix:

1) **RATE LIMITS**. You can now change maximum requests the script sends per minute or per second (`--rpm`, `--rps` respectfully).

2) The script now updates the output file (specified by `-o` flag) on *EVERY* found email. No more losing your whole progress because of accidental Ctrl+C.

3) **No** `log.Fatal()` when receiving timeout (and losing your progress). Just wait for a reasonable amount of time (1 hour by default).

4) Random `User-Agent` header on every request to reduce likelihood of a ban.

5) Also *auto-throttling* when receiving 429 response status code (in case Exchange is protected by some kind of WAF).

6) **BETTER** logging which includes timestamps and an optional debug log.

###### Bonus

The whole project was rewritten following the guidelines of the modern Go programming getting rid of ALL the shitty code.

I mean just look at this

```go 
if webRequestCodeResponse(url1) == 401 {
		urlToHarvest = url1
	} else if webRequestCodeResponse(url2) == 401 {
		urlToHarvest = url2
	} else if webRequestCodeResponse(url3) == 401 {
		urlToHarvest = url3
	} else if webRequestCodeResponse(url4) == 401 {
		urlToHarvest = url4
	} else if webRequestCodeResponse(url5) == 401 {
		urlToHarvest = url5
	} else if webRequestCodeResponse(url6) == 401 {
		urlToHarvest = url6
	}
```

or this 

```go
    mux.Lock()
    fmt.Printf(BrightGreen, "[+] "+user+" - ")
    fmt.Printf(BrightGreen, elapsedTime)
    fmt.Println("")
```

## Getting Started


If you want to download and compile the simple, non-dependant code, you must first install GoLang! I will let the incredible documentation, and other online resources help you with this task.

https://golang.org/doc/install

You may also download the compiled release [here](https://github.com/yellowphil/msmailprobe2/releases).

## Syntax

List examples of commands for this applications, but simply running the binary with the `examples` command:

```bash
msmailprobe examples
```

You can also get more specific help by running the binary with the arguments you are interested in:

```bash
msmailprobe identify
msmailprobe userenum
msmailprobe userenum onprem
msmailprobe userenum o365
```

## Usage

#### Identify Command
* Used for gathering information about a host that may be pointed towards an Exchange or o365 tied domain
* Queries for specific DNS records related to Office 365 integration
* Attempts to extract internal domain name for onprem instance of Exchange
* Identifies services vulnerable to time-based user enumeration for onprem Exchange
* Lists password-sprayable services exposed for onprem Exchange host

```
Flag to use:
	-t to specify target host

Example:
	./msmailprobe identify -t mail.target.com
```

#### Userenum (o365) Command
* Error-based user enumeration for Office 365 integrated email addresses

```
Flags to use:
	-E for email list OR -e for single email address
	-o [optional]to specify an out file for valid emails identified
	--threads [optional] for setting amount of requests to be made concurrently

Examples:
	./msmailprobe userenum --o365 -E emailList.txt -o validemails.txt --threads 25
	./msmailprobe userenum --o365 -e admin@target.com
```

#### Userenum (onprem) Command
* Time-based user enumeration against multiple onprem Exchange services

```
Flags to use:
	-t to specify target host
	-U for user list OR -u for single username
	-o [optional]to specify an out file for valid users identified
	--threads [optional] for setting amount of requests to be made concurrently

Examples:
	./msmailprobe userenum --onprem -t mail.target.com -U userList.txt -o validusers.txt --threads 25
	./msmailprobe userenum --onprem -t mail.target.com -u admin
```

## Acknowledgments

* [**poptart**](https://github.com/HosakaCorp) - *For a truck load of golang assistance, poking of Exchange services, and help testing timing of responses*
* [**jlarose**](https://github.com/jordanlarose) - *Parsing decimal data within NTLMSSP authentication reponse for internal domain name*
* [**Vincent Yui**](https://github.com/vysec)  - *Office 365 check python script*
* **grimhacker** - *Discovery/disclosure of error-based user enumeration within Office 365* [blog post](https://grimhacker.com/2017/07/24/office365-activesync-username-enumeration/)
* **Nate Power** - *Discovery and disclosure of OWA time-based user enumeration*

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/customsync/msmailprobe/blob/master/LICENSE) file for details
