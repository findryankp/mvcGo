package apps

func CreateEnv() string {
	var text = `DBUSERNAME=findryankp
DBPASS=findryankp
DBHOST=localhost
DBPORT=3306
DBNAME=testdb
	`
	return text
}
