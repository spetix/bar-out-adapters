#!/usr/bin/env python3
import json
import os
from datetime import datetime
from pathlib import Path

root = Path(__file__).resolve().parent.parent
versions_file = root / 'docs-site' / 'static' / 'versions.json'

is_tag = os.environ.get('IS_TAG', 'false').lower() == 'true'
version = os.environ.get('VERSION', '')

if not versions_file.exists():
    data = {'versions': []}
else:
    with versions_file.open() as f:
        data = json.load(f)

if is_tag and version:
    if not any(v.get('version') == version for v in data.get('versions', [])):
        data.setdefault('versions', []).insert(0, {
            'version': version,
            'date': datetime.now().strftime('%Y-%m-%d'),
            'current': False,
        })
else:
    for v in data.get('versions', []):
        v['current'] = False

versions_file.parent.mkdir(parents=True, exist_ok=True)
with versions_file.open('w') as f:
    json.dump(data, f, indent=2)
