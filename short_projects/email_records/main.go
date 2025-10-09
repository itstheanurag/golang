package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    log.Println("Program started. Awaiting domain input (one per line)...")
    fmt.Printf("domain, hasMxRecord, hasSPFRecord, sprRecord, hashDMARC, dmarcRecord \n")

    for scanner.Scan() {
        domain := scanner.Text()
        log.Printf("Checking domain: %s\n", domain)
        checkDomain(domain)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error: something went wrong while reading input: %v\n", err)
    }

    log.Println("Program finished.")
}


func checkDomain(domain string) {
    var hasMxRecord, hasSPFRecord, hasDMARC bool
    var spfRecord, dmarcRecord string

    // A records
    ipRecords, err := net.LookupHost(domain)

    if err != nil {
        log.Printf("Error fetching A/AAAA records: %v\n", err)
    }

    fmt.Printf("A/AAAA Records: %v\n", ipRecords)

    // MX records
    mxRecords, err := net.LookupMX(domain)
    if err != nil {
        log.Printf("Error fetching MX records: %v\n", err)
    }

    if len(mxRecords) > 0 {
        hasMxRecord = true
    }

    for _, mx := range mxRecords {
      log.Printf("MX Record - Host: %s, Preference: %d\n", mx.Host, mx.Pref)
    }

    // TXT records
    txtRecords, err := net.LookupTXT(domain)

    if err != nil {
        log.Printf("Error fetching TXT records: %v\n", err)
    }

    for _, record := range txtRecords {
        if strings.HasPrefix(record, "v=spf1") {
            hasSPFRecord = true
            spfRecord = record
        }
    }
    fmt.Printf("TXT Records: %v\n", txtRecords)

    // DMARC record
    dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
    if err != nil {
        log.Printf("Error fetching DMARC records: %v\n", err)
    }

    for _, record := range dmarcRecords {
        if strings.HasPrefix(record, "v=DMARC1") {
            hasDMARC = true
            dmarcRecord = record

			break
        }
    }

	 for index, record := range dmarcRecords {
		 fmt.Printf("%v DMARC Record: %v\n", index, record)
    }

    // CNAME record
    cname, err := net.LookupCNAME(domain)
    if err != nil {
        log.Printf("Error fetching CNAME record: %v\n", err)
    }
    fmt.Printf("CNAME Record: %v\n", cname)

    // NS records
    nsRecords, err := net.LookupNS(domain)
    if err != nil {
        log.Printf("Error fetching NS records: %v\n", err)
    }
    fmt.Printf("NS Records: %v\n", nsRecords)

    // SRV records (example for _sip._tcp)
    _, srvs, err := net.LookupSRV("sip", "tcp", domain)
    if err != nil {
        log.Printf("Error fetching SRV records: %v\n", err)
    }
    fmt.Printf("SRV Records (_sip._tcp): %v\n", srvs)

    fmt.Printf("%s, %t, %t, %s, %t, %s\n", domain, hasMxRecord, hasSPFRecord, spfRecord, hasDMARC, dmarcRecord)
}

