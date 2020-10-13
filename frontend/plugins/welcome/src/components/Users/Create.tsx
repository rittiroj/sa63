import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import React, { useEffect, FC } from 'react';
import { EntUser } from '../../api/models/EntUser';
import { EntRequisition } from '../../api/models/EntRequisition';
import { EntRegisterStore } from '../../api/models/EntRegisterStore';
import { EntDrug } from '../../api/models/EntDrug';
import { DefaultApi } from '../../api/apis';
import { Alert } from '@material-ui/lab';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    textField: {
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
      width: 200,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  }),
);
interface requisition {
  user: number;
  drug: number;
  registerstore: number;
  requisition: string;
}

const Requisition: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [requisition, setRequisition] = React.useState<Partial<requisition>>({});
  const [drugs, setDrugs] = React.useState<EntDrug[]>([]);
  const [registerstores, setRegisterstores] = React.useState<EntRegisterStore[]>([]);
  const [users, setUsers] = React.useState<EntUser[]>([]);
  const [alert, setAlert] = React.useState(true);


  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof Requisition;
    const { value } = event.target;
    setRequisition({ ...requisition, [name]: value });
  };

  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUsers(res);
  };

  const getDrugs = async () => {
    const res = await http.listDrug({ limit: 10, offset: 0 });
    setDrugs(res);
  };

  const getregisterstore = async () => {
    const res = await http.listRegisterstore({ limit: 10, offset: 0 });
    setRegisterstores(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getDrugs();
    getUsers();
    getregisterstore();
  }, []);

  function save() {
    console.log(requisition)
  }

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          {/* <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
          </IconButton> */}
          <Typography variant="h6" className={classes.title}>
            ระบบเบิกยาจากคลังสำหรับห้องยา
          </Typography>
        </Toolbar>
      </AppBar>
      <Content>
        <Typography variant="h2" className={classes.title}>
          ใบบันทึกการเบิกยาสำหรับห้องยา
          </Typography>

      </Content>
      <div className={classes.root}>
        <form noValidate autoComplete="off"></form>
        <Grid
          container
          direction="column"
          justify="space-around"
          alignItems="center"
        >
          {/* ชื่อยา */}
          <Grid item xs>
            <FormControl className={classes.formControl}>
              <InputLabel >ชื่อยา</InputLabel>
              <Select
                value={requisition.drug}
                onChange={handleChange}
              >
                {drugs.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          {/* คลังยา */}
          <Grid item xs>
            <FormControl className={classes.formControl}>
              <InputLabel>คลังยา</InputLabel>
              <Select
                value={requisition.registerstore}
                onChange={handleChange}
              >
                {registerstores.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          {/* จำนวนยา */}
          <Grid item xs>
            <form className={classes.root} noValidate autoComplete="off">
              <TextField
                onChange={handleChange}

                label="จำนวนยา"
                variant="outlined"
              />
            </form>
          </Grid>

          {/* วันที่-เวลา */}
          <Grid item xs>
            <form className={classes.container} noValidate>
              <TextField
                label="วันที่-เวลา"
                type="datetime-local"
                defaultValue="2020-08-27T08:00"
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
              />
            </form>
          </Grid>

          {/* ชื่อเภสัช */}
          <Grid item xs>

            <FormControl className={classes.formControl}>
              <InputLabel>ชื่อเภสัชกร</InputLabel>
              <Select
                value={requisition.user}
                onChange={handleChange}
              >
                {users.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          {/* ปุ่ม */}
          <Grid container direction="row" justify="center" alignItems="center">
            <Grid item>
              <Button
                variant="contained"
                color="primary"
                onClick={() => {

                }}
              >
                บันทึกข้อมูล
              </Button>
            </Grid>
            <Grid item>
              <Button variant="contained" color="secondary">
                ยกเลิก
              </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                กลับ
              </Button>
            </Grid>
          </Grid>
        </Grid>

      </div>
    </div>
  );
}
export default Requisition;