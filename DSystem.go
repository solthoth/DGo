package DGo

import(
  "os"
)

func AssignFile(f *os.File, filename string) error{
  f, err := os.Open(filename)
  return err
}

func CloseFile(f *os.File) {
  f.Close()
}

func ReadBuffer(f *os.File, count int) string {
  buffer := make([]byte, count)
  f.Read(buffer)
  return string(buffer)
}

func ReadLn(f *os.File, value *string) string {
  buffer := ""
  buff   := ""
  for ; buff !="\n"; {
    buff = ReadBuffer(f,1)
    buffer = buffer + buff
  }
  *value = buffer
  return buffer
}

func Reset(f *os.File){
  f.Seek(0,0)
}
