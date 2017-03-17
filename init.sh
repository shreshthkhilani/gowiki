#!/bin/bash

# Bash script to install Goland and PostgreSQL on RHEL 7.3

### Set up Golang

sudo yum install golang

# put these in ~/.bash_profile
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
export GOBIN="$HOME/go/bin"
export GOPATH="$HOME/go/src"
export GOROOT="/usr/local"

### Set up PostgreSQL

sudo yum install postgresql-server postgresql-contrib
sudo postgresql-setup initdb

# change the `ident`s to `trust`s for a pw-free experience
# or to `md5` to do pws in the following file:
# $ sudo vi /var/lib/pgsql/data/pg_hba.conf
# then restart if already started
# $ sudo systemctl restart postgresql

sudo systemctl start postgresql
sudo systemctl enable postgresql

### to access a db called `ipcmdb` with a user called `postgres`

sudo -u postgres psql ipcmdb

# create a new user with a password, please