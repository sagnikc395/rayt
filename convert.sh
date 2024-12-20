#!/bin/bash
for i in *.ppm;
  do name=`echo "$i" | cut -d'.' -f1`
  echo "$name"
  ffmpeg -i "$i" "${name}.png"
done