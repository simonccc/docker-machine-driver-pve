#!/bin/bash

# Node dependencies
yarn licenses generate-disclaimer > ./licenses.txt

# Go dependencies
rm -rf ./tmp/licenses
mkdir -p ./tmp/licenses
./tools/go-licenses save ./... --include_tests --force --save_path ./tmp/licenses

for FILE_PATH in $(find "./tmp/licenses" -type f | LC_ALL=C sort); do
  echo -e "-----\n\nThe following software may be included in this product: $(dirname ${FILE_PATH#./tmp/licenses/}). This software contains the following license and notice below:\n" >> ./licenses.txt \
  && while read -r LINE; do echo "$LINE" >> ./licenses.txt; done < $FILE_PATH \
  && echo -e "" >> ./licenses.txt
done
