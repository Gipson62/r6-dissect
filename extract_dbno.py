import json

with open('match_dbno_fixed.json') as f:
    data = json.load(f)

for r in data['rounds']:
    rnd = r['roundNumber']
    if rnd in [11, 12, 13]:
        print(f"\n=== Round {rnd} ===")
        for e in r['matchFeedback']:
            if e['type']['name'] in ['Kill', 'DBNO']:
                hs = e.get('headshot', False)
                dbno_by = e.get('dbnoBy', '')
                event_type = e['type']['name']
                line = f"  [{event_type}] {e['username']} -> {e.get('target','')} @ {e['time']}"
                if hs:
                    line += ' (HS)'
                if dbno_by:
                    line += f' [knocked by {dbno_by}]'
                print(line)
