Table of Contents
=================

   * [Introduction](#introduction)
   * [Prerequisite](#prerequisite)
   * [Installation](#installation)
   * [Setup](#setup)
   * [Usage](#usage)
   * [Examples](#examples)
        * [Create](#create)
        * [list](#list)
        * [ssh](#ssh)
        * [Up](#up)
        * [Stop](#stop)
        * [Status](#status)
        * [Restart](#restart)
        * [Destroy](#destroy)
        * [Update Configuration](#update-configuration)
        * [Delete Configuration](#delete-configuration)
   * [Demo (vm provisioning)](#demo-vm-provisioning)
   * [Developers / Contributors](#developers--contributors)

# Introduction

Datalab CLI is a wrapper to help you create and manage the vagrant VM with ease

# Prerequisite

Please follow the [instruction and prerequisite steps](https://github.com/pivotal-gss/go-gpdb#prerequisite) which are needed for the datalab cli to work.

# Installation

+ Git clone this repository to the location of your choice
    ```
    cd /Users/xxxx/Document
    git clone https://github.com/pivotal-gss/go-gpdb.git 
    ```
+ Download the latest version of the datalab CLI from the [release link](https://github.com/pivotal-gss/go-gpdb/releases).
+ Run the setup as mentioned on the [setup](#setup) section and you are good to go.

# Setup

Run the below steps to start using the datalabs

+ Copy the datalab CLI to the bin location so that it can be accessed from anywhere
    ```
    chmod +x <location-to-where-you-downloaded-datalab>/datalab
    cp <location-to-where-you-downloaded-datalab>/datalab /usr/local/bin/
    ```
+ Now run the below command to setup up the network API and Repository location
    ```
    datalab update-config -t <pivnet-token> -l <location-to-the-repository eg.s /Users/xxxx/Document/go-gpdb>
    ```
    To obtain the pivnet-token, Navigate to [PivNet Edit Profile](https://network.pivotal.io/users/dashboard/edit-profile) and scroll to the bottom of the page near “UAA API TOKEN” & click on the button “Request New API Token”, copy the token (**PLEASE NOTE:** This token will change if you click on the “Request New API Token” again)

# Usage

The usage information of datalab software

```
The program manages like create / destroy / stop / list and helps to login to vagrant installation

Usage:
  datalab [command] [flags]
  datalab [command]

Available Commands:
  create        Create the vagrant environment
  delete-config Delete the configuration from the datalab config file
  destroy       Destroy the vagrant environment
  help          Help about any command
  list          list all the configuration from the datalab config file
  restart       Restart / Reload of the vagrant environment
  ssh           SSH to the vagrant environment
  status        Status of the vagrant environment
  stop          Stop the vagrant environment
  up            Bring up the vagrant environment
  update-config Update the configuration of the datalab tool

Flags:
      --developer   Setup the virtual machine with developer tools to build the go-gpdb binaries
  -h, --help        help for datalab
  -v, --verbose     Enable verbose or debug logging
      --version     version for datalab

Use "datalab [command] --help" for more information about a command.
```

# Examples

### Create

+ To create a single node vagrant VM (the default vagrant name is "gpdb")
    ```
    datalab create
    ```
+ To create a single node vagrant VM with specific name
    ```
    datalab create -n <name>
    ```
+ To create a multi node vagrant VM with specific name
    ```
    datalab create -n <name> -s 2
    ```
+ To create a multi node vagrant VM & a standby host with specific name
    ```
    datalab create -n <name> -s 2 --standby
    ```
+ To create a multi node vagrant VM & a standby host with specific name and also configure cpu and memory ( i.e customize )
    ```
    datalab create -n <name> -s 2 -c 2 -m 2048 --standby
    ```
+ Run the below command to get all the help of create subcommand
    ```
    datalab help create
    ```

### List

+ To list all the VM installed
    ```
    datalab list
    ```
+ To list all the VM installed along with the vagrant global-status
    ```
    datalab list -g
    ```
+ Run the below command to get all the help of list subcommand
    ```
    datalab help list
    ```

### Ssh

+ To ssh to the default VM (i.e gpdb)
    ```
    datalab ssh
    ```
+ To ssh to a VM with non default name 
    ```
    datalab ssh -n <name>
    ```
+ Run the below command to get all the help of ssh subcommand
    ```
    datalab help ssh
    ```

### Up

+ To bring up the default VM (i.e gpdb)
    ```
    datalab up
    ```
+ To bring up a VM with non default name 
    ```
    datalab up -n <name>
    ```
+ Run the below command to get all the help of up subcommand
    ```
    datalab help up
    ```

### Stop

+ To stop the default VM (i.e gpdb)
    ```
    datalab stop
    ```
+ To stop a VM with non default name 
    ```
    datalab stop -n <name>
    ```
+ Run the below command to get all the help of stop subcommand
    ```
    datalab help stop
    ```

### Status

+ To get status of the default VM (i.e gpdb)
    ```
    datalab status
    ```
+ To get status of a VM with non default name 
    ```
    datalab status -n <name>
    ```
+ Run the below command to get all the help of status subcommand
    ```
    datalab help status
    ```

### Restart

+ To restart the default VM (i.e gpdb)
    ```
    datalab restart
    ```
+ To restart a VM with non default name 
    ```
    datalab restart -n <name>
    ```
+ Run the below command to get all the help of restart subcommand
    ```
    datalab help restart
    ```

### Destroy

+ To destroy a VM with a name 
    ```
    datalab destroy -n <name>
    ```
+ Run the below command to get all the help of destroy subcommand
    ```
    datalab help destroy
    ```

### Update Configuration

+ To update the datalab configuration of API token ( for eg.s if API token Change ) 
    ```
    datalab update-config -t <new-token>
    ```
+ To update the datalab configuration of go-gpdb repository 
    ```
    datalab update-config -l <new-location>
    ```
+ Run the below command to get all the help of update config subcommand
    ```
    datalab help update-config
    ```

### Delete Configuration

+ To delete the datalab configuration for example if VM is removed manually or some other reason 
    ```
    datalab delete-config -n <vm-name>
    ```
+ Run the below command to get all the help of delete config subcommand
    ```
    datalab help delete-config
    ```
# Demo (vm provisioning)

[![asciicast](https://asciinema.org/a/zCHod2oqujIBqR6dR5e5fa8Hd.svg)](https://asciinema.org/a/zCHod2oqujIBqR6dR5e5fa8Hd)

# Developers / Contributors

1. Clone the git repository
2. Export the GOPATH
    ```
    export GOPATH=<path to the clone repository>
    ```
3. cd to project folder
    ```
    cd $GOPATH/src/github.com/pivotal-gss/go-gpdb/datalab
    ```
4. Install all the dependencies. If you don't have dep installed, follow the instruction from [here](https://github.com/golang/dep)
    ```
    dep ensure
    ```
5. You are all set, you can run it locally using
    ```
    go run *.go <commands>
    ```
6. To build the package use
    ```
    go build -o datalab
    ```