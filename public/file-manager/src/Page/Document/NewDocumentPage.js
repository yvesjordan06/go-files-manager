import API from "../../Infrastructure/network";
import React, {useEffect, useState} from "react";
import {Paper, Select, Typography, Button} from "@material-ui/core";
import Swal from "sweetalert2"
import TextField from "@material-ui/core/TextField";
import makeStyles from "@material-ui/core/styles/makeStyles";
import LinearProgress from "@material-ui/core/LinearProgress";
import {useHistory} from "react-router";

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
    const [document, setDocument] = useState({title: "", reference: "", object: "", receiver_id: 0, file_id: 0})
    const [others, setOthers] = useState([])
    const history = useHistory()

    useEffect(() => {
        API.get("/documents/users").then((response) => {
            setOthers(response.data)
        }).catch(console.log)
    }, [])

    async function doUpload(e) {
        let formData = new FormData();
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
            setDocument((state) => ({...state, file_id: response.data.ID, title: response.data.name}))
        }).catch((e) => {
            Swal.fire({
                titleText: "Could not upload file",
                text: "Please try again",
                icon: "error"
            })
        })
    }


    let [progress, setProgress] = useState(0.0)

    function createDocument(evt){
        evt.preventDefault()


        Swal.fire("Creating document", ).then()
        Swal.showLoading()
        API.post('/document', document,).then(
            (response) => {
                history.replace('/')
                Swal.close()
            }
        )
            .catch((e) => {
                console.log(e)
                let data = e.response.data
                Swal.hideLoading()
                Swal.update({
                    title: "Oops",
                    text: data.error,
                    icon: "error"
                })
            })
    }

    return (
        <Paper className={classes.root}>
            <Typography>{file ? "Uploaded" : `Progress ${progress} %`}</Typography>
            <LinearProgress hidden={!progress} variant={"determinate"} style={{width: '100%'}} value={progress}/>
            <form onSubmit={createDocument} id="uploadForm" encType="multipart/form-data" style={{marginTop: 8, width: '100%'}}>
                <input type="file" id="file" name="file" hidden onChange={doUpload}/>
                <label htmlFor="file">
                    <Button variant="contained" color="primary" component="span">
                        {file ? "Change file" : "Choose file"}
                    </Button>
                </label>
                {file && <React.Fragment>
                    <TextField required hidden fullWidth margin={"normal"} style={{marginTop: 8}} value={document.file_id}
                               label={"Document ID"} aria-readonly={"true"} readonly disabled/>
                    <TextField required value={document.title} margin={"normal"} style={{marginTop: 8}} fullWidth
                               label={"Document title"}
                               onChange={(evt) => setDocument((state) => ({...state, title: evt.target.value}))}/>
                    <TextField required margin={"normal"} style={{marginTop: 8}} fullWidth
                               label={"Document Reference"}
                               onChange={(evt) => setDocument((state) => ({...state, reference: evt.target.value}))}/>

                    <TextField required margin={"normal"} style={{marginTop: 8}} fullWidth
                               label={"Document Object"}
                               onChange={(evt) => setDocument((state) => ({...state, object: evt.target.value}))}/>

                    <Select required margin={"normal"} style={{marginTop: 8}} fullWidth
                            label={"Receiver"}
                            onChange={(evt) => setDocument((state) => ({...state, receiver_id: evt.target.value}))}>

                        {others.map((u) => (<option value={u.id}>{u.name || u.username}</option>))}
                    </Select>

                    <Button type={"submit"} color={"primary"} style={{marginTop: 16}}  variant={"contained"}> Create Document</Button>
                </React.Fragment>}
            </form>
        </Paper>
    )
}

export default NewDocumentPage