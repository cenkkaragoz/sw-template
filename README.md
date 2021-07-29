
## swagger 


orginal post ==> https://github.com/cirocosta/hello-swagger/

go get -u -v github.com/go-swagger/go-swagger/cmd/swagger

depencies 
go get github.com/alexflint/go-arg



# 
$ make
--> install $GOPATH/bin/sw-template  

$ sw-template --help


## Define the API and generate the code
.
├── Makefile            # makefile that summarizes the
│                       # build and code generation procedures
│                    
├── main.go             # our application entrypoint
│       
├── swagger             # swagger definition and generated code
│   │       
│   └── swagger.yml
└── vendor              # dependencies
    ├── github.com
    └── ...