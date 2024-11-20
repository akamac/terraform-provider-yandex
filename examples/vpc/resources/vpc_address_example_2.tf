resource "yandex_vpc_address" "vpnaddr" {
  name = "vpnaddr"

  external_ipv4_address {
    zone_id                  = "ru-central1-a"
    ddos_protection_provider = "qrator"
  }
}