/*
output "ip_collection" {
  value = ochk_ip_collection.default
}

output "subtenant_1" {
  value = ochk_subtenant.subtenant-1
}

output "subtenant_2" {
  value = ochk_subtenant.subtenant-2
}

output "subtenant_3" {
  value = ochk_subtenant.subtenant-3
}

output "subtenant_4" {
  value = ochk_subtenant.subtenant-4
}
*/

/*
output "subtenant-vpc" {
  value = ochk_router.subtenant-vpc
}
*/

output "network" {
  value = data.ochk_network.subtenant-network
}

output "user1" {
  value = data.ochk_user.bg_manager_user
}

output "vnet" {
  value = ochk_virtual_network.network_for_vm
}

output "billing-tag" {
  value = ochk_billing_tag.res-bt-cc2
}

output "system-tag" {
  value = ochk_system_tag.res-st-os2
}
/*
output "vm" {
  value = ochk_virtual_machine.default
}
*/

/*
output "aes-generated" {
  value = ochk_kms_key.aes-generated
  sensitive = true
}

output "rsa-to-unwrap" {
  value = ochk_kms_key.rsa-to-unwrap
  sensitive = true
}

output "aes-enc" {
  value = ochk_kms_key.aes-enc
  sensitive = true
}
*/
/*
output "encrypted-with-own-encrypted-key" {
  value = ochk_virtual_machine.encrypted-with-own-encrypted-key
}

output "encrypted-own-key" {
  value = ochk_virtual_machine.encrypted
}
*/


