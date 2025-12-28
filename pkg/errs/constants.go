package errs

// Error Type
type ErrorType string

const (
	OTHER               ErrorType = "OTHER"               // Unclassified error. This value is not printed in the error message.
	INVALID             ErrorType = "INVALID"             // Invalid operation for this type of item.
	IO                  ErrorType = "IO"                  // External I/O error such as network failure.
	EXIST               ErrorType = "EXIST"               // Item already exists.
	NOT_EXIST           ErrorType = "NOT_EXIST"           // Item does not exist.
	PRIVATE             ErrorType = "PRIVATE"             // Information withheld.
	INTERNAL            ErrorType = "INTERNAL"            // Internal error or inconsistency.
	BROKEN_LINK         ErrorType = "BROKEN_LINK"         // Link target does not exist.
	DATABASE            ErrorType = "DATABASE"            // Error from database.
	VALIDATION          ErrorType = "VALIDATION"          // Input validation error.
	UNANTICIPATED       ErrorType = "UNANTICIPATED"       // Unanticipated error.
	INVALID_REQUEST     ErrorType = "INVALID_REQUEST"     // Invalid Request
	SMTPERROR           ErrorType = "SMTP_ERROR"          // SMTP Error
	EMAILERROR          ErrorType = "EMAIL_ERROR"         // SMTP Error
	STORAGE             ErrorType = "STORAGE"             // Storage Error
	INSUFFICIENT_ACCESS ErrorType = "INSUFFICIENT_ACCESS" // Insufficient Access

	//Business Message

	CUSTOMER_DUPLICATE_TRANSACTION ErrorType = "CUSTOMER_DUPLICATE_TRANSACTION" // Duplicate Transaction
)

// Error Message
const (
	MALFORMED_REQUEST = "MALFORMED_REQUEST" //Malformed Request Body
	VALIDATION_ERROR  = "VALIDATION_ERROR"  //Validation Error in Request
	INVALID_CREDS     = "INVALID_CREDS"     //Incorrect Email / Password
	MAPPING_ERROR     = "MAPPING_ERROR"     //Error while mapping data to model
	SMTP_ERROR        = "SMTP_ERROR"        //Error while mapping data to model
	INVALID_USER      = "INVALID_USER"      //Invalid User
	LOGIN_DISABLED    = "LOGIN_DISABLED"    //Invalid User

	// FORMAT ERRORS
	INVALID_AMOUNT_FORMAT = "INVALID_AMOUNT_FORMAT" //Incorrect Email / Password
	INVALID_DATE_FORMAT   = "INVALID_DATE_FORMAT"   //Incorrect Email / Password
	INCORRECT_PARAMETER   = "INCORRECT_PARAMETER"   //Missing/Incorrect parameter

	//ERROR
	JWT_ERROR = "JWT_ERROR" //JWT Signing Error

	//Business Messgages
	DUPLICATE_CANDIDATE = "DUPLICATE_CANDIDATE" // Duplicate Candidate
	FILE_UPLOAD_ERROR   = "FILE_UPLOAD_ERROR"   // File Upload Error
)
