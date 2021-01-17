import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false

router.beforeEach((to, from, next) => {

  if (to.meta.auth && !router.app.isAuth) return next("/signin")
  if ((to.path == "/signin" || to.path == "/signup") && router.app.isAuth) return next("/")

  return next()
});

new Vue({
  data() {
    return {
      isAuth: false,
    }
  },
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
