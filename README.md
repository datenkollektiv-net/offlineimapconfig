# offlineimapconfig

This is a tool to generate a offlineimaprc for bulk synchronization. It reads data from an input csv and configuration parameters as flags to produce the config file.

## Usage

To build the project run `go build`.

```
Usage of ./offlineimapconfig:
      --in string                   input csv file with all source/remote account/password combinations (default "./input.csv")
      --out string                  file to store config to (default "./offlineimaprc")
      --remoteCreateFolders         create folders on source (default true)
      --remoteFolderFilter string   remote folder filter lambda
      --remoteHost string           remote host of all remote repositories
      --remoteNameTrans string      remote nametrans lambda
      --remoteReadOnly              remote is read only
      --remoteType string           type of all remote repositories (default "IMAP")
      --remoteUseSSL                use SSL for all remote repositories
      --remoteUseStartTLS           use StartTLS for all remote repositories (default true)
      --sourceCreateFolders         create folders on source
      --sourceFolderFilter string   source folder filter lambda
      --sourceHost string           remote host of all source repositories
      --sourceNameTrans string      source nametrans lambda
      --sourceReadOnly              source is read only
      --sourceType string           type of all source repositories (default "IMAP")
      --sourceUseSSL                use SSL for all source repositories
      --sourceUseStartTLS           use StartTLS for all source repositories (default true)
      --sslCertPath string          path to ssl certificates (default "/etc/ssl/certs/ca-certificates.crt")
```

Simple execution: `./offlineimapconfig --remoteHost imap.destinationhost.net --sourceHost imap.sourcehost.de --in input.csv`
