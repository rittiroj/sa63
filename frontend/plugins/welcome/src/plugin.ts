import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/Users';
import Login from './components/Login';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/login', Login);
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/user', CreateUser);



  },
});
