import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Requisition from './components/Requisition';
import Login from './components/Login';
import Tables from './components/Table';



export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/requisition', Requisition);
    router.registerRoute('/tables', Tables);
    


  },
});
