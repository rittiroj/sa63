import React, { FC, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import { makeStyles } from '@material-ui/core/styles';
import { EntUser } from '../../api/models/EntUser';
import {
  // Content,
  Header,
  Page,
  pageTheme,
  // ContentHeader,
} from '@backstage/core';
import { DefaultApi } from '../../api';

const HeaderCustom = {
  minHeight: '120px',
};

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

// interface user {
//   email: string;
//   password: string;
// }

const SignIn: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  // const [users, setUsers] = React.useState<EntUser[]>([])
  const [email, setEmail] = React.useState<EntUser[]>([]);
  const [password, setPassword] = React.useState<EntUser[]>([]);
  
  // const getUsers = async () => {
  //   const u = await http.listUser({ limit: 10, offset: 0 });
  //   setUsers(u);
  // };
  // useEffect(() => {
  //   getUsers();
  // }, []);

  
  return (
    <div className={classes.paper}>
      <Page theme={pageTheme.website}>
        <Header style={HeaderCustom} title={`Medicine Room System`}
          subtitle="กรุณาบันทึกข้อมูลก่อนเข้าสู่ระบบ">
        </Header>
        <form className = {classes.submit} noValidate> 
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            value={email}
           
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            
          >
            Sign In
          </Button>
          <Grid container>
            <Grid item>
              <Link href="/user" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
        </form>
      </Page>
    </div>
  );
};
export default SignIn;