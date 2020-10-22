import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntRequisition } from '../../api/models/EntRequisition';
import moment from 'moment';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

export default function ComponentsTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [requisitions, setRequisitions] = useState<EntRequisition[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getRequisitions = async () => {
      const res = await http.listRequisition({ limit: 10, offset: 0 });
      setLoading(false);
      setRequisitions(res);
    };
    getRequisitions();
  }, [loading]);

  const deleteRequisitions = async (id: number) => {
    const res = await http.deleteRequisition({ id: id });
    setLoading(true);
  };

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">No.</TableCell>
            <TableCell align="center">ชื่อยา</TableCell>
            <TableCell align="center">ชื่อคลังยา</TableCell>
            <TableCell align="center">จำนวนยา</TableCell>
            <TableCell align="center">วันที่-เวลา</TableCell>
            <TableCell align="center">ชื่อเภสัชกร</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {requisitions.map(item => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>
              <TableCell align="center">{item.edges?.drug?.name}</TableCell>
              <TableCell align="center">{item.edges?.registerstore?.name}</TableCell>
              <TableCell align="center">{item.amount}</TableCell>
              <TableCell align="center">{moment(item.addedTime).format('DD/MM/YYYY HH:mm')}</TableCell>
              <TableCell align="center">{item.edges?.user?.name}</TableCell>
              <TableCell align="center">
                <Button
                  onClick={() => {
                    deleteRequisitions(item.id);
                  }}
                  style={{ marginLeft: 10 }}
                  variant="contained"
                  color="secondary"
                >
                  Delete
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}