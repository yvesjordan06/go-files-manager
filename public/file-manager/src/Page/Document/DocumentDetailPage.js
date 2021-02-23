import {useParams} from "react-router";
import React, {useEffect, useState} from "react"
import API, {baseURL} from "../../Infrastructure/network";
import {
    Avatar,
    Button,
    CircularProgress,
    Container,
    Divider,
    Grid,
    List,
    ListItem,
    ListItemAvatar,
    ListItemText,
    Paper,
    TextField,
    Typography
} from "@material-ui/core";
import {CloseOutlined} from "@material-ui/icons";
import getReadableFileSizeString from "../../Infrastructure/readableFileSize";
import Moment from "moment";

function DocumentDetailPage() {
    let {id: documentID} = useParams()
    const [document, setDocument] = useState({
        title: "", reference: "", object: "", user: "", receiver: "", id: "",
        expeditor: null
    })

    const [loading, setLoading] = useState(true)
    const [loadingComment, setLoadingComment] = useState(true)
    const [comments, setComments] = useState([])
    const [commentText, setCommentText] = useState("")


    function sendComment(){
        if (!commentText) return

        API.post("/documents/" + documentID + "/comment", {text: commentText}).then(
            (response) => {
                setCommentText("")
                setComments((state) => [response.data, ...state])
            },
            console.log
        )
    }

    useEffect(() => {
        API.get("/documents/" + documentID).then(response => {
                setDocument(response.data)
                setLoading(false)
            },
            (err) => {
                console.log(err)
            })


        API.get("/documents/" + documentID + "/comments").then(response => {
                setComments(response.data)
                setLoadingComment(false)
            },
            (err) => {
                console.log(err)
            })
    }, [])

    return (
        loading ? <CircularProgress variant={"indeterminate"}/> : <React.Fragment>
            <Paper style={{padding: "8px 16px"}}>
                <Typography variant={"h3"}>Document : {document.title} <Typography
                    variant={"caption"}>({document.id})</Typography></Typography>


                <Grid container>
                    <Grid item xs={12}>
                        <Typography><b>Object</b> : {document.object}</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Typography><b>Reference</b> : {document.reference}</Typography>
                    </Grid>
                </Grid>

                <Grid style={{marginTop: 16}} spacing={3} container justify={"flex-end"}>
                    <Grid item>
                        <Button variant={"contained"} color={"primary"} download
                                href={baseURL + "/files/" + document.file_id}>Download File</Button>
                    </Grid>
                    <Grid item>
                        <Button variant={"contained"} color={"secondary"} download
                                href={baseURL + "/files/" + document.file_id}>Delete File <CloseOutlined/></Button>
                    </Grid>
                </Grid>
            </Paper>

            <Paper>
                <h1>File Information</h1>
                <Container>
                    <Grid container spacing={3}>
                        <Grid item xs={6}>
                            <Grid container spacing={3} justify={"flex-start"}>
                                <Grid item xs={3}><b>From</b></Grid>
                                <Grid item>{document.expeditor.name || document.expeditor.username}</Grid>
                            </Grid>

                            <Grid container spacing={3} justify={"flex-start"}>
                                <Grid item xs={3}><b>Object</b></Grid>
                                <Grid item>{document.object}</Grid>
                            </Grid>

                            <Grid container spacing={3} justify={"flex-start"}>
                                <Grid item xs={3}><b>File size</b></Grid>
                                <Grid item>{getReadableFileSizeString(document.file.size)}</Grid>
                            </Grid>


                            <Grid container spacing={3} justify={"flex-start"}>
                                <Grid item xs={3}><b>Uploaded on</b></Grid>
                                <Grid item>{Moment(document.created_at).format("MMM DD, YYYY - HH:mm")}</Grid>
                            </Grid>
                        </Grid>

                        <Grid item xs={6}>
                            <Grid container spacing={3} justify={"flex-start"}>
                                <Grid item xs={3}><b>To</b></Grid>
                                <Grid item>{document.receiver.name || document.receiver.username}</Grid>
                            </Grid>

                            <Grid container spacing={3} justify={"flex-start"}>
                                <Grid item xs={3}><b>Reference</b></Grid>
                                <Grid item>{document.reference}</Grid>
                            </Grid>

                            <Grid container spacing={3} justify={"flex-start"}>
                                <Grid item xs={3}><b>File name</b></Grid>
                                <Grid item>{document.file.name}</Grid>
                            </Grid>

                            <Grid container spacing={3} justify={"flex-start"}>
                                <Grid item xs={3}><b>Status</b></Grid>
                                <Grid item>{document.status.toUpperCase()}</Grid>
                            </Grid>
                        </Grid>
                    </Grid>

                </Container>

            </Paper>

            <Paper>
                <h1>Comments</h1>
                <Container>
                    {loadingComment ? <CircularProgress variant={"indeterminate"}/> :
                        <List>
                            <ListItem alignItems="flex-start">
                                <ListItemAvatar>
                                    <Avatar
                                        alt={JSON.parse(localStorage.getItem("user")).name || JSON.parse(localStorage.getItem("user")).username}/>
                                </ListItemAvatar>
                                <ListItemText
                                    primary={
                                        <TextField
                                            id="standard-full-width"
                                            label="Your comment"
                                            style={{margin: 8}}
                                            value={commentText}
                                            onChange={(evt) => {
                                                setCommentText(evt.target.value);
                                            }
                                            }
                                            fullWidth
                                            margin="normal"
                                            InputLabelProps={{
                                                shrink: true,
                                            }}
                                        />
                                    }

                                    secondary={
                                        <React.Fragment>
                                            <Grid container justify={"flex-end"}>
                                                <Button variant={"contained"} color={"primary"} onClick={sendComment}>Send</Button>
                                            </Grid>
                                             </React.Fragment>
                                    }

                                />
                            </ListItem>
                            {
                                comments.map(item => (
                                    <React.Fragment>
                                        <Divider variant="inset" component="li"/>
                                        <ListItem alignItems="flex-start">
                                            <ListItemAvatar>
                                                <Avatar alt={item.user.name}/>
                                            </ListItemAvatar>
                                            <ListItemText
                                                primary={<b>{item.user.username.toUpperCase() || item.user.name}</b>}
                                                secondary={
                                                    <React.Fragment>
                                                        <Typography
                                                            component="span"
                                                            variant="caption"

                                                            color="textPrimary"
                                                        >
                                                            {Moment(item.created_at).format("MMM DD, YYYY - HH:mm")}
                                                        </Typography>
<div>
                                                        <Typography
                                                            component="span"
                                                            variant="body1"

                                                            color="textSecondary"
                                                        >
                                                            {item.text}
                                                        </Typography>
</div>
                                                    </React.Fragment>
                                                }
                                            />
                                        </ListItem>

                                    </React.Fragment>
                                ))
                            }


                        </List>}
                </Container>

            </Paper>
        </React.Fragment>

    )


}

export default DocumentDetailPage