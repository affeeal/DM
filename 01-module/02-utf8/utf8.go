package main

func encode(utf32 []rune) []byte {
    var image []byte
    for _, sym := range utf32 {
        switch {
            case sym < (1 << 7):
                image = append(image, byte(sym))
            case sym < (1 << 11):
                image = append(image, byte((sym >> 6)   + (((1 << 2) - 1) << 6)))
                image = append(image, byte((sym         & ((1 << 6) - 1)) + (1 << 7)))
            case sym < (1 << 16):
                image = append(image, byte((sym >> 12)  + (((1 << 3) - 1) << 5)))
                image = append(image, byte((sym >> 6)   & ((1 << 6) - 1) + (1 << 7)))
                image = append(image, byte((sym         & ((1 << 6) - 1)) + (1 << 7)))
            default:
                image = append(image, byte((sym >> 18)  + (((1 << 4) - 1) << 4)))
                image = append(image, byte(((sym >> 12) & ((1 << 6) - 1)) + (1 << 7)))
                image = append(image, byte(((sym >> 6)  & ((1 << 6) - 1)) + (1 << 7)))
                image = append(image, byte((sym         & ((1 << 6) - 1)) + (1 << 7)))
        }
    }
    return image
}

func decode(utf8 []byte) []rune {
    var image []rune
    for i := 0; i < len(utf8); i++ {
        s1 := rune(utf8[i])
        switch {
            case s1 >> 7 == 0:
                image = append(image, s1)
            case s1 >> 5 == 6:
                s2 := rune(utf8[i+1])
                image = append(image, ((s1 & ((1 << 5) - 1)) << 6) +
                                       (s2 & ((1 << 6) - 1)))
                i += 1
            case s1 >> 4 == 14:
                s2 := rune(utf8[i+1])
                s3 := rune(utf8[i+2])
                image = append(image, ((s1 & ((1 << 4) - 1)) << 12) +
                                      ((s2 & ((1 << 6) - 1)) << 6) +
                                       (s3 & ((1 << 6) - 1)))
                i += 2
            default:
                s2 := rune(utf8[i+1])
                s3 := rune(utf8[i+2])
                s4 := rune(utf8[i+3])
                image = append(image, ((s1 & ((1 << 3) - 1)) << 18) +
                                      ((s2 & ((1 << 6) - 1)) << 12) +
                                      ((s3 & ((1 << 6) - 1)) << 6) +
                                       (s4 & ((1 << 6) - 1)))
                i += 3
        }
    }
    return image
}

func main() {

}