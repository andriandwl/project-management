package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(value interface{}) error {
	var str string

	switch v := value.(type){
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("Failed to parse UUIDArray: unsupported type")
	}

	str = strings.TrimPrefix(str, "{}")
	str = strings.TrimSuffix(str, "{}")
	parts := strings.Split(str, ",")

	*a = make(UUIDArray, 0, len(parts))

	for _, s := range parts {
		s = strings.TrimSpace(strings.Trim(s, `"`))
		if s == ""{
			continue
		}

		u, error := uuid.Parse(s)

		if error != nil{
			return fmt.Errorf("Failed to parse UUIDArray: invalid UUID %s", s)
		}

		*a = append(*a, u)
	}

	return nil
}

func (a UUIDArray) Value()(driver.Value, error){
	if len(a) == 0 {
		return "{}", nil
	}

	postgreFormat := make([]string, 0, len(a))

	for _, value := range a {
		postgreFormat = append(postgreFormat, fmt.Sprintf(`"%s"`, value.String()))
	}

	return "{"+ strings.Join(postgreFormat, ",") +"}", nil
}

func (UUIDArray) GoremDataType() string {
	return "uuid[]"
}