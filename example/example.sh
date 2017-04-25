#!/bin/bash

base_url="http://someforum.org/forumdisplay.php?f=123"
pages=10
browser="firefox"

declare -a res=()

for pageno in $(seq 1 $pages); do
     url="$base_url&page=$pageno"
     echo "$url..."

     declare -a out=$(./vbmatch -forum-url "$url")
     res+=($out)
done

if [ ${#res[@]} -gt 0 ]; then
    $browser ${res[@]}
fi

