# Email Domain DNS Record Checker

This Go program reads domain names from standard input and checks various DNS records using Go's `net` package. It prints a summary of key DNS records for each domain, which are essential for email delivery, security, and domain configuration.

## Checked DNS Record Types

- **A/AAAA Records**: Map a domain to its IPv4/IPv6 address, allowing browsers and services to locate the server.
- **MX Records**: Specify mail servers responsible for receiving email on behalf of the domain.
- **TXT Records**: Hold arbitrary text data, commonly used for SPF, DKIM, and other verification mechanisms.
- **SPF Records**: A type of TXT record that defines which mail servers are allowed to send email for the domain, helping prevent email spoofing.
- **DMARC Records**: TXT records that specify policies for handling email that fails SPF or DKIM checks, improving email security.
- **CNAME Record**: Maps a domain to another canonical domain name, often used for subdomain redirection.
- **NS Records**: Indicate the authoritative name servers for the domain, which manage DNS queries for it.
- **SRV Records**: Define the location (hostname and port) of servers for specific services, such as SIP or XMPP.

## Why These Records Matter

These DNS records are crucial for:

- Ensuring reliable email delivery and preventing spam or spoofing.
- Properly routing web and service traffic to the correct servers.
- Verifying domain ownership and configuration for various services.

## Usage

```sh
go run main.go
```

Then enter domain names, one per line. Press `Ctrl+D` (Linux/macOS) or `Ctrl+Z` (Windows) to end input.

The program will output a summary of the DNS records found for each domain.
