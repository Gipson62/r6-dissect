import json

with open('match_dbno.json') as f:
    data = json.load(f)

for r in data['rounds']:
    rnd = r['roundNumber']
    if rnd in [11, 12, 13]:
        print(f"\n=== Round {rnd} ===")
        for e in r['matchFeedback']:
            if e['type']['name'] == 'Kill':
                hs = e.get('headshot', False)
                print(f"  {e['username']} -> {e.get('target','')} {'(HS)' if hs else ''} @ {e['time']}")
