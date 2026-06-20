#!/usr/bin/env python3
import json
import re
from datetime import datetime, timezone
from pathlib import Path

root = Path(__file__).resolve().parent.parent
result_file = root / 'test-results.json'
stats_file = root / 'docs-site' / 'static' / 'project-stats.json'

if not result_file.exists():
    raise SystemExit('test-results.json not found; run `make test` first')

with result_file.open() as f:
    results = [json.loads(line) for line in f if line.strip()]

package_coverages = []
packages_tested = 0
packages_passed = 0
packages_failed = 0
packages_skipped = 0

for entry in results:
    pkg = entry.get('Package', '')
    if not pkg:
        continue

    action = entry.get('Action')
    if action == 'pass':
        packages_passed += 1
    elif action == 'fail':
        packages_failed += 1
    elif action == 'skip':
        packages_skipped += 1

    if action in {'pass', 'fail', 'skip'}:
        packages_tested += 1

    output = entry.get('Output', '')
    if 'coverage:' in output:
        match = re.search(r'coverage: ([0-9.]+%)', output)
        if match:
            package_coverages.append({
                'package': pkg,
                'coverage': match.group(1),
            })

stats_file.parent.mkdir(parents=True, exist_ok=True)
with stats_file.open('w') as f:
    json.dump({
        'packagesTested': packages_tested,
        'packagesPassed': packages_passed,
        'packagesFailed': packages_failed,
        'packagesSkipped': packages_skipped,
        'coverageReportUrl': '/coverage/index.html',
        'rawTestResultsUrl': '/test-results.json',
        'generatedAt': datetime.now(timezone.utc).isoformat(),
        'packageCoverages': package_coverages,
    }, f, indent=2)
