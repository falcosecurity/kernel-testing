#!/usr/bin/env bash
set -euo pipefail

calc_md5_hash() {
    RUN_ID="$1"
    VM_ID="$2"

    printf "%s" "${RUN_ID}-${VM_ID}" | md5sum | cut -d ' ' -f1
}

compute_tap() {
  RUN_ID="$1"
  VM_ID="$2"

  # Compute hash and get the first 12 characters.
  ID=$(calc_md5_hash "$RUN_ID" "$VM_ID" | cut -c1-12)

  # Add "tap" prefix.
  echo "tap$ID"
}

compute_addresses() {
  RUN_ID="$1"
  VM_ID="$2"

  hash=$(calc_md5_hash "$RUN_ID" "$VM_ID")

  # Convert last 4 hex chars to integer and mask to 14 bits.
  subnet_idx=$(( 0x${hash:28:4} & 0x3FFF ))

  # Compute the third and fourth octet (each /30 advances by 4 in the last octet).
  THIRD_OCTET=$(( subnet_idx / 64 ))
  FORTH_OCTET=$(( (subnet_idx % 64) * 4 ))
  HOST_IP="172.16.$THIRD_OCTET.$(( FORTH_OCTET + 1 ))"
  GUEST_IP="172.16.$THIRD_OCTET.$(( FORTH_OCTET + 2 ))"
  echo "$HOST_IP $GUEST_IP"
}

RUN_ID="$1"
VM_IDS="$2"

# Create arrays to pass to jq later.
JQ_ARGS=()
JQ_CODE="{}"

for VM_ID in $VM_IDS; do
  read -r HOST_IP GUEST_IP <<< "$(compute_addresses "$RUN_ID" "$VM_ID")"
  TAP=$(compute_tap "$RUN_ID" "$VM_ID")

  SAFE_VM_ID="${VM_ID//[^a-zA-Z0-9_]/_}"

  # Prepare named jq args.
  JQ_ARGS+=( --arg "name_$SAFE_VM_ID" "$TAP" )
  JQ_ARGS+=( --arg "host_ip_$SAFE_VM_ID" "$HOST_IP" )
  JQ_ARGS+=( --arg "guest_ip_$SAFE_VM_ID" "$GUEST_IP" )

  # Extend jq program. E.g.: .["1"] = {name: $name_1, host_ip: $host_ip_1, guest_ip: $guest_ip_1}.
  JQ_CODE+=" | .[\"$VM_ID\"] = {name: \$name_$SAFE_VM_ID, host_ip: \$host_ip_$SAFE_VM_ID, guest_ip: \$guest_ip_$SAFE_VM_ID}"
done

# Build the final JSON result.
jq -n "${JQ_ARGS[@]}" "$JQ_CODE"