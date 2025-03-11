# Connection to Proxmox VE
pve_url          = "https://<PROXMOX ADDRESS>:8006"
pve_insecure_tls = false
pve_token_id     = ""
pve_token_secret = ""

# Proxmox VE node to create the template on
pve_node = ""

# Proxmox VE resource pool to create the template in
pve_pool = ""

# Optional: Proxmox VE ID to use for the template
# pve_id = null

# Optional: Proxmox VE name to use for the template and OS hostname
# pve_name = "ubuntu-server-24.04-cloudinit"

# Proxmox VE storage to use for OS ISO
pve_storage_iso = ""

# Proxmox VE storage to use for template's HDD
pve_storage_disk = ""

# Proxmox VE network bridge to use for the template's NIC
pve_network_bridge = "vmbr0"

# Username for SSH user
ssh_username = "service"

# SSH Public Key (as string) that will be added to SSH user's ~/.ssh/authorized_keys
# NOTE: corresponding SSH identity must be present in the local SSH agent when building the template
ssh_authorized_key = ""
