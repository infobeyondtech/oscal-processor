# oscal-processor

# create a profile

curl -i \
    -H "Content-Type: application/json" \
    -X POST -d  \
    '{"controls": ["ac-1"], "baseline": "Fedramp", "catalogs": ["800-53"]}'
    http://localhost:8080/profile/create

# test
  - change to the directory where the test files are
  - Test
    > go test -v

# check download or upload files
   - go to /home/infobeyondtech7/oscal_processing_space/download

# deploy 

  - Find any running instance 
    
    > ps -eaf | grep oscal-processor
    > kill xxx

  - change to the directory ~/go/src/github.com/infobeyondtech/oscal-processor/host
  - Build 
    > go build -o oscal-processor-engine main.go
  - Run 
    > nohup ./oscal-processor-engine &
  
# docker 

  - Build 
    > docker build -t main .
  - Run 
    > docker run -dp 9050:9050 main .

