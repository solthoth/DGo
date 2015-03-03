package DGo

import "strconv"

func IntToStr(value int) string {
  return strconv.FormatInt(int64(value),10)
}
