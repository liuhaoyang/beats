// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_aix.go

package aix

type prcred struct {
	Euid    uint64
	Ruid    uint64
	Suid    uint64
	Egid    uint64
	Rgid    uint64
	Sgid    uint64
	X_pad   [8]uint64
	X_pad1  uint32
	Ngroups uint32
	Groups  [1]uint64
}

type pstatus struct {
	Flag            uint32
	Flag2           uint32
	Flags           uint32
	Nlwp            uint32
	Stat            uint8
	Dmodel          uint8
	X_pad1          [6]uint8
	Sigpend         prSigset
	Brkbase         uint64
	Brksize         uint64
	Stkbase         uint64
	Stksize         uint64
	Pid             uint64
	Ppid            uint64
	Pgid            uint64
	Sid             uint64
	Utime           prTimestruc64
	Stime           prTimestruc64
	Cutime          prTimestruc64
	Cstime          prTimestruc64
	Sigtrace        prSigset
	Flttrace        fltset
	Sysentry_offset uint32
	Sysexit_offset  uint32
	X_pad           [8]uint64
	Lwp             lwpstatus
}
type prTimestruc64 struct {
	Sec    int64
	Nsec   int32
	X__pad uint32
}
type prSigset struct {
	Set [4]uint64
}
type fltset struct {
	Set [4]uint64
}
type lwpstatus struct {
	Lwpid    uint64
	Flags    uint32
	X_pad1   [1]uint8
	State    uint8
	Cursig   uint16
	Why      uint16
	What     uint16
	Policy   uint32
	Clname   [8]uint8
	Lwppend  prSigset
	Lwphold  prSigset
	Info     prSiginfo64
	Altstack prStack64
	Action   prSigaction64
	X_pad2   uint32
	Syscall  uint16
	Nsysarg  uint16
	Sysarg   [8]uint64
	Errno    int32
	Ptid     uint32
	X_pad    [9]uint64
	Reg      prgregset
	Fpreg    prfpregset
	Family   pfamily
}
type prSiginfo64 struct {
	Signo   int32
	Errno   int32
	Code    int32
	Imm     int32
	Status  int32
	X__pad1 uint32
	Uid     uint64
	Pid     uint64
	Addr    uint64
	Band    int64
	Value   [8]byte
	X__pad  [4]uint32
}
type prStack64 struct {
	Sp     uint64
	Size   uint64
	Flags  int32
	X__pad [5]int32
}
type prSigaction64 struct {
	Union  [8]byte
	Mask   prSigset
	Flags  int32
	X__pad [5]int32
}
type prgregset struct {
	X__iar    uint64
	X__msr    uint64
	X__cr     uint64
	X__lr     uint64
	X__ctr    uint64
	X__xer    uint64
	X__fpscr  uint64
	X__fpscrx uint64
	X__gpr    [32]uint64
	X__pad1   [8]uint64
}
type prfpregset struct {
	X__fpr [32]float64
}
type pfamily struct {
	Extoff  uint64
	Extsize uint64
	Pad     [14]uint64
}
