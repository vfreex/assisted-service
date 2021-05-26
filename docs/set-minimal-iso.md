# Generating a minimal ISO
The default discovery ISO generated by the Assisted Installer is a complete RHCOS ISO with a custom ignition file.

Serving this large file through OOBM (e.g. BMC, iDRAC) can sometimes lead to unexpected issues.

To overcome this, the Assisted Installer offers an option to generate a smaller ISO file that dynamically downloads the rest of the RHCOS rootfs via the Internet.

The following script provides an example of how to use the Assisted Installer API in order to set the type of the discovery ISO to "minimal-iso".

For more information about the API and its various authentication methods, see [this document](cloud.md).

# Set discovery ISO type
```
#!/bin/bash

set -euo pipefail

if ! ocm token 2>/dev/null >/dev/null; then
    echo "Failed to run 'ocm token' command, please see the assisted_service/docs/cloud.md doc for authentication information"
    exit 1
fi

if [ -z ${SSH_KEY+x} ]; then
	echo 'Please set SSH_KEY to your SSH public key.'
	echo 'For example: export SSH_KEY=$(cat ~/.ssh/id_rsa.pub)'
	exit 1
fi

if [ -z ${CLUSTER_ID+x} ]; then
	echo 'Please set CLUSTER_ID to your cluster ID, which can be found in the Assisted Installer URL'
	exit 1
fi

## User specific configuration <-----------
TOKEN=$(ocm token)
OCM_API_ENDPOINT="https://api.openshift.com/api/" 
###############################

function log() {
    if [[ ! $? == 0 ]]; then
        echo "Script enountered an error"
        exit 1
    fi
}

trap log EXIT

echo Telling service to generate a minimal ISO with our public SSH key file
curl -fail -s $OCM_API_ENDPOINT/assisted-install/v1/clusters/$CLUSTER_ID/downloads/image -H "Authorization: Bearer $TOKEN" --request POST --header "Content-Type: application/json" --data @<(echo '{"image_type": "minimal-iso", "ssh_public_key": ""}' | jq --rawfile pubkey <(echo -n $SSH_KEY) '.ssh_public_key = $pubkey')

echo "Done, retrieving ISO url..."

curl -s $OCM_API_ENDPOINT/assisted-install/v1/clusters/$CLUSTER_ID -H "Authorization: Bearer $TOKEN" --request GET | jq '.image_info.download_url' -r
```