package qbranch

type BondDataType uint8

const (
	BT_STOP        BondDataType = 0
	BT_STOP_BASE   BondDataType = 1
	BT_BOOL        BondDataType = 2
	BT_UINT8       BondDataType = 3
	BT_UINT16      BondDataType = 4
	BT_UINT32      BondDataType = 5
	BT_UINT64      BondDataType = 6
	BT_FLOAT       BondDataType = 7
	BT_DOUBLE      BondDataType = 8
	BT_STRING      BondDataType = 9
	BT_STRUCT      BondDataType = 10
	BT_LIST        BondDataType = 11
	BT_SET         BondDataType = 12
	BT_MAP         BondDataType = 13
	BT_INT8        BondDataType = 14
	BT_INT16       BondDataType = 15
	BT_INT32       BondDataType = 16
	BT_INT64       BondDataType = 17
	BT_WSTRING     BondDataType = 18
	BT_UNAVAILABLE BondDataType = 127
)
