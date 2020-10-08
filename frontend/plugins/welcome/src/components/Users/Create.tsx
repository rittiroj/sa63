import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
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
import { Content } from '@backstage/core';
import Card from '@material-ui/core/Card';
import { Link as RouterLink } from 'react-router-dom';
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      '& > *': {
        margin: theme.spacing(1),
        flexGrow: 1,
      },
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
    card: {
      textAlign: 'center',
      justifyContent: 'center',
      alignContent: 'center',
      height: 705,
    },
  }),
);

export default function ButtonAppBar() {
  const classes = useStyles();
  const [drug, setDrug] = React.useState('');
  const [stock, setStock] = React.useState('');
  const [employee, setEmployee] = React.useState('');
  const [unit, setUnit] = React.useState(0);

  const handleChangeDrug = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDrug(event.target.value as string);
  };

  const handleChangeStock = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStock(event.target.value as string);
  };

  const handleChangeEmployee = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setEmployee(event.target.value as string);
  };

  const handleChangeUnit = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUnit(event.target.value as number);
  };

  const handleSubmit = () => {
    alert('บันทึกสำเร็จ');
    console.log(drug, stock, employee, unit);
  };

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
      <Card className={classes.card}>
        <Content>
          <Typography variant="h2" className={classes.title}>
            ใบบันทึกการเบิกยาสำหรับห้องยา
          </Typography>
        </Content>
        <Grid
          container
          direction="column"
          justify="space-around"
          alignItems="center"
        >
          {/* ชื่อยา */}
          <Grid item xs>
            <FormControl className={classes.formControl}>
              <InputLabel id="demo-simple-select-label">ชื่อยา</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={drug}
                onChange={handleChangeDrug}
              >
                <MenuItem value={'para'}>para</MenuItem>
                <MenuItem value={'ยาแก้ไอ'}>ยาแก้ไอ</MenuItem>
                <MenuItem value={'ยาแก้แพ้'}>ยาแก้แพ้</MenuItem>
              </Select>
            </FormControl>
          </Grid>

          {/* คลังยา */}
          <Grid item xs>
            <FormControl className={classes.formControl}>
              <InputLabel id="demo-simple-select-label">คลังยา</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={stock}
                onChange={handleChangeStock}
              >
                <MenuItem value={'คลังภายใน'}>คลังภายใน</MenuItem>
                <MenuItem value={'คลังภายนอก'}>คลังภายนอก</MenuItem>
              </Select>
            </FormControl>
          </Grid>

          {/* จำนวนยา */}
          <Grid item xs>
            <form className={classes.root} noValidate autoComplete="off">
              <TextField
                value={unit}
                onChange={handleChangeUnit}
                id="outlined-basic"
                label="จำนวนยา"
                variant="outlined"
              />
            </form>
          </Grid>

          {/* วันที่-เวลา */}
          <Grid item xs>
            <form className={classes.container} noValidate>
              <TextField
                id="datetime-local"
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
              <InputLabel id="demo-simple-select-label">ชื่อเภสัชกร</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={employee}
                onChange={handleChangeEmployee}
              >
                <MenuItem value={'ปลั๊ก'}>ปลั๊ก</MenuItem>
                <MenuItem value={'พู่'}>พู่</MenuItem>
                <MenuItem value={'บอล'}>บอล</MenuItem>
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
                  handleSubmit();
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
      </Card>
    </div>
  );
}
