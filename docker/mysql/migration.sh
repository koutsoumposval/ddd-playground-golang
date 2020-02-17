#!/bin/bash
set -e
service mysql start
mysql < /mysql/schema.sql
service mysql stop