package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProposedCertificateRevocationKeyPrefix is the prefix to retrieve all ProposedCertificateRevocation.
	ProposedCertificateRevocationKeyPrefix = "ProposedCertificateRevocation/value/"
)

// ProposedCertificateRevocationKey returns the store key to retrieve a ProposedCertificateRevocation from the index fields.
func ProposedCertificateRevocationKey(
	subject string,
	subjectKeyId string,
) []byte {
	var key []byte

	subjectBytes := []byte(subject)
	key = append(key, subjectBytes...)
	key = append(key, []byte("/")...)

	subjectKeyIdBytes := []byte(subjectKeyId)
	key = append(key, subjectKeyIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
