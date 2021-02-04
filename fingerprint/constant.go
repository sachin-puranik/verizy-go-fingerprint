package fingerprint

const (

	//FingerPrintStartCode - Baotou start byte
	FINGERPRINT_STARTCODE = 0xEF01

	//FingerPrintCommandPacket - Packet identification
	FINGERPRINT_COMMANDPACKET = 0x01

	FINGERPRINT_ACKPACKET     = 0x07
	FINGERPRINT_DATAPACKET    = 0x02
	FINGERPRINT_ENDDATAPACKET = 0x08

	//Instruction codes
	FINGERPRINT_VERIFYPASSWORD      = 0x13
	FINGERPRINT_SETPASSWORD         = 0x12
	FINGERPRINT_SETADDRESS          = 0x15
	FINGERPRINT_SETSYSTEMPARAMETER  = 0x0E
	FINGERPRINT_GETSYSTEMPARAMETERS = 0x0F
	FINGERPRINT_TEMPLATEINDEX       = 0x1F
	FINGERPRINT_TEMPLATECOUNT       = 0x1D

	FINGERPRINT_READIMAGE = 0x01

	//Note: The documentation mean upload to host computer.
	FINGERPRINT_DOWNLOADIMAGE = 0x0A

	FINGERPRINT_CONVERTIMAGE = 0x02

	FINGERPRINT_CREATETEMPLATE = 0x05
	FINGERPRINT_STORETEMPLATE  = 0x06
	FINGERPRINT_SEARCHTEMPLATE = 0x04
	FINGERPRINT_LOADTEMPLATE   = 0x07
	FINGERPRINT_DELETETEMPLATE = 0x0C

	FINGERPRINT_CLEARDATABASE          = 0x0D
	FINGERPRINT_GENERATERANDOMNUMBER   = 0x14
	FINGERPRINT_COMPARECHARACTERISTICS = 0x03

	//Note: The documentation mean download from host computer.
	FINGERPRINT_UPLOADCHARACTERISTICS = 0x09

	//Note: The documentation mean upload to host computer.
	FINGERPRINT_DOWNLOADCHARACTERISTICS = 0x08

	//Parameters of setSystemParameter()
	FINGERPRINT_SETSYSTEMPARAMETER_BAUDRATE       = 4
	FINGERPRINT_SETSYSTEMPARAMETER_SECURITY_LEVEL = 5
	FINGERPRINT_SETSYSTEMPARAMETER_PACKAGE_SIZE   = 6

	//Packet reply confirmations
	FINGERPRINT_OK                  = 0x00
	FINGERPRINT_ERROR_COMMUNICATION = 0x01

	FINGERPRINT_ERROR_WRONGPASSWORD = 0x13

	FINGERPRINT_ERROR_INVALIDREGISTER = 0x1A

	FINGERPRINT_ERROR_NOFINGER  = 0x02
	FINGERPRINT_ERROR_READIMAGE = 0x03

	FINGERPRINT_ERROR_MESSYIMAGE       = 0x06
	FINGERPRINT_ERROR_FEWFEATUREPOINTS = 0x07
	FINGERPRINT_ERROR_INVALIDIMAGE     = 0x15

	FINGERPRINT_ERROR_CHARACTERISTICSMISMATCH = 0x0A

	FINGERPRINT_ERROR_INVALIDPOSITION = 0x0B
	FINGERPRINT_ERROR_FLASH           = 0x18

	FINGERPRINT_ERROR_NOTEMPLATEFOUND = 0x09

	FINGERPRINT_ERROR_LOADTEMPLATE = 0x0C

	FINGERPRINT_ERROR_DELETETEMPLATE = 0x10

	FINGERPRINT_ERROR_CLEARDATABASE = 0x11

	FINGERPRINT_ERROR_NOTMATCHING = 0x08

	FINGERPRINT_ERROR_DOWNLOADIMAGE           = 0x0F
	FINGERPRINT_ERROR_DOWNLOADCHARACTERISTICS = 0x0D

	//Unknown error codes
	FINGERPRINT_ADDRCODE   = 0x20
	FINGERPRINT_PASSVERIFY = 0x21

	FINGERPRINT_PACKETRESPONSEFAIL = 0x0E

	FINGERPRINT_ERROR_TIMEOUT   = 0xFF
	FINGERPRINT_ERROR_BADPACKET = 0xFE

	//Char buffers
	FINGERPRINT_CHARBUFFER1       = 0x01 //Char buffer 1
	FINGERPRINT_CHARBUFFER2       = 0x02 //Char buffer 2
	SMALLEST_RESPONSE_PACKET_SIZE = 12
)
