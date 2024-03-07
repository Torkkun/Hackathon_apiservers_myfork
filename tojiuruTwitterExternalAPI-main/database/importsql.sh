#!/bin/bash

function exit_trap() {
    if [ $? != 0 ]; then
        echo "Command [$BASH_COMMAND] is failed"
        exit 1
    fi
}
trap exit_trap ERR

touch ~/.sqliterc
echo "PRAGMA foreign_keys = ON;" >> ~/.sqliterc
dbfile="/go/src/script/tojiuru.db"
touch $dbfile
ddl_path="/go/src/database/sqlfile/ddl/*"
dml_path="/go/src/database/sqlfile/dml/*"

ddls=`find $ddl_path -maxdepth 0 -type f -name *.sql`
for ddl in $ddls;
do
    sqlite3 $dbfile < $ddl
done

dmls=`find $dml_path -maxdepth 0 -type f -name *.sql`

for dml in $dmls;
do
    sqlite3 $dbfile < $dml
done