package schema

const (
	SIOC  = "http://rdfs.org/sioc/ns#"
	SIOCT = "http://rdfs.org/sioc/types#"
)

var (
	SIOCT_MessageBoard = Define(SIOCT + "MessageBoard")
	SIOCT_MailMessage  = Define(SIOCT + "MailMessage")

	SIOC_Post    = Define(SIOC + "Post")
	SIOC_content = Define(SIOC + "content")
	DCT_subject  = Define(DCT + "subject")
	DCT_created  = Define(DCT + "created")
	DCT_creator  = Define(DCT + "creator")
)
