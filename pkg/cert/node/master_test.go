package node

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseKubeadmCertsCheckExpiration(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect map[string]time.Time
	}{
		{
			name: "kubeadm 1.15.2 and 1.16.2",
			input: `
CERTIFICATE                EXPIRES                  RESIDUAL TIME   EXTERNALLY MANAGED
admin.conf                 May 12, 2021 02:29 UTC   364d            no
apiserver                  May 12, 2021 02:29 UTC   364d            no
apiserver-etcd-client      May 12, 2021 02:29 UTC   364d            no
apiserver-kubelet-client   May 12, 2021 02:29 UTC   364d            no
controller-manager.conf    May 12, 2021 02:29 UTC   364d            no
etcd-healthcheck-client    May 12, 2021 02:29 UTC   364d            no
etcd-peer                  May 12, 2021 02:29 UTC   364d            no
etcd-server                May 12, 2021 02:29 UTC   364d            no
front-proxy-client         May 12, 2021 02:29 UTC   364d            no
scheduler.conf             May 12, 2021 02:29 UTC   364d            no		
			`,
			expect: map[string]time.Time{
				"admin.conf":               time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"apiserver":                time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"apiserver-etcd-client":    time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"apiserver-kubelet-client": time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"controller-manager.conf":  time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"etcd-healthcheck-client":  time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"etcd-peer":                time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"etcd-server":              time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"front-proxy-client":       time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"scheduler.conf":           time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
			},
		},
		{
			name: "kubeadm 1.17.4",
			input: `
CERTIFICATE                EXPIRES                  RESIDUAL TIME   CERTIFICATE AUTHORITY   EXTERNALLY MANAGED
admin.conf                 May 12, 2021 02:29 UTC   364d                                    no
apiserver                  May 12, 2021 02:29 UTC   364d            ca                      no
apiserver-etcd-client      May 12, 2021 02:29 UTC   364d            etcd-ca                 no
apiserver-kubelet-client   May 12, 2021 02:29 UTC   364d            ca                      no
controller-manager.conf    May 12, 2021 02:29 UTC   364d                                    no
etcd-healthcheck-client    May 12, 2021 02:29 UTC   364d            etcd-ca                 no
etcd-peer                  May 12, 2021 02:29 UTC   364d            etcd-ca                 no
etcd-server                May 12, 2021 02:29 UTC   364d            etcd-ca                 no
front-proxy-client         May 12, 2021 02:29 UTC   364d            front-proxy-ca          no
scheduler.conf             May 12, 2021 02:29 UTC   364d                                    no

CERTIFICATE AUTHORITY   EXPIRES                  RESIDUAL TIME   EXTERNALLY MANAGED
ca                      Mar 11, 2030 01:51 UTC   9y              no
etcd-ca                 Mar 11, 2030 01:51 UTC   9y              no
front-proxy-ca          Mar 11, 2030 01:51 UTC   9y              no
					`,
			expect: map[string]time.Time{
				"admin.conf":               time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"apiserver":                time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"apiserver-etcd-client":    time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"apiserver-kubelet-client": time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"ca":                       time.Date(2030, time.March, 11, 01, 51, 00, 00, time.UTC),
				"controller-manager.conf":  time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"etcd-ca":                  time.Date(2030, time.March, 11, 01, 51, 00, 00, time.UTC),
				"etcd-healthcheck-client":  time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"etcd-peer":                time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"etcd-server":              time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"front-proxy-ca":           time.Date(2030, time.March, 11, 01, 51, 00, 00, time.UTC),
				"front-proxy-client":       time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
				"scheduler.conf":           time.Date(2021, time.May, 12, 02, 29, 00, 00, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := parseKubeadmCertsCheckExpiration(tt.input)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("got %v is not equals to expected", got)
			}
		})
	}
}
