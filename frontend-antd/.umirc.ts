import { defineConfig } from '@umijs/max';

export default defineConfig({
  antd: {},
  access: {},
  model: {},
  initialState: {},
  request: {},
  layout: {
    title: 'XAVIER',
  },
  routes: [
    {
      path: '/',
      redirect: '/Token',
    },
    {
      name: 'Token',
      path: '/token',
      component: './Token',
    },
    {
      name: 'Call Log',
      path: '/call-log',
      component: './CallLog',
    },
    {
      name: 'Alias',
      path: '/alias',
      component: './Alias',
    },
  ],
  npmClient: 'npm',
  proxy: {
    '/api/v1': {
      // target: 'http://192.168.123.11:8081',
      target: 'http://127.0.0.1:4523/m1/5903528-5590418-default',
      changeOrigin: true,
    },
  },
  favicons: ['https://www.zxavier.com/avatar.webp'],
});
