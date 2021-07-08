package main

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

type RepoConfig struct {
	Type          string
	StartTLS      bool
	SSL           bool
	RemoteHost    string
	SSLCertPath   string
	FolderFilter  string
	NameTrans     string
	ReadOnly      bool
	CreateFolders bool
}

func (c RepoConfig) AddConfig(s *ini.Section) error {
	var err error
	if _, err = s.NewKey("type", c.Type); err != nil {
		return err
	}
	if _, err = s.NewKey("starttls", yesNo(c.StartTLS)); err != nil {
		return err
	}
	if _, err = s.NewKey("readonly", yesNo(c.ReadOnly)); err != nil {
		return err
	}
	if _, err = s.NewKey("ssl", yesNo(c.SSL)); err != nil {
		return err
	}
	if c.SSL || c.StartTLS {
		s.NewKey("ssl_version", "tls1_2")
		s.NewKey("sslcacertfile", c.SSLCertPath)
	}
	if _, err = s.NewKey("remotehost", c.RemoteHost); err != nil {
		return err
	}
	if _, err = s.NewKey("createfolders", yesNo(c.CreateFolders)); err != nil {
		return err
	}
	if c.FolderFilter != "" {
		if _, err = s.NewKey("folderfilter", c.FolderFilter); err != nil {
			return err
		}
	}
	if c.NameTrans != "" {
		if _, err = s.NewKey("nametrans", c.NameTrans); err != nil {
			return err
		}
	}
	return nil

}

type Accounts []*SourceRemote

func (a Accounts) Write(f *ini.File, sourceConf, remoteConf RepoConfig) error {
	var err error
	if err = a.WriteGeneral(f); err != nil {
		return err
	}
	if err = a.WriteAccounts(f, sourceConf, remoteConf); err != nil {
		return err
	}
	if err = a.WriteRepositories(f, sourceConf, remoteConf); err != nil {
		return err
	}
	return nil
}

func (a Accounts) WriteGeneral(f *ini.File) error {
	var (
		general *ini.Section
		err     error
		ids     = make([]string, 0, len(a))
	)
	if general, err = f.NewSection("general"); err != nil {
		return err
	}
	for _, acc := range a {
		ids = append(ids, acc.ID)
	}
	general.NewKey("accounts", strings.Join(ids, ","))
	return nil
}

func (a Accounts) WriteAccounts(f *ini.File, sourceRepoConf, remoteRepoConf RepoConfig) error {
	var err error
	for _, acc := range a {
		if err = acc.WriteAccount(f); err != nil {
			return err
		}
	}
	return nil
}

func (a Accounts) WriteRepositories(f *ini.File, sourceRepoConf, remoteRepoConf RepoConfig) error {
	var err error
	for _, acc := range a {
		if err = acc.WriteRepositories(f, sourceRepoConf, remoteRepoConf); err != nil {
			return err
		}
	}
	return nil
}

func NewSourceRemote(line []string, name string, idx int) *SourceRemote {
	return &SourceRemote{
		ID:             fmt.Sprintf("%v-%v", name, idx),
		sourceAccount:  line[0],
		sourcePassword: line[1],
		remoteAccount:  line[2],
		remotePassword: line[3],
	}
}

type SourceRemote struct {
	ID             string
	sourceAccount  string
	sourcePassword string
	remoteAccount  string
	remotePassword string
}

func (sr SourceRemote) GetRepoSource() string {
	return "RepoSource-" + sr.ID
}

func (sr SourceRemote) GetRepoRemote() string {
	return "RepoRemote-" + sr.ID
}

func (sr SourceRemote) WriteAccount(f *ini.File) error {
	var (
		section *ini.Section
		err     error
	)
	if section, err = f.NewSection("Account " + sr.ID); err != nil {
		return err
	}
	if _, err = section.NewKey("localrepository", sr.GetRepoSource()); err != nil {
		return err
	}
	if _, err = section.NewKey("remoterepository", sr.GetRepoRemote()); err != nil {
		return err
	}
	if _, err = section.NewKey("synclabels", "yes"); err != nil {
		return err
	}
	if _, err = section.NewKey("utf8foldernames", "yes"); err != nil {
		return err
	}
	return nil
}

func (sr SourceRemote) WriteRepositories(f *ini.File, sourceRepoConfig, remoteRepoConfig RepoConfig) error {
	var (
		source, remote *ini.Section
		err            error
	)
	if source, err = f.NewSection("Repository " + sr.GetRepoSource()); err != nil {
		return err
	}
	if remote, err = f.NewSection("Repository " + sr.GetRepoRemote()); err != nil {
		return err
	}
	if err = sourceRepoConfig.AddConfig(source); err != nil {
		return err
	}
	if err = remoteRepoConfig.AddConfig(remote); err != nil {
		return err
	}
	if _, err = source.NewKey("remoteuser", sr.sourceAccount); err != nil {
		return err
	}
	if _, err = source.NewKey("remotepass", sr.sourcePassword); err != nil {
		return err
	}
	if _, err = remote.NewKey("remoteuser", sr.remoteAccount); err != nil {
		return err
	}
	if _, err = remote.NewKey("remotepass", sr.remotePassword); err != nil {
		return err
	}
	return nil
}

func yesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
