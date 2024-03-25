#!/bin/sh

aqua -c /app/aqua.yaml i \
&& mage -d /app/server dev \
& sleep infinity
