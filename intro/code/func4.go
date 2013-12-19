// +build OMIT

package main

import (
	"fmt"
)

// STARTMAIN OMIT
func main() {
        addIfNotNilUint16 := func(name, param string) {
    value, ok := post[param].([]interface{})
                if ok {
                        values := make([]uint16, len(value))
                        for i, v := range value {
                                values[i] = uint16(v.(float64))
                        }
                        post[param] = values
                        match[name] = bson.M{`$in`: value}
                }
        }
        // Optional filters.
        addIfNotNilString(`sex`, `sex`)
        addIfNotNilString(`mstatus`, `mstatus`)
        addIfNotNilUint8(`influence_zones`, `influence`)
        addIfNotNilString(`type`, `proposal_types`)
        addIfNotNilUint16(`catid`, `categories`)
        addIfNotNilUint16(`comid`, `commerces`)
	f := Generator(3)
	fmt.Println(f())
}

// STOPMAIN OMIT
