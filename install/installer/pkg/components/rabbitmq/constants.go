// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package rabbitmq

import (
	"github.com/gitpod-io/gitpod/installer/pkg/common"
)

const (
	Component             = "rabbitmq"
	CookieSecret          = "messagebus-erlang-cookie"
	InClusterMsgBusSecret = common.InClusterMessageQueueName
	TLSSecret             = common.InClusterMessageQueueTLS
	rabbitMQUsername      = "gitpod"
	InitContainerImage    = "library/alpine"
	InitContainerTag      = "3.16"
	passwordReplaceString = "%PASSWORD%"
)