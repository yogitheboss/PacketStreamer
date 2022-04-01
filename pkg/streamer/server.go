package streamer

import "crypto/tls"

func getTlsConfig(certParam string, keyParam string, rootParam string) (*tls.Config, error) {

	config := &tls.Config{}
	certInfo, err := loadCertificates(certParam, keyParam, rootParam)
	if err != nil {
		return nil, err
	}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0] = certInfo
	config.InsecureSkipVerify = true
	config.CipherSuites = []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256}
	//Use only TLS v1.2
	config.MinVersion = tls.VersionTLS12
	//Don't allow session resumption
	config.SessionTicketsDisabled = true
	return config, nil
}

func loadCertificates(certParam string, keyParam string, rootParam string) (tls.Certificate, error) {
	mycert, err := tls.LoadX509KeyPair(certParam, keyParam)
	if err != nil {
		return mycert, err
	}
	return mycert, nil
}
