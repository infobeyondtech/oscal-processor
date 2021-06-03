# Create Profile Test
#curl -i -H "Content-Type: application/json" \
#-X POST -d  \
#'{"controls": ["ac-1"], "baseline": "Fedramp", "catalogs": ["800-53"]}' \
#http://localhost:8080/profile/create

# Profile Navigator Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:8080/profile/navigator/e47dd5bd-ee87-433b-acb4-877877079ea9

# Resolve Profile Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:8080/profile/resolve/e47dd5bd-ee87-433b-acb4-877877079ea9

# Upload
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#    -d @
#http://localhost:8080/upload

# Control Test
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:8080/control/ac-1

# Enhancement Test
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:8080/control_enhancement/ac-2.2

# Get Param Test
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:8080/getparam/fileid1/paramid1

# Set Param Test
curl -i -H "Content-Type: application/json" \
    -X GET \
http://localhost:8080/setparam/fileid2/paramid2/value3
