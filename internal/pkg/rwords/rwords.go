package rwords

import (
	"bufio"
	crand "crypto/rand"
	"errors"
	"math"
	"math/big"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const UnixWordsPath = "/usr/share/dict/words"
const DefaultWordsPath = "./conf/words.txt"

var words = sync.OnceValue(func() []string {
	filePath := UnixWordsPath
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		filePath = DefaultWordsPath
	}

	file, err := os.Open(filePath)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}
	}
	return lines
})

var random = sync.OnceValue(func() *mrand.Rand {
	return mrand.New(mrand.NewSource(time.Now().UnixNano()))
})

func GetGetRWordsDefault() (string, error) {
	return GetRWords(false, true)
}

// GetRWords 获取随机单词
// capitalize 首字母大写, includeNumber 包含数字
func GetRWords(capitalize, includeNumber bool) (string, error) {
	wordSlice := words()
	if len(wordSlice) == 0 {
		return "", errors.New("no words found")
	}
	r := mrand.New(random())
	result := wordSlice[r.Int()%len(wordSlice)]
	for len(result) < 4 {
		result = wordSlice[r.Int()%len(wordSlice)]
	}
	if capitalize {
		result = strings.ToUpper(string(result[0])) + strings.ToLower(result[1:])
	}
	if includeNumber {
		result = result + strconv.Itoa(r.Intn(9000)+1000)
	}
	return result, nil
}

func cryptoInt() (int, error) {
	result, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return 0, err
	}
	return int(result.Uint64()), nil
}

func GetRWordsCryptoDefault() (string, error) {
	return GetRWordsCrypto(false, true)
}

// GetRWordsCrypto 获取安全随机单词
// capitalize 首字母大写, includeNumber 包含数字
func GetRWordsCrypto(capitalize, includeNumber bool) (string, error) {
	wordSlice := words()
	if len(wordSlice) == 0 {
		return "", errors.New("no words found")
	}

	i, err := cryptoInt()
	if err != nil {
		return "", err
	}
	result := wordSlice[i%len(wordSlice)]
	for len(result) < 4 {
		i, err = cryptoInt()
		if err != nil {
			return "", err
		}
		result = result + strconv.Itoa(i)
	}
	if capitalize {
		result = strings.ToUpper(string(result[0])) + strings.ToLower(result[1:])
	}
	if includeNumber {
		ci, err := crand.Int(crand.Reader, big.NewInt(9000))
		if err != nil {
			return "", err
		}
		return strconv.Itoa(int(ci.Uint64()) + 1000), nil
	}
	return result, nil
}
