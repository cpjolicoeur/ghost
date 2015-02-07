# Ghost

![Ghost](http://pixabay.com/static/uploads/photo/2013/07/13/10/18/ghost-156969_640.png)


## Installation

    $ go get github.com/cpjolicoeur/ghost


## Usage

    $ ghost
    
## Options

### Port
    
`ghost` defaults to running on port 8000, but the port can be configured a few different ways.

    $ ghost --port 1234
    $ PORT=1234 ghost
    $ GHOST_PORT=1234 ghost
    
### Directory

`ghost` defaults to serving up the current working directory, but the directory can be configured via a flag or ENV var

    $ ghost --dir /my/dir/to/serve
    $ GHOST_DIR=/my/dir/to/serve ghost
