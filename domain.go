package incsw

import "errors"

func readDomainMap(result *Domain, buffer []byte, mapSize uint, offset uint) (uint, error) {
	var key []byte
	var err error
	for i := uint(0); i < mapSize; i++ {
		key, offset, err = readMapKey(buffer, offset)
		if err != nil {
			return 0, err
		}
		switch string(key) {
		case "domain":
			result.Domain, offset, err = readString(buffer, offset)
			if err != nil {
				return 0, err
			}
		default:
			return 0, errors.New("unknown domain key: " + string(key))
		}
	}
	return offset, nil
}
