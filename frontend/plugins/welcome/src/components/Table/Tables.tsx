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
import { EntDrug } from '../../api/models/EntDrug';
import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [drugs, setDrugs] = useState<EntDrug[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getDrugs = async () => {
      const res = await api.listDrug({ limit: 10, offset: 0 });
      setLoading(false);
      setDrugs(res);
    };
    getDrugs();
  }, [loading]);

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            {/* <TableCell align="center">ชื่อเภสัช</TableCell> */}
            <TableCell align="center">ชื่อยา</TableCell>
            {/* <TableCell align="center">จำนวนยา</TableCell>
            <TableCell align="center">เภสัชกร</TableCell> */}
          </TableRow>
        </TableHead>
        <TableBody>
          {drugs.map(item => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.drugsName}</TableCell>
              {/* <TableCell align="center">{item.value}</TableCell> */}

              <TableCell align="center">
                <Button
                  onClick={() => {}}
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
