#!/usr/bin/env bash

RESULTS=imgs

if [ ! -d $RESULTS ]; then
  echo $RESULTS
  mkdir -p $RESULTS
fi

for VAR in {0..255}
do
    ./blackcab -pop 400 -generations 400 -file "$RESULTS"/ca_rule_"$VAR".png -rule $VAR
done