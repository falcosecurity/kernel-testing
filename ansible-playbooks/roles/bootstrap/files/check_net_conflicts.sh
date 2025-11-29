#!/usr/bin/env bash
set -euo pipefail

check_duplicates() {
    FIELD="$1"
    DUPLICATES=""

    DUPLICATES=$(jq -r "
        group_by(.${FIELD})
        | map(select(length > 1))
        | map(.[0].${FIELD})
        | .[]
    " <<< "$TAP_DEV_MAP" || true)

    if [[ -n "$DUPLICATES" ]]; then
        echo "Error: Duplicate ${FIELD} values detected: $DUPLICATES" >&2
        exit 1
    fi
}

TAP_DEV_MAP="$1"

# Iterate over JSON objects one by one and check for any external conflict.
jq -r '.[] | "\(.name) \(.host_ip)"' <<<"$TAP_DEV_MAP" | while read -r TAP_NAME HOST_IP; do

  # Check tap name conflict.
  if ip link show "$TAP_NAME" >/dev/null 2>&1; then
      echo "Error: TAP device '$TAP_NAME' already exists" >&2
      exit 1
  fi

  # Check host IP address conflict.
  if ! ip -o addr show to "$HOST_IP" >/dev/null 2>&1; then
      echo "Error: Host already has an IP $HOST_IP" >&2
      exit 1
  fi
done

# Check for any internal conflict in names and host IPs, separately.
check_duplicates "name"
check_duplicates "host_ip"