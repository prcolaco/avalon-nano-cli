package nano

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

const (
	CMD_PORT = "4028"

	LEDRGB_STEP = 5
)

func GetVersion(host string) []string {
	result, err := command(host, "version")
	if err != nil {
		fmt.Println(err)
	}

	parts := strings.Split(result, "|")
	versions := strings.Split(parts[1], ",")
	return versions[1:]
}

func GetWifi(host string) []string {
	result, err := command(host, "wifi|get")
	if err != nil {
		fmt.Println(err)
	}

	params := make([]string, 0, 10)
	parts := strings.Split(result, " ")
	for _, part := range parts {
		params = append(params, strings.ReplaceAll(strings.TrimSuffix(part, "]"), "[", "="))
	}
	return params
}

func GetWorkLevel(host string) string {
	result, err := command(host, "ascset|0,worklevel,get")
	if err != nil {
		fmt.Println(err)
	}

	levels := map[string]string{"0": "low", "1": "med", "2": "high"}
	level := strings.Split(result, " ")[1]
	return levels[level]
}

func SetWorkLevel(host string, level string) bool {
	level = strings.ToLower(level)
	levels := map[string]int{"low": 0, "med": 1, "high": 2}
	if _, ok := levels[level]; !ok {
		fmt.Println("Error: invalid work level, use 'low', 'med' or 'high'...")
		return false
	}

	result, err := command(host, "ascset|0,worklevel,set,"+fmt.Sprint(levels[level]))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return result == "OK"
}

func GetLed(host string) []string {
	result, err := command(host, "ascset|0,led,get")
	if err != nil {
		fmt.Println(err)
	}

	info := strings.Split(strings.TrimPrefix(strings.TrimSuffix(result, "]"), "led["), "-")
	// info[2] is always 100, no idea for what, maybe for brightness max? don't know...
	modes := map[string]string{"0": "off", "1": "fixed", "2": "flash", "3": "pulse", "4": "loop"}
	return []string{modes[info[0]], info[1], rgbSettingToHex(info[3:])}
}

func SetLedMode(host string, mode string) bool {
	mode = strings.ToLower(mode)
	modes := map[string]string{"off": "0", "fixed": "1", "flash": "2", "pulse": "3", "loop": "4"}
	if _, ok := modes[mode]; !ok {
		fmt.Println("Error: invalid led power, use 'off', 'fixed', 'flash', 'pulse' or 'loop'...")
		return false
	}

	result, err := command(host, "ascset|0,led,setmode,"+modes[mode])
	if err != nil {
		fmt.Println(err)
		return false
	}

	return result == "led setmode ok"
}

func SetLedColor(host string, brightness string, rgbhex string) bool {
	params := fmt.Sprintf("%v-100-%v", brightness, strings.Join(rgbHexToSetting(rgbhex), "-"))
	// params: 10-100-1-2-51
	result, err := command(host, "ascset|0,led,setrgb,"+params)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return result == "led setrgb ok"
}

func GetMining(host string) []string {
	result, err := command(host, "pools")
	if err != nil {
		fmt.Println(err)
	}

	pools := strings.Split(result, "|")
	return pools
}

func SetMiningPool(host string, password string, pool string, poolurl string, pooluser string, poolpass string) bool {
	// user is always root, so not requiring
	index, _ := strconv.Atoi(pool)
	index--
	passhash := hashPassword(password)
	cmd := fmt.Sprintf("setpool|root,%v,%v,%v,%v,%v", passhash, index, poolurl, pooluser, poolpass)
	result, err := command(host, cmd)
	if err != nil {
		fmt.Println(err)
	}

	return result == "\n"
}

func GetStats(host string) []string {
	result, err := command(host, "estats|")
	if err != nil {
		fmt.Println(err)
	}

	stats := strings.Split(result, "|")
	details := strings.Split(stats[1], ",")
	return details
}

func ChangePassword(host string, currentPassword, newPassword string) bool {
	curhash := hashPassword(currentPassword)
	newhash := hashPassword(newPassword)
	result, err := command(host, "ascset|0,password,"+curhash+","+newhash)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return result == "new password success set."
}

func Reboot(host string) bool {
	result, err := command(host, "ascset|0,reboot,all")
	if err != nil {
		fmt.Println(err)
		return false
	}

	return result == "OK"
}

func command(host string, command string) (string, error) {
	// Connect to the server
	conn, err := net.Dial("tcp", targetFromHost(host))
	if err != nil {
		return "", err
	}

	// Send some data to the server
	_, err = conn.Write([]byte(command))
	if err != nil {
		return "", err
	}

	res, _ := bufio.NewReader(conn).ReadString('\n')

	// Close the connection
	conn.Close()

	return cleanupResult(res), nil
}

func targetFromHost(host string) string {
	return host + ":" + CMD_PORT
}

func cleanupResult(res string) (result string) {
	// remove last chr(0)
	result = strings.TrimSuffix(res, string(0))

	// remove trailing pipe
	result = strings.TrimSuffix(result, "|")

	// split result from detail, if present
	parts := strings.Split(result, "|")
	result = parts[0]
	details := ""
	if len(parts) > 1 {
		details = strings.Join(parts[1:], "|")
	}

	// cleanup result
	result = strings.Split(result, ",")[3]
	result = strings.TrimPrefix(result, "Msg=")
	result = strings.TrimPrefix(result, "ASC 0 set ")
	infoPrefix := regexp.MustCompile("^.*info: ")
	result = infoPrefix.ReplaceAllString(result, "")

	// add detail if present
	if len(details) > 0 {
		result = strings.Join([]string{result, details}, "|")
	}

	return
}

func rgbSettingToHex(setting []string) string {
	if len(setting) != 3 {
		fmt.Println("Error: invalid color setting...")
		return "invalid"
	}

	values := make([]byte, 0, 3)
	for _, v := range setting {
		i, _ := strconv.ParseUint(v, 10, 64) // 0 if failed
		b := byte(i) * LEDRGB_STEP
		values = append(values, b)
	}
	return strings.ToUpper(hex.EncodeToString(values))
}

func rgbHexToSetting(hexrgb string) []string {
	values, err := hex.DecodeString(hexrgb)
	if err != nil || len(values) != 3 {
		fmt.Println("Error: invalid color hex...")
		return []string{"0", "0", "0"}
	}

	setting := make([]string, 0, 3)
	for _, v := range values {
		b := int(v / LEDRGB_STEP)
		setting = append(setting, strconv.Itoa(b))
	}

	return setting
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return "ff0000ff" + hex.EncodeToString(hash[0:12])
}
