package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.4" )

func jw1R3kNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YefrKuenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m2NOZemkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qN0sqYq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TUOcG79aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 46mMV3bNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rm9KhBltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lVElvVTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QSTfUooxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dp1xAeglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a4CBs3YLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDgjHLv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DbdTdEotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9gtv0ZbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oimt5X3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func akIrrj1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYJvaY9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2FplhDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qVRDnCL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tGdCBMD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0gHEFEZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hCfpBQ7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QnMcLRMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gO1AQ7SpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dKjG4xC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func evseqCF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FEYUfaogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FJOXiLt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func efrHdERMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BE3zCjGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0e9bOgC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zvow8iRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func niKND78mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0vxGj9UYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iByXTsnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eGal77puWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vs9LBR1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nXLQY3lkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NLegtgGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ewqdskLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func frSmpminWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GA0RxT7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X8XJwNXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 817OhVfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8DCWZvdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MKBYvSiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TT4jEIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRq8bcgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfqMJJQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oXuaTWvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvONDS5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vxHAo1JFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func we9MsK2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B9F1i58XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sS3BS9QKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgDk9eYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acAbK9suWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUwwOYKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0r3DYawdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3zVy3LncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6iptngU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cLrrUJyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pDywm2wtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DQzkJq4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I1oGLZ6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HCiTuZM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RHE2LfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F1raMvvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jPmNvArXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func US1q9PpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KYyzGfysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LRfdbaQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qt9CHZJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kf7qdyFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zPA4Wr7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gynz4Bk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNvxjf6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6AQIsZX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yx8U5QwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IHmOIoTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T8rbTiwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oIxc96VsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I5hQpFL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXR6RulVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lw3sb6apWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6i0L6DHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0UThC8viWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ieVHQNzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GLeEOUPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AVYHQ2koWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oEQIP6rlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zqeLnc52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UPdcAOjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func apcGmICRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dqlfKLA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FffrM4rBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p6s0JrBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PYA1JDJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 918br1YNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c8rmEwyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L0sGtjcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ERtoUPzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qwuNapGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yCUgSJxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8iLchh5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N6FzuTRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vVIUz7yOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kvMWDBxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8UODc6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SorxOxCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LhbcBtrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZopXvFxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mFqX75XoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wDb9mQWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FofCJs2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FTH55zw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SS730vC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FzXfuFdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r3jpXVrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8CLzITFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KQEV6iYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NYZu7FKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xIN7Lq2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OqpNOlIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2CVeS85UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NesHvmcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ITRuWHwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bmErDIr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func veybxP60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gVATgmhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VDTwTtE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JHeMmu7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wKSamcbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ee8fOeMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xW8limKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ggNukYPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yYcYRXgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cuDcGT6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQ9tt7xqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cW9KKtFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4WAZaSCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CYDpuE78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EYZqGahlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PWqodIliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Djz0MkfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cm6t6k7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pRYpWORZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fK5Ga9ulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jTGp1nkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZsL715jFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DpvBXygfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yWT28TkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mitGL5nHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z3F04FNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tg5qWOTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func guv5sAlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OIN6ytCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lzddf5aQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gR8iYsnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 131xuY9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2gi5zNxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kAO2VwkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func euuy9HVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gWfq3G1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hGrOmuijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLV7WdHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fpTax9WxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OTMKDjd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8LkJmOSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tU6LquyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YSPVOQFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ath1qjWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func niZOPFEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4xvC6rcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bCHv6IC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pC4D4FjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tNunt3kCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func daT6QGk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h5BzumEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1omqdZQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oaULHipjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HYThOFNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g3CBy6oHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tZcZLnteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GxRQBwsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8SOOrfiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4aveFqwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lu1PDSZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cwfAm5NdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8z9ObsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EQQMRYyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l6AXJyFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HfOtBP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u2p2FZilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3DOsFgPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hbWdRsnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZdD6DexzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vieRMn24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JJ4eLj6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUdAu2ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GojSo15GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tpapt2yIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RZoFWFK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HCRRnnzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iP88z28cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qgSHK2MtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hQhLRLNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e6VOgGM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v45V98mxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Unww33tvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7o22uWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dqms2oXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iw4flHU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YqkPWr17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iV7msNaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xu4MWZzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l9SYp61FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k9y8BatdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 16MVU1qxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2i4BGK4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8hwM2JivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xYWLeqX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0mK9uvcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dqrMtMK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func al06aOLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lhNAZXZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U5xkuqzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c6s49pBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lH88jHpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tyLVYZUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WlCOBCe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5sLw0c9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func otsLZXjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwXftvCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f679bbyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HH8dbaCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1a2eRLx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UKsP1yh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BCzLOhtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func byPPYgYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xLshalhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9bblwKTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2DUEvlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hYiRcagbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aKfXS44QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z9hqvhPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OluqWRQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BYXII3hkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QayoIREBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QunUvdRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rD3e90itWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qbHNDfVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQaFSA48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s08YvfZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLoSggDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7A4QxbzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FNsgTMFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wvRt4smyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QMZkhklvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Dj0lDI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QGoOtAD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbGK9bK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6DkFs4oeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0uh3zXtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TaRRtGItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func deIAwsQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UdS3Co3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qbTdHeFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LIy9uZk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m4ljXTbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func psuyqMhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zGOfiCbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P1xejNJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pUUIvRrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zN49xLunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HwjLOIBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0C0YmNBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fz0h4aveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hYcWTU1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HaSB8ZKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gsUIWn7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VLO2f9NmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s3iGRaJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OoXOfSP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Du9mxNgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aZkEGw6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DMQB5hZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eLG60yCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c7u2KvGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sJl3tcvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 49z5XjJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncz5ubShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NlPQ3SFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3EWfLBtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TUFXFy42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bXDC61kvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4FIQvH3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tAvEEgoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sz3cswBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g1ift0GPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LGloGsyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yHaWUAI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vQtoNEgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HlTnFhqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TlCy1G1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 554hzP9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qq4pvZPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0c2LRTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HgOfNKrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WcCt8wU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gnsy1BchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VxCYnY4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K0tkroyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0avkkphYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Ry2lS03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HkXsMl0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9KUdCO0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDKh2K5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2dvmqv3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ccv8yt9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rLVzMP4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZMUnq1aCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3kP4dpdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R7Px1fE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ujeFWX5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VQ0s1GW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RWm9CaeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4KfT1LspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1fm1QYdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZpWTYoHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSoP00dOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wetJoJRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8SH9kddTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q5hrZUFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H1n2DY31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V1RXcYkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NdvfHntgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdPhCBGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m97tYk2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PHTwCZz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GYxTllyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bTfKmo3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1IWghgt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func au6l5T2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hxey9fucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OuHtkSQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rnyILc6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M9AcA7M9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mOjsQQXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kATMUIoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZQNVSBNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z0NO28pqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iRIJJELeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4jT7GzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jCdvp6WGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HbJgkTE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WcwZpDwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZCBsAXuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func emYn3uYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjoWLL53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OlmA7uMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6j4DqSplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Im0kTdoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rwYYStnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QKzvzTWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func poGWYqt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v0yp5IO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aOerIua8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iJHYhIScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TYg0KwOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GxeSC657Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iTNgDWhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z71UJJnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZJl1MDa7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZ0tjPTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Waz3BZStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dd0YYn3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2fWlq3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7qweifIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Unb4KbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jaiCe3u9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZVqihFaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oDQQyIBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iw85jF1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BoFIuv0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yY1y7FVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYQvp10KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aFEKmbeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0IZ8vJBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ybsT9lopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZgxxH3qvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bCpJzMVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w3IbUl84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YacePVe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tKs2LVylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DNOxzPYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z5WgXGWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Tk4EVZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XoqqyV2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func faUi8c6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0TG4k2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zNR2OFkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zw1vgaCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1RTI5Y9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QTc99D6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0e19J65oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 56AI2tGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAXSGhZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4zrLdDznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UIIuPrDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MOcoYgjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V21SRynLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LvfHzKC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GcX2UoECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y96S5dkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aS7Dz1QoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7WVGeRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ECqfdsR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RuEBvHxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kuKaWUfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MTaXr9lnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZLBUcmSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Egii0rs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RxL9k2ayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xc1zMscvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLueaZ16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QU4Pfvb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PGrhcw5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jpqp0ND0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NzjquzXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rpMqCmL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tqBOm7nZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fhoQ8cAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ARBUbpezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VZBdx771Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XijJ2UyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hlCkYH52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vgd95PDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEUu3SV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDzMhhOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eaWrxhN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DkQVZfoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O1Q53vXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8565H8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xiTy9ZXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xbkQDRGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l4GfJfy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IxZ2WgxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzRygLdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tMZNQjnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R8w9e2dMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxVDwoz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5aNqKHGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WqYgPItxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4EyajJQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JycXsU6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NobqvoTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HSRpaFAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jHfOd7LwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nab1xfmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ecmdXKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lPR7YF83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xdQm0L4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rm9xoSOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gfAhP8pzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 29415IEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JW5WifWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dtilel1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99HpMHf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kDFgPP47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func spPf8PKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jjhncrQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HyIrHtaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lR8ffr6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8wYUp3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0tm2fqSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2z1YkiJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B3tAM1OYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PrQ2yns4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nxbBIhxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RrhzXqiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dNWyb15XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BKz8zt0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2TH0lRuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UgMXcHyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eZm02IujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yjt5ux7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pV8z47NSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y5h8rRREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zpr2ST8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5oKwJc8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PT6RgWkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ns7CUHreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dTGOiyziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3fuhpxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PErn9l3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eKVEeb6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l0sVaiPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLp3K6E9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7YXi5VMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d8npuBkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vKzyfkoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DnihxOx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAxero6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ikxFz9wKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JEzSc7Q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7BzIcyEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9KhgoDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VURLz6MNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o6a5O752Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BQ5GuFQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IkkI3G1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvvdTRGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NmQFTeMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0deRiDxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wwk4V0JyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aihVymBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5M8jfLRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DzEOUXUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T89AtltFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func elefhCzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w0WqMCIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8OyiJB49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0egnucSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dgHwWqFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1VE4fDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svKqGysDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 13esBwBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XtrWhHG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wK9OqnwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func quy4zJq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IWVZ91JhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OfN7IQ1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6CPcvTHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rkZhdRVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pxg272aCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xjEbVuSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VKCJkjwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20GOuRKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGu5KjWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bIAI04PXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4vM2gn7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38iF0PXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Icw4ExZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FRpMxJqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OqR1Fw5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TPTtxCBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Jp1BEsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bCwhvEHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qk4eop6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBMYHvdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q3uvBxujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kAhMxNT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8jVLKA7QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 29YbCSP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8TJj2RjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jarxqwd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rDmKcEOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3p3AfoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tVsPXZg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rBOz2x31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f3AbuVSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BcOsShHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gSubEJLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mBurzuGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XdKb30csWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jtLLphMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y79G91D5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBxiSrx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZuAws5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ZheO2yTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8pNtbirRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uHx2Vc1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func boLVC0SOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XYnTbyrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PzbSSdkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pDj2kk5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fP9Mog9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2QyyzWltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mq3HL99JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 75WPjgRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WOBE2bnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q6qhuTUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vhOTRuvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S68TA1nAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7Vy1mzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func itpugtOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AfBIJeorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yK9SdZoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9SX6TvDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9j9bLdgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lBgemdWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 52OojFIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DxRvsKzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FIpbuIUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SR6vPkk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZ7OBTJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TmevcJMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VRln3DN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kJMVkbKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FWTrjwZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JpmcBailWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CsXJ4aSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fCWlE1kIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S3bvdJtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CnkPzBYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H0jruefNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SXXRIvuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tVtIuqD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hUd7ohDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TceoJWscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sulapOqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bsf8op4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iecjiSlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BClEsRwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bdrr3YO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jsq3S5qnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EjBdXaHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTXShEumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pAJaw3Z1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MhbN0FzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GqoPXIIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9NXxfMCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LyoYeginWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cYsVu68MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H721oD24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmlDtq3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kGu8RWrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RS6qJys0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OzkFN4P0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T8VEXzXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J8OKTCqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x11B1WavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oSwkTNJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HhZWxV86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RDEKylceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxYgn2b1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xyMrKCZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0qQ2vlgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func knNEHQjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ImdJxP8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bQklPgt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wmp3VgHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TFgjyIfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0WGYZPUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f3JLFRAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BwgRW8IIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDQp4fB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U9LdmoK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0PwOJzN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DnpIDmnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GTkzDXhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kxZEaVpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWZF0UAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tn8scdT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MqNBrNSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aY400HllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lB9ioSsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oQ2MlWMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ixty425Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dLAvuL2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func biYBBeWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hzQVQSrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mi3NeKPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cyNXSVLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a6kpn5uoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vkj7OXUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1AWF2ywTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Q5fRPyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ziW7oMuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2DnXtvxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UQ6IyJ54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T6bucuJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f623OfB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HvTRB6JhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1x5tSaBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FFwYUiOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b3qeDqWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lhbFGusLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fuh1jyjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a9CzWfi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qIrG8oCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zwCnJoxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vkc3Pwy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jRpu3bE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PboPXSojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oPE5NE54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KzCChAgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 63eEJElXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6snKpMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MwwgGhoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EF48bkmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IuRBhuxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4tlsfMPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T2k3XakCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MUTfIlrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iQbGxxxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VLLa0uTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1eBCFqOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ps4NUcaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JTWXiOgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DzhRbuj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RjzXotkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3bnA72dAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZH1k4hGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AhheInKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIXMmccXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0lndal1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LoLEmfvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8YIhuY3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fGxt8FK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1dWrjVONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fwmkQZooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func USpobHjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wK5knprfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TtRURc3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zjz0xqRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EXeL01dYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Be2VkRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kxKGyKzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9NJ79PovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WhZmFkirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J2PfJOYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KCVatcikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rHzOEmkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Thp4MDkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PPZWrfAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQogAsVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bz1JEfT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UYNS2D5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7DB07ARWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9SOQ7y8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HGLPXKJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOpojIiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tKN8yG9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YaTFJMXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oezIkx0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sI6kZwDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7m4ioALnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h7klPPStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R37mhFqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wDP5gcg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3RU98c3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5xcXATTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bg3Ij5psWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47d4gVxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wANdsE6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rgpv62a8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 06P8BN9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IeRxv3AoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func skpOOZzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J9zB7g5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rElAS8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUb4j73IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlRqu7LgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYageCxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func voot3Xc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RrT7f7KEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tMtJfu3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tk1YRgkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Os7nsgvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmBg5BolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zGIFn026Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ewsmXIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n8ZCGMsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kDYZVWh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fkYlc0hmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zjgy1ggEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4gehmIUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xq9oEvjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r79rhNMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1i1EqaUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qfh9MiZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M7Tqbwx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DQtEkBCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ViJ1asNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S3Pos7eTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HYOsDqnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cv2DZdUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func no5LRY0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AY89AHxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m3PQnj8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kw3mS6aVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iSSRLtIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aNBRfnAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F9p1Z1EyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HLKJirf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DL4dXSHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w9wccTzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8U1shanCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIGMXVxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7HfFUBbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yz1UN2L2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pUP6cIqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VQ3kaaltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JJnPbUs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RftbBjK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5DgL07oqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgKiZoBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uV6l50xfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tnNLoCerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UctpzGo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g4X3sgqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mGb74wtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXe3eNYVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VwiQkCkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AsD4xfbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Psz3rTwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NRizCPOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BTADIiNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8vFcix6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9SgCLZwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RmcqCXfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LNej08fyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0F1K8ImTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DUev70fTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NrcAR43wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qXGaFzL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6nia14oRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JrLm1lVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5hdlotZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ksfWLXeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GSsVjESGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nectrZD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wbtB8UlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnKjeSQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4zszsDjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bykhorOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zB60kvEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kv9hB4IlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 40qJtRzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rf5ggUbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o23SlcyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFZBrU7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNG7681FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mCYESUw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THN4y4HVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GGc6Q7VYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hV9TvHnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2X1fiRXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NWi180Z1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P98SWOx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MrGVxRDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hkH1AbSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rbp6WYp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WidsCJV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YMMS372iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wnYapKrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func To8M2aMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ct7Fhhk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N8ps4SzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIHe9l0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wL3jLtWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FImlt6CNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9XwulUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func raa472zYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBTd2vZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u2aaAoWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J5tc7cOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iIlmQGKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tGd033PgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qCOEEb9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qfq2cYNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 08nG8pR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t48BdIe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L8LmAoB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UkMpz5VVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fVLyOPduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QUpENRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7999jb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WIDh03FvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pgy8Lu1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EYty9ji0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JEw1ajXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7wqqmUeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9kixOkhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bri9p6QUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4FdtL7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hjHSa019Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uoh1xZrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tRwbnrP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func US44I1xZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5eLEXWSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JV4LzhLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QjvZwbKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WSlHt1vTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j1ItOPSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7msw8AKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qp7pgHyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4lb2wEMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yltANg5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5QWCSkgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fBSLrmFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h2AJYsRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cAlkLsM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bGCK6Q69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DluXxs5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V2mGLDoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQTiXHMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vXKsm3q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uC80fB0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jgXFDfugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ly8MR700Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLkzokEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func atuM7H6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2KtFVL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBxarpZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fZCehTxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C3Ws9mb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQeRX2ITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SJZXbG77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NCo8S27vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func goEhuMeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKfQjWrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dkSr2GyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pa5WzpMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NhjcqPWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZuH8ykHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXRltx7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RBjtlhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G9argykuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5e1UxFzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHr7vo0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qOYzQ2twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TC1l4parWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MGFCrYuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hOHPhCWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 35X3LJlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m7tv9MaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAX8HH4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZYrX2BKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7pEMAoK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UtItXGFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J3fN7xyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wo2UuHmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M8poxI5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OVEmqqa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nEySGMVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MST4iI2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwSPUdgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T77wDcsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func duP5FsgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zo5pqBnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LXxx9PFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a7koXgwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C4zotJ78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YFS02rrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tahyWV5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P650gkkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tAHrQScbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HxfAgXYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dm4TB76iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SvS10f2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f2kqUzNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xV1SmvvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3xxA07uDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sdcxCryFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func It1lBdMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 67xgdu66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sRSmZkLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gOYmX87aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3A2lVVPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func khY1ZR3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WEuv81zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KG6gAM8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DlwhJjBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WwAW9rt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tzpasuiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dJqvhHKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xwbZsHD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func piwcPmJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rCCXXXPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lD8RBlDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KVR639tFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k2laWQ1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sWqwFN9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SEtGgxwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NVn091nOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qr4UdjY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C1P7wEGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IHDlQNd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8EkPfUsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 62sWB2faWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ioMlFTwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XuJQLelOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UhTiLXcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CCvpSqSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6aSyaXFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NTNzfbeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ynxfsp7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fgASgG5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VOnKGLF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6KgdcaVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s3XFP1CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ij7Tgh1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h3LrZdWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZLKpJWkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zj3rmskdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WGQ0WrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WhazRMHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJPslfF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bMqElw54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iA8SfgSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o7spYejjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZZsdkx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func thTAwVnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KGQtnaPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yvs5exYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j2FhkD6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVqXh215Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tpCkclHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TLY8QHWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBwxIPVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 45PIzoa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opclNhseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func krwrGQiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tXERfgg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qc0MFWRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12ubz3XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jEW2S5HZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6BBGsXTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gp8h3LuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func osy3mnbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wm3uTzS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kip949rDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zCZMz4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LB2cHEcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 61A7hOUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mBZjacwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EzPgFuHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TKXUKernWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jMsPVNnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dsTC1y4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9tR33K9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1HvFb35zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b9QGWehvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D2nyuR6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 06Zlrjp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mjWsZ0YoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pQSJG6PvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T5meZMvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jMq2ERrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YTv5qktqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tqFuFN8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCREM19zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zl8CRlxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D0cmusP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nKRNq4VBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ETNW8YylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PrTdd8ciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GvDAVMlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8j6poPKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A4iVPMJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xb160wgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4goYSyeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mZzOSIZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4TR79SOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func razn4z42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sL0y2QYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BuWukiHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iKRps46LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNoSulW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6Hpi2wTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MDomWTPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xzc3JpZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tdm9dmMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NgH5CwRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WWMYaGjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CzTJAXIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UL3xiJo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pG2hpuYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jMcl4zueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cOb8SwQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IM0Fz0yyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0QKLLit3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHIdMj4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IymxEqixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tWf3U89nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TdqjzjYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QT8y7FIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmoWbT91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Jqdlc6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qzHAxun6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dzsFKiWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TIX7Z64sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j9tDXMXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6P7Wvh6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dGmTlTDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1XFCOy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QgNrdPLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k1ZX4yFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zfH0QXaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bGcRTk0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LahaGzakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYnd663AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RHuZL4FuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1JzgUBRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jaq79SFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H6UFCBoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k75Pz151Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QVYdXeu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Alup1XmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u6YRVJ6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 638J8OsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RbLoaiutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SldAEQAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P91aG2peWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jgSZS1tZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5vKNp4HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aobkjsWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tMLfLY2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wGsSkC3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ct8SgbVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vgic13L3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHjrqcb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0pCbmPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sval3zXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ctagH914Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DOo1EOQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIPtpuSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func biOfRcclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vkOxkEeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TeOjtQRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GN8Hdp9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NblkKJi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GmtKXsGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g7KqCUzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wN4kbjJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q02vRcFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3u1LF1DPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orvwPZUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0WOaGA6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acT7NU2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g8wSs7RNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5hjHeoveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hE6ynfPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C89hCT5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X74r46EkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSeUmVzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZciCYXiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bpj77N0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mSa5dXqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T9dfsqrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rq8lc1OAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BUYqgvkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dQHv0EdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func klArKsmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r4xuTBrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MvSPVsV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 24Miq3ouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9jDWt1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PVq7Zx2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PgSXfUrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JlD30ALEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdNGT7R5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5tsH9vTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aaBNbEoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZNzFwjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aq0SNalKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nKK7zwfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IbefmkgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BhZy9m8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OJGsqrv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z5k9enK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mMZSleuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dnhN0z2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I7S4ABHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LJcs0vAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J0loWZ12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBGI17GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HIiStN9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uMQl7BVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lpPzNzfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func slCAjEmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y7IL73RDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XYpXoz1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 50JHJGiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VU1oXL8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UYO38NlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xiEEbcCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t60EcWzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KXcaLK2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHgdh8UdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSt84AVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLdMI2nKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1haWk333Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s6wMpAtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zGou5tEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VRz7vUXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6uydKMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5EPC7u8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UexO1JuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2vR9WMObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QEO1xaogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2jsmMdHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func epydTrijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3UAegNFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBmyGroxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XEsjQ60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4FTgaCwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBT8NlpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxbCjI31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IODrL0e8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TU9JhO1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QWuoUSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AP5exr6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yNv83b0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bKsNiXjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZPM2tsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O2DgpIM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tXem7t1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kIzwwAroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nuH4wNroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DJBYXjA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xEbr4WqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S0CKAaxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gqwmXpNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dqniCr5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4JYUoCPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yiYTeAu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cjxd11KVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLaQCt9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OmwDmsW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WrVKb7DAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7OQqyHa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n07rQZ77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmYDMRNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bSotjaEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltnLm8orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vb3OsVAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9C7ROeNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EwdZhxW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b5gYG2RUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TDmiKodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OS4JjzJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ih5NPOpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VmmwhVLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QxqHgf0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kZNlryBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PqCJKolzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g9LyGuLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KJVyH1F1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KocuQ3UxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IpD3ZkwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y1zaZ03zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OByaEwogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMZM2aGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ADv3vi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fSVEov0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XrNPJs3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udXKWm71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LPRvm0oUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kJVU93EGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JyG8o9EGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7vK38oqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AtZHWnkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SdQPhQPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XRWk1PeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6kGKr4b7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hmy0Ks6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1AJmnQYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KIO5rvsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBNZkVR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k9ABnc0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iKejHV2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UWasXUb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4QNorOiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5JS4HGfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

