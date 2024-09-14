package dyscinfo

import (
	"fmt"
	"log/slog"

	"github.com/shirou/gopsutil/disk"
)

type DriverInfo struct {
	DriveLetter string
	VolumeName  string
	FileSystem  string
	TotalSize   uint64
	FreeSpace   uint64
}

func New() *DriverInfo {
	return &DriverInfo{}
}

func (d *DriverInfo) GetLogicalDrivesInfo() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		slog.Error("Errors", "err", err)
		panic(err)
	}
	for _, partition := range partitions {
		fmt.Printf("Диск: %s\n", partition.Device)
		fmt.Printf("Точка монтирования: %s\n", partition.Mountpoint)
		fmt.Printf("Файловая система: %s\n", partition.Fstype)

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			slog.Error("Errors", "err", err)
			panic(err)
		}
		fmt.Printf("Метка тома: %s\n", usage.Fstype)
		fmt.Printf("Размер: %.2f GB\n", float64(usage.Total)/1024/1024/1024)
		fmt.Printf("Свободно: %.2f GB\n", float64(usage.Free)/1024/1024/1024)
		fmt.Printf("Используется: %.2f GB\n\n", float64(usage.Used)/1024/1024/1024)
	}

}
