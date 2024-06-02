package main

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"net"
)

func main() {
	//resultMap := make(map[string]interface{})

	//gosnmp.Default.Target = os.Args[1]
	//gosnmp.Default.Target = "172.16.9.44"
	//gosnmp.Default.Target = "172.16.8.3"
	gosnmp.Default.Target = "192.168.0.1"
	gosnmp.Default.Community = "public"
	gosnmp.Default.Port = 161
	err := gosnmp.Default.Connect()
	if err != nil {
		fmt.Println("Error connecting to SNMP agent:", err)
		return
	}
	defer gosnmp.Default.Conn.Close()

	//// Get system name
	//sysName, err := gosnmp.Default.Get([]string{"1.3.6.1.2.1.1.5.0"})
	//if err != nil {
	//	fmt.Println("Error getting system name:", err)
	//	return
	//}
	//resultMap["system.name"] = string(sysName.Variables[0].Value.([]uint8))
	//
	//// Get system description
	//sysDesc, err := gosnmp.Default.Get([]string{"1.3.6.1.2.1.1.1.0"})
	//if err != nil {
	//	fmt.Println("Error getting system description:", err)
	//	return
	//}
	//resultMap["system.description"] = fmt.Sprintf("%s", sysDesc.Variables[0].Value)
	//
	//// Get system location
	//sysLoc, err := gosnmp.Default.Get([]string{"1.3.6.1.2.1.1.6.0"})
	//if err != nil {
	//	fmt.Println("Error getting system location:", err)
	//	return
	//}
	//resultMap["system.location"] = fmt.Sprintf("%s", sysLoc.Variables[0].Value)
	//
	//// Get system objectId
	//sysObjId, err := gosnmp.Default.Get([]string{"1.3.6.1.2.1.1.2.0"})
	//if err != nil {
	//	fmt.Println("Error getting system objectId:", err)
	//	return
	//}
	//resultMap["system.objectId"] = fmt.Sprintf("%s", sysObjId.Variables[0].Value)
	//
	//// Get system uptime
	//sysUptime, err := gosnmp.Default.Get([]string{"1.3.6.1.2.1.1.3.0"})
	//if err != nil {
	//	fmt.Println("Error getting system uptime:", err)
	//	return
	//}
	//resultMap["system.uptime"] = sysUptime.Variables[0].Value.(uint32)
	//
	//// Get system interfaces
	//sysInterfaces, err := gosnmp.Default.Get([]string{"1.3.6.1.2.1.2.1.0"})
	//if err != nil {
	//	fmt.Println("Error getting system uptime:", err)
	//	return
	//}
	//resultMap["system.interfaces"] = sysInterfaces.Variables[0].Value.(int)

	// "interface.index": "1.3.6.1.2.1.2.1.1.1"
	//oids := map[string]string{
	//	"interface.alias":        ".1.3.6.1.2.1.31.1.1.1.18",
	//	"interface.admin.status": ".1.3.6.1.2.1.2.2.1.7",
	//	"test":                   ".1.3.6.1.2.1.2.2.1.6",
	//}
	//_ = oids

	//for _, oid := range oids {
	_ = gosnmp.Default.BulkWalk(".1.3.6.1.2.1.2.2.1.6", printValue)

	//for _, v := range packet {
	//	fmt.Println(v.Name, "\t", string(v.Value.([]byte)), "\t", v.Type)
	//
	//}
	//}

	//for _, v := range oids {
	//err = gosnmp.Default.Walk(v, printValue)
	//packet, err := gosnmp.Default.GetNext(oid)
	//if err != nil {
	//	fmt.Println("Error getting interface descriptions:", err)
	//	return
	//}
	//for _, variable := range packet.Variables {
	//	fmt.Println(variable.Name, "\t", variable.Value)
	//}
	//}

	//err = gosnmp.Default.Walk("1.3.6.1.2.1.31.1.1.1", handleInterface)
	//if err != nil {
	//	fmt.Println("Error getting interface descriptions:", err)
	//	return
	//}

}

//	func handleInterface(variable gosnmp.SnmpPDU) error {
//		fmt.Printf("%s = %v\n", variable.Name, variable.Value)
//		return nil
//	}
func printValue(pdu gosnmp.SnmpPDU) error {
	fmt.Printf("%s = ", pdu.Name)

	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		fmt.Println(net.HardwareAddr(b).String())
	default:
		fmt.Printf("TYPE %d: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
	}
	return nil
}
