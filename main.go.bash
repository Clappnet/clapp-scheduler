#!/bin/bash
export DSN=postgres://postgres:postgres@127.0.0.1:5432/scheduler?sslmode=disable
/usr/bin/go run /opt/clapp-scheduler/main.go &
