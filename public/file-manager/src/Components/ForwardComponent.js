import React, {useEffect, useState} from "react";
import Paper, {Select, Typography} from "@material-ui/core";
import API from "../Infrastructure/network";
import Swal from "sweetalert2";

export default function ForwardComponent({documentID}){

    const [receiver, setReceiver] = useState(null)
    const [others, setOthers] = useState([])

    useEffect(() => {
        API.get("/documents/users").then((response) => {
            setOthers(response.data)
        }).catch(console.log)
    }, [])

    function createDocument(evt){
        evt.preventDefault()


        Swal.fire("Sharing document", ).then()
        Swal.showLoading()
        API.post("/documents/"+documentID+"/share", {to: receiver},).then(
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
        <Paper>
            <Typography>Foward document to</Typography>
            <Select required margin={"normal"} style={{marginTop: 8}} fullWidth
                    label={"Receiver"}
                    onChange={(evt) => setReceiver(evt.target.value)}>

                {others.map((u) => (<option value={u.id}>{u.name || u.username}</option>))}
            </Select>
        </Paper>
    )
}