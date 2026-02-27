package exitCodes

const (
	// SuccessCode is the exit code for a successful execution of the program
	SuccessCode = 0;

	// RequestFailureCode is the exit code for any network bound error, timeout, dns resolution failed, etc
	// Note: a non 2xx response code from the api also count in this category
	RequestFailureCode int = 1;
	
	// CommandErrorCode is the exit code for any error related to the user input, such as an invalid card category 
	// or an invalid command line argument
	CommandErrorCode int = 2;

	// ApiResponseParsingFailureCode is the exit code for any failure to parse the api response as json
	ApiResponseParsingFailureCode int = 3;

	// FileIOErrorCode is the exit code for any File I/O bound error, unable to read/write to a file
	// Note: permission denied is the only error balafetch might get for this
	FileIOErrorCode int = 4;

	// LogFileSetupFailureCode is the exit code for any error related to the setup of the log file, 
	// such as unable to create or open the log file
	LogFileSetupFailureCode int = 50;
)