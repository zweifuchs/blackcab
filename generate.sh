#!/usr/bin/env bash

$RESULTS = "results"

if [ ! -d "$RESULTS" ]; then
  mkdir -p "$RESUTS"
fi

for VAR in {0..255}
do
    ./ca5 -pop 401 -generations 401 -file imgs/ca_rule_"$VAR"_pop401.png -rule $VAR
done