# oscal-processor

# create a profile

curl -i \
    -H "Content-Type: application/json" \
    -X POST -d  \
    '{"controls": ["ac-1"], "baseline": "Fedramp", "catalogs", ["800-53"]}'
    http://localhost:8080/profile/create

# deploy 

  - Find any running instance 
    
    > ps -eaf | grep oscal-processor
    > kill xxx

  - change to the directory oscal-processor
  - Build 
    > go build -o oscal-processor-engine main.go
  - Run 
    > nohup ./oscal-processor-engine &
  
