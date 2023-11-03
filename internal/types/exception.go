package types

type Exception any

type (
	Fatal Exception

	CryptError Exception
	IOError    Exception

	JsonBindingError        Exception
	QueryBindingError       Exception
	RequestDataInvalidError Exception
	ResourceNotFoundError   Exception
	NotPermittedError       Exception

	JWTSignError  Exception
	JWTParseError Exception

	DBConnectError       Exception
	DBIOError            Exception
	DBDuplicateDataError Exception
	HttpUnknownError     Exception

	StatusInternalServerError Exception
	StatusConflict            Exception
	StatusUnauthorized        Exception
)
