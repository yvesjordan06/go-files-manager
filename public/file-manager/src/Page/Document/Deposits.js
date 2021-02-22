import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Title from './Title';
import {Link} from "react-router-dom";

function preventDefault(event) {
  event.preventDefault();
}

const useStyles = makeStyles({
  depositContext: {
    flex: 1,
  },
});

export default function Deposits() {
  const classes = useStyles();
  return (
    <React.Fragment>
      <Title>Welcome</Title>
      <Typography component="p" variant="h4">
          {JSON.parse(localStorage.getItem("user")).username}
      </Typography>
      <Typography color="textSecondary" className={classes.depositContext}>
        Today : {new Date().getDate() < 10 ? '0' : ''}{new Date().getDate()}-{new Date().getMonth() < 9 ? '0' : ''}{new Date().getMonth()+1}-{new Date().getFullYear()}
      </Typography>
      <div>
        <Link color="primary" to={"/new"}>
          Add a new document
        </Link>
      </div>
    </React.Fragment>
  );
}
