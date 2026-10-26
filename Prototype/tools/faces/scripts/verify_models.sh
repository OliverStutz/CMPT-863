#!/usr/bin/env bash
# Verify every downloaded model file against the publisher's SHA-256 (from the Hugging Face LFS metadata)
# and check that each safetensors header parses and matches the file size.
set -u
LAB=/home/oliver/.cache/cmpt863-faces; ok=0; bad=0
while read -r sha size name; do [ -z "$sha" ] && continue;
  f=$(find "$LAB/models" -name "$name" | head -1)
  if [ -z "$f" ]; then echo "MISSING $name"; bad=$((bad+1)); continue; fi
  actual_size=$(stat -c %s "$f")
  if [ "$actual_size" != "$size" ]; then echo "SIZE MISMATCH $name: $actual_size vs $size"; bad=$((bad+1)); continue; fi
  actual=$(sha256sum "$f" | cut -d' ' -f1)
  if [ "$actual" != "$sha" ]; then echo "SHA MISMATCH $name"; bad=$((bad+1)); continue; fi
  hdr=$(python3 - "$f" <<'PY'
import sys, json, struct
f = open(sys.argv[1], 'rb'); n = struct.unpack('<Q', f.read(8))[0]
h = json.loads(f.read(n)); import os
end = max((v['data_offsets'][1] for k, v in h.items() if k != '__metadata__'), default=0)
print('ok' if 8 + n + end == os.path.getsize(sys.argv[1]) else f'header/size mismatch {8+n+end} vs {os.path.getsize(sys.argv[1])}', len(h), 'tensors')
PY
)
  echo "OK   $name  sha256 matches, $(( size / 1048576 )) MiB, header $hdr"; ok=$((ok+1))
done < "${MODELS_SHA256:-$(dirname "$0")/../models.sha256}"
echo "verified=$ok failed=$bad"; [ "$bad" = 0 ]
