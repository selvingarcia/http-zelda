#!/bin/bash

set -e


for f in *.jpg; do
    echo "Processing $f"
    magick "$f" -resize 400x "$f"
done

echo "Done"