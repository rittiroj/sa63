// import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  // ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import TextField from '@material-ui/core/TextField'
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Grid from '@material-ui/core/Grid';
import React, { useEffect, FC } from 'react';
import { EntUser } from '../../api/models/EntUser';
import { EntRegisterStore } from '../../api/models/EntRegisterStore';
import { EntDrug } from '../../api/models/EntDrug';
import { DefaultApi } from '../../api/apis';
import Swal from 'sweetalert2';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 300,
      minHeight: 100,
      marginTop: '5%',
      marginLeft: 40,

    },
    container: {
      display: 'flex',
      flexWrap: 'wrap',
      minWidth: 300,


    },
    textField: {
      marginLeft: 40,
      marginRight: theme.spacing(1),
      width: 200,
      minWidth: 300,
      minHeight: 100,

    },
    menuButton: {

      marginLeft: 500,
    },
    title: {
      flexGrow: 1,

    },
    Bottom: {
      minWidth: 300,
      minHeight: 100,
    }

  }),
);

interface requisition {
  User: number;
  Registerstore: number;
  Drug: number;
  Amount: number;
  Added: String;
}

const Requisition: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [requisition, setRequisition] = React.useState<Partial<requisition>>({});
  const [drugs, setDrugs] = React.useState<EntDrug[]>([]);
  const [registerstores, setRegisterstores] = React.useState<EntRegisterStore[]>([]);
  const [users, setUsers] = React.useState<EntUser[]>([]);
  // const [alert, setAlert] = React.useState(true);


  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'center',
    showConfirmButton: false,
    timer: 1500,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });


  const handleAmountChange = (
    event: React.ChangeEvent<{ name: string; value: number }>,) => {
    const name = event.target.name as keyof typeof Requisition;
    const { value } = event.target;
    setRequisition({ ...requisition, [name]: +value });
    // console.log(requisition);
  };
  const handleDateChange = (
    event: React.ChangeEvent<{ name: string; value: string }>,) => {
    const name = event.target.name as keyof typeof Requisition;
    const { value } = event.target;
    console.log('date select: ', value, typeof value) // show date from event.target.value
    setRequisition({ ...requisition, [name]: value });
    // console.log(requisition);
  };
  const handleChange = (
    event: React.ChangeEvent<{ name: string; value: unknown }>,) => {
    // console.log(requisition);
    const name = event.target.name as keyof typeof Requisition;
    const { value } = event.target;
    setRequisition({ ...requisition, [name]: value });
    // console.log(requisition);
  };


  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUsers(res);
  };

  const getDrugs = async () => {
    const res = await http.listDrug({ limit: 10, offset: 0 });
    setDrugs(res);
  };

  const getRegisterstores = async () => {
    const res = await http.listRegisterstore({ limit: 10, offset: 0 });
    setRegisterstores(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getDrugs();
    getUsers();
    getRegisterstores();
    console.log(requisition)
  }, [requisition]);

  // clear input form
  function clear() {
    setRequisition({});
  }

  function save() {
    requisition.Added += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/requisitions';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(requisition),
    };
    // console.log('date: ', requisition.Added)
    console.log('requisition: ', requisition)
    // console.log(requisition); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        // setRequisition({ ...requisition, Added: requisition.Added });
        // console.log(added_time)
        console.log(data); // last data
        // alert(JSON.stringify(requisition))
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }


  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            ระบบเบิกยาจากคลังสำหรับห้องยา
            </Typography>
          <Typography variant="h6" className={classes.title}>
            {/* <div style={{ marginLeft: 50 }}>ไทยแลนด์ อิอิ</div> */}

          </Typography>

        </Toolbar>
      </AppBar>

      <Content>
        <Typography variant="h3" className={classes.title}>
          ระบบเบิกยาจากคลังสำหรับห้องยา
            </Typography>
        {/* ยา */}
        <div className={classes.formControl}>
          <form noValidate autoComplete="off">
            <Grid item xs>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ชื่อยา</InputLabel>
                <Select
                  name="Drug"
                  label="ชื่อยา"
                  value={requisition.Drug || ''}
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

              <FormControl variant="outlined" className={classes.formControl} >
                <InputLabel>คลังยา</InputLabel>
                <Select
                  name="Registerstore"
                  label="คลังยา"
                  value={requisition.Registerstore || ''}
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
              <FormControl variant="outlined" className={classes.textField} >
                {/* <form className={classes.textField} noValidate autoComplete="off"> */}
                <TextField
                  name="Amount"
                  label="จำนวนยา"
                  id="Amount"
                  type="number"
                  value={requisition.Amount || ''}
                  onChange={handleAmountChange}

                />
                {/* </form> */}
              </FormControl>
            </Grid>

            {/* วันที่ */}
            <Grid item xs>

              {/* <form className={classes.container} noValidate> */}
              <TextField
                name="Added"
                label="วันที่-เวลา"
                type="datetime-local"
                value={requisition.Added || ''} // (undefined || '') = ''
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                onChange={handleDateChange}
              />
              {/* </form> */}
            </Grid>

            <Grid item xs>
              {/* ชื่อเภสัชกร */}
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ชื่อเภสัชกร</InputLabel>
                <Select
                  name="User"
                  label="ชื่อเภสัชกร"
                  value={requisition.User || ''}
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

            <Button variant="contained" color="primary" onClick={save} style={{ marginLeft: 100 }}>บันทึกข้อมูล</Button>
            <Button variant="contained" color="secondary" onClick={clear} style={{ marginLeft: 10 }}>ยกเลิก</Button>

          </form>
        </div>
      </Content>
    </div>

  );
};

export default Requisition;
