#!/usr/bin/env bash
# Downloads the model files listed in ../models.txt into $FACES_LAB/models/<dir>/ (resumable, three at a time)
# and verifies every file against ../models.sha256 (SHA-256 from the Hugging Face LFS metadata).
set -u
LAB=${FACES_LAB:-/home/oliver/.cache/cmpt863-faces}; HERE=$(dirname "$(readlink -f "$0")"); export M="$LAB/models"
mkdir -p "$M"/{diffusion_models,text_encoders,vae,audio_encoders,loras,checkpoints}
xargs -P 3 -L 1 bash -c 'dir=$0; url=$1; f="$M/$dir/$(basename "$url")"; curl -sfL --retry 5 --retry-delay 5 -C - -o "$f" "$url" && echo "OK   $f" || echo "FAIL $url"' < "$HERE/../models.txt"
MODELS_SHA256="$HERE/../models.sha256" FACES_LAB="$LAB" bash "$HERE/verify_models.sh"
