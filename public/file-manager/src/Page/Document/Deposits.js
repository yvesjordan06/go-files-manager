import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Title from './Title';
import {Link} from "react-router-dom";
import Moment from "moment"
import Button from "@material-ui/core/Button";

function preventDefault(event) {
  event.preventDefault();
}

const useStyles = makeStyles({
  depositContext: {
    flex: 1,
  },
});

export default function Deposits() {
    Moment.locale("en")
  const classes = useStyles();
  return (
    <React.Fragment>
      <Title>Welcome</Title>
      <Typography component="p" variant="h4">
          {JSON.parse(localStorage.getItem("user")).name ||JSON.parse(localStorage.getItem("user")).username}
      </Typography>
      <Typography color="textSecondary" className={classes.depositContext}>
        Today : {Moment(new Date()).format("dddd MMM DD, YYYY")}
      </Typography>
      <div>
        <Link color="primary" to={"/new"}>
            <Button variant={"contained"} color={"primary"}>
                Add a new document
            </Button>

        </Link>
      </div>
    </React.Fragment>
  );
}
