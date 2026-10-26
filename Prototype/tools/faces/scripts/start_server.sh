#!/usr/bin/env bash
# Start ComfyUI headless on localhost, pinned to the first R9700, with the integrated GPU hidden,
# cloud nodes off, and all inputs and outputs inside the lab folder.
LAB=${FACES_LAB:-/home/oliver/.cache/cmpt863-faces}
export ROCR_VISIBLE_DEVICES=${ROCR_VISIBLE_DEVICES:-0}
export HF_HUB_OFFLINE=1 TRANSFORMERS_OFFLINE=1
cd "$LAB/ComfyUI"
exec "$LAB/venv/bin/python" main.py --listen 127.0.0.1 --port "${PORT:-8188}" --offline --disable-auto-launch --cache-none \
  --extra-model-paths-config "$LAB/extra_model_paths.yaml" \
  --output-directory "$LAB/out" --input-directory "$LAB/inputs"
