#!/bin/bash

dist=/home/grevin/dist
mkdir -p $dist
rsync -av index.html *.js img $dist
