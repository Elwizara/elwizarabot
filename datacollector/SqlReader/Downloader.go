package main

import (
	"database/sql"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

//UpdateUsersPageToCrawl :
func UpdateUsersPageToCrawl(dbsource *sql.DB, source DBSource, DestinationConnectionString string) error {
	NeedToCrawlCount, err := GetUsersNeedToCrawlCount(dbsource)
	if err != nil {
		logger.Critical(err)
		return err
	}
	if NeedToCrawlCount < source.MinimumNeedToCrawlCount {
		scriptPath := fmt.Sprintf(config.BackupscriptPath, "UsersPageToCrawl")
		output, err := exec.Command(scriptPath, DestinationConnectionString, fmt.Sprint(source.PageNeedToCrawlCount), source.ConnectionString).Output()
		if err != nil {
			logger.Critical(err)
			return err
		}
		if !strings.HasPrefix(fmt.Sprintf("%s", output), "COMMIT") {
			logger.Criticalf("Unexpected output:'%s'", output)
		} else {
			logger.Infof("%v COMMIT UsersPage Successfull ", source.Name)
		}
	} else {
		logger.Infof("%v No need to upload users page to crawl MinimumNeedToCrawlCount:%v SourceNeedToCrawlCount:%v", source.Name, source.MinimumNeedToCrawlCount, NeedToCrawlCount)
	}
	return nil
}

//ExecScript :
func ExecScript(dbsource *sql.DB, configuration *Configuration, source DBSource, table DBTable) {
	/*MAXUpdatedLimit, err := MAXUpdatedAtLimit(dbsource, table.Name, source.DownloadLimit)
	if err != nil {
		logger.Criticalf("%v:'%v' MAXUpdatedAt Error : %v", source.Name, table.Name, err)
		return
	}
	if MAXUpdatedLimit == 0 {
		logger.Warningf("%v:'%v' MAXUpdatedAt is 0", source.Name, table.Name)
		return
	}*/
	MAXUpdatedLimit := 1000
	output, err := exec.Command(table.ScriptPath, source.ConnectionString, configuration.DBDestination.ConnectionString, fmt.Sprintf("%v.gz", MAXUpdatedLimit), fmt.Sprint(MAXUpdatedLimit)).Output()
	if err != nil {
		logger.Criticalf("%v:'%v' Command failed:'%v'  output:%s", source.Name, table.Name, err, output)
	} else {
		logger.Infof("COMMIT %v:%v:%v\n%s", source.Name, table.Name, MAXUpdatedLimit, output)
	}
}

//StartDownloadSource :
func StartDownloadSource(configuration *Configuration, source DBSource) {
	dbsource := initDatabase(source.ConnectionString, source.Name)
	duration := time.Duration(source.DownloadDuration)
	for {
		for _, table := range configuration.DBDestination.Tables {
			ExecScript(dbsource, configuration, source, table)
		}
		<-time.After(duration * time.Second)
	}
}
