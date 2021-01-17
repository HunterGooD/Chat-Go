import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Messages from '../views/Messages.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    meta: {
      auth: true,
    },
    component: Home
  },
  {
    path: '/messages',
    name: 'Messages',
    meta: {
      auth: true,
    },
    component: Messages
  },
  {
    path: '/signin',
    name: 'Signin',
    meta: {
      auth: false,
    },
    component: Home
  },
  {
    path: '/signup',
    name: 'Signup',
    meta: {
      auth: false,
    },
    component: Home
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
