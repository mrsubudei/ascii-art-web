# ASCII ART WEB

### DESCRIPTION    
The program allows users to draw ASCII Art from the input and save result as txt file.  
Users are supposed to input only ASCII characters.  
Users are given three options of banner type i.e. *standard*, *shadow*, *thinkertoy*  
  
### USAGE  
1. Run the program
```
go run cmd/main.go
```
2. Open http://localhost:8087/ in a browser
3. Type any text in english layout
4. Choose the preferred banner
5. To save file click on export button

_________________________________________________

To create docker image:
```
make build
```
To run image
```
make run
```
To stop container
```
make stop
```
To remove image:
```
make remove
```
Stops and removes container:
```
make kill
```

### AUTHORS
@Subudei
