package computer

import (
	"reflect"
	"testing"
)

func TestReadPacket(t *testing.T) {
	t.Run("Literal Packet", func(t *testing.T) {
		got := ReadBitsPacket("D2FE28")
		want := BitsPacket{Version: "110", Type: "100", Contents: "011111100101"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestReadPacket() = %v, want %v", got, want)
		}
	})
	t.Run("Bits Length Operator", func(t *testing.T) {
		got := ReadBitsPacket("38006F45291200")
		want := BitsPacket{Version: "001", Type: "110", SubPackets: []BitsPacket{
			{Version: "110", Type: "100", Contents: "1010"},
			{Version: "010", Type: "100", Contents: "00010100"},
		}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestReadPacket() = %v, want %v", got, want)
		}
	})
	t.Run("Packet Length Operator", func(t *testing.T) {
		got := ReadBitsPacket("EE00D40C823060")
		want := BitsPacket{Version: "111", Type: "011", SubPackets: []BitsPacket{
			{Version: "010", Type: "100", Contents: "0001"},
			{Version: "100", Type: "100", Contents: "0010"},
			{Version: "001", Type: "100", Contents: "0011"},
		}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestReadPacket() = %v, want %v", got, want)
		}
	})
}

func TestVersionCheckSum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "8A004A801A8002F478", want: 16},
		{input: "620080001611562C8802118E34", want: 12},
		{input: "C0015000016115A2E0802F182340", want: 23},
		{input: "A0016C880162017C3686B18A3D4780", want: 31},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			packet := ReadBitsPacket(tt.input)
			if got := BitsPacketVersionChecksum(packet); got != tt.want {
				t.Errorf("BitsPacketVersionChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "C200B40A82", want: 3},
		{input: "04005AC33890", want: 54},
		{input: "880086C3E88112", want: 7},
		{input: "CE00C43D881120", want: 9},
		{input: "D8005AC2A8F0", want: 1},
		{input: "F600BC2D8F", want: 0},
		{input: "9C005AC2F8F0", want: 0},
		{input: "9C0141080250320F1802104A08", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			packet := ReadBitsPacket(tt.input)
			if got := packet.Value(); got != tt.want {
				t.Errorf("BitsPacketVersionChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
