package main

import "strings"
import "strconv"
// import "fmt"
import "os"
import "flag"
import "github.com/utamaro/rsraid"

func main(){
  mPtr := flag.Int("m", 1, "Number of redundant files (coding devices)")
  nPtr := flag.Int("n", 2, "Number of files to split across (data devices)")

  // encodeNamePtr := flag.String("encode", "", "Encode (Split)")
  encodePtr := flag.Bool("e", false, "Encode (Split)")

  decodePtr := flag.Bool("d", false, "Decode (Combine)")
  outputPtr := flag.String("o", "", "Output filename (for decoding)")
  sizePtr := flag.Int("s", 0, "Size of output filename (for decoding)")

  flag.Parse()
  
  if *encodePtr && *decodePtr {
    panic("You can't encode and decode")
  }

  if ! (*encodePtr || *decodePtr) {
    panic("You must encode or decode")
  }

  if *decodePtr && *outputPtr == "" {
    panic("You must specify an output file")
  }

  if *encodePtr{
    rsraid.EncodeFile(*mPtr, *nPtr, flag.Args()[0])
  } else {
    fileNames := make([]string, *mPtr + *nPtr)
    for _,i := range flag.Args() {
      if i != "" {
        file_parts := strings.Split(i, "_")
        file_part := file_parts[len(file_parts)-1]
        last_parts := strings.Split(file_part, ".")
        num, _ := strconv.ParseInt(last_parts[0], 10, 64)
        fileNames[num] = i
      }
    }

    var size int64

    // if the size isn't specified
    if *sizePtr == 0 {
      // GUESS! You have a 1 in n chance of being right! :D I thought about doing
      // All n possibilities, but the utamaro/rsraid is too bugged to care
      s, _ := os.Stat(flag.Args()[0])
      size = int64(*nPtr) * s.Size()
    } else {
      size = int64(*sizePtr)
    }

    rsraid.DecodeFile(size, *nPtr, fileNames, *outputPtr)
  }
}