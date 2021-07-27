# Create Profile Test
#curl -i -H "Content-Type: application/json" \
#-X POST -d  \
#'{
#    "baseline": "Fedramp",
#    "controls": ["ac-1"],
#    "catalogs": ["800-53"],
#    "title": "test_title",
#    "orgUuid": "test_orgUuid",
#    "orgName": "test_orgName",
#    "orgEmail": "test_orgEmail"
#}' \
#http://localhost:8080/profile/create

# Profile Navigator Test
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:9050/profile/navigator/testprofile.xml

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
#curl -i -H "Content-Type: application/json"      -X GET \
#http://localhost:9050/control/ac-7

# Enhancement Test
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:9050/control_enhancement/ac-2.2

# Get Param Test
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:8080/getparam/fileid1/paramid1

# Set Param Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:9050/setparam/fileid2/paramid2/value5

# Get Param Info
curl -i -H "Content-Type: application/json" \
    -X GET \
http://localhost:9050/getparaminfo/ac-7_prm_3

# Get Param Info
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:9050/getparam/1

# Update Param Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:9050/setparam/1/boom

# Create Param Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:9050/createparam/4/component_id_2/param_id_2/value_2

# Delete Param Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:9050/deleteparam/28

# Get Param 
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:9050/getparam/1

# Get Params
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:9050/get-params/1

# Create Component Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:9050/createcomponent/23/statement_id_23/component_id_23

# Update Component Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:9050/setcomponent/1/updated_component_id

# Delete Component Test
#curl -i -H "Content-Type: application/json" \
#    -X POST \
#http://localhost:9050/deletecomponent/2

# Get Component Value
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:9050/getcomponent/1

# Get Components
#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:9050/get-components/23

#curl -i -H "Content-Type: application/json" \
#    -X GET \
#http://localhost:9050/information/get-component/795533ab-9427-4abe-820f-0b571bacfe6d




