package DGo

import "strconv"

func IntToStr(value int) string {
  return strconv.FormatInt(int64(value),10)
}

func StrToInt(value string) int32 {
  iVal, err := strconv.ParseInt(value,10,32)
  if (err != nil){}
  return int32(iVal)
}
