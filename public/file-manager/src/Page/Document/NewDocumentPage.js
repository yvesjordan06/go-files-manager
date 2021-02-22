import API from "../../Infrastructure/network";
import React, {useState} from "react";
import {Container, Paper, Typography} from "@material-ui/core";
import Swal from "sweetalert2"
import TextField from "@material-ui/core/TextField";
import makeStyles from "@material-ui/core/styles/makeStyles";
import CircularProgress from "@material-ui/core/CircularProgress";
import LinearProgress from "@material-ui/core/LinearProgress";
const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
        flexWrap: 'wrap',
        padding: 16
    },
    textField: {
        marginLeft: theme.spacing(1),
        marginRight: theme.spacing(1),
        width: '25ch',
    },
}));
function NewDocumentPage() {
    const classes = useStyles();
    const [file, setFileID] = useState(null)
    const [document, setDocument]= useState({})

    async function doUpload(e) {
        var formData = new FormData();
        formData.append("file", e.target.files[0]);
        let headers = {
            'Content-Type': 'multipart/form-data'
        }

        API.post("files/upload", formData, {
            onUploadProgress: progressEvent => {
                console.log(progressEvent)
                setProgress(
                    parseFloat(((progressEvent.loaded / progressEvent.total) * 100).toFixed(2))
                )
            }
        }).then((response) => {
            console.log(response.data)
            setFileID(response.data)
            setDocument({fileID: response.data.ID, name: response.data.name})
        }).catch((e) => {
            Swal.fire({
                titleText: "Could not upload file",
                text: "Please try again",
                icon: "error"
            })
        })
    }


    let [progress, setProgress] = useState(0.0)

    return (
        <Paper className={classes.root}>
            <Typography>Progress {progress} %</Typography>
            <LinearProgress hidden={!progress} variant={"determinate"} style={{width:'100%'}} value={progress} />
            <form id="uploadForm" encType="multipart/form-data" style={{marginTop:8, width:'100%'}}>
                <input type="file" id="file" name="file"  onChange={doUpload}/>

                {file && <React.Fragment>
                    <TextField fullWidth margin={"normal"} style={{marginTop: 8}} value={file.ID} label={"Document ID"} aria-readonly={"true"} readonly disabled/>
                    <TextField value={file.name} margin={"normal"} style={{marginTop: 8}} fullWidth label={"FileName"}
    onChange={(evt) => setFileID({...file, name: evt.target.value})}/>
                </React.Fragment>}
            </form>
        </Paper>
    )
}

export default NewDocumentPage