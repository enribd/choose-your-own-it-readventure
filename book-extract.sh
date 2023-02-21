#!/bin/bash

for row in "$(grep img book-index.md)"; do

  echo "$row"
  COVER=$(echo $row | awk -F'|' '{print $2}' | awk -F '/' '{print $NF}' | tr -d ')')
  TITLE=$(echo $row | awk -F'|' '{print $3}' | grep -Eo '\*\*.*\*\*' | cut -d ':' -f1 | tr -d '*')
  SUBTITLE=$(echo $row | awk -F'|' '{print $3}' | grep -Eo '\*\*.*\*\*' | cut -d ':' -f2 | tr -d '*')
  URL=$(echo $row | awk -F'|' '{print $3}' | grep -Eo '\(.*\)' | tr -d '('| tr -d ')')
  AUTHORS=$(echo $row | awk -F'|' '{print $3}' | grep -Eo '\<br\>.*\<br\>' | tr -d '*')
  RELEASE=$(echo $row | awk -F'|' '{print $3}' | grep -Eo 'Published in 2[0-9][0-9][0-9]' | sed 's/Published in //g')
  PAGES=$(echo $row | awk -F'|' '{print $3}' | grep -Eo '[0-9][0-9][0-9] pages' | sed 's/ pages//g')
  LPS=$(echo $row | awk -F'|' '{print $4}' | sed -e 's/<ul>//g' -e 's/<li>//g' -e 's/<\/ul>//g' -e 's/<\/li>/,/g')

  envsubst << EOF >> books-extracted.yaml
  - title: $TITLE
    subtitle: $SUBTITLE
    cover: $COVER
    order: TODO
    draft: true
    url: $URL
    authors:
      - $AUTHORS
    release: $RELEASE
    pages: $PAGES
    desc: |-
      TODO
    learning_paths:
      - $LPS
    badges:
      - TODO
EOF
break
done
