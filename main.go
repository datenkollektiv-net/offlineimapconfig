package main

import (
	"log"

	flag "github.com/spf13/pflag"

	"gopkg.in/ini.v1"
)

var (
	sourceType          = flag.String("sourceType", "IMAP", "type of all source repositories")
	sourceUseStartTLS   = flag.Bool("sourceUseStartTLS", true, "use StartTLS for all source repositories")
	sourceUseSSL        = flag.Bool("sourceUseSSL", false, "use SSL for all source repositories")
	sourceHost          = flag.String("sourceHost", "", "remote host of all source repositories")
	sourceFolderFilter  = flag.String("sourceFolderFilter", "", "source folder filter lambda")
	sourceNameTrans     = flag.String("sourceNameTrans", "", "source nametrans lambda")
	sourceReadOnly      = flag.Bool("sourceReadOnly", false, "source is read only")
	sourceCreateFolders = flag.Bool("sourceCreateFolders", false, "create folders on source")
	remoteType          = flag.String("remoteType", "IMAP", "type of all remote repositories")
	remoteUseStartTLS   = flag.Bool("remoteUseStartTLS", true, "use StartTLS for all remote repositories")
	remoteUseSSL        = flag.Bool("remoteUseSSL", false, "use SSL for all remote repositories")
	remoteHost          = flag.String("remoteHost", "", "remote host of all remote repositories")
	remoteFolderFilter  = flag.String("remoteFolderFilter", "", "remote folder filter lambda")
	remoteNameTrans     = flag.String("remoteNameTrans", "", "remote nametrans lambda")
	remoteReadOnly      = flag.Bool("remoteReadOnly", false, "remote is read only")
	remoteCreateFolders = flag.Bool("remoteCreateFolders", true, "create folders on source")
	inputFile           = flag.String("in", "./input.csv", "input csv file with all source/remote account/password combinations")
	outputFile          = flag.String("out", "./offlineimaprc", "file to store config to")
	sslCertPath         = flag.String("sslCertPath", "/etc/ssl/certs/ca-certificates.crt", "path to ssl certificates")
)

func main() {
	var (
		err      error
		f        *ini.File
		accounts Accounts
	)
	flag.Parse()
	f = ini.Empty()
	if accounts, err = ReadCSV(*inputFile); err != nil {
		log.Fatalf("could not read input csv: %v", err)
	}
	if err = accounts.Write(f, RepoConfig{
		Type:          *sourceType,
		StartTLS:      *sourceUseStartTLS,
		SSL:           *sourceUseSSL,
		SSLCertPath:   *sslCertPath,
		RemoteHost:    *sourceHost,
		FolderFilter:  *sourceFolderFilter,
		NameTrans:     *sourceNameTrans,
		ReadOnly:      *sourceReadOnly,
		CreateFolders: *sourceCreateFolders,
	}, RepoConfig{
		Type:          *remoteType,
		StartTLS:      *remoteUseStartTLS,
		SSL:           *remoteUseSSL,
		SSLCertPath:   *sslCertPath,
		RemoteHost:    *remoteHost,
		FolderFilter:  *remoteFolderFilter,
		NameTrans:     *remoteNameTrans,
		ReadOnly:      *remoteReadOnly,
		CreateFolders: *remoteCreateFolders,
	}); err != nil {
		log.Fatalf("Could not write accounts to ini: %v", err)
	}
	if err = f.SaveTo(*outputFile); err != nil {
		log.Fatalf("Could not save accounts to ini: %v", err)
	}

}
