#!/bin/bash

# map-stream@0.1.0 -> MIT according to https://www.npmjs.com/package/map-stream/v/0.1.0

yarn exec -- license-checker \
  --excludePackages "map-stream@0.1.0" \
  --onlyAllow "Unlicense;MIT;BSD;ISC;Apache-2.0;OFL-1.1;Python-2.0;CC-BY-3.0;CC0-1.0;CC-BY-4.0;BlueOak-1.0.0;MPL-2.0;Artistic-2.0" \
  --summary
