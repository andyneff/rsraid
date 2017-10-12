package main

// import "fmt"
import "flag"
import "github.com/utamaro/rsraid"

func main(){
  mPtr := flag.Int("m", 1, "Number of redundant files (coding devices)")
  nPtr := flag.Int("n", 2, "Number of files to split across (data devices)")

  // encodeNamePtr := flag.String("encode", "", "Encode (Split)")
  encodePtr := flag.Bool("encode", false, "Encode (Split)")
  decodePtr := flag.Bool("decode", false, "Decode (Combine)")

  outputPtr := flag.String("out", "", "Output filename")
  output := *outputPtr

  flag.Parse()
  
  if *encodePtr && *decodePtr {
    panic("You can't encode and decode")
  }

  if ! (*encodePtr || *decodePtr) {
    panic("You must encode or decode")
  }

  if *decodePtr && output == "" {
    panic("You must specify an output file")
  }

  if *encodePtr{
    rsraid.EncodeFile(*mPtr, *nPtr, flag.Args()[0])
  } else {
    panic("The encoding wasn't what I wanted. So I moved on")
    //fnames := flag.Args()
    //s:=int64(43545234) //size of test.dat
    //rsraid.DecodeFile(s, *nPtr, fnames, *outputPtr)
  }
}