# Klippa assessment
This project is commissioned by Klippa as a internship assessment. The basic goals are to have cli application that can receive a few flags from the commandline and based on those send a request to their own API.

## The assignment
Create a CLI tool that allows you to call our OCR API. <br>
The CLI tool must be able to do the following: <br>
 - Provide the API key as an option
 - Provide the template as an option
 - Provide PDF text extraction fast or full as an option.
 - Provide a file that it processes via the OCR API (PDF or image).
 - Display the result of the processing in a nicely formatted way.
 - Include the option to save the json output to the file as {file name}.json

Bonus points for: <br>
- Being able to process a folder instead of 1 file and being able to batch process the entire folder
 - For making it run in Docker
- Keep track of the totals of the folder processing, such as VAT percentages and total amount and display that after processing
- Monitor a folder, so when a new file is added to the folder it will be processed automatically.
- Even more bonus points if you manage to do the above things concurrently, so if you process a whole folder you process several at once
- Provide options for API key, template and PDF text extraction through a global config file (e.g. like Docker and Git: ~/.docker/config.json and ~/.gitconfig)

## How to build
go build cmd/klippa-assessment/main.go

## How to use

### flags:
'-api=' This is where you put your api key of the klippa api. no default value <br>
'-template=' The template of the document you want to scan. the templates your api key has access to can be found using the /templates route. default: 'financial_full' <br>
'-textextractiontype=' which extraction type do you want to use, there are 2 options, fast and full. default: fast <br>
'-file=' which file do you want to get scanned. no default value <br>
'-save=' if you want to save the json response to a file specify a name or a path with a name. no default value <br>
'-debug=' enable the debug mode, this is for debugging the output of the program, not(yet?) for the api functionality. its a boolean, so either true of false as value. default value false <br>
'-fulloutput=' get the full output of the api response, even the values that have no value will get printed. boolean type, so either true or false as values. default value is false. <br>

#### cli exmples:
./bin/main -api=[api-key] -file=some/path/to/file.pdf -template=financial_full -save=save -fulloutput=true <br>
./bin/main -debug=true

## docker:
As seen in the repo, there is a Dockerfile. The cli version does work in there, but due to not getting volumes/mounts working to my satisfaction i decided to not have it documented.