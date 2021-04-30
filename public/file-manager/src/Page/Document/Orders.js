import React, {useEffect, useState} from 'react';
import {makeStyles} from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Title from './Title';
import API, {baseURL} from "../../Infrastructure/network";
import IconButton from "@material-ui/core/IconButton";
import {Clear, CloudDownload, Reply} from "@material-ui/icons";
import Moment from "moment"
import {Link} from "react-router-dom";
// Generate Order Data
function createData(id, date, name, shipTo, paymentMethod, amount) {
    return {id, date, name, shipTo, paymentMethod, amount};
}

const rows = [
    createData(0, '16 Mar, 2019', 'Elvis Presley', 'Tupelo, MS', 'VISA ⠀•••• 3719', 312.44),
    createData(1, '16 Mar, 2019', 'Paul McCartney', 'London, UK', 'VISA ⠀•••• 2574', 866.99),
    createData(2, '16 Mar, 2019', 'Tom Scholz', 'Boston, MA', 'MC ⠀•••• 1253', 100.81),
    createData(3, '16 Mar, 2019', 'Michael Jackson', 'Gary, IN', 'AMEX ⠀•••• 2000', 654.39),
    createData(4, '15 Mar, 2019', 'Bruce Springsteen', 'Long Branch, NJ', 'VISA ⠀•••• 5919', 212.79),
];

function preventDefault(event) {
    event.preventDefault();
}

const useStyles = makeStyles((theme) => ({
    seeMore: {
        marginTop: theme.spacing(3),
    },
}));

export default function Orders() {
    Moment.locale("en")
    const classes = useStyles();
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState([]);

    // Note: the empty deps array [] means
    // this useEffect will run once
    // similar to componentDidMount()
    useEffect(() => {
        API.get("documents")

            .then(
                (result) => {
                    setIsLoaded(true);
                    setItems(result.data);
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
    }, [])


    return (
        <React.Fragment>
            <Title>Created Documents</Title>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Date</TableCell>
                        <TableCell>Title</TableCell>
                        <TableCell>Object</TableCell>
                        <TableCell>Reference</TableCell>
                        <TableCell align="right">Actions</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {items.map((row) => (
                        <TableRow key={row.id}>
                            <TableCell>{Moment(row.created_at).format("MMM DD, YYYY - HH:mm")}</TableCell>
                            <TableCell><Link to={"documents/"+row.id} style={{textDecoration: "none"}}>{row.title}</Link></TableCell>
                            <TableCell>{row.object}</TableCell>
                            <TableCell>{row.reference}</TableCell>


                            <TableCell align="right">
                                <IconButton size={"small"} variant={"contained"} style={{marginLeft: 4}}
                                            color={"primary"}><Reply/></IconButton>
                                <IconButton size={"small"} variant={"contained"} color={"secondary"}><Clear/></IconButton>
                                <IconButton size={"small"} download href={baseURL + "/files/" + row.file_id}
                                            variant={"contained"} color={"primary"}><CloudDownload/></IconButton>
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
            <div className={classes.seeMore}>

            </div>
        </React.Fragment>
    );
}
