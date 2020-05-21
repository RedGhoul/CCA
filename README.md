##### NAME:
   Code Compare Analysis (CCA) - Compare code files & calculate metrics. Useful when trying to merge two code bases

##### USAGE:
   CCA.exe [global options] command [command options] [arguments...]

##### COMMANDS:
   diff, d   Finds differences between files in two folders.
                         First argument is the first folder {Full Path}
                         Second argument is the second folder {Full Path}
   patch, p  Creates a patch between two files.
                       First argument is the first file {Full Path}
                       Second argument is the second file {Full Path}
   help, h   Shows a list of commands or help for one command

##### GLOBAL OPTIONS:
   --filetype value      --filetype *. Used to filter by file type (default: "*")
   --graphtitle value    --graphtitle string. Sets the graph title (default: "Default Title")
   --outputfolder value  --outputfolder string. Sets the output location of chart & report
   --chart               --chart true. Used to create a chart as html (default: true)
   --report              --report true. Used to create a txt report (default: true)
   --help, -h            show help (default: false)
