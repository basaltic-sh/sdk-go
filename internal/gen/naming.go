package main

import (
	"slices"
	"strings"
	"unicode"
)

func sortStrings(s []string) { slices.Sort(s) }

// initialisms are rendered all-caps in Go identifiers, matching what a Go
// reviewer would write by hand. Order matters only for readability here; the
// matcher works on whole words.
var initialisms = map[string]string{
	"id": "ID", "ids": "IDs", "url": "URL", "urls": "URLs",
	"uri": "URI", "api": "API", "http": "HTTP", "https": "HTTPS",
	"ip": "IP", "ips": "IPs", "cpu": "CPU", "ram": "RAM", "ttl": "TTL",
	"dns": "DNS", "vpc": "VPC", "vpcs": "VPCs", "nat": "NAT", "acl": "ACL", "cidr": "CIDR",
	"mac": "MAC", "vm": "VM", "os": "OS", "ssh": "SSH", "tls": "TLS",
	"ssl": "SSL", "sql": "SQL", "json": "JSON", "yaml": "YAML",
	"uuid": "UUID", "crn": "CRN", "arn": "ARN", "kms": "KMS", "iam": "IAM",
	"sts": "STS", "mfa": "MFA", "cors": "CORS", "sla": "SLA", "eol": "EOL",
	"pdf": "PDF", "csv": "CSV", "html": "HTML", "smtp": "SMTP",
	"nic": "NIC", "nics": "NICs", "gb": "GB", "mb": "MB", "kb": "KB",
	"tb": "TB", "gib": "GiB", "mib": "MiB", "iops": "IOPS", "otlp": "OTLP",
	"grpc": "GRPC", "rbd": "RBD", "asn": "ASN", "bgp": "BGP", "vlan": "VLAN",
	"l7": "L7", "l4": "L4", "sni": "SNI", "cn": "CN", "san": "SAN",
	"ca": "CA", "csr": "CSR", "pem": "PEM", "kek": "KEK", "dek": "DEK",
	"aes": "AES", "rsa": "RSA", "ec": "EC", "jwt": "JWT", "oauth": "OAuth",
	"s3": "S3", "db": "DB", "ha": "HA", "az": "AZ", "ntp": "NTP",
	"pkcs": "PKCS", "hmac": "HMAC", "sha": "SHA", "md5": "MD5",
	"ipv4": "IPv4", "ipv6": "IPv6", "dnssec": "DNSSEC", "icmp": "ICMP",
	"icmpv6": "ICMPv6", "nvme": "NVMe", "ssd": "SSD", "hdd": "HDD",
}

// splitWords breaks an identifier written in snake_case, kebab-case,
// camelCase or PascalCase into its lowercase words.
func splitWords(s string) []string {
	var words []string
	var cur strings.Builder
	flush := func() {
		if cur.Len() > 0 {
			words = append(words, strings.ToLower(cur.String()))
			cur.Reset()
		}
	}
	runes := []rune(s)
	for i, r := range runes {
		switch {
		case !unicode.IsLetter(r) && !unicode.IsDigit(r):
			// Any punctuation separates words. Query parameters written in
			// the repeated-array style — "by[]", "match[]" — reach here, and
			// the brackets must not survive into an identifier.
			flush()
		case unicode.IsUpper(r):
			// A new word starts at a lower→upper transition, and at the end
			// of a run of capitals followed by a lowercase letter, so that
			// "HTTPServer" splits as "http", "server" rather than one word.
			prevLower := i > 0 && unicode.IsLower(runes[i-1])
			nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			prevUpper := i > 0 && unicode.IsUpper(runes[i-1])
			if prevLower || (prevUpper && nextLower) {
				flush()
			}
			cur.WriteRune(r)
		case unicode.IsDigit(r):
			// Keep digits with the word they trail ("l7", "sha256") but start
			// a word when they follow a separator.
			cur.WriteRune(r)
		default:
			cur.WriteRune(r)
		}
	}
	flush()
	return words
}

// exportedName renders s as an exported Go identifier.
func exportedName(s string) string {
	words := splitWords(s)
	var b strings.Builder
	for _, w := range words {
		if up, ok := initialisms[w]; ok {
			b.WriteString(up)
			continue
		}
		b.WriteString(strings.ToUpper(w[:1]))
		b.WriteString(w[1:])
	}
	out := b.String()
	if out == "" {
		return "Field"
	}
	// An identifier cannot start with a digit.
	if unicode.IsDigit(rune(out[0])) {
		out = "N" + out
	}
	return out
}

// goDoc renders a specification description as a Go doc comment body,
// prefixed with the identifier it documents so the comment reads the way Go
// doc comments are supposed to.
func goDoc(name, description string, extra ...string) []string {
	var lines []string
	description = strings.TrimSpace(description)
	if description == "" {
		if len(extra) == 0 {
			return nil
		}
	} else {
		para := strings.Split(description, "\n")
		first := true
		for _, ln := range para {
			ln = strings.TrimRight(ln, " \t")
			if first && name != "" {
				trimmed := strings.TrimSpace(ln)
				// Only prepend the name when the sentence does not already
				// start with it; "Instance Instance is..." reads badly.
				if trimmed != "" && !strings.HasPrefix(trimmed, name+" ") {
					ln = name + " " + lowerFirstWord(trimmed)
				}
				first = false
			}
			lines = append(lines, ln)
		}
	}
	for _, e := range extra {
		if e == "" {
			continue
		}
		if len(lines) > 0 {
			lines = append(lines, "")
		}
		lines = append(lines, strings.Split(e, "\n")...)
	}
	return lines
}

// lowerFirstWord lowercases the first word unless it looks like a proper noun
// or an acronym, so "Create a new instance" becomes "creates..." material
// without mangling "IAM policy ...".
func lowerFirstWord(s string) string {
	if s == "" {
		return s
	}
	word, rest, _ := strings.Cut(s, " ")
	// All-caps or mixed-caps words are acronyms or names; leave them.
	if word == strings.ToUpper(word) && len(word) > 1 {
		return s
	}
	upperCount := 0
	for _, r := range word {
		if unicode.IsUpper(r) {
			upperCount++
		}
	}
	if upperCount > 1 {
		return s
	}
	// "Cloud Resource Name" is a name, not a sentence — a capitalised word
	// followed by another capitalised word is left alone.
	if rest != "" {
		next, _, _ := strings.Cut(rest, " ")
		if next != "" && unicode.IsUpper(rune(next[0])) {
			return s
		}
	}
	lowered := strings.ToLower(word[:1]) + word[1:]
	if rest == "" {
		return lowered
	}
	return lowered + " " + rest
}

// commentLines renders doc lines as "// ..." at the given indent.
func commentLines(indent string, lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	var b strings.Builder
	for _, ln := range lines {
		b.WriteString(indent)
		if ln == "" {
			b.WriteString("//\n")
			continue
		}
		b.WriteString("// ")
		b.WriteString(ln)
		b.WriteByte('\n')
	}
	return b.String()
}

// wrapText reflows prose to width columns.
//
// Consecutive prose lines are joined before wrapping, so a description that
// arrives already broken at 72 columns does not end up broken twice at
// different points. Pre-formatted lines — fenced code, indented blocks,
// tables, list items, headings — are passed through untouched.
func wrapText(s string, width int) string {
	var out []string
	var para []string

	flush := func() {
		if len(para) == 0 {
			return
		}
		words := strings.Fields(strings.Join(para, " "))
		para = para[:0]
		cur := ""
		for _, w := range words {
			switch {
			case cur == "":
				cur = w
			case len(cur)+1+len(w) > width:
				out = append(out, cur)
				cur = w
			default:
				cur += " " + w
			}
		}
		if cur != "" {
			out = append(out, cur)
		}
	}

	inFence := false
	for _, line := range strings.Split(s, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") {
			flush()
			inFence = !inFence
			out = append(out, line)
			continue
		}
		verbatim := inFence || trimmed == "" ||
			strings.HasPrefix(line, "  ") ||
			strings.HasPrefix(line, "\t") ||
			strings.HasPrefix(trimmed, "- ") ||
			strings.HasPrefix(trimmed, "* ") ||
			strings.HasPrefix(trimmed, "|") ||
			strings.HasPrefix(trimmed, "#")
		if verbatim {
			flush()
			out = append(out, line)
			continue
		}
		para = append(para, trimmed)
	}
	flush()
	return strings.Join(out, "\n")
}
