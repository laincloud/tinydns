package dns

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"
)

func TestPrepareData(t *testing.T) {
	coreInfoData := `
    {
  "console.web.web": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [\"console.lain.local\", \"console.lain\"], \"https_only\": false}",
        "ContainerInfos": [
          {
            "Command": [
              "./entry.sh"
            ],
            "ContainerId": "bbcba375ed9b6076f2201d92f1d3a17cce394fcb78c26bb174b64f05eb5be4c5",
            "ContainerIp": "172.20.0.4",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=console",
              "LAIN_APP_RELEASE_VERSION=1457588625-f8351bccae8d56ac32873aa4afc086ffa66b3033",
              "LAIN_PROCNAME=web",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=bfafd813d7ea65ee4db1f09d7c8ffbf4"
            ],
            "Expose": 8000,
            "HostInterfaceName": "cali20625806fae",
            "Image": "registry.lain.local/console:release-1457588625-f8351bccae8d56ac32873aa4afc086ffa66b3033",
            "Memory": 268435456,
            "NodeIp": "192.168.77.21",
            "NodeName": "node1",
            "Volumes": [
              "/lain/app/logs"
            ]
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      }
    ]
  },
  "deploy.web.web": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [\"deploy.lain.local\", \"deploy.lain\"], \"https_only\": false}",
        "ContainerInfos": [
          {
            "Command": [
              "/entry.sh"
            ],
            "ContainerId": "a791237d013db9d97afaa52f02868780804b9716a7f0eef319b9c02277c119bb",
            "ContainerIp": "172.20.0.2",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=deploy",
              "LAIN_APP_RELEASE_VERSION=1454387342-7e606525daa0659e9aff9a45b0d5c8ca57a03741",
              "LAIN_PROCNAME=web",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=078f40fa23e0672777adc7c05d4773dd"
            ],
            "Expose": 9000,
            "HostInterfaceName": "cali04c7c360fae",
            "Image": "registry.lain.local/deploy:release-1454387342-7e606525daa0659e9aff9a45b0d5c8ca57a03741",
            "Memory": 134217728,
            "NodeIp": "192.168.77.21",
            "NodeName": "node1",
            "Volumes": []
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      }
    ]
  },
  "lvault.web.web": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [\"lvault.lain.local\", \"lvault.lain\"], \"https_only\": false}",
        "ContainerInfos": [
          {
            "Command": [
              "sh",
              "run.sh"
            ],
            "ContainerId": "88375e621e7645a436c74d10ae5f0e24df1ea7c5efcae0c481750ce2d7209d22",
            "ContainerIp": "172.20.0.8",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=lvault",
              "LAIN_APP_RELEASE_VERSION=1453431538-d601791354f5b589bfa0e00c415b33701e0fa0e5",
              "LAIN_PROCNAME=web",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=e7f0b4806ffacca30a3dc8a294ee2dcc"
            ],
            "Expose": 8001,
            "HostInterfaceName": "cali452f9e1efae",
            "Image": "registry.lain.local/lvault:release-1453431538-d601791354f5b589bfa0e00c415b33701e0fa0e5",
            "Memory": 33554432,
            "NodeIp": "192.168.77.21",
            "NodeName": "node1",
            "Volumes": []
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      }
    ]
  },
  "perf.worker.loader": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [], \"https_only\": true}",
        "ContainerInfos": [
          {
            "Command": [
              "tail",
              "-f",
              "/etc/hosts"
            ],
            "ContainerId": "d7332b577fa497ab1a217a9ef5a817e9ddaa8d7245d34dc5ce53d2bc04fa313b",
            "ContainerIp": "172.20.0.28",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=perf",
              "LAIN_APP_RELEASE_VERSION=1459406828-c1acaaaeda48e323de5915bcd78c33bed7877b33",
              "LAIN_PROCNAME=loader",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=3a0b03ba8f3b6c0cd97ca15da2d57731"
            ],
            "Expose": 0,
            "HostInterfaceName": "calia995563afae",
            "Image": "registry.lain.local/perf:release-1459406828-c1acaaaeda48e323de5915bcd78c33bed7877b33",
            "Memory": 33554432,
            "NodeIp": "192.168.77.22",
            "NodeName": "node2",
            "Volumes": []
          }
        ],
        "Dependencies": [
          {
            "PodName": "resource.hello-server.perf.portal.portal-hello",
            "Policy": 0
          }
        ],
        "InstanceNo": 1
      }
    ]
  },
  "registry.web.web": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [\"registry.lain.local\", \"registry.lain\"], \"https_only\": false}",
        "ContainerInfos": [
          {
            "Command": [
              "./entry.sh"
            ],
            "ContainerId": "20858a896d6dbddea19bdf6d9ffd1bc82afb10495b43e9f8ef25426f2ba828d0",
            "ContainerIp": "172.20.0.1",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=registry",
              "LAIN_APP_RELEASE_VERSION=1457409719-821051f5ecc4d6a3b9f9d2165fb01f99c038003b",
              "LAIN_PROCNAME=web",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=a9205dcfd4a6f7c2cbe8be01566ff84a"
            ],
            "Expose": 5000,
            "HostInterfaceName": "cali1bc07ba4fae",
            "Image": "registry.lain.local/registry:release-1457409719-821051f5ecc4d6a3b9f9d2165fb01f99c038003b",
            "Memory": 134217728,
            "NodeIp": "192.168.77.21",
            "NodeName": "node1",
            "Volumes": []
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      }
    ]
  },
  "resource.hello-server.perf.worker.hello": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [], \"https_only\": true}",
        "ContainerInfos": [
          {
            "Command": [
              "hello"
            ],
            "ContainerId": "b3dec934ace1f83d00b7bf96d95b45f7876e1e046a7e012224060403d63b5fed",
            "ContainerIp": "172.20.0.25",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=resource.hello-server.perf",
              "LAIN_APP_RELEASE_VERSION=1459831485-c09832803a81a9b8d5e7275fe10065ffae0ce266",
              "LAIN_PROCNAME=hello",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=57f2f3cb023b5458a4b75e9bd522a62a"
            ],
            "Expose": 80,
            "HostInterfaceName": "cali8a3a106efae",
            "Image": "registry.lain.local/hello-server:release-1459831485-c09832803a81a9b8d5e7275fe10065ffae0ce266",
            "Memory": 104857600,
            "NodeIp": "192.168.77.22",
            "NodeName": "node2",
            "Volumes": []
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      },
      {
        "Annotation": "{\"mountpoint\": [], \"https_only\": true}",
        "ContainerInfos": [
          {
            "Command": [
              "hello"
            ],
            "ContainerId": "6952ccc46fa2554d0f1f321b466e3c211671a167f0a597b65eebd2d68b4cd7e3",
            "ContainerIp": "172.20.0.26",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=resource.hello-server.perf",
              "LAIN_APP_RELEASE_VERSION=1459831485-c09832803a81a9b8d5e7275fe10065ffae0ce266",
              "LAIN_PROCNAME=hello",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=57f2f3cb023b5458a4b75e9bd522a62a"
            ],
            "Expose": 80,
            "HostInterfaceName": "cali8af9242cfae",
            "Image": "registry.lain.local/hello-server:release-1459831485-c09832803a81a9b8d5e7275fe10065ffae0ce266",
            "Memory": 104857600,
            "NodeIp": "192.168.77.22",
            "NodeName": "node2",
            "Volumes": []
          }
        ],
        "Dependencies": [],
        "InstanceNo": 2
      }
    ]
  },
  "tinydns.worker.worker": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [], \"logs\": [\"tinydns.log\", \"tinydns.err\", \"dns-creator.log\", \"dns-creator.err\"], \"https_only\": true}",
        "ContainerInfos": [
          {
            "Command": [
              "/usr/bin/supervisord",
              "-c",
              "/lain/app/supervisord.conf"
            ],
            "ContainerId": "2cce3ec16348c7d6ce6b5a354e3245791a354935fa3a33643d4919f1ddfb0b78",
            "ContainerIp": "172.20.0.3",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=tinydns",
              "LAIN_APP_RELEASE_VERSION=1458208956-d914d56fc289917378709f7e5a18d1646790a8b4",
              "LAIN_PROCNAME=worker",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=fe92dd4d9af20cbfa4ba98f7e29d856a"
            ],
            "Expose": 53,
            "HostInterfaceName": "cali16c1cdfefae",
            "Image": "registry.lain.local/tinydns:release-1458208956-d914d56fc289917378709f7e5a18d1646790a8b4",
            "Memory": 268435456,
            "NodeIp": "192.168.77.21",
            "NodeName": "node1",
            "Volumes": [
              "/var/log/supervisor",
              "/lain/logs"
            ]
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      }
    ]
  },
  "vault.web.vaultProxy": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [\"vault.lain.local/proxy\", \"vault.lain/proxy\"], \"https_only\": false}",
        "ContainerInfos": [
          {
            "Command": [
              "sh",
              "proxy.sh"
            ],
            "ContainerId": "36e2e5c2f8680583dc56e01f6d9b5feab780390d016f9ef10def7ff30d630e7c",
            "ContainerIp": "172.20.0.7",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=vault",
              "LAIN_APP_RELEASE_VERSION=1453431538-d601791354f5b589bfa0e00c415b33701e0fa0e5",
              "LAIN_PROCNAME=vaultProxy",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=184aa077df08b90ac9fe282cceaa325e"
            ],
            "Expose": 8200,
            "HostInterfaceName": "cali3fb047eafae",
            "Image": "registry.lain.local/vault:release-1453431538-d601791354f5b589bfa0e00c415b33701e0fa0e5",
            "Memory": 33554432,
            "NodeIp": "192.168.77.21",
            "NodeName": "node1",
            "Volumes": []
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      }
    ]
  },
  "vault.web.web": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [\"vault.lain.local\", \"vault.lain\"], \"https_only\": false}",
        "ContainerInfos": [
          {
            "Command": [
              "sh",
              "run.sh"
            ],
            "ContainerId": "503ce3d9d4797874562f23903f718bb19e9824767518fde17b339f2e2309cf35",
            "ContainerIp": "172.20.0.6",
            "Cpu": 0,
            "Env": [
              "VAULT_ADDR=http://127.0.0.1:8200",
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=vault",
              "LAIN_APP_RELEASE_VERSION=1453431538-d601791354f5b589bfa0e00c415b33701e0fa0e5",
              "LAIN_PROCNAME=web",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=184aa077df08b90ac9fe282cceaa325e"
            ],
            "Expose": 8200,
            "HostInterfaceName": "cali3eee739afae",
            "Image": "registry.lain.local/vault:release-1453431538-d601791354f5b589bfa0e00c415b33701e0fa0e5",
            "Memory": 33554432,
            "NodeIp": "192.168.77.21",
            "NodeName": "node1",
            "Volumes": [
              "/lain/app/log"
            ]
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      }
    ]
  },
  "webrouter.worker.worker": {
    "PodInfos": [
      {
        "Annotation": "{\"mountpoint\": [], \"https_only\": true}",
        "ContainerInfos": [
          {
            "Command": [
              "/usr/bin/supervisord"
            ],
            "ContainerId": "6d0646b42184bd693a327c58da9225d99aba81a04960fe198c99a536f2d96ea8",
            "ContainerIp": "172.20.0.5",
            "Cpu": 0,
            "Env": [
              "TZ=Asia/Shanghai",
              "LAIN_APPNAME=webrouter",
              "LAIN_APP_RELEASE_VERSION=1458717572-03198de50fe90803a8379277fb52a8a51203d2ce",
              "LAIN_PROCNAME=worker",
              "LAIN_DOMAIN=lain.local",
              "CALICO_IP=AUTO",
              "CALICO_PROFILE=3015557ac95c5bf5fbd8bcea08c58a2d"
            ],
            "Expose": 0,
            "HostInterfaceName": "cali30dece08fae",
            "Image": "registry.lain.local/webrouter:release-1458717572-03198de50fe90803a8379277fb52a8a51203d2ce",
            "Memory": 268435456,
            "NodeIp": "192.168.77.21",
            "NodeName": "node1",
            "Volumes": [
              "/etc/nginx/nginx.conf",
              "/etc/nginx/proxy.conf",
              "/etc/nginx/default.conf",
              "/etc/nginx/conf.d",
              "/etc/nginx/upstreams",
              "/etc/nginx/locations",
              "/etc/nginx/buffer",
              "/etc/nginx/ssl",
              "/var/log/nginx",
              "/var/log/watcher",
              "/var/log/supervisor"
            ]
          }
        ],
        "Dependencies": [],
        "InstanceNo": 1
      }
    ]
  }
}
    `
	dependsData := `
    {
  "resource.hello-server.perf.portal.portal-hello": {
    "node2": {
      "perf": {
        "Annotation": "{\"mountpoint\": [], \"https_only\": true, \"service_name\": \"hello\"}",
        "Containers": [
          {
            "NodeIP": "192.168.77.22",
            "IP": "172.20.0.29",
            "Port": 10000
          }
        ]
      }
    }
  }
}
    `
	fqdnsData := `
    {
        "tinydns_fqdns/lain": "[\".lain:192.168.77.202:a:300\"]",
        "tinydns_fqdns/20.172.in-addr.arpa": "[\".20.172.in-addr.arpa:192.168.77.202:a:300\"]",
        "tinydns_fqdns/webrouter.lain": "[\"+webrouter.lain:192.168.77.201:300::\"]"
    }
    `
	expect := []string{
		"=worker-1.webrouter.lain:172.20.0.5:300",
		"=web-1.deploy.lain:172.20.0.2:300",
		"=loader-1.perf.lain:172.20.0.28:300",
		"=hello-2.perf.hello-server.resource.lain:172.20.0.26:300",
		"=web-1.registry.lain:172.20.0.1:300",
		"=hello-1.perf.hello-server.resource.lain:172.20.0.25:300",
		"=web-1.console.lain:172.20.0.4:300",
		"=worker-1.tinydns.lain:172.20.0.3:300",
		"=vaultProxy-1.vault.lain:172.20.0.7:300",
		"=web-1.vault.lain:172.20.0.6:300",
		"=web-1.lvault.lain:172.20.0.8:300",
		"=hello.perf.lain:172.20.0.29:300::00",
		"+webrouter.lain:192.168.77.201:300::",
		".lain:192.168.77.202:a:300",
		".20.172.in-addr.arpa:192.168.77.202:a:300",
		"%00:192.168.77.22",
	}
	mockCreator := &Creator{
		depends:  make(Depends),
		coreInfo: make(CoreInfo),
		fqdns:    make(Fqdns),
	}
	json.Unmarshal([]byte(dependsData), &(mockCreator.depends))
	json.Unmarshal([]byte(coreInfoData), &(mockCreator.coreInfo))
	json.Unmarshal([]byte(fqdnsData), &(mockCreator.fqdns))
	actual := mockCreator.prepareData()
	sort.Strings(actual)
	sort.Strings(expect)
	if !reflect.DeepEqual(actual, expect) {
		t.Error("Array of lines is not correct\n")
	}
}
