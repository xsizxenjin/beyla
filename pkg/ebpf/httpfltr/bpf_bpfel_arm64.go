// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64
// +build arm64

package httpfltr

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfAcceptArgsT struct {
	Addr       uint64
	AcceptTime uint64
}

type bpfHttpConnectionInfoT struct {
	S_h    uint64
	S_l    uint64
	D_h    uint64
	D_l    uint64
	S_port uint16
	D_port uint16
	Flags  uint32
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	KretprobeSockAlloc  *ebpf.ProgramSpec `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4 *ebpf.ProgramSpec `ebpf:"kretprobe_sys_accept4"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	ActiveAcceptArgs    *ebpf.MapSpec `ebpf:"active_accept_args"`
	Events              *ebpf.MapSpec `ebpf:"events"`
	FilteredConnections *ebpf.MapSpec `ebpf:"filtered_connections"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	ActiveAcceptArgs    *ebpf.Map `ebpf:"active_accept_args"`
	Events              *ebpf.Map `ebpf:"events"`
	FilteredConnections *ebpf.Map `ebpf:"filtered_connections"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.ActiveAcceptArgs,
		m.Events,
		m.FilteredConnections,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	KretprobeSockAlloc  *ebpf.Program `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4 *ebpf.Program `ebpf:"kretprobe_sys_accept4"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.KretprobeSockAlloc,
		p.KretprobeSysAccept4,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_arm64.o
var _BpfBytes []byte
