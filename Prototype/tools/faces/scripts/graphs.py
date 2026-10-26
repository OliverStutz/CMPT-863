"""API-format graphs for ComfyUI: a Flux schnell portrait and the Wan 2.2 S2V talking head.
The S2V graph mirrors the official template video_wan2_2_14B_s2v.json (Comfy-Org/workflow_templates):
first chunk WanSoundImageToVideo -> KSampler, then N-1 WanSoundImageToVideoExtend -> KSampler -> LatentConcat,
first latent frame doubled before decoding (template's VAE first-frame fix) and dropped after, CreateVideo at 16 fps."""
import json, os

NEG = json.load(open(os.path.join(os.path.dirname(os.path.abspath(__file__)), "negative.json")))["negative"]

def portrait(prompt, seed, width=768, height=768, batch=4, prefix="portrait/ciso"):
    return {
        "1": {"class_type": "CheckpointLoaderSimple", "inputs": {"ckpt_name": "flux1-schnell-fp8.safetensors"}},
        "2": {"class_type": "CLIPTextEncode", "inputs": {"text": prompt, "clip": ["1", 1]}},
        "3": {"class_type": "CLIPTextEncode", "inputs": {"text": "", "clip": ["1", 1]}},
        "4": {"class_type": "EmptySD3LatentImage", "inputs": {"width": width, "height": height, "batch_size": batch}},
        "5": {"class_type": "KSampler", "inputs": {"model": ["1", 0], "seed": seed, "steps": 4, "cfg": 1.0,
               "sampler_name": "euler", "scheduler": "simple", "positive": ["2", 0], "negative": ["3", 0],
               "latent_image": ["4", 0], "denoise": 1.0}},
        "6": {"class_type": "VAEDecode", "inputs": {"samples": ["5", 0], "vae": ["1", 2]}},
        "7": {"class_type": "SaveImage", "inputs": {"images": ["6", 0], "filename_prefix": prefix}},
    }

def s2v(audio, image, prompt, chunks, steps=4, cfg=1.0, lora=True, seed=7, width=640, height=640,
        length=77, prefix="s2v/clip"):
    g = {
        "unet": {"class_type": "UNETLoader", "inputs": {"unet_name": "wan2.2_s2v_14B_fp8_scaled.safetensors", "weight_dtype": "default"}},
        "clip": {"class_type": "CLIPLoader", "inputs": {"clip_name": "umt5_xxl_fp8_e4m3fn_scaled.safetensors", "type": "wan", "device": "default"}},
        "vae": {"class_type": "VAELoader", "inputs": {"vae_name": "wan_2.1_vae.safetensors"}},
        "aenc": {"class_type": "AudioEncoderLoader", "inputs": {"audio_encoder_name": "wav2vec2_large_english_fp16.safetensors"}},
        "audio": {"class_type": "LoadAudio", "inputs": {"audio": audio}},
        "aemb": {"class_type": "AudioEncoderEncode", "inputs": {"audio_encoder": ["aenc", 0], "audio": ["audio", 0]}},
        "img": {"class_type": "LoadImage", "inputs": {"image": image}},
        "pos": {"class_type": "CLIPTextEncode", "inputs": {"text": prompt, "clip": ["clip", 0]}},
        "neg": {"class_type": "CLIPTextEncode", "inputs": {"text": NEG, "clip": ["clip", 0]}},
    }
    model_src = ["unet", 0]
    if lora:
        g["lora"] = {"class_type": "LoraLoaderModelOnly", "inputs": {"model": model_src,
                     "lora_name": "wan2.2_t2v_lightx2v_4steps_lora_v1.1_high_noise.safetensors", "strength_model": 1.0}}
        model_src = ["lora", 0]
    g["ms"] = {"class_type": "ModelSamplingSD3", "inputs": {"model": model_src, "shift": 8.0}}
    model = ["ms", 0]
    g["c0"] = {"class_type": "WanSoundImageToVideo", "inputs": {"positive": ["pos", 0], "negative": ["neg", 0], "vae": ["vae", 0],
               "width": width, "height": height, "length": length, "batch_size": 1,
               "audio_encoder_output": ["aemb", 0], "ref_image": ["img", 0]}}
    g["k0"] = {"class_type": "KSampler", "inputs": {"model": model, "seed": seed, "steps": steps, "cfg": cfg,
               "sampler_name": "uni_pc", "scheduler": "simple", "positive": ["c0", 0], "negative": ["c0", 1],
               "latent_image": ["c0", 2], "denoise": 1.0}}
    acc = ["k0", 0]
    for i in range(1, chunks):
        g[f"c{i}"] = {"class_type": "WanSoundImageToVideoExtend", "inputs": {"positive": ["pos", 0], "negative": ["neg", 0],
                      "vae": ["vae", 0], "length": length, "video_latent": acc, "audio_encoder_output": ["aemb", 0],
                      "ref_image": ["img", 0]}}
        g[f"k{i}"] = {"class_type": "KSampler", "inputs": {"model": model, "seed": seed + i, "steps": steps, "cfg": cfg,
                      "sampler_name": "uni_pc", "scheduler": "simple", "positive": [f"c{i}", 0], "negative": [f"c{i}", 1],
                      "latent_image": [f"c{i}", 2], "denoise": 1.0}}
        g[f"cat{i}"] = {"class_type": "LatentConcat", "inputs": {"samples1": acc, "samples2": [f"k{i}", 0], "dim": "t"}}
        acc = [f"cat{i}", 0]
    g["cut"] = {"class_type": "LatentCut", "inputs": {"samples": acc, "dim": "t", "index": 0, "amount": 1}}
    g["dbl"] = {"class_type": "LatentConcat", "inputs": {"samples1": ["cut", 0], "samples2": acc, "dim": "t"}}
    g["dec"] = {"class_type": "VAEDecode", "inputs": {"samples": ["dbl", 0], "vae": ["vae", 0]}}
    g["trim"] = {"class_type": "ImageFromBatch", "inputs": {"image": ["dec", 0], "batch_index": 4, "length": 4096}}
    g["vid"] = {"class_type": "CreateVideo", "inputs": {"images": ["trim", 0], "fps": 16.0, "audio": ["audio", 0]}}
    g["save"] = {"class_type": "SaveVideo", "inputs": {"video": ["vid", 0], "filename_prefix": prefix, "format": "auto", "codec": "auto"}}
    return g
