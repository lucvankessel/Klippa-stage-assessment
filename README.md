# Klippa assessment

## How to build
go build cmd/klippa-assessment/main.go

## How to use

### flags:
'-api=' This is where you put your api key of the klippa api. no default value
'-template=' The template of the document you want to scan. the templates your api key has access to can be found using the /templates route. default: 'financial_full'
'-textextractiontype=' which extraction type do you want to use, there are 2 options, fast and full. default: fast
'-file=' which file do you want to get scanned. no default value
'-save=' if you want to save the json response to a file specify a name or a path with a name. no default value
'-debug=' enable the debug mode, this is for debugging the output of the program, not(yet?) for the api functionality. its a boolean, so either true of false as value. default value false
'-fulloutput=' get the full output of the api response, even the values that have no value will get printed. boolean type, so either true or false as values. default value is false.

#### cli exmples:
./main -api=[api-key] -file=testDocs/testDocPDF.pdf -template=financial_full -save=save
./main -debug=true

#### docker
to build the image:
docker build -t klippa-assessment .

to run the image:
docker run klippa-assessment {any flags you want to add}

for example: docker run klippa-assessment -debug=true

