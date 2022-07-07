package fluidrelayservers

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FluidRelayServerUpdate struct {
	Identity   *identity.SystemAndUserAssignedMap `json:"identity,omitempty"`
	Location   *string                            `json:"location,omitempty"`
	Properties *FluidRelayServerUpdateProperties  `json:"properties,omitempty"`
	Tags       *map[string]string                 `json:"tags,omitempty"`
}
