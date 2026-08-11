#!/bin/bash

set -e


for f in /static/img/*.jpg; do
    echo "Processing $f"
    magick "$f" -resize 400x "$f"
done

echo "Done"