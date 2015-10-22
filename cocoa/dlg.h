typedef struct {
	char* msg;
	char* title;
} AlertDlgParams;

typedef struct {
	int save; /* non-zero => save dialog, zero => open dialog */
	char* buf; /* buffer to store selected file */
	int nbuf; /* number of bytes allocated at buf */
	char* title; /* title for dialog box (can be nil) */
} FileDlgParams;

typedef enum {
	DLG_OK,
	DLG_CANCEL,
	DLG_URLFAIL,
} DlgResult;

DlgResult alertDlg(AlertDlgParams*);
DlgResult fileDlg(FileDlgParams*);
