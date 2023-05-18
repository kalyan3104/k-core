package core

// HeaderType defines the type to be used for the header that is sent
type HeaderType string

const (
	// MetaHeader defines the type of *block.MetaBlock
	MetaHeader HeaderType = "MetaBlock"
	// ShardHeaderV1 defines the type of *block.Header
	ShardHeaderV1 HeaderType = "Header"
	// ShardHeaderV2 defines the type of *block.HeaderV2
	ShardHeaderV2 HeaderType = "HeaderV2"
)

// NodeType represents the node's role in the network
type NodeType string

// NodeTypeObserver signals that a node is running as observer node
const NodeTypeObserver NodeType = "observer"

// NodeTypeValidator signals that a node is running as validator node
const NodeTypeValidator NodeType = "validator"

// pkPrefixSize specifies the max numbers of chars to be displayed from one publc key
const pkPrefixSize = 12

// FileModeUserReadWrite represents the permission for a file which allows the user for reading and writing
const FileModeUserReadWrite = 0600

// FileModeReadWrite represents the permission for a file which allows reading and writing for user and group and read
// for others
const FileModeReadWrite = 0664

// MetachainShardId will be used to identify a shard ID as metachain
const MetachainShardId = uint32(0xFFFFFFFF)

// AllShardId will be used to identify that a message is for all shards
const AllShardId = uint32(0xFFFFFFF0)

// MegabyteSize represents the size in bytes of a megabyte
const MegabyteSize = 1024 * 1024

// BuiltInFunctionClaimDeveloperRewards is the key for the claim developer rewards built-in function
const BuiltInFunctionClaimDeveloperRewards = "ClaimDeveloperRewards"

// BuiltInFunctionChangeOwnerAddress is the key for the change owner built in function built-in function
const BuiltInFunctionChangeOwnerAddress = "ChangeOwnerAddress"

// BuiltInFunctionSetUserName is the key for the set user name built-in function
const BuiltInFunctionSetUserName = "SetUserName"

// BuiltInFunctionSaveKeyValue is the key for the save key value built-in function
const BuiltInFunctionSaveKeyValue = "SaveKeyValue"

// BuiltInFunctionDCTTransfer is the key for the electronic standard digital token transfer built-in function
const BuiltInFunctionDCTTransfer = "DCTTransfer"

// BuiltInFunctionDCTBurn is the key for the electronic standard digital token burn built-in function
const BuiltInFunctionDCTBurn = "DCTBurn"

// BuiltInFunctionDCTFreeze is the key for the electronic standard digital token freeze built-in function
const BuiltInFunctionDCTFreeze = "DCTFreeze"

// BuiltInFunctionDCTUnFreeze is the key for the electronic standard digital token unfreeze built-in function
const BuiltInFunctionDCTUnFreeze = "DCTUnFreeze"

// BuiltInFunctionDCTWipe is the key for the electronic standard digital token wipe built-in function
const BuiltInFunctionDCTWipe = "DCTWipe"

// BuiltInFunctionDCTPause is the key for the electronic standard digital token pause built-in function
const BuiltInFunctionDCTPause = "DCTPause"

// BuiltInFunctionDCTUnPause is the key for the electronic standard digital token unpause built-in function
const BuiltInFunctionDCTUnPause = "DCTUnPause"

// BuiltInFunctionSetDCTRole is the key for the electronic standard digital token set built-in function
const BuiltInFunctionSetDCTRole = "DCTSetRole"

// BuiltInFunctionUnSetDCTRole is the key for the electronic standard digital token unset built-in function
const BuiltInFunctionUnSetDCTRole = "DCTUnSetRole"

// BuiltInFunctionDCTSetLimitedTransfer is the key for the electronic standard digital token built-in function which sets the property
// for the token to be transferable only through accounts with transfer roles
const BuiltInFunctionDCTSetLimitedTransfer = "DCTSetLimitedTransfer"

// BuiltInFunctionDCTUnSetLimitedTransfer is the key for the electronic standard digital token built-in function which unsets the property
// for the token to be transferable only through accounts with transfer roles
const BuiltInFunctionDCTUnSetLimitedTransfer = "DCTUnSetLimitedTransfer"

// BuiltInFunctionDCTLocalMint is the key for the electronic standard digital token local mint built-in function
const BuiltInFunctionDCTLocalMint = "DCTLocalMint"

// BuiltInFunctionDCTLocalBurn is the key for the electronic standard digital token local burn built-in function
const BuiltInFunctionDCTLocalBurn = "DCTLocalBurn"

// BuiltInFunctionDCTNFTTransfer is the key for the electronic standard digital token NFT transfer built-in function
const BuiltInFunctionDCTNFTTransfer = "DCTNFTTransfer"

// BuiltInFunctionDCTNFTCreate is the key for the electronic standard digital token NFT create built-in function
const BuiltInFunctionDCTNFTCreate = "DCTNFTCreate"

// BuiltInFunctionDCTNFTAddQuantity is the key for the electronic standard digital token NFT add quantity built-in function
const BuiltInFunctionDCTNFTAddQuantity = "DCTNFTAddQuantity"

// BuiltInFunctionDCTNFTCreateRoleTransfer is the key for the electronic standard digital token create role transfer function
const BuiltInFunctionDCTNFTCreateRoleTransfer = "DCTNFTCreateRoleTransfer"

// BuiltInFunctionDCTNFTBurn is the key for the electronic standard digital token NFT burn built-in function
const BuiltInFunctionDCTNFTBurn = "DCTNFTBurn"

// BuiltInFunctionDCTNFTAddURI is the key for the electronic standard digital token NFT add URI built-in function
const BuiltInFunctionDCTNFTAddURI = "DCTNFTAddURI"

// BuiltInFunctionDCTNFTUpdateAttributes is the key for the electronic standard digital token NFT update attributes built-in function
const BuiltInFunctionDCTNFTUpdateAttributes = "DCTNFTUpdateAttributes"

// BuiltInFunctionMultiDCTNFTTransfer is the key for the electronic standard digital token multi transfer built-in function
const BuiltInFunctionMultiDCTNFTTransfer = "MultiDCTNFTTransfer"

// BuiltInFunctionSetGuardian is the key for setting a guardian built-in function
const BuiltInFunctionSetGuardian = "SetGuardian"

// BuiltInFunctionFreezeAccount is the built-in function key for freezing an account
const BuiltInFunctionFreezeAccount = "FreezeAccount"

// BuiltInFunctionUnfreezeAccount is the built-in function key for unfreezing an account
const BuiltInFunctionUnfreezeAccount = "UnfreezeAccount"

// DCTRoleLocalMint is the constant string for the local role of mint for DCTtokens
const DCTRoleLocalMint = "DCTRoleLocalMint"

// DCTRoleLocalBurn is the constant string for the local role of burn for DCTtokens
const DCTRoleLocalBurn = "DCTRoleLocalBurn"

// DCTRoleNFTCreate is the constant string for the local role of create for DCTNFT tokens
const DCTRoleNFTCreate = "DCTRoleNFTCreate"

// DCTRoleNFTCreateMultiShard is the constant string for the local role of create for DCTNFT tokens multishard
const DCTRoleNFTCreateMultiShard = "DCTRoleNFTCreateMultiShard"

// DCTRoleNFTAddQuantity is the constant string for the local role of adding quantity for existing DCTNFT tokens
const DCTRoleNFTAddQuantity = "DCTRoleNFTAddQuantity"

// DCTRoleNFTBurn is the constant string for the local role of burn for DCTNFT tokens
const DCTRoleNFTBurn = "DCTRoleNFTBurn"

// DCTRoleNFTAddURI is the constant string for the local role of adding a URI for DCTNFT tokens
const DCTRoleNFTAddURI = "DCTRoleNFTAddURI"

// DCTRoleNFTUpdateAttributes is the constant string for the local role of updating attributes for DCTNFT tokens
const DCTRoleNFTUpdateAttributes = "DCTRoleNFTUpdateAttributes"

// DCTRoleTransfer is the constant string for the local role to transfer DCT , only for special tokens
const DCTRoleTransfer = "DCTTransferRole"

// DCTType defines the possible types in case of DCTtokens
type DCTType uint32

const (
	// Fungible defines the token type for DCTfungible tokens
	Fungible DCTType = iota
	// NonFungible defines the token type for DCTnon fungible tokens
	NonFungible
)

// FungibleDCT defines the string for the token type of fungible DCT
const FungibleDCT = "FungibleDCT"

// NonFungibleDCT defines the string for the token type of non fungible DCT
const NonFungibleDCT = "NonFungibleDCT"

// SemiFungibleDCT defines the string for the token type of semi fungible DCT
const SemiFungibleDCT = "SemiFungibleDCT"

// MaxRoyalty defines 100% as uint32
const MaxRoyalty = uint32(10000)

// RelayedTransaction is the key for the electronic meta/gassless/relayed transaction standard
const RelayedTransaction = "relayedTx"

// RelayedTransactionV2 is the key for the optimized electronic meta/gassless/relayed transaction standard
const RelayedTransactionV2 = "relayedTxV2"

// SCDeployInitFunctionName is the key for the function which is called at smart contract deploy time
const SCDeployInitFunctionName = "_init"

// ProtectedKeyPrefix is the key prefix which is protected from writing in the trie - only for special builtin functions
const ProtectedKeyPrefix = "E" + "L" + "R" + "O" + "N" + "D"

// DelegationSystemSCKey is the key under which there is data in case of system delegation smart contracts
const DelegationSystemSCKey = "delegation"

// DCTKeyIdentifier is the key prefix for dct tokens
const DCTKeyIdentifier = "dct"

// DCTRoleIdentifier is the key prefix for dct role identifier
const DCTRoleIdentifier = "role"

// DCTNFTLatestNonceIdentifier is the key prefix for dct latest nonce identifier
const DCTNFTLatestNonceIdentifier = "nonce"

// GuardiansKeyIdentifier is the key prefix for guardians
const GuardiansKeyIdentifier = "guardians"

// MaxNumShards represents the maximum number of shards possible in the system
const MaxNumShards = 256

// MinMetaTxExtraGasCost is the constant defined for minimum gas value to be sent in meta transaction
const MinMetaTxExtraGasCost = uint64(1_000_000)

// MaxLeafSize represents maximum amount of data which can be saved under one leaf
const MaxLeafSize = uint64(1 << 26) //64MB

// MaxBufferSizeToSendTrieNodes represents max buffer size to send in bytes used when resolving trie nodes
// Every trie node that has a greater size than this constant is considered a large trie node and should be split in
// smaller chunks
const MaxBufferSizeToSendTrieNodes = 1 << 18 //256KB

// MaxUserNameLength represents the maximum number of bytes a UserName can have
const MaxUserNameLength = 32

// MinLenArgumentsDCTTransfer defines the min length of arguments for the DCTtransfer
const MinLenArgumentsDCTTransfer = 2

// MinLenArgumentsDCTNFTTransfer defines the minimum length for dct nft transfer
const MinLenArgumentsDCTNFTTransfer = 4

// MaxLenForDCTIssueMint defines the maximum length in bytes for the issued/minted balance
const MaxLenForDCTIssueMint = 100

// BaseOperationCostString represents the field name for base operation costs
const BaseOperationCostString = "BaseOperationCost"

// BuiltInCostString represents the field name for built in operation costs
const BuiltInCostString = "BuiltInCost"

// DCTSCAddress is the hard-coded address for dct issuing smart contract
var DCTSCAddress = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 255, 255}

// SCDeployIdentifier is the identifier for a smart contract deploy
const SCDeployIdentifier = "SCDeploy"

// SCUpgradeIdentifier is the identifier for a smart contract upgrade
const SCUpgradeIdentifier = "SCUpgrade"

// WriteLogIdentifier is the identifier for the information log that is generated by a smart contract call/dct transfer
const WriteLogIdentifier = "writeLog"

// SignalErrorOperation is the identifier for the log that is generated when a smart contract is executed but return an error
const SignalErrorOperation = "signalError"

// CompletedTxEventIdentifier is the identifier for the log that is generated when the execution of a smart contract call is done
const CompletedTxEventIdentifier = "completedTxEvent"

// GasRefundForRelayerMessage is the return message for to the smart contract result with refund for the relayer
const GasRefundForRelayerMessage = "gas refund for relayer"
