package labs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportLabVirtualMachineRequest struct {
	DestinationVirtualMachineName  *string `json:"destinationVirtualMachineName,omitempty"`
	SourceVirtualMachineResourceId *string `json:"sourceVirtualMachineResourceId,omitempty"`
}
