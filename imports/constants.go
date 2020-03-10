package includes

/*
 *  This file is used to hold things that should be included  in multiple files.
 *  In our case these are common to both the client and server.
 *
 *  Do not forget that ONLY things that start with a capital letter are exported
 *
 */
const (
	ListenHost = "localhost" // The default name of the target host
	ListenPort = "50001"     // The port the server is listening on
)
