#!/usr/bin/env bash
cat "$1" \
| pup -p -i 0 'dl.mgtM table.tblNfud[summary*="주요중심가"]' \
| awk '{ printf "%s", $0 }' \
| sed -E "s/ scope=\"row\">/><p>=====<\/p>/g;s/<br>/ /g" \
| pup -p 'tbody tr text{}'

cat "$1" \
| pup -p -i 0 'dl.mgtM table.tblNfud[summary*="KTX"]' \
| awk '{ printf "%s", $0 }' \
| sed -E "s/<strong>([^\s]+)역<\/strong>/<p>=====<\/p><strong>\1역<\/strong>/g;s/<br>/ /g" \
| pup -p 'tbody tr text{}'