# oscal-processor

# create a profile

curl -i \
    -H "Content-Type: application/json" \
    -X POST -d  \
    '{"controls": ["ac-1"], "baseline": "Fedramp", "catalogs", ["800-53"]}'
    http://localhost:8080/profile/create
