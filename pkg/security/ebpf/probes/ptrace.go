// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// +build linux

package probes

import manager "github.com/DataDog/ebpf-manager"

// ptraceProbes holds the list of probes used to track ptrace events
var ptraceProbes []*manager.Probe

func getPTraceProbes() []*manager.Probe {
	ptraceProbes = append(ptraceProbes, ExpandSyscallProbes(&manager.Probe{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID: SecurityAgentUID,
		},
		SyscallFuncName: "ptrace",
	}, EntryAndExit)...)
	return ptraceProbes
}