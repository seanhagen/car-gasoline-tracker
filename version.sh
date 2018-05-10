#!/bin/bash

if [ -z "$SOURCE_VERSION" ]; then echo `git rev-parse HEAD`; else echo "$SOURCE_VERSION"; fi
