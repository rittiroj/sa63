import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';

const WelcomePage: FC<{}> = () => {
  const profile = { givenName: 'ห้องยา' };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ระบบ${profile.givenName || ''}`}
        subtitle="ระบบเบิกยาจากคลังสำหรับห้องยา."
      ></Header>
      <Content>
        <ContentHeader title="บันทึกรายการเบิกยาจากคลังสำหรับห้องยา">
          <Link component={RouterLink} to="/user">
            <Button variant="contained" color="primary">
              เพิ่มข้อมูล
            </Button>
          </Link>
        </ContentHeader>
        <ComponanceTable></ComponanceTable>
      </Content>
    </Page>
  );
};

export default WelcomePage;
