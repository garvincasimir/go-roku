# go-roku

The **rokuclient** package is an implementation of the protocol outlined in the [Roku External Control Guide](https://sdkdocs.roku.com/display/sdkdoc/External+Control+Guide). The **rokuremote** console app uses the library to discovery and control roku devices on the local network. This code was ported over from the [ .net core version](https://github.com/garvincasimir/roku-remote) 

## Testing
Until I find a better way, testing the discovery function requires being on a network with an actual roku device.

## Building
There are no special requirements for building. Simply *go build*

## Contributing
Go masters/novices out there please feel free to send a pull request. 

## 3 things I love about #GO#
* The approach to concurrency 
* The general purpose template package
* go fmt
