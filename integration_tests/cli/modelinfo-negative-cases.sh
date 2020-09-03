#!/bin/bash
set -e
source integration_tests/cli/common.sh

# Preparation of Actors

echo "Create regular account"
create_new_account user_account ""

echo "Create Vendor account"
create_new_account vendor_account "Vendor"

# Body

# Ledger side errors

vid=$RANDOM
pid=$RANDOM

echo "Add Model with VID: $vid PID: $pid: Not Vendor"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $user_account --yes)
check_response_and_report "$result" "\"success\": false"
check_response_and_report "$result" "\"code\": 4"
echo "$result"

echo "Add Model with VID: $vid PID: $pid: Twice"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes)
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Device #1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes)
check_response_and_report "$result" "\"success\": false"
check_response_and_report "$result" "\"code\": 101"
echo "$result"

# CLI side errors

echo "Add Model with VID: $vid PID: $pid: Unknown account"
result=$(zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Name 1" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from "Unknown"  2>&1) || true
check_response_and_report "$result" "Key Unknown not found"

echo "Add model with invalid VID/PID"
for i in "0" "string"
do
  result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$i --pid=$pid --name="Name" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes 2>&1) || true
  check_response_and_report "$result" "Invalid VID"

  result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$i --name="Name" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes 2>&1) || true
  check_response_and_report "$result" "Invalid PID"
done

echo "Add model with epmty name"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid Name"

echo "Add model with epmty description"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Name" --description="" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid Description"

echo "Add model with epmty SKU"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Name" --description="Description" --sku="" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid SKU"

echo "Add model with epmty Firmwere Version"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Name" --description="Description" --sku="SKU12FS" --firmware-version="" --hardware-version="2.0" --tis-or-trp-testing-completed=true --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid FirmwareVersion"

echo "Add model with epmty Hardware Version"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Name" --description="Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="" --tis-or-trp-testing-completed=true --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid HardwareVersion"

echo "Add model with Invalid TIS flag"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Name" --description="Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed="string" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid Tis-or-trp-testing-completed"

echo "Add model without --from flag"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="Name" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --tis-or-trp-testing-completed=true --yes 2>&1) || true
check_response_and_report "$result" "required flag(s) \"from\" not set"

echo "Add model with not enought parameters"
result=$(echo "test1234" | zblcli tx modelinfo add-model --vid=$vid --pid=$pid --name="" --description="Device Description" --sku="SKU12FS" --firmware-version="1.0" --hardware-version="2.0" --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "required flag(s) \"tis-or-trp-testing-completed\" not set"