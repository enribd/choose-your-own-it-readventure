#!/bin/bash

rm -rf /book-extract.yaml
rm -rf /tmp/books
grep img book-index.md > /tmp/books

while IFS= read -r line; do
  # echo "$line"
  COVER=$(echo $line | awk -F'|' '{print $2}' | awk -F '/' '{print $NF}' | tr -d ')')
  TITLE=$(echo $line | awk -F'|' '{print $3}' | grep -Eo '\*\*.*\*\*' | cut -d ':' -f1 | tr -d '*')
  SUBTITLE=$(echo $line | awk -F'|' '{print $3}' | grep -Eo '\*\*.*\*\*' | cut -d ':' -f2 | tr -d '*')
  URL=$(echo $line | awk -F'|' '{print $3}' | grep -Eo '\(.*\)' | tr -d '('| tr -d ')')
  AUTHORS=$(echo $line | awk -F'|' '{print $3}' | grep -Eo '\<br\>.*\<br\>' |  sed 's/Published.*//g' | tr -d '*' | sed -e 's/br>//g' | tr -d '<')
  RELEASE=$(echo $line | awk -F'|' '{print $3}' | grep -Eo 'Published in 2[0-9][0-9][0-9]' | sed 's/Published in //g')
  PAGES=$(echo $line | awk -F'|' '{print $3}' | grep -Eo '[0-9][0-9][0-9] pages' | sed 's/ pages//g')
  LPS=$(echo $line | awk -F'|' '{print $4}' | sed -e 's/<ul>//g' -e 's/<li>/|/g' -e 's/<\/ul>//g' -e 's/<\/li>//g' | awk '{print tolower($0)}' | tr -s ' ' | sed -e 's/|/#      - /g' | tr '#' '\n' | grep -v '  architecture'  | awk 'NF')

  envsubst << EOF >> book-extract.yaml
  - title: $TITLE
    subtitle: $SUBTITLE
    cover: $COVER
    order: TODO
    draft: true
    url: $URL
    authors:$AUTHORS
    release: $RELEASE
    pages: $PAGES
    desc: |-
      TODO
    learning_paths:
$LPS
    badges:
      - TODO
EOF
done < /tmp/books

sed -i 's/:  /: /g' book-extract.yaml
sed -i 's/- system design/- system-design/g' book-extract.yaml
sed -i 's/- software architecture/- software-architecture/g' book-extract.yaml
sed -i 's/- event driven architecture/- event-driven-architecture/g' book-extract.yaml
sed -i 's/- software development/- software-development/g' book-extract.yaml
sed -i 's/- software delivery/- software-delivery/g' book-extract.yaml
sed -i 's/- architecture//g' book-extract.yaml
sed -i 's/- api design/- api-design/g' book-extract.yaml
sed -i 's/- organization design/- organization-design/g' book-extract.yaml
sed -i 's/- product management/- product-management/g' book-extract.yaml
sed -i '/^[[:blank:]]*$/ d' book-extract.yaml
